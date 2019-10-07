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

// Package vdf implements the VDF engine.
package vdf

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"sync"

	lru "github.com/hashicorp/golang-lru"

	"github.com/ethereum/go-ethereum/log"
	"github.com/harmony-one/vdf/src/vdf_go"
)

const (
	inmemoryVDFCache = 32
)

var (
	cache     *lru.ARCCache // vdfInput => []byte
	cacheOnce sync.Once
)

type vdfInput struct {
	seed      []byte
	iteration uint64
	bitSize   uint64
}

func inputKey(seed []byte, iteration uint64, bitSize uint64) uint64 {
	hasher := fnv.New64a()
	hasher.Write(seed)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, iteration)
	hasher.Write(b)
	binary.BigEndian.PutUint64(b, bitSize)
	hasher.Write(b)
	return hasher.Sum64()
}

// Verify verifies the generated output against the seed
func Verify(seed, output []byte, iteration uint64, bitSize uint64) (valid bool) {
	cacheOnce.Do(func() {
		cache, _ = lru.NewARC(inmemoryVDFCache)
	})
	key := inputKey(seed, iteration, bitSize)
	if value, ok := cache.Get(key); ok {
		return bytes.Compare(output, value.([]byte)) == 0
	}
	defer func() {
		if x := recover(); x != nil {
			log.Error("vdf.Verify: verification process panic", "reason", x)
			valid = false
		}
	}()
	if !vdf_go.VerifyVDF(seed, output, int(iteration), int(bitSize)) {
		return false
	}
	// verify success, cache the output
	cache.Add(key, output)
	return true
}
