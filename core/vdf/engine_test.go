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
	vdfGen = "vdf-cli"

	input0     = []byte{0x01, 0x02, 0x03, 0x04}
	input1     = []byte{0x04, 0x03, 0x01, 0x01}
	iteration0 = ITERATION_0
	iteration1 = ITERATION_1
)

func UseGoVDF() {
	vdfGen = vdfGenInternal
	iteration0 = ITERATION_0 / 20
	iteration1 = ITERATION_1 / 20
}

func TestGenerateVerify(t *testing.T) {
	bitSize := uint64(127)
	stopCh := make(chan struct{})
	output, err := Generator(vdfGen).Generate(input0, iteration0, bitSize, stopCh)
	if err != nil {
		t.Error("Error", "err", err)
		return
	}
	if !Verify(input0, output, iteration0, bitSize) {
		t.Error("Invalid ouput", "output", common.Bytes2Hex(output))
	}
}

func TestGenerateVerifyGo(t *testing.T) {
	UseGoVDF()
	TestGenerateVerify(t)
}

func TestVerifyCache(t *testing.T) {
	input := common.Hex2Bytes("0123456789abcdef")
	output := common.Hex2Bytes("31f44950a7467e3eb03d5e2c1513eea580ed5de129d9603ed1a34a2921555e59713deaaccb01fc8ca5c7bc4e017bc6a6520b2de17ae3065d5f7fc8023c663f6efe9a2e04fd7bfe28722c27fe0ca849c134f106f77b06375bd7b9c681d28f5da8a3404447c42082e0e5c87d57acd89d45f8de23ce10a811f61a68a80632a8326a309f2e569475fad67b52c59ae4c8a435e44c34f3e4ba08bd77319a916f820f2bc5bd830e76de7b6d4383a28769420fc0b01b8c39ad29f578bd3e957374597f69f7137ab26c2424297d7e5ef7c2ac62a802641496b9abfa28ad4185ec65ac2696e2882f07e8222333ee8b11e1562436632fb2900b0ecbc915d8f952407c203b470f6fa753415bf33621ecb09614dfe7a8f8ac37e3d8e51d97aacfdfdb821979c4241459b235d6e6f158cae17c5ffb284b2ab05317059eea9a4f471a7cde7e791107db93e2c61616cb6120e1c2143b498237623aa962551501bb5520eefe822da1bca7296ae803bc775708e2cccfddbc35652a09ec3c0350369a77bd1ce103a8acfd0b43c663f1e93ceec32e1ea05f6d7e9e47a5036caa7aaf6dfc7d443084dfe06997c15e126d7c9786a32ba94c764a42c742c85decd02b46f0e504f0e4a7923d68349009df078804535b5f95a62647487dd018a89469938db478d4b54190dddd22c8ec70ebae842b2422ff159afee0e245dc933241ebf9874e0a8c56c7158539")
	for i := 0; i < 999999; i++ {
		if !Verify(input, output, 1234567, 2047) {
			t.Error("Invalid ouput", "output", common.Bytes2Hex(output))
		}
	}
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
	output, err := Generator(vdfGen).Generate(input0, iteration0, bitSize, stopCh)
	if err != nil {
		t.Error("Error", "err", err)
		return
	}
	if output != nil {
		t.Error("Invalid ouput", "output", common.Bytes2Hex(output))
	}
}

func TestInterruptedGeneratorGo(t *testing.T) {
	UseGoVDF()
	TestInterruptedGenerator(t)
}

func TestLeak(t *testing.T) {
	for i := uint64(1); i < iteration0/300; i++ {
		output, err := Generator(vdfGen).Generate(input0, i, 127, nil)
		if err != nil {
			t.Error("error", "err", err)
		}
		if output != nil {
			if Verify(input0, output, i, 127) {
				t.Log("success", "output", common.Bytes2Hex(output))
			} else {
				t.Error("failed", "output", common.Bytes2Hex(output))
			}
		}
	}
}

func TestLeakGo(t *testing.T) {
	UseGoVDF()
	TestLeak(t)
}
