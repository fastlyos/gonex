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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errInvalidCrosslinkLength     = errors.New("invalid crosslink digest length")
	errInvalidSealerDigestLength  = errors.New("invalid sealer digest length")
	errInvalidSealerAddressLength = errors.New("invalid sealer address length")
	errInvalidExtraLinkLength     = errors.New("invalid extra link data length")
)

// A single byte right after extraVanity indicates the DataType, allow multiple
// structures and/or versions of RLP can be decoded from the extra bytes.
const (
	ExtendedDataTypeNone        byte = 0x00
	ExtendedDataTypeVDF         byte = 0x01
	ExtendedDataTypeSealerJoin  byte = 0xF0
	ExtendedDataTypeSealerLeave byte = 0xF1
	ExtendedDataTypeAnchor      byte = 0xFF
)

// ExtendedData is the data encoded in header.Extra[extraVanity:-extraSeal]
type ExtendedData struct {
	anchor *AnchorData // always comes first
	random RandomData
}

func (e *ExtendedData) bytes() []byte {
	if e == nil {
		return []byte{}
	}
	var bytes []byte
	bytes = append(bytes, e.anchor.bytes()...)
	return bytes
}

func getExtDataFromExtra(extBytes []byte) (*ExtendedData, error) {
	anchorData, n, err := anchorDataFromExtraBytes(extBytes)
	if err != nil {
		return nil, err
	}
	extBytes = extBytes[n:]
	vdfOutput, n := getRandomDataFromExtra(extBytes)
	extBytes = extBytes[n:]
	extData := ExtendedData{
		anchor: anchorData,
		random: vdfOutput,
	}
	return &extData, nil
}

func (d *Dccs) getExtData(header *types.Header) (*ExtendedData, error) {
	if len(header.Extra) <= extraVanity+extraSeal {
		return nil, nil
	}
	hash := header.Hash()
	if e, ok := d.extDataCache.Get(hash); ok {
		// in-memory SealingQueue found
		ed := e.(*ExtendedData)
		return ed, nil
	}
	ext, err := getExtDataFromExtra(header.Extra[extraVanity : len(header.Extra)-extraSeal])
	if err != nil || ext == nil {
		return nil, err
	}
	d.extDataCache.Add(hash, ext)
	return ext, nil
}

type RandomData []byte

func getRandomDataFromExtra(extra []byte) (RandomData, int) {
	size := len(extra)
	if size == 0 {
		return nil, 0
	}
	if size < 1+randomSeedSize {
		return nil, 0
	}
	if extra[0] != ExtendedDataTypeVDF {
		return nil, 0
	}
	output := extra[1 : 1+randomSeedSize]
	return append(output[:0:0], output...), 1 + randomSeedSize
}

func (r RandomData) bytes() []byte {
	if len(r) > 0 {
		return append([]byte{ExtendedDataTypeVDF}, r...)
	}
	return nil
}

func (d *Dccs) getRandomData(header *types.Header) (RandomData, error) {
	extData, err := d.getExtData(header)
	if err != nil {
		return nil, err
	}
	if extData == nil {
		return nil, nil
	}
	return extData.random, nil
}

// SealerApplication packs the sealer address and its application action.
type SealerApplication struct {
	action byte
	sealer common.Address
}

func (a *SealerApplication) isJoined() bool {
	return a.action == ExtendedDataTypeSealerJoin
}

type AnchorData struct {
	destHash      common.Hash         // hash of the anchor destination block
	sealersDigest common.Hash         // digest of the ordered active sealer for this block
	applications  []SealerApplication // sealer applications confirmed in the parent block
}

func (e *AnchorData) bytes() []byte {
	if e == nil {
		return []byte{}
	}
	size := 1 + 2*common.HashLength
	size += (1 + common.AddressLength) * len(e.applications) // sealer applications
	extra := make([]byte, 0, size)
	extra = append(extra, ExtendedDataTypeAnchor)
	extra = append(extra, e.destHash.Bytes()...)
	extra = append(extra, e.sealersDigest.Bytes()...)
	for _, app := range e.applications {
		extra = append(extra, app.action)
		extra = append(extra, app.sealer.Bytes()...)
	}
	return extra
}

func anchorDataFromExtraBytes(extra []byte) (*AnchorData, int, error) {
	size := len(extra)
	if size == 0 {
		return nil, 0, nil
	}
	if extra[0] != ExtendedDataTypeAnchor {
		return nil, 0, nil
	}
	if size < 1+2*common.HashLength {
		return nil, 0, errInvalidExtraLinkLength
	}
	anchorData := AnchorData{
		destHash:      common.BytesToHash(extra[1 : 1+common.HashLength]),
		sealersDigest: common.BytesToHash(extra[1+common.HashLength : 1+2*common.HashLength]),
	}
	for i := 1 + 2*common.HashLength; i < size; {
		kind := extra[i]
		i++
		switch kind {
		case ExtendedDataTypeSealerJoin, ExtendedDataTypeSealerLeave:
			if i+common.AddressLength > size {
				log.Error("bytesToLinkData", "extra", common.Bytes2Hex(extra), "i", i)
				return nil, i, errInvalidSealerAddressLength
			}
			anchorData.applications = append(anchorData.applications, SealerApplication{
				sealer: common.BytesToAddress(extra[i : i+common.AddressLength]),
				action: kind,
			})
			i += common.AddressLength
		default:
			return &anchorData, i - 1, nil
		}
	}
	return &anchorData, size, nil
}

func verifyAnchorBytes(extBytes []byte, expectedAnchorBytes []byte) error {
	n := len(expectedAnchorBytes)
	if len(extBytes) < n {
		return errInvalidExtendedDataLength
	}
	if n > 0 {
		// anchor data always comes first in the extended extra
		if bytes.Compare(extBytes[:n], expectedAnchorBytes) != 0 {
			return errInvalidExtendedData
		}
	}
	return nil
}

func (d *Dccs) getAnchorData(header *types.Header) (*AnchorData, error) {
	extData, err := d.getExtData(header)
	if err != nil {
		return nil, err
	}
	if extData == nil {
		return nil, nil
	}
	return extData.anchor, nil
}

func hasAnchorData(header *types.Header) bool {
	// if d.config.CoLoaBlock.Cmp(header.Number) == 0 {
	// 	return true
	// }
	if len(header.Extra) <= extraVanity+extraSeal {
		return false
	}
	// link data always at the start of the extended data
	return header.Extra[extraVanity] == ExtendedDataTypeAnchor
}

// AnchorData is recorded when the block is a cross-link, which happens when either:
// + the parent block has sealer application tx(s), or
// + the new SealingQueue super-majority continuity is broken
func (d *Dccs) assembleAnchorExtra(parent *types.Header, parents []*types.Header, chain consensus.ChainReader) ([]byte, error) {
	parentHash := parent.Hash()
	if a, ok := d.anchorExtraCache.Get(parentHash); ok {
		// in-memory SealingQueue found
		ab := a.([]byte)
		return ab, nil
	}
	queue, err := d.getSealingQueue(parentHash, parents, chain)
	if err != nil {
		return nil, err
	}

	if parent.Number.Uint64()+1 == d.config.CoLoaBlock.Uint64() {
		// special handling for hardfork block
		anchorData := AnchorData{
			destHash:      common.Hash{},
			sealersDigest: queue.sealersDigest(),
		}
		ext := ExtendedData{
			anchor: &anchorData,
		}
		return ext.bytes(), nil
	}

	linkHeader := getLinkDest(parent, parents, chain)
	anchorHeader, err := d.getAnchorDest(linkHeader, parents, chain)
	if err != nil {
		return nil, err
	}

	anchorData := AnchorData{
		destHash: anchorHeader.Hash(),
	}
	newAnchor := false

	apps, err := d.fetchSealerApplications(parent, chain)
	if err != nil {
		log.Error("failed to get changed sealer from log", "parent", parent.Number, "err", err)
		return nil, err
	}
	// parent block has sealer application tx(s)
	if len(apps) > 0 {
		log.Info("sealers", "applications", apps)
		anchorData.applications = apps
		newAnchor = true
	}

	anchorQueue, err := d.getSealingQueue(anchorHeader.ParentHash, parents, chain)
	if err != nil {
		return nil, err
	}

	anchorRatio, broken := queue.commonRatio(anchorQueue)
	if broken {
		log.Info("Anchor continuity is broken", "anchor ratio", anchorRatio.RatString(), "anchor number", anchorHeader.Number)
		// anchor continuity is broken, compare the current link to anchor
		linkQueue, err := d.getSealingQueue(linkHeader.ParentHash, parents, chain)
		if err != nil {
			return nil, err
		}
		linkRatio, broken := queue.commonRatio(linkQueue)
		if !broken || linkRatio.Cmp(anchorRatio) > 0 {
			// link continuity is preserved, or atleast better than anchor continuity
			log.Info("Link is promoted to Anchor", "link ratio", linkRatio.RatString(), "link number", linkHeader.Number)
			anchorData.destHash = parent.MixDigest
			newAnchor = true
		} else {
			// link continuity is also broken and worse than anchor continuity
			log.Warn("Broken anchor has no better alternative", "link ratio", linkRatio.RatString())
		}
	}

	if !newAnchor {
		return nil, nil
	}

	// only calculate sealersDigest when nessesary
	anchorData.sealersDigest = queue.sealersDigest()

	anchorExtra := anchorData.bytes()
	d.anchorExtraCache.Add(parentHash, anchorExtra)
	return anchorExtra, nil
}
