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
	errInvalidExtExtraKind        = errors.New("invalid extended extra kind")
)

// A single byte right after extraVanity indicates the DataType, allow multiple
// structures and/or versions of RLP can be decoded from the extra bytes.
const (
	ExtendedDataTypeNone         byte = 0x00
	ExtendedDataTypeSealerJoin   byte = 0xF0
	ExtendedDataTypeSealerLeave  byte = 0xF1
	ExtendedDataTypeSealerDigest byte = 0xFE
	ExtendedDataTypeCrossLink    byte = 0xFF
)

type extData struct {
	anchor        *common.Hash
	sealersDigest *common.Hash
	applications  []sealerApplication
}

type sealerApplication struct {
	sealer common.Address
	action bool // isJoined
}

func (e *extData) Bytes() []byte {
	size := len(e.applications) * (1 + common.AddressLength) // sealer applications
	if e.anchor != nil {
		size += 1 + common.HashLength
	}
	if e.sealersDigest != nil {
		size += 1 + common.HashLength
	}
	extra := make([]byte, size)
	i := 0
	if e.anchor != nil {
		extra[i] = ExtendedDataTypeCrossLink
		i++
		copy(extra[i:i+common.HashLength], e.anchor.Bytes())
		i += common.HashLength
	}
	if e.sealersDigest != nil {
		extra[i] = ExtendedDataTypeSealerDigest
		i++
		copy(extra[i:i+common.HashLength], e.sealersDigest.Bytes())
		i += common.HashLength
	}
	for _, app := range e.applications {
		if app.action {
			extra[i] = ExtendedDataTypeSealerJoin
		} else {
			extra[i] = ExtendedDataTypeSealerLeave
		}
		i++
		copy(extra[i:i+common.AddressLength], app.sealer.Bytes())
		i += common.AddressLength
	}
	return extra
}

func bytesToExtData(extra []byte) (*extData, error) {
	var extData extData
	size := len(extra)
	for i := 0; i < size; {
		kind := extra[i]
		i++
		switch kind {
		case ExtendedDataTypeSealerJoin, ExtendedDataTypeSealerLeave:
			if i+common.AddressLength > size {
				log.Error("bytesToExtData", "extra", common.Bytes2Hex(extra), "i", i)
				return nil, errInvalidSealerAddressLength
			}
			extData.applications = append(extData.applications, sealerApplication{
				sealer: common.BytesToAddress(extra[i : i+common.AddressLength]),
				action: kind == ExtendedDataTypeSealerJoin,
			})
			i += common.AddressLength
		case ExtendedDataTypeSealerDigest:
			if i+common.HashLength > size {
				return nil, errInvalidSealerDigestLength
			}
			digest := common.BytesToHash(extra[i : i+common.HashLength])
			extData.sealersDigest = &digest
			i += common.HashLength
		case ExtendedDataTypeCrossLink:
			if i+common.HashLength > size {
				return nil, errInvalidCrosslinkLength
			}
			digest := common.BytesToHash(extra[i : i+common.HashLength])
			extData.anchor = &digest
			i += common.HashLength
		default:
			return nil, errInvalidExtExtraKind
		}
	}
	return &extData, nil
}

func (d *Dccs) getExtData(header *types.Header) (*extData, error) {
	if len(header.Extra) <= extraVanity+extraSeal {
		return nil, nil
	}
	hash := header.Hash()
	if e, ok := d.extDataCache.Get(hash); ok {
		// in-memory sealingQueue found
		ed := e.(*extData)
		return ed, nil
	}
	ext, err := bytesToExtData(header.Extra[extraVanity : len(header.Extra)-extraSeal])
	if err != nil || ext == nil {
		return nil, err
	}
	d.extDataCache.Add(hash, ext)
	return ext, nil
}
