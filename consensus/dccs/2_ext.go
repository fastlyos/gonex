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

// type extData struct {
// 	data []byte
// }

// func (e *extData) Kind() byte {
// 	return e.data[0]
// }

// func (e *extData) Bytes() []byte {
// 	return e.data[1:]
// }

// func (e *extData) Address() common.Address {
// 	return common.BytesToAddress(e.Bytes())
// }

// func (e *extData) Hash() common.Hash {
// 	return common.BytesToHash(e.Bytes())
// }

type extExtra struct {
	majorityLink  *common.Hash
	sealersDigest *common.Hash
	applications  []sealerApplication
}

type sealerApplication struct {
	sealer common.Address
	action bool // isJoined
}

// func decodeSealerApplications(buf []byte) ([]sealerApplication, int) {
// 	if len(buf) == 0 {
// 		return nil, 0
// 	}
// 	log.Error("decodeSealerApplications", "buf", common.Bytes2Hex(buf))
// 	count := len(buf) / (common.AddressLength + 1)
// 	apps := make([]sealerApplication, count)
// 	for i := 0; i < count; i++ {
// 		offset := i * (common.AddressLength + 1)
// 		if buf[offset]&0xF0 == 0 {
// 			// not sealer application
// 			return apps, i
// 		}
// 		var action bool
// 		if buf[offset] == ExtendedDataTypeSealerJoin {
// 			action = true
// 		}
// 		apps[i] = sealerApplication{
// 			sealer: common.BytesToAddress(buf[offset+1 : offset+1+common.AddressLength]),
// 			action: action,
// 		}
// 	}
// 	return apps, len(buf)
// }

// func encodeSealerApplications(applications []sealerApplication) []byte {
// 	if len(applications) == 0 {
// 		return nil
// 	}
// 	log.Error("encodeSealerApplications", "applications", applications)
// 	buf := make([]byte, len(applications)*(common.AddressLength+1))
// 	for i, app := range applications {
// 		if app.action {
// 			buf[i*(common.AddressLength+1)] = ExtendedDataTypeSealerJoin
// 		} else {
// 			buf[i*(common.AddressLength+1)] = ExtendedDataTypeSealerLeave
// 		}
// 		copy(buf[i*(common.AddressLength+1)+1:], app.sealer[:])
// 	}
// 	return buf
// }

func (e *extExtra) Bytes() []byte {
	size := len(e.applications) * (1 + common.AddressLength) // sealer applications
	if e.majorityLink != nil {
		size += 1 + common.HashLength
	}
	if e.sealersDigest != nil {
		size += 1 + common.HashLength
	}
	extra := make([]byte, size)
	i := 0
	if e.majorityLink != nil {
		extra[i] = ExtendedDataTypeCrossLink
		i++
		copy(extra[i:i+common.HashLength], e.majorityLink.Bytes())
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

func bytesToExtExtra(extra []byte) (*extExtra, error) {
	var extExtra extExtra
	size := len(extra)
	for i := 0; i < size; {
		kind := extra[i]
		i++
		switch kind {
		case ExtendedDataTypeSealerJoin, ExtendedDataTypeSealerLeave:
			if i+common.AddressLength > size {
				log.Error("bytesToExtExtra", "extra", common.Bytes2Hex(extra), "i", i)
				return nil, errInvalidSealerAddressLength
			}
			extExtra.applications = append(extExtra.applications, sealerApplication{
				sealer: common.BytesToAddress(extra[i : i+common.AddressLength]),
				action: kind == ExtendedDataTypeSealerJoin,
			})
			i += common.AddressLength
		case ExtendedDataTypeSealerDigest:
			if i+common.HashLength > size {
				return nil, errInvalidSealerDigestLength
			}
			digest := common.BytesToHash(extra[i : i+common.HashLength])
			extExtra.sealersDigest = &digest
			i += common.HashLength
		case ExtendedDataTypeCrossLink:
			if i+common.HashLength > size {
				return nil, errInvalidCrosslinkLength
			}
			digest := common.BytesToHash(extra[i : i+common.HashLength])
			extExtra.majorityLink = &digest
			i += common.HashLength
		default:
			return nil, errInvalidExtExtraKind
		}
	}
	return &extExtra, nil
}
