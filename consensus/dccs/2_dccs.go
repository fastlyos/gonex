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
	"context"
	"errors"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errInvalidNonce = errors.New("Invalid block nonce as distant from the last sealer application block")
)

// Init the second hardfork of DCCS consensus
func (d *Dccs) init2() *Dccs {
	d.init1()
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
	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	// checkpoint := d.config.IsCheckpoint(number)
	// signersBytes := len(header.Extra) - extraVanity - extraSeal
	// if !checkpoint && signersBytes != 0 {
	// 	return errExtraSigners
	// }
	// if checkpoint && signersBytes%common.AddressLength != 0 {
	// 	return errInvalidCheckpointSigners
	// }
	// // Ensure that the mix digest is zero as we don't have fork protection currently
	// if header.MixDigest != (common.Hash{}) {
	// 	return errInvalidMixDigest
	// }
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
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
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+d.config.Period > header.Time {
		return ErrInvalidTimestamp
	}

	// verify the distant from the last sealer application block
	var nonce types.BlockNonce
	if isSealerApplicationBlock(parent) {
		nonce = types.EncodeNonce(1)
	} else {
		nonce = types.EncodeNonce(parent.Nonce.Uint64() + 1)
	}
	if header.Nonce != nonce {
		log.Error("invalid nonce", "expected", nonce, "actual", header.Nonce)
		return errInvalidNonce
	}

	// TODO
	// // Retrieve the snapshot needed to verify this header and cache it
	// snap, err := d.snapshot1(chain, header, parents)
	// if err != nil {
	// 	return err
	// }
	// // If the block is a checkpoint block, verify the signer list
	// if d.config.IsCheckpoint(number) {
	// 	signers := make([]byte, len(snap.Signers)*common.AddressLength)
	// 	for i, signer := range snap.signers1() {
	// 		copy(signers[i*common.AddressLength:], signer.Address[:])
	// 	}
	// 	extraSuffix := len(header.Extra) - extraSeal
	// 	if !bytes.Equal(header.Extra[extraVanity:extraSuffix], signers) {
	// 		return errInvalidCheckpointSigners
	// 	}
	// }
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

	// Retrieve the authorized sealers list to verify this header (TODO: and cache it)
	sealers, err := d.getSealers(number, chain, parents, true)
	if err != nil {
		return err
	}
	// Verify mix digest as sealers digest
	sealersBytes := serializeSealers(sealers)
	sealersDigest := crypto.Keccak256Hash(sealersBytes)
	// Mix digest is record the hash digest of all authorized sealer for this block
	if header.MixDigest != sealersDigest {
		log.Error("invalid mix digest as sealers digest", "expected", sealersDigest, "actual", header.MixDigest)
		return errors.New("invalid mix digest as sealers digest")
	}

	// Retrieve the snapshot needed to verify this header and cache it
	// snap, err := d.snapshot1(chain, header, parents)
	// if err != nil {
	// 	return err
	// }

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, d.signatures)
	if err != nil {
		return err
	}

	if !isAuthorized(signer, sealers) {
		return errUnauthorizedSigner
	}

	headers, err := d.GetRecentHeaders(len(sealers)/2, chain, header, parents)
	if err != nil {
		return err
	}
	for _, h := range headers {
		sig, err := ecrecover(h, d.signatures)
		if err != nil {
			return err
		}
		if signer == sig {
			// Signer is among recents, only fail if the current block doesn't shift it out
			return errRecentlySigned
		}
	}

	// var parent *types.Header
	// if len(headers) > 0 {
	// 	parent = headers[0]
	// }

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	// signerDifficulty := snap.difficulty(signer, parent)
	// if header.Difficulty.Uint64() != signerDifficulty {
	// 	return errInvalidDifficulty
	// }
	return nil
}

type LogFilterBackend struct {
	chain consensus.ChainReader
	db    ethdb.Reader
}

func (b *LogFilterBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	return b.chain.GetHeaderByNumber(uint64(blockNr.Int64())), nil
}

func (b *LogFilterBackend) HeaderByHash(ctx context.Context, blockHash common.Hash) (*types.Header, error) {
	return b.chain.GetHeaderByHash(blockHash), nil
}

func (b *LogFilterBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	header := b.chain.GetHeaderByHash(blockHash)
	if header == nil {
		return nil, nil
	}
	receipts := rawdb.ReadReceipts(b.db, blockHash, header.Number.Uint64(), b.chain.Config())
	return receipts, nil
}

func (b *LogFilterBackend) GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error) {
	receipts, _ := b.GetReceipts(ctx, blockHash)
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

// This nil assignment ensures compile time that LogFilterBackend implements filters.SimpleBackend.
var _ filters.SimpleBackend = (*LogFilterBackend)(nil)

// Keccak256("Joined(address,address)")
var joinedTopic = common.HexToHash("7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea")

// Keccak256("Left(address,address)")
var leftTopic = common.HexToHash("4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5")

type sealerApplication struct {
	sealer common.Address
	action bool // isJoined
}

// A single byte right after extraVanity indicates the DataType, allow multiple
// structures and/or versions of RLP can be decoded from the extra bytes.
const (
	ExtendedDataTypeNone         byte = 0x00
	ExtendedDataTypeSealerJoin   byte = 0xF0
	ExtendedDataTypeSealerLeave  byte = 0xF1
	ExtendedDataTypeSealerLeak   byte = 0xF2 // reserved
	ExtendedDataTypeSealerBanned byte = 0xF3 // reserved
)

// consensusRange holds the consensus info for a range of block
// in a header chain
// type consensusRange struct {
// 	hash     common.Hash // hash of the last block hash
// 	sigCount map[common.Address]int
// }

func isAuthorized(signer common.Address, sealers []common.Address) bool {
	for _, sealer := range sealers {
		if signer == sealer {
			return true
		}
	}
	return false
}

func serializeSealers(sealers []common.Address) []byte {
	buf := make([]byte, len(sealers)*common.AddressLength)
	for i, sealer := range sealers {
		copy(buf[i:], sealer[:])
	}
	return buf
}

// TODO: UNOPTIMIZED
// getSealers gets the authorized sealers for a block number
// the list is deterministically sorted by sealer addresses
func (d *Dccs) getSealers(number uint64, chain consensus.ChainReader, parents []*types.Header, sorted bool) ([]common.Address, error) {
	found := map[common.Address]struct{}{}
	list := []common.Address{}
	addSealer := func(sealer common.Address) {
		if _, exists := found[sealer]; !exists {
			log.Error("+++", "sealer", sealer)
			found[sealer] = struct{}{}
			list = append(list, sealer)
		}
	}
	remSealer := func(sealer common.Address) {
		if _, exists := found[sealer]; exists {
			log.Error("+++", "sealer", sealer)
			delete(found, sealer)
			for i, s := range list {
				if s == sealer {
					list[i] = list[len(list)-1]
					list = list[:len(list)-1]
					break
				}
			}
		}
	}
	// scan from number-LeakDuration to number-1
	for i := d.config.LeakDuration; i > 0; i-- {
		if number <= i {
			// skip any 'negative' number before genesis block
			continue
		}
		header := getAvailableHeader(number-i, nil, parents, chain)
		sealer, err := ecrecover(header, d.signatures)
		if err != nil {
			return nil, err
		}
		addSealer(sealer)
		log.Error("loop", "i", i, "n", number-i, "nonce", header.Nonce.Uint64())
		// confirmed application
		if header.Nonce.Uint64() == d.config.ApplicationConfirmation {
			if number <= i+d.config.ApplicationConfirmation {
				log.Error("invalid block nonce as application confirmation", "nonce", header.Nonce.Uint64())
				return nil, errors.New("invalid block nonce as application confirmation")
			}
			var p []*types.Header
			if int(i) < len(parents) {
				p = parents[:len(parents)-int(i)]
			}
			appHeader := getAvailableHeader(number-i-d.config.ApplicationConfirmation, header, p, chain)
			if appHeader == nil {
				log.Error("sealer application block not available", "number", appHeader.Number)
				return nil, errors.New("sealer application block not available")
			}
			if len(appHeader.Extra) <= extraVanity+extraSeal {
				log.Error("no sealer application data in header extra", "app number", appHeader.Number, "number", number, "i", i)
				return nil, errors.New("no sealer application data in header extra")
			}
			apps, _ := decodeSealerApplications(appHeader.Extra[extraVanity : len(appHeader.Extra)-extraSeal])
			for _, app := range apps {
				if app.action {
					log.Error("++++++++++++++++++++++++++")
					addSealer(app.sealer)
				} else {
					log.Error("--------------------------")
					remSealer(app.sealer)
				}
			}
		}
	}
	if sorted {
		sort.Sort(signersAscending(list))
	}
	return list, nil
}

func decodeSealerApplications(buf []byte) ([]sealerApplication, int) {
	if len(buf) == 0 {
		return nil, 0
	}
	log.Error("decodeSealerApplications", "buf", common.Bytes2Hex(buf))
	apps := make([]sealerApplication, len(buf)/(common.AddressLength+1))
	for i := 0; i < len(buf); i += common.AddressLength + 1 {
		if buf[i]&0xF0 == 0 {
			// not sealer application
			return apps, i
		}
		var action bool
		if buf[i] == ExtendedDataTypeSealerJoin {
			action = true
		}
		apps = append(apps, sealerApplication{
			sealer: common.BytesToAddress(buf[i+1 : i+1+common.AddressLength]),
			action: action,
		})
	}
	return apps, len(buf)
}

func encodeSealerApplications(applications []sealerApplication) []byte {
	if len(applications) == 0 {
		return nil
	}
	log.Error("encodeSealerApplications", "applications", applications)
	buf := make([]byte, len(applications)*(common.AddressLength+1))
	for i, app := range applications {
		if app.action {
			buf[i*(common.AddressLength+1)] = ExtendedDataTypeSealerJoin
		} else {
			buf[i*(common.AddressLength+1)] = ExtendedDataTypeSealerLeave
		}
		copy(buf[i*(common.AddressLength+1)+1:], app.sealer[:])
	}
	return buf
}

// fetchSealerApplications filters the block for any joining or leaving sealer.
// Newly joined sealer is mapped to it's coinbase address, left sealer is mapped to nil.
// Sealer joining and leaving txs can be confirmed in the same block, and the order of the requests
// is preserved.
func (d *Dccs) fetchSealerApplications(header *types.Header, chain consensus.ChainReader) ([]sealerApplication, error) {
	logs, err := filters.BlockLogs(header,
		[]common.Address{d.config.Contract},
		[][]common.Hash{{joinedTopic, leftTopic}},
		&LogFilterBackend{
			chain: chain,
			db:    d.db,
		})

	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, nil
	}

	applications := make([]sealerApplication, len(logs))

	for i, l := range logs {
		// len(log.Data) must be 32 * 2 here
		sealer := common.BytesToAddress(l.Data[32:])
		var joined bool
		if l.Topics[0] == joinedTopic {
			joined = true
			staker := common.BytesToAddress(l.Data[:32])
			log.Error("Sealer joined", "sealer", sealer, "coinbase", staker)
		} else {
			log.Error("Sealer left", "sealer", sealer)
		}
		applications[i] = sealerApplication{
			sealer: sealer,
			action: joined,
		}
	}
	return applications, nil
}

func isSealerApplicationBlock(header *types.Header) bool {
	if len(header.Extra) <= extraVanity+extraSeal {
		return false
	}
	return header.Extra[extraVanity]&0xF0 != 0
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
	if isSealerApplicationBlock(parent) {
		header.Nonce = types.EncodeNonce(1)
	} else {
		header.Nonce = types.EncodeNonce(parent.Nonce.Uint64() + 1)
	}

	d.prepareBeneficiary(header, chain)

	// Set the correct difficulty
	parents := []*types.Header{parent}
	snap, err := d.snapshot1(chain, header, parents)
	if err != nil {
		return err
	}
	header.Difficulty = CalcDifficulty1(snap, d.signer, parent)
	log.Trace("header.Difficulty", "difficulty", header.Difficulty)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	// if d.config.IsCheckpoint(number) {
	// 	for _, signer := range snap.signers1() {
	// 		header.Extra = append(header.Extra, signer.Address[:]...)
	// 	}
	// }

	applications, err := d.fetchSealerApplications(parent, chain)
	if err != nil {
		log.Error("failed to get changed sealer from log", "parent", parent.Number, "err", err)
		return err
	}
	log.Error("sealers", "applications", applications)
	extExtra := encodeSealerApplications(applications)
	if len(extExtra) > 0 {
		header.Extra = append(header.Extra, extExtra...)
	}
	log.Error("prepare2", "extExtra", common.Bytes2Hex(extExtra))

	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	sealers, err := d.getSealers(number, chain, parents, true)
	if err != nil {
		return err
	}
	sealersBytes := serializeSealers(sealers)
	// Mix digest is record the hash digest of all authorized sealer for this block
	header.MixDigest = crypto.Keccak256Hash(sealersBytes)

	// Ensure the timestamp has the correct delay
	header.Time = parent.Time + d.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}
	return nil
}

// finalize2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalize2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// finalizeAndAssemble2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalizeAndAssemble2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

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

	// Bail out if we're unauthorized to sign a block
	snap, err := d.snapshot1(chain, header, nil)
	if err != nil {
		return err
	}
	if _, authorized := snap.Signers[signer]; !authorized {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	headers, err := d.GetRecentHeaders(len(snap.Signers)/2, chain, header, nil)
	if err != nil {
		return err
	}
	for _, h := range headers {
		sig, err := ecrecover(h, d.signatures)
		if err != nil {
			return err
		}
		if signer == sig {
			// Signer is among recents
			log.Info("Signed recently, must wait for others")
			return nil
		}
	}
	var parent *types.Header
	if len(headers) > 0 {
		parent = headers[0]
	}
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if !snap.inturn1(signer, parent) {
		// It's not our turn explicitly to sign, delay it a bit
		offset, err := snap.offset(signer, parent)
		if err != nil {
			return err
		}
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
