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

func TestAddressHash(t *testing.T) {
	want := common.HexToHash("00d98964e677cd2f659f61db176cc55a6f06f74e07cc1e78664ae690c6452763")
	adrs := []common.Address{
		common.HexToAddress("cf5995b244355f7ba6dc3dbcf40fda37fb11a068"),
	}
	have := addressesHash(adrs)
	if have != want {
		t.Errorf("Failed: want=%x, have=%x", want, have)
	}
}

func TestAddressesHash(t *testing.T) {
	want := common.HexToHash("5a900ea7d7054c0962bda38d7fe0a32300c577662f357cfb2ef2f6b3714105de")
	adrs := []common.Address{
		common.HexToAddress("1c050030ec6979fa0099403fb83372896981bce2"),
		common.HexToAddress("35fbcac4fce527290a55c793bf17437d95f5038d"),
		common.HexToAddress("3742a4a59260da9dedb41f08d35fb3097c9bd658"),
		common.HexToAddress("cf5995b244355f7ba6dc3dbcf40fda37fb11a068"),
		common.HexToAddress("398d4a1cc5bb07df81e185aa2937346a3f756012"),
		common.HexToAddress("fadf93a13f5c0fced717fb8543ca59b3e90a2edc"),
		common.HexToAddress("b7c4a482e69a93f4a4d7c59d4687bd23cfd2cf4f"),
	}
	have := addressesHash(adrs)
	if have != want {
		t.Errorf("Failed: want=%x, have=%x", want, have)
	}
}
