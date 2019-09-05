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

const (
	ITERATION_0 uint64 = 123456
	ITERATION_1 uint64 = 135246
)

var (
	input0     = []byte{0x01, 0x02, 0x03, 0x04}
	input1     = []byte{0x04, 0x03, 0x01, 0x01}
	iteration0 = ITERATION_0
	iteration1 = ITERATION_1
)

func NoCLI() {
	InitCLI("non-exist") // use the vdf-go instead
	iteration0 = ITERATION_0 / 20
	iteration1 = ITERATION_1 / 20
}

func TestGenerateVerify(t *testing.T) {
	bitSize := uint64(127)
	stopCh := make(chan struct{})
	output, err := Instance().Generate(input0, iteration0, bitSize, stopCh)
	if err != nil {
		t.Error("Error", "err", err)
		return
	}
	if !Instance().Verify(input0, output, iteration0, bitSize) {
		t.Error("Invalid ouput", "output", common.Bytes2Hex(output))
	}
}

func TestGenerateVerifyGo(t *testing.T) {
	NoCLI()
	TestGenerateVerify(t)
}

func TestInterruptedGenerator(t *testing.T) {
	bitSize := uint64(127)
	stopCh := make(chan struct{})
	go func() {
		time.Sleep(time.Second / 2)
		select {
		case stopCh <- struct{}{}:
		default:
		}
	}()
	output, err := Instance().Generate(input0, iteration0, bitSize, stopCh)
	if err != nil {
		t.Error("Error", "err", err)
		return
	}
	if output != nil {
		t.Error("Invalid ouput", "output", common.Bytes2Hex(output))
	}
}

func TestInterruptedGeneratorGo(t *testing.T) {
	NoCLI()
	TestInterruptedGenerator(t)
}
