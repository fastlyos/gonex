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
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func TestGenerateVerify(t *testing.T) {
	input := []byte{0x01, 0x02, 0x03, 0x04}
	iteration := uint64(12345)
	bitSize := uint64(127)
	stopCh := make(chan struct{})
	output, err := Instance().Generate(input, iteration, bitSize, stopCh)
	if err != nil {
		t.Error("Error", "err", err)
		return
	}
	if !Instance().Verify(input, output, iteration, bitSize) {
		t.Error("Invalid ouput", "output", common.Bytes2Hex(output))
	}
}

func TestInterruptedGenerator(t *testing.T) {
	input := []byte{0x01, 0x02, 0x03, 0x04}
	iteration := uint64(123456789)
	bitSize := uint64(127)
	stopCh := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		stopCh <- struct{}{}
	}()
	output, err := Instance().Generate(input, iteration, bitSize, stopCh)
	if err != nil {
		t.Error("Error", "err", err)
		return
	}
	if output != nil {
		t.Error("Invalid ouput", "output", common.Bytes2Hex(output))
	}
}
