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
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errInvalidNonce          = errors.New("Invalid block nonce as distant from the last sealer application block")
	errUnknownPreviousSealer = errors.New("Unknown previous block sealer")
)

// Init the second hardfork of DCCS consensus
func (d *Dccs) init2() *Dccs {
	d.init1()
	return d
}

type authorization struct {
	parent     common.Hash
	active     map[common.Address]struct{}
	recent     map[common.Address]struct{}
	sortedOnce sync.Once
	sorted     []common.Address
	digestOnce sync.Once
	digest     common.Hash
}

func (a *authorization) Digest() common.Hash {
	a.digestOnce.Do(func() {
		active := make([]common.Address, 0, len(a.active))
		for adr := range a.active {
			active = append(active, adr)
		}
		sort.Sort(signersAscending(active))
		a.digest = types.RLPHash(active)
	})
	return a.digest
}

func (a *authorization) sortedQueue() []common.Address {
	a.sortedOnce.Do(func() {
		queue := make([]common.Address, 0, len(a.active)-len(a.recent))
		for adr := range a.active {
			_, recentlySigned := a.recent[adr]
			if !recentlySigned {
				queue = append(queue, adr)
			}
		}
		sort.Sort(signersAscending(queue))
		a.sorted = queue
	})
	return a.sorted
}

func (a *authorization) isRecentlySigned(address common.Address) bool {
	_, recentlySigned := a.recent[address]
	return recentlySigned
}

func (a *authorization) isActive(address common.Address) bool {
	_, active := a.active[address]
	return active
}

func (a *authorization) offset(address common.Address,
	getHeaderByHash func(common.Hash) *types.Header,
	sigCache *lru.ARCCache) (int, error) {
	queue := a.sortedQueue()
	return offset2(address, a.parent, queue, getHeaderByHash, sigCache)
}

func (a *authorization) difficulty(address common.Address,
	getHeaderByHash func(common.Hash) *types.Header,
	sigCache *lru.ARCCache) uint64 {
	offset, err := a.offset(address, getHeaderByHash, sigCache)
	if err != nil {
		return 0
	}

	n := len(a.active)

	return uint64(n - offset)
}

func (d *Dccs) getAuthorization(parentHash common.Hash, parents []*types.Header, chain consensus.ChainReader) (*authorization, error) {
	log.Error("getAuthorization", "parentHash", parentHash)
	parent := getAvailableHeaderByHash(parentHash, nil, parents, chain)
	if parent == nil {
		return nil, errUnknownPreviousSealer
	}
	a := authorization{
		parent: parentHash,
		active: map[common.Address]struct{}{},
		recent: map[common.Address]struct{}{},
	}

	number := parent.Number.Uint64()

	// temporary queue for recents
	var recents []common.Address

	addRecent := func(sealer common.Address) {
		if _, exists := a.recent[sealer]; !exists {
			log.Error("***", "sealer", sealer)
			a.recent[sealer] = struct{}{}
			recents = append(recents, sealer)
		}
	}
	addActive := func(sealer common.Address) {
		if _, exists := a.active[sealer]; !exists {
			log.Error("+++", "sealer", sealer)
			a.active[sealer] = struct{}{}
		}
	}
	remActive := func(sealer common.Address) {
		if _, exists := a.active[sealer]; exists {
			log.Error("---", "sealer", sealer)
			delete(a.active, sealer)
		}
	}

	var maxDiff uint64

	// scan backward atmost LeakDurations blocks from number
	for i := uint64(0); i < d.config.LeakDuration; i++ {
		if number <= i {
			break // stop at the genesis block
		}

		// OPTIMIZATION: for non-leakage case
		// correct if there's at least one in-turn block in the last 16 blocks
		const minBlockToScan = 16
		if i > minBlockToScan && len(a.active) >= int(maxDiff) {
			break // all active sealers has probably been collected
		}
		// END OF OPTIMIZATION

		// TODO: optimization for leakage case

		n := number - i
		header := getAvailableHeader(n, nil, parents, chain)
		if header == nil {
			log.Error("getAuthorization: getAvailableHeader returns nil", "n", n, "len(parents)", len(parents))
			return nil, errUnknownBlock
		}
		sealer, err := ecrecover(header, d.signatures)
		if err != nil {
			return nil, err
		}

		addActive(sealer)

		// use the difficulty for total number of recently active sealer count
		if header.Difficulty.Uint64() > maxDiff {
			maxDiff = header.Difficulty.Uint64()
		}
		// somewhat probabilistically optimized, fairly safe nonetheless
		if i < minBlockToScan || len(recents) < int(maxDiff)/2 {
			addRecent(sealer)
		}
	}

	// crawl back the sealer applications skip-list
	allApps := []sealerApplication{}
	for appHeader := getAvailableHeaderByHash(parent.MixDigest, nil, parents, chain); appHeader != nil; appHeader = getAvailableHeaderByHash(appHeader.MixDigest, nil, parents, chain) {
		log.Error("crawling", "appNumber", appHeader.Number, "appNumber.Hash", appHeader.Hash(), "cross-link", appHeader.MixDigest)
		if (appHeader.MixDigest == common.Hash{}) {
			// reach the CoLoa hardfork (new genesis)
			break
		}
		appConfirmedNumber := appHeader.Number.Uint64() + d.config.ApplicationConfirmation
		// condition: appConfirmedNumber in (number-LeakDuration;number]
		if appConfirmedNumber+d.config.LeakDuration <= number {
			// any applications from this would be too old
			break
		}
		if appConfirmedNumber > number {
			// not enough confirmation
			continue
		}
		if len(appHeader.Extra) <= extraVanity+extraSeal {
			log.Error("no sealer application data in header extra", "app number", appHeader.Number, "number", number)
			return nil, errors.New("no sealer application data in header extra")
		}
		extExtra, err := bytesToExtExtra(appHeader.Extra[extraVanity : len(appHeader.Extra)-extraSeal])
		if err != nil {
			return nil, err
		}
		if len(extExtra.applications) > 0 {
			allApps = append(extExtra.applications, allApps...)
		}
	}
	for _, app := range allApps {
		if app.action {
			log.Error("++++++++++ joined")
			addActive(app.sealer)
		} else {
			log.Error("---------- left")
			remActive(app.sealer)
		}
	}

	// truncate the extra recents
	if len(a.active)/2 < len(recents) {
		for i := len(a.active) / 2; i < len(recents); i++ {
			delete(a.recent, recents[i])
		}
	}

	return &a, nil
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
	return types.BlockNonce{}
	// if d.config.CoLoaBlock.Cmp(number) == 0 {
	// 	return types.BlockNonce{}
	// } else if isSealerApplicationBlock(parent) {
	// 	return types.EncodeNonce(1)
	// } else {
	// 	return types.EncodeNonce(parent.Nonce.Uint64() + 1)
	// }
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
	nonce := d.getBlockNonce(header.Number, parent)
	if header.Nonce != nonce {
		log.Error("invalid nonce", "expected", nonce, "actual", header.Nonce)
		return errInvalidNonce
	}

	// verify the cross-link reference to the last sealer application block
	var expectedMixDigest common.Hash
	if d.config.CoLoaBlock.Cmp(header.Number) == 0 {
		expectedMixDigest = common.Hash{}
	} else if d.isSealerApplicationBlock(parent) {
		expectedMixDigest = parent.Hash()
	} else {
		expectedMixDigest = parent.MixDigest
	}
	if header.MixDigest != expectedMixDigest {
		return fmt.Errorf("invalid cross-link digest: number=%v, want=%v, have=%v", header.Number, expectedMixDigest, header.MixDigest)
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

	// Retrieve the authorized sealers list to verify this header (TODO: and cache it)
	author, err := d.getAuthorization(header.ParentHash, parents, chain)
	if err != nil {
		return err
	}
	// // Verify mix digest as sealers digest
	// sealersDigest := types.RLPHash(sealers)
	// // Mix digest is record the hash digest of all authorized sealer for this block
	// if header.MixDigest != sealersDigest {
	// 	log.Error("invalid mix digest as sealers digest", "expected", sealersDigest, "actual", header.MixDigest)
	// 	return errors.New("invalid mix digest as sealers digest")
	// }

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, d.signatures)
	if err != nil {
		return err
	}

	if !author.isActive(signer) {
		return errUnauthorizedSigner
	}

	if author.isRecentlySigned(signer) {
		return errRecentlySigned
	}

	// TODO: verify extExtra here

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

func (d *Dccs) isSealerApplicationBlock(header *types.Header) bool {
	if d.config.CoLoaBlock.Cmp(header.Number) == 0 {
		return true
	}
	if len(header.Extra) <= extraVanity+extraSeal {
		return false
	}
	return header.Extra[extraVanity]&0xF0 == 0xF0
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

	d.prepareBeneficiary(header, chain)

	// set the cross-link reference to the last sealer application block
	if d.config.CoLoaBlock.Cmp(header.Number) == 0 {
		header.MixDigest = common.Hash{}
	} else if d.isSealerApplicationBlock(parent) {
		header.MixDigest = parent.Hash()
	} else {
		header.MixDigest = parent.MixDigest
	}

	author, err := d.getAuthorization(header.ParentHash, nil, chain)
	if err != nil {
		return err
	}
	// // Mix digest is record the hash digest of all authorized sealer for this block
	// header.MixDigest = types.RLPHash(sealers)

	// Set the correct difficulty
	difficulty := author.difficulty(d.signer, chain.GetHeaderByHash, d.signatures)
	header.Difficulty = new(big.Int).SetUint64(difficulty)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	applications, err := d.fetchSealerApplications(parent, chain)
	if err != nil {
		log.Error("failed to get changed sealer from log", "parent", parent.Number, "err", err)
		return err
	}
	if len(applications) > 0 {
		log.Error("sealers", "applications", applications)
		digest := author.Digest()
		extExtra := extExtra{
			sealersDigest: &digest,
			applications:  applications,
		}
		header.Extra = append(header.Extra, extExtra.Bytes()...)
		log.Error("prepare2", "extExtra", extExtra)
	}

	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

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

	// // Bail out if we're unauthorized to sign a block
	// snap, err := d.snapshot1(chain, header, nil)
	// if err != nil {
	// 	return err
	// }

	author, err := d.getAuthorization(header.ParentHash, nil, chain)
	if err != nil {
		return err
	}

	if !author.isActive(signer) {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	if author.isRecentlySigned(signer) {
		// TBD: return err?
		return nil
	}
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	// Find the signer offset
	offset, err := author.offset(signer, chain.GetHeaderByHash, d.signatures)
	if err != nil {
		return err
	}
	if !d.inturn2(offset) {
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

// inturn2 returns if a signer at a given block height is in-turn or not.
func (d *Dccs) inturn2(offset int) bool {
	return offset == 0
}

func signerPosition2(signer common.Address, signers []common.Address) (int, bool) {
	for i, sig := range signers {
		if sig == signer {
			return i, true
		}
	}
	return -1, false
}

func offset2(signer common.Address, parentHash common.Hash, signers []common.Address, getHeaderByHash func(common.Hash) *types.Header, sigCache *lru.ARCCache) (int, error) {
	n := len(signers)
	if n <= 1 {
		// no competition
		return 0, nil
	}

	pos, ok := signerPosition2(signer, signers)
	if !ok {
		// unable to find the signer position
		log.Error("eee", "signer", signer, "sealers", signers)
		return n, errUnauthorizedSigner
	}

	prevPos, err := func() (int, error) {
		for {
			parent := getHeaderByHash(parentHash)
			if parent == nil {
				return 0, errUnknownPreviousSealer
			}
			// Resolve the last authorization key and check against signer
			prevSigner, err := ecrecover(parent, sigCache)
			if err != nil {
				return 0, err
			}
			prevPos, ok := signerPosition2(prevSigner, signers)
			if ok {
				// previous signer position found
				return prevPos, nil
			}
			// previous signer could just left the syndicate
			parentHash = parent.ParentHash
		}
	}()
	if err != nil {
		return 0, nil
	}

	offset := pos - prevPos - 1
	if offset < 0 {
		offset += n
	}

	log.Debug("offset", "signer position", pos, "previous signer position", prevPos, "len(signers)", n, "offset", offset)

	return offset, nil
}

// difficulty2 returns the block weight at a given block height for a signer.
// Turn-ness is the directional distant from a signer to the previous one,
// following a circular order of the signers list.
// @return maximum value = len(signers) if signer is right after the prevSigner (circularly)
// @return minimum value = 1 if the signer is right before the prevSigner (circularly)
// @return invalid value = 0 if the signer or parent signer is not on the sealer list
func (d *Dccs) difficulty2(signer common.Address, parentHash common.Hash, signers []common.Address, getHeaderByHash func(common.Hash) *types.Header) uint64 {
	offset, err := offset2(signer, parentHash, signers, getHeaderByHash, d.signatures)
	if err != nil {
		return 0
	}

	n := len(signers)

	return uint64(n - offset)
}
