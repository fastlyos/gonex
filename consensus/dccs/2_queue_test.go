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
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestMerkleHash1(t *testing.T) {
	expected := common.HexToHash("126f53715820d5d122efe9aee15ba2b6428fa410")
	leafs := [][]byte{
		common.Hex2Bytes("126f53715820d5d122efe9aee15ba2b6428fa410"),
	}
	root := merkleHash(leafs)
	if root != expected {
		t.Errorf("Failed: want: %x, have: %x", root, expected)
	}
}

func TestMerkleHash2(t *testing.T) {
	leafs := [][]byte{
		common.Hex2Bytes("126f53715820d5d122efe9aee15ba2b6428fa410"),
		common.Hex2Bytes("1367fc3b5c3ce52d61347c0fe2216e576cb2060e"),
	}
	expected := common.HexToHash("d57bbe64ccc96359b7b2dd146dfbcfb5a185b1d5f9e12b6d585d5e56fb18e38b")
	root := merkleHash(leafs)
	if root != expected {
		t.Errorf("Failed: want: %x, have: %x", root, expected)
	}
}

func TestMerkleHash3(t *testing.T) {
	expected := common.HexToHash("3a21e8cd2fe615f399f7edded825277cadf2c7593e6e8d08f3b81c998d3d075a")
	leafs := [][]byte{
		common.Hex2Bytes("1367fc3b5c3ce52d61347c0fe2216e576cb2060e"),
		common.Hex2Bytes("126f53715820d5d122efe9aee15ba2b6428fa410"),
		common.Hex2Bytes("14ee8236a2dbd16d8107c16f156481c01b8424f5"),
	}
	root := merkleHash(leafs)
	if root != expected {
		t.Errorf("Failed: want: %x, have: %x", root, expected)
	}
}

func TestMerkleHash4(t *testing.T) {
	expected := common.HexToHash("9d61bdc11b5ae2090da613ff1a140c248c4baf135d1f14d7b208ac416f540ed7")
	leafs := [][]byte{
		common.Hex2Bytes("126f53715820d5d122efe9aee15ba2b6428fa410"),
		common.Hex2Bytes("14ee8236a2dbd16d8107c16f156481c01b8424f5"),
		common.Hex2Bytes("1367fc3b5c3ce52d61347c0fe2216e576cb2060e"),
		common.Hex2Bytes("15f8317b560489705d8aa3f5a2cba32b8b4d14da"),
	}
	root := merkleHash(leafs)
	if root != expected {
		t.Errorf("Failed: want: %x, have: %x", root, expected)
	}
}

func TestMerkleHash13(t *testing.T) {
	expected := common.HexToHash("c8650020797651420f7d83b1a45815898f136ac3876dbc34e16b83552dda7f36")
	leafs := [][]byte{
		common.Hex2Bytes("1c050030ec6979fa0099403fb83372896981bce2"),
		common.Hex2Bytes("35fbcac4fce527290a55c793bf17437d95f5038d"),
		common.Hex2Bytes("3742a4a59260da9dedb41f08d35fb3097c9bd658"),
		common.Hex2Bytes("cf5995b244355f7ba6dc3dbcf40fda37fb11a068"),
		common.Hex2Bytes("398d4a1cc5bb07df81e185aa2937346a3f756012"),
		common.Hex2Bytes("fadf93a13f5c0fced717fb8543ca59b3e90a2edc"),
		common.Hex2Bytes("b7c4a482e69a93f4a4d7c59d4687bd23cfd2cf4f"),
		common.Hex2Bytes("3b73b4c06523f9e05cbd7c10c025529226fb72e7"),
		common.Hex2Bytes("9612658163b41d03a09622599ba74112757bcc85"),
		common.Hex2Bytes("3c09c95848158d41d03f699671d576f4c4bdf39a"),
		common.Hex2Bytes("b0ad402eadbdf0d474d8ffacf6d3a42393ce4a88"),
		common.Hex2Bytes("15f8317b560489705d8aa3f5a2cba32b8b4d14da"),
		common.Hex2Bytes("e5caa37c25c80e2eb309a51619cc441aa2e0419d"),
	}
	root := merkleHash(leafs)
	if root != expected {
		t.Errorf("Failed: want: %x, have: %x", root, expected)
	}
}
