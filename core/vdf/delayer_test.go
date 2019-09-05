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
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func TestOutput(t *testing.T) {
	delayer := NewDelayer(32)
	output := delayer.Get(input0, iteration0)
	t.Log("main routine", "output", common.Bytes2Hex(output))
}

func TestOutputGo(t *testing.T) {
	NoCLI()
	TestOutput(t)
}

func TestSequentialOutput(t *testing.T) {
	var wg sync.WaitGroup
	delayer := NewDelayer(32)
	wg.Add(1)
	go func() {
		defer wg.Done()
		output := delayer.Get(input0, iteration0)
		if output == nil {
			t.Log("sub routine: replaced")
			return
		}
		if delayer.Verify(input0, output, iteration0) {
			t.Log("sub routine", "output", common.Bytes2Hex(output))
		} else {
			t.Error("sub routine: invalid", "output", common.Bytes2Hex(output))
		}
	}()
	wg.Wait()
	output := delayer.Get(input1, iteration1)
	if output == nil {
		t.Log("main routine: replaced")
		return
	}
	if delayer.Verify(input1, output, iteration1) {
		t.Log("main routine", "output", common.Bytes2Hex(output))
	} else {
		t.Error("main routine: invalid", "output", common.Bytes2Hex(output))
	}
}

func TestSequentialOutputGo(t *testing.T) {
	NoCLI()
	TestSequentialOutput(t)
}

func TestReplacedOutput(t *testing.T) {
	var wg sync.WaitGroup
	delayer := NewDelayer(32)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second / 4)
		output := delayer.Get(input0, iteration0)
		if output == nil {
			t.Log("sub routine: replaced")
			return
		}
		if delayer.Verify(input0, output, iteration0) {
			t.Log("sub routine", "output", common.Bytes2Hex(output))
		} else {
			t.Error("sub routine: invalid", "output", common.Bytes2Hex(output))
		}
	}()
	defer wg.Wait()
	output := delayer.Get(input1, iteration1)
	if output == nil {
		t.Log("main routine: replaced")
		return
	}
	if delayer.Verify(input1, output, iteration1) {
		t.Log("main routine", "output", common.Bytes2Hex(output))
	} else {
		t.Error("main routine: invalid", "output", common.Bytes2Hex(output))
	}
}

func TestReplacedOutputGo(t *testing.T) {
	NoCLI()
	TestReplacedOutput(t)
}

func TestRacingOutput(t *testing.T) {
	var wg sync.WaitGroup
	delayer := NewDelayer(32)
	wg.Add(1)
	go func() {
		defer wg.Done()
		output := delayer.Get(input0, iteration0)
		if output == nil {
			t.Log("sub routine: replaced")
			return
		}
		if delayer.Verify(input0, output, iteration0) {
			t.Log("sub routine", "output", common.Bytes2Hex(output))
		} else {
			t.Error("sub routine: invalid", "output", common.Bytes2Hex(output))
		}
	}()
	defer wg.Wait()
	output := delayer.Get(input1, iteration1)
	if output == nil {
		t.Log("main routine: replaced")
		return
	}
	if delayer.Verify(input1, output, iteration1) {
		t.Log("main routine", "output", common.Bytes2Hex(output))
	} else {
		t.Error("main routine: invalid", "output", common.Bytes2Hex(output))
	}
}

func TestRacingOutputGo(t *testing.T) {
	NoCLI()
	TestRacingOutput(t)
}

func TestOutputBroadcasting(t *testing.T) {
	var wg sync.WaitGroup
	delayer := NewDelayer(32)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			output := delayer.Get(input0, iteration0)
			if output == nil {
				t.Log("routine: replaced", "i", i)
			} else if delayer.Verify(input0, output, iteration0) {
				t.Log("routine", "i", i, "output", common.Bytes2Hex(output))
			} else {
				t.Error("routine: invalid", "i", i, "output", common.Bytes2Hex(output))
			}
		}(i)
	}
	output := delayer.Get(input0, iteration0)
	if delayer.Verify(input0, output, iteration0) {
		t.Log("main routine", "output", common.Bytes2Hex(output))
	} else {
		t.Error("main routine: invalid", "output", common.Bytes2Hex(output))
	}
	wg.Wait()
}

func TestOutputBroadcastingGo(t *testing.T) {
	NoCLI()
	TestOutputBroadcasting(t)
}

func TestOutputReplacedBroadcasting(t *testing.T) {
	var wg sync.WaitGroup
	delayer := NewDelayer(32)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			output := delayer.Get(input0, iteration0)
			if output == nil {
				t.Log("routine: replaced", "i", i)
			} else if delayer.Verify(input0, output, iteration0) {
				t.Log("routine", "i", i, "output", common.Bytes2Hex(output))
			} else {
				t.Error("routine: invalid", "i", i, "output", common.Bytes2Hex(output))
			}
		}(i)
	}
	time.Sleep(time.Second / 2)
	for i := 3; i < 6; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			output := delayer.Get(input1, iteration1)
			if delayer.Verify(input1, output, iteration1) {
				t.Log("routine", "i", i, "output", common.Bytes2Hex(output))
			} else {
				t.Error("routine: invalid", "i", i, "output", common.Bytes2Hex(output))
			}
		}(i)
	}
	output := delayer.Get(input1, iteration1)
	if delayer.Verify(input1, output, iteration1) {
		t.Log("main routine", "output", common.Bytes2Hex(output))
	} else {
		t.Error("main routine: invalid", "output", common.Bytes2Hex(output))
	}
	wg.Wait()
}

func TestOutputReplacedBroadcastingGo(t *testing.T) {
	NoCLI()
	TestOutputReplacedBroadcasting(t)
}

func TestCache(t *testing.T) {
	delayer := NewDelayer(32)
	output := delayer.Peek(input0, iteration0)
	t.Log("peek", "output", common.Bytes2Hex(output))
	output = delayer.Get(input0, iteration0)
	t.Log("first get", "output", common.Bytes2Hex(output))
	output = delayer.Peek(input0, iteration0)
	t.Log("peek", "output", common.Bytes2Hex(output))
	output = delayer.Get(input0, iteration0)
	t.Log("second get", "output", common.Bytes2Hex(output))
}

func TestCacheGo(t *testing.T) {
	NoCLI()
	TestCache(t)
}

func TestRequest(t *testing.T) {
	delayer := NewDelayer(32)
	delayer.Request(input0, iteration0)
	delayer.Request(input1, iteration1)
	time.Sleep(time.Second / 4)
	output := delayer.Get(input0, iteration0)
	if delayer.Verify(input0, output, iteration0) {
		t.Log("success", "output", common.Bytes2Hex(output))
	} else {
		t.Error("failed", "output", common.Bytes2Hex(output))
	}
}

func TestRequestGo(t *testing.T) {
	NoCLI()
	TestRequest(t)
}
