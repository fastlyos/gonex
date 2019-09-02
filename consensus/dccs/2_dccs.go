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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
)

const (
	inmemorySealingQueues = 16
	inmemoryExtDatas      = 16
	inmemoryAnchorExtras  = 4
)

var (
	errInvalidNonce              = errors.New("Invalid block nonce as distant from the last sealer application block")
	errUnknownPreviousSealer     = errors.New("Unknown previous block sealer")
	errInvalidExtendedDataLength = errors.New("Invalid extended data length")
	errExtendedDataMismatch      = errors.New("Extended data does not matches")
)

// Init the second hardfork of DCCS consensus
func (d *Dccs) init2() *Dccs {
	d.init1()
	d.sealingQueueCache, _ = lru.NewARC(inmemorySealingQueues)
	d.extDataCache, _ = lru.NewARC(inmemoryExtDatas)
	d.anchorExtraCache, _ = lru.NewARC(inmemoryAnchorExtras)
	return d
}

// verifyHeader2 checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (d *Dccs) verifyHeader2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
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
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return d.verifyCascadingFields2(chain, header, parents)
}

func (d *Dccs) getBlockNonce(number *big.Int, parent *types.Header) types.BlockNonce {
	if d.config.CoLoaBlock.Cmp(number) == 0 {
		return types.BlockNonce{}
	}
	nonce := parent.Nonce
	b := parent.Hash().Bytes()[:8]
	if b[0] < 256/16 {
		// change the nonce once in 16 blocks, probabilistically
		copy(nonce[:], b)
		log.Trace("New random seed", "nonce", nonce)
	}
	return nonce
}

// verifyCascadingFields2 verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (d *Dccs) verifyCascadingFields2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	parent := getAvailableHeaderByHash(header.ParentHash, header, parents, chain)
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+d.config.Period > header.Time {
		return ErrInvalidTimestamp
	}

	// verify the random seed
	nonce := d.getBlockNonce(header.Number, parent)
	if header.Nonce != nonce {
		log.Error("invalid nonce", "expected", nonce, "actual", header.Nonce)
		return errInvalidNonce
	}

	// verify the cross-link reference to the last sealer application block
	var expectedMixDigest common.Hash
	if d.config.CoLoaBlock.Cmp(header.Number) == 0 {
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
	expectedAnchorBytes, err := d.assembleAnchorExtra(parent, parents, chain)
	if err != nil {
		return err
	}
	expectedAnchorBytesLength := len(expectedAnchorBytes)

	extBytes := header.Extra[extraVanity : len(header.Extra)-extraSeal]
	if len(extBytes) < expectedAnchorBytesLength {
		return errInvalidExtendedDataLength
	}
	if expectedAnchorBytesLength > 0 {
		// anchor data always comes first in the extended extra
		if bytes.Compare(extBytes[:expectedAnchorBytesLength], expectedAnchorBytes) != 0 {
			return errExtendedDataMismatch
		}
	}

	// All basic checks passed, verify the seal and return
	return d.verifySeal2(chain, header, parents)
}

// verifySeal2 checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (d *Dccs) verifySeal2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// Retrieve the sealing queue verify this header
	queue, err := d.getSealingQueue(header.ParentHash, parents, chain)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, d.signatures)
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
	signerDifficulty := queue.difficulty(signer, func(hash common.Hash) *types.Header {
		return getAvailableHeaderByHash(hash, header, parents, chain)
	}, d.signatures)
	if header.Difficulty.Uint64() != signerDifficulty {
		return errInvalidDifficulty
	}
	return nil
}

// getAvailableHeaderByHash returns either:
// + the input header, if hash == header.Hash()
// + chain.GetHeaderByHash(hash) if available
// + the header in parents if available (nessesary for batch headers processing)
func getAvailableHeaderByHash(hash common.Hash, header *types.Header, parents []*types.Header, chain consensus.ChainReader) *types.Header {
	if header != nil && header.Hash() == hash {
		return header
	}
	if h := chain.GetHeaderByHash(hash); h != nil {
		return h
	}
	for _, parent := range parents {
		if parent.Hash() == hash {
			return parent
		}
	}
	return nil
}

func getLinkDest(header *types.Header, parents []*types.Header, chain consensus.ChainReader) *types.Header {
	if header.MixDigest == (common.Hash{}) {
		// hardfork block is linked to itself
		return header
	}
	return getAvailableHeaderByHash(header.MixDigest, header, parents, chain)
}

func (d *Dccs) getAnchorDest(header *types.Header, parents []*types.Header, chain consensus.ChainReader) (*types.Header, error) {
	linkHeader := header
	if !hasAnchorData(header) {
		linkHeader = getLinkDest(header, parents, chain)
	}
	anchorData, err := d.getAnchorData(linkHeader)
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
	return getAvailableHeaderByHash(anchorData.destHash, header, parents, chain), nil
}

// prepareBeneficiary2 gets the beneficiary of signer from smart contract and
// set to header's coinbase to give sealing reward later.
// + check the contract of current state first
// + trace back the previous header for 'just left sealer'
// + if all else fails, the sealer address is kept as reward beneficiary
func (d *Dccs) prepareBeneficiary2(header *types.Header, chain consensus.ChainReader) {
	index := common.BigToHash(common.Big1).String()[2:]
	sealer := "0x000000000000000000000000" + header.Coinbase.String()[2:]
	key := crypto.Keccak256Hash(hexutil.MustDecode(sealer + index))

	number := header.Number.Uint64()

	// try the current active state first
	state, err := chain.State()
	if err == nil && state != nil {
		hash := state.GetState(chain.Config().Dccs.Contract, key)
		if (hash != common.Hash{}) {
			header.Coinbase = common.HexToAddress(hash.Hex())
			return
		}
	}

	// scan the previous signed blocks
	for n := number - 1; n >= number-d.config.LeakDuration; n-- {
		if n < 1 {
			break
		}
		h := chain.GetHeaderByNumber(n)
		if h == nil {
			break
		}
		s, err := ecrecover(h, d.signatures)
		if err != nil {
			log.Error("Unable to recover signature", "err", err)
			return
		}
		if s == d.signer {
			// found the previous sealed block
			header.Coinbase = h.Coinbase
			return
		}
	}
}

// prepare2 implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) prepare2(chain consensus.ChainReader, header *types.Header) error {
	number := header.Number.Uint64()
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// record the distant from the last sealer application block
	header.Nonce = d.getBlockNonce(header.Number, parent)

	d.prepareBeneficiary2(header, chain)

	queue, err := d.getSealingQueue(header.ParentHash, nil, chain)
	if err != nil {
		return err
	}

	// Set the correct difficulty
	difficulty := queue.difficulty(d.signer, chain.GetHeaderByHash, d.signatures)
	header.Difficulty = new(big.Int).SetUint64(difficulty)

	// Ensure the timestamp has the correct delay
	header.Time = parent.Time + d.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	// set the cross-link reference to the last block with anchor data
	if d.config.CoLoaBlock.Cmp(header.Number) == 0 {
		// special handling for hardfork block
		header.MixDigest = common.Hash{}
	} else if hasAnchorData(parent) {
		header.MixDigest = parent.Hash()
	} else {
		header.MixDigest = parent.MixDigest
	}

	anchorBytes, err := d.assembleAnchorExtra(parent, nil, chain)
	if err != nil {
		return err
	}

	// Prepare the start of the header.Extra
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	header.Extra = append(header.Extra, anchorBytes...)
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)
	return nil
}

// finalize2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalize2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.EmptyUncleHash
}

// finalizeAndAssemble2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalizeAndAssemble2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (block *types.Block, err error) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.EmptyUncleHash

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// seal2 implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) seal2(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if d.config.Period == 0 && len(block.Transactions()) == 0 {
		return errWaitTransactions
	}
	// Don't hold the signer fields for the entire sealing procedure
	d.lock.RLock()
	signer, signFn := d.signer, d.signFn
	d.lock.RUnlock()

	queue, err := d.getSealingQueue(header.ParentHash, nil, chain)
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
	offset, err := queue.offset(signer, chain.GetHeaderByHash, d.signatures)
	if err != nil {
		return err
	}
	if offset > 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := d.calcDelayTimeForOffset(offset)
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
