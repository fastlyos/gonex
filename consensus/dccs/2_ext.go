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
	"errors"

	"github.com/ethereum/go-ethereum/common"
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
	ExtendedDataTypeSealerJoin  byte = 0xF0
	ExtendedDataTypeSealerLeave byte = 0xF1
	ExtendedDataTypeAnchor      byte = 0xFF
)

// ExtendedData is the data encoded in header.Extra[extraVanity:-extraSeal]
type ExtendedData struct {
	anchor *AnchorData // always comes first
}

type AnchorData struct {
	destHash      common.Hash         // hash of the anchor destination block
	sealersDigest common.Hash         // digest of the ordered active sealer for this block
	applications  []SealerApplication // sealer applications confirmed in the parent block
}

// SealerApplication packs the sealer address and its application action.
type SealerApplication struct {
	action byte
	sealer common.Address
}

func (a *SealerApplication) isJoined() bool {
	return a.action == ExtendedDataTypeSealerJoin
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

func bytesToAnchorData(extra []byte) (*AnchorData, int, error) {
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

func (e *ExtendedData) bytes() []byte {
	if e == nil {
		return []byte{}
	}
	var bytes []byte
	bytes = append(bytes, e.anchor.bytes()...)
	return bytes
}

func bytesToExtData(extra []byte) (*ExtendedData, error) {
	anchorData, _, err := bytesToAnchorData(extra)
	if err != nil {
		return nil, err
	}
	extData := ExtendedData{
		anchor: anchorData,
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
	ext, err := bytesToExtData(header.Extra[extraVanity : len(header.Extra)-extraSeal])
	if err != nil || ext == nil {
		return nil, err
	}
	d.extDataCache.Add(hash, ext)
	return ext, nil
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
