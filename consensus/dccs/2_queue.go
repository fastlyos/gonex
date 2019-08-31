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
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

type sealingQueue struct {
	hash       common.Hash                 // hash of the header
	sealer     common.Address              // sealer address of the header
	seed       []byte                      // random seed
	active     map[common.Address]struct{} // active sealers for the next block
	recent     map[common.Address]struct{} // recently signed sealers
	sorted     []common.Address            // sorted queue ([active]-[recent]+sealer)
	digest     common.Hash                 // hash(sort(active))
	sortedOnce sync.Once
	digestOnce sync.Once
}

func (q *sealingQueue) sealersDigest() common.Hash {
	q.digestOnce.Do(func() {
		active := make([]common.Address, 0, len(q.active))
		for adr := range q.active {
			active = append(active, adr)
		}
		sort.Sort(signersAscending(active))
		q.digest = types.RLPHash(active)
	})
	return q.digest
}

// sealerShuffling implements the sort interface to allow sorting a list of addresses
type seedShuffle struct {
	seed  []byte
	queue []common.Address
}

func (s seedShuffle) Len() int      { return len(s.queue) }
func (s seedShuffle) Swap(i, j int) { s.queue[i], s.queue[j] = s.queue[j], s.queue[i] }
func (s seedShuffle) Less(i, j int) bool {
	// TODO cache these hashes
	sha := func(address common.Address, seed []byte) []byte {
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(address[:])
		hasher.Write(seed)
		return hasher.Sum(nil)
	}
	si := sha(s.queue[i], s.seed)
	sj := sha(s.queue[j], s.seed)
	return bytes.Compare(si, sj) < 0
}

// include the last sealer in for positioning even when the last sealer is just left/leaked
func (q *sealingQueue) sortedQueue() []common.Address {
	q.sortedOnce.Do(func() {
		size := len(q.active) - len(q.recent) + 1
		s := seedShuffle{
			seed:  q.seed,
			queue: make([]common.Address, 1, size),
		}
		s.queue[0] = q.sealer
		for adr := range q.active {
			_, recentlySigned := q.recent[adr]
			if !recentlySigned {
				s.queue = append(s.queue, adr)
				// TODO: pre-calculate the shuffling hash here
			}
		}
		sort.Sort(s)
		q.sorted = s.queue
	})
	return q.sorted
}

func (q *sealingQueue) isRecentlySigned(address common.Address) bool {
	_, recentlySigned := q.recent[address]
	return recentlySigned
}

func (q *sealingQueue) isActive(address common.Address) bool {
	_, active := q.active[address]
	return active
}

func (q *sealingQueue) offset(signer common.Address,
	getHeaderByHash func(common.Hash) *types.Header,
	sigCache *lru.ARCCache) (int, error) {

	activeLen := len(q.active)

	if activeLen-len(q.recent) <= 1 {
		// no competition
		return 0, nil
	}

	if signer == q.sealer {
		return activeLen, errRecentlySigned
	}

	queue := q.sortedQueue()

	signerPosition := func(signer common.Address) (int, bool) {
		for i, sig := range queue {
			if sig == signer {
				return i, true
			}
		}
		return -1, false
	}

	pos, ok := signerPosition(signer)
	if !ok {
		// unable to find the signer position
		return activeLen, errUnauthorizedSigner
	}

	prevPos, ok := signerPosition(q.sealer)
	if !ok {
		// unable to find the previous signer position: should never happen
		return activeLen, errUnknownPreviousSealer
	}

	offset := pos - prevPos - 1
	if offset < 0 {
		offset += len(queue)
	}

	log.Trace("sealingQueue.offset",
		"offset", offset,
		"signer position", pos,
		"previous signer position", prevPos,
		"len(queue)", len(queue),
		"len(active)", len(q.active),
		"len(recent)", len(q.recent))

	return offset, nil
}

func (q *sealingQueue) difficulty(address common.Address,
	getHeaderByHash func(common.Hash) *types.Header,
	sigCache *lru.ARCCache) uint64 {

	offset, err := q.offset(address, getHeaderByHash, sigCache)
	if err != nil {
		return 0
	}

	n := len(q.active)

	return uint64(n - offset)
}

// recents len is MIN(lastActiveLen,activeLen)/2
func (d *Dccs) getSealingQueue(parentHash common.Hash, parents []*types.Header, chain consensus.ChainReader) (*sealingQueue, error) {
	if q, ok := d.sealingQueueCache.Get(parentHash); ok {
		// in-memory sealingQueue found
		queue := q.(*sealingQueue)
		return queue, nil
	}
	log.Error("getSealingQueue", "parentHash", parentHash)
	parent := getAvailableHeaderByHash(parentHash, nil, parents, chain)
	if parent == nil {
		return nil, errUnknownPreviousSealer
	}
	sealer, err := ecrecover(parent, d.signatures)
	if err != nil {
		return nil, err
	}
	queue := sealingQueue{
		hash:   parentHash,
		sealer: sealer,
		seed:   parent.Nonce[:],
		active: map[common.Address]struct{}{},
		recent: map[common.Address]struct{}{},
	}

	number := parent.Number.Uint64()

	// temporary queue for recents
	var recents []common.Address

	addRecent := func(sealer common.Address) {
		if _, exists := queue.recent[sealer]; !exists {
			log.Error("***", "sealer", sealer)
			queue.recent[sealer] = struct{}{}
			recents = append(recents, sealer)
		}
	}
	addActive := func(sealer common.Address) {
		if _, exists := queue.active[sealer]; !exists {
			log.Error("+++", "sealer", sealer)
			queue.active[sealer] = struct{}{}
		}
	}
	remActive := func(sealer common.Address) {
		if _, exists := queue.active[sealer]; exists {
			log.Error("---", "sealer", sealer)
			delete(queue.active, sealer)
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
		if i > minBlockToScan && len(queue.active) >= int(maxDiff) {
			break // all active sealers has probably been collected
		}
		// END OF OPTIMIZATION

		// TODO: optimization for leakage case

		n := number - i
		header := getAvailableHeader(n, nil, parents, chain)
		if header == nil {
			log.Error("getSealingQueue: getAvailableHeader returns nil", "n", n, "len(parents)", len(parents))
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

	apps, err := d.crawlSealerApplications(parent, parents, chain)
	if err != nil {
		return nil, err
	}
	for _, app := range apps {
		if app.action {
			log.Error("++++++++++ joined")
			addActive(app.sealer)
		} else {
			log.Error("---------- left")
			remActive(app.sealer)
		}
	}

	// truncate the extra recents
	if len(queue.active)/2 < len(recents) {
		for i := len(queue.active) / 2; i < len(recents); i++ {
			delete(queue.recent, recents[i])
		}
	}

	// Store found snapshot into mem-cache
	d.sealingQueueCache.Add(queue.hash, &queue)
	return &queue, nil
}

// crawl back the sealer applications skip-list
func (d *Dccs) crawlSealerApplications(header *types.Header, parents []*types.Header, chain consensus.ChainReader) ([]sealerApplication, error) {
	number := header.Number.Uint64()
	apps := []sealerApplication{}
	for header := getAvailableHeaderByHash(header.MixDigest, nil, parents, chain); header != nil; header = getAvailableHeaderByHash(header.MixDigest, nil, parents, chain) {
		log.Error("crawling", "appNumber", header.Number, "appNumber.Hash", header.Hash(), "cross-link", header.MixDigest)
		if (header.MixDigest == common.Hash{}) {
			// reach the CoLoa hardfork (new genesis)
			break
		}
		appConfirmedNumber := header.Number.Uint64() + d.config.ApplicationConfirmation
		// condition: appConfirmedNumber in (number-LeakDuration;number]
		if appConfirmedNumber+d.config.LeakDuration <= number {
			// any applications from this would be too old
			break
		}
		if appConfirmedNumber > number {
			// not enough confirmation
			continue
		}
		if len(header.Extra) <= extraVanity+extraSeal {
			log.Error("no sealer application data in header extra", "app number", header.Number, "number", number)
			return nil, errors.New("no sealer application data in header extra")
		}
		extData, err := bytesToExtData(header.Extra[extraVanity : len(header.Extra)-extraSeal])
		if err != nil {
			return nil, err
		}
		if len(extData.applications) > 0 {
			apps = append(extData.applications, apps...)
		}
	}
	return apps, nil
}
