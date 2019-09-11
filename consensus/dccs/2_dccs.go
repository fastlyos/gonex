// Copyright 2019 The gonex Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package dccs implements the proof-of-foundation consensus engine.
package dccs

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vdf"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
)

const (
	inmemorySealingQueues = 16
	inmemoryExtDatas      = 16
	inmemoryAnchorExtras  = 16
	randomSeedSize        = 32
)

var (
	errInvalidNonce              = errors.New("Invalid block nonce as distant from the last seed block")
	errUnknownPreviousSealer     = errors.New("Unknown previous block sealer")
	errInvalidExtendedDataLength = errors.New("Invalid extended data length")
	errInvalidExtendedData       = errors.New("Extended data does not matches")
	errInvalidRandomData         = errors.New("Invalid random data in extra data")
	errInvalidRandomDataSize     = errors.New("Invalid random data size from relayer")
)

// Init the second hardfork of DCCS consensus
func (d *Dccs) init2() *Dccs {
	d.init1()
	d.sealingQueueCache, _ = lru.NewARC(inmemorySealingQueues)
	d.extDataCache, _ = lru.NewARC(inmemoryExtDatas)
	d.anchorExtraCache, _ = lru.NewARC(inmemoryAnchorExtras)
	d.queueShuffler = vdf.NewDelayer(randomSeedSize)
	return d
}

// Context represents the context of a consensus request.
type Context struct {
	head    *types.Header         // the verifying header, nil for preparation
	parents []*types.Header       // the previous headers are being parallel verified, empty for preparation
	chain   consensus.ChainReader // the underlining chain
	engine  *Dccs                 // shared config and caches
}

func NewContext(engine *Dccs, chain consensus.ChainReader) *Context {
	context := Context{
		chain:  chain,
		engine: engine,
	}
	return &context
}

// verifyHeader2 checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (c *Context) verifyHeader2() error {
	header := c.head
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}

	// Don't verify the UncleHash to allow possible velvet upgrade later
	// // Ensure that the block doesn't contain any uncles which are meaningless in Dccs
	// if header.UncleHash != types.EmptyUncleHash {
	// 	return errInvalidUncleHash
	// }

	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil {
			return errInvalidDifficulty
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(c.chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return c.verifyCascadingFields2()
}

func (c *Context) getBlockNonce(parent *types.Header) types.BlockNonce {
	if parent.Number.Uint64()+1 == c.engine.config.CoLoaBlock.Uint64() {
		return types.BlockNonce{}
	}
	return types.EncodeNonce(parent.Nonce.Uint64() + 1)
}

// verifyCascadingFields2 verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (c *Context) verifyCascadingFields2() error {
	header := c.head
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	parent := c.getHeaderByHash(header.ParentHash)
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+c.engine.config.Period > header.Time {
		return ErrInvalidTimestamp
	}

	// verify the cross-link reference to the last sealer application block
	var expectedMixDigest common.Hash
	if c.engine.config.CoLoaBlock.Cmp(header.Number) == 0 {
		expectedMixDigest = common.Hash{}
	} else if hasAnchorData(parent) {
		expectedMixDigest = parent.Hash()
	} else {
		expectedMixDigest = parent.MixDigest
	}
	if header.MixDigest != expectedMixDigest {
		return fmt.Errorf("invalid cross-link digest: number=%v, want=%v, have=%v", header.Number, expectedMixDigest, header.MixDigest)
	}

	// Verify the extended extra data in header.Extra
	expectedAnchorBytes, err := c.assembleAnchorExtra(parent)
	if err != nil {
		return err
	}
	extBytes := header.Extra[extraVanity : len(header.Extra)-extraSeal]
	err = verifyAnchorBytes(extBytes, expectedAnchorBytes)
	if err != nil {
		return err
	}
	extBytes = extBytes[len(expectedAnchorBytes):]
	randomData, n := bytesToRandomData(extBytes)

	// reset the seed block ref on valid vdf output
	nonce := types.BlockNonce{}
	if len(randomData) > 0 {
		if len(randomData) != randomSeedSize {
			return errInvalidRandomDataSize
		}
		// Verify VDF ouput here
		input := c.getChainRandomInput(parent)
		if !c.engine.queueShuffler.Verify(input[:], randomData, c.engine.config.RandomSeedIteration) {
			return errInvalidRandomData
		}
		log.Info("New random data received", "random data", common.Bytes2Hex(randomData))
	} else {
		nonce = c.getBlockNonce(parent)
	}
	extBytes = extBytes[n:]
	// TODO: verify price value in extBytes here

	// verify the seed block ref
	if header.Nonce != nonce {
		log.Error("invalid nonce as seed ref", "want", nonce, "have", header.Nonce)
		return errInvalidNonce
	}

	// All basic checks passed, verify the seal and return
	return c.verifySeal2()
}

// verifySeal2 checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (c *Context) verifySeal2() error {
	header := c.head
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// Retrieve the sealing queue verify this header
	queue, err := c.getSealingQueue(header.ParentHash)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against signers
	signer, err := c.ecrecover(header)
	if err != nil {
		return err
	}
	if !queue.isActive(signer) {
		return errUnauthorizedSigner
	}
	if queue.isRecentlySigned(signer) {
		return errRecentlySigned
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	signerDifficulty := queue.difficulty(signer, c.getHeaderByHash, c.engine.signatures)
	if header.Difficulty.Uint64() != signerDifficulty {
		return errInvalidDifficulty
	}
	return nil
}

// getHeaderByNumber returns either:
// + the context head, if number == head.Number
// + the header in parents if available (nessesary for batch headers processing)
// + chain.GetHeaderByNumber(number), if all else fail
func (c *Context) getHeaderByNumber(number uint64) *types.Header {
	var headerNumber uint64
	if c.head != nil {
		headerNumber = c.head.Number.Uint64()
		if number == headerNumber {
			return c.head
		}
		if number > headerNumber {
			return c.chain.GetHeaderByNumber(number)
		}
	} else {
		if len(c.parents) == 0 {
			return c.chain.GetHeaderByNumber(number)
		}
		headerNumber = c.parents[len(c.parents)-1].Number.Uint64() + 1
	}
	idx := len(c.parents) - int(headerNumber) + int(number)
	if idx >= 0 {
		header := c.parents[idx]
		if header.Number.Uint64() == number {
			return header
		}
		log.Error("invalid parrent number", "expected", number, "actual", header.Number)
	}
	return c.chain.GetHeaderByNumber(number)
}

// getHeaderByHash returns either:
// + the input header, if hash == header.Hash()
// + chain.GetHeaderByHash(hash) if available
// + the header in parents if available (nessesary for batch headers processing)
func (c *Context) getHeaderByHash(hash common.Hash) *types.Header {
	if c.head != nil && c.head.Hash() == hash {
		return c.head
	}
	if h := c.chain.GetHeaderByHash(hash); h != nil {
		return h
	}
	for _, parent := range c.parents {
		if parent.Hash() == hash {
			return parent
		}
	}
	return nil
}

func (c *Context) getChainRandomHeader(parent *types.Header) *types.Header {
	seedNumber := parent.Number.Uint64() - parent.Nonce.Uint64()
	return c.getHeaderByNumber(seedNumber)
}

func (c *Context) getChainRandomInput(parent *types.Header) common.Hash {
	seedHeader := c.getChainRandomHeader(parent)
	return seedHeader.Hash()
}

func (c *Context) getChainRandomSeed(parent *types.Header) (RandomData, error) {
	seedHeader := c.getChainRandomHeader(parent)
	if c.engine.config.CoLoaBlock.Cmp(seedHeader.Number) == 0 {
		// use the sealer digest for hardfork block
		anchorData, err := c.getAnchorData(seedHeader)
		if err != nil {
			return nil, err
		}
		return anchorData.sealersDigest[:], nil
	}
	return c.getRandomData(seedHeader)
}

func (c *Context) getLinkDest(header *types.Header) *types.Header {
	if header.MixDigest == (common.Hash{}) {
		// hardfork block is linked to itself
		return header
	}
	return c.getHeaderByHash(header.MixDigest)
}

func (c *Context) getAnchorDest(header *types.Header) (*types.Header, error) {
	linkHeader := header
	if !hasAnchorData(header) {
		linkHeader = c.getLinkDest(header)
	}
	anchorData, err := c.getAnchorData(linkHeader)
	if err != nil {
		return nil, err
	}
	if anchorData == nil {
		// should never happen
		log.Error("getAnchorHeader returns nil", "number", header.Number)
		return nil, errors.New("getAnchorHeader returns nil")
	}
	if anchorData.destHash == (common.Hash{}) {
		// hardfork block is anchored to itself
		return linkHeader, nil
	}
	return c.getHeaderByHash(anchorData.destHash), nil
}

// prepareBeneficiary2 gets the beneficiary of signer from smart contract and
// set to header's coinbase to give sealing reward later.
// + check the contract of current state first
// + trace back the previous header for 'just left sealer'
// + if all else fails, the sealer address is kept as reward beneficiary
func (c *Context) prepareBeneficiary2(header *types.Header) {
	index := common.BigToHash(common.Big1).String()[2:]
	sealer := "0x000000000000000000000000" + header.Coinbase.String()[2:]
	key := crypto.Keccak256Hash(hexutil.MustDecode(sealer + index))

	number := header.Number.Uint64()

	// try the current active state first
	state, err := c.chain.State()
	if err == nil && state != nil {
		hash := state.GetState(c.chain.Config().Dccs.Contract, key)
		if (hash != common.Hash{}) {
			header.Coinbase = common.HexToAddress(hash.Hex())
			return
		}
	}

	// scan the previous signed blocks
	for n := number - 1; n >= number-c.engine.config.LeakDuration; n-- {
		if n < 1 {
			break
		}
		h := c.chain.GetHeaderByNumber(n)
		if h == nil {
			break
		}
		s, err := c.ecrecover(h)
		if err != nil {
			log.Error("Unable to recover signature", "err", err)
			return
		}
		if s == c.engine.signer {
			// found the previous sealed block
			header.Coinbase = h.Coinbase
			return
		}
	}
}

// prepare2 implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *Context) prepare2(header *types.Header) error {
	number := header.Number.Uint64()
	parent := c.chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// request the new vdf calculation after each zero block nonce
	if parent.Nonce == (types.BlockNonce{}) {
		input := parent.Hash()
		log.Info("Requesting for new random seed calculation", "input", input)
		c.engine.queueShuffler.Request(input[:], c.engine.config.RandomSeedIteration)
	} else {
		// request the first VDF task after the node started
		c.engine.queueShufflerOnce.Do(func() {
			input := c.getChainRandomInput(parent)
			log.Info("Requesting for random seed calculation", "input", input)
			c.engine.queueShuffler.Request(input[:], c.engine.config.RandomSeedIteration)
		})
	}

	c.prepareBeneficiary2(header)

	queue, err := c.getSealingQueue(header.ParentHash)
	if err != nil {
		return err
	}

	// Set the correct difficulty
	difficulty := queue.difficulty(c.engine.signer, c.chain.GetHeaderByHash, c.engine.signatures)
	header.Difficulty = new(big.Int).SetUint64(difficulty)

	// Ensure the timestamp has the correct delay
	header.Time = parent.Time + c.engine.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	// set the cross-link reference to the last block with anchor data
	if c.engine.config.CoLoaBlock.Cmp(header.Number) == 0 {
		// special handling for hardfork block
		header.MixDigest = common.Hash{}
	} else if hasAnchorData(parent) {
		header.MixDigest = parent.Hash()
	} else {
		header.MixDigest = parent.MixDigest
	}

	anchorBytes, err := c.assembleAnchorExtra(parent)
	if err != nil {
		return err
	}

	input := c.getChainRandomInput(parent)
	randomData := RandomData(c.engine.queueShuffler.Peek(input[:], c.engine.config.RandomSeedIteration))
	if len(randomData) > 0 {
		log.Trace("prepare2", "vdfOutput", common.Bytes2Hex(randomData))
		if len(randomData) != randomSeedSize {
			return errInvalidRandomDataSize
		}
		header.Nonce = types.BlockNonce{}
	} else {
		// record the distant from the last sealer application block
		header.Nonce = c.getBlockNonce(parent)
	}

	// Prepare the start of the header.Extra
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	header.Extra = append(header.Extra, anchorBytes...)
	header.Extra = append(header.Extra, randomData.bytes()...)
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)
	return nil
}

// finalize2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (c *Context) finalize2(header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Calculate any block reward for the sealer and commit the final state root
	c.engine.calculateRewards(c.chain, state, header)
	header.Root = state.IntermediateRoot(c.chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.EmptyUncleHash
}

// finalizeAndAssemble2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (c *Context) finalizeAndAssemble2(header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (block *types.Block, err error) {
	// Calculate any block reward for the sealer and commit the final state root
	c.engine.calculateRewards(c.chain, state, header)
	header.Root = state.IntermediateRoot(c.chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.EmptyUncleHash

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// seal2 implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (c *Context) seal2(block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if c.engine.config.Period == 0 && len(block.Transactions()) == 0 {
		return errWaitTransactions
	}
	// Don't hold the signer fields for the entire sealing procedure
	c.engine.lock.RLock()
	signer, signFn := c.engine.signer, c.engine.signFn
	c.engine.lock.RUnlock()

	queue, err := c.getSealingQueue(header.ParentHash)
	if err != nil {
		return err
	}

	if !queue.isActive(signer) {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	if queue.isRecentlySigned(signer) {
		return errRecentlySigned
	}
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	// Find the signer offset
	offset, err := queue.offset(signer, c.chain.GetHeaderByHash, c.engine.signatures)
	if err != nil {
		return err
	}
	if offset > 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := c.engine.calcDelayTimeForOffset(offset)
		delay += wiggle
		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
	}
	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeClique, DccsRLP(header))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

func (c *Context) ecrecover(header *types.Header) (common.Address, error) {
	return ecrecover(header, c.engine.signatures)
}
