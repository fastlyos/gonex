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
	"fmt"
	"runtime"
	"sync"

	lru "github.com/hashicorp/golang-lru"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Delayer is a single process/routine for delay task of <seed, iteration>.
// The same task request will be ignored; a new task request will interrupt
// and replace the current one.
//
// bitSize(bit) will be outputSize(byte)<<2 - 1
//
// ORIGINAL:
//   bitSize value of 2^n-1 is recommened
//   output size (in bytes) will be ((bitSize+16)>>4)*4
type Delayer struct {
	bitSize  uint64
	loopOnce sync.Once
	stopCh   chan struct{}        // to stop all running vdf routines
	reqCh    chan task            // to request new delay task
	resChCh  chan (<-chan []byte) // to get the chan that will return the vdf output

	outputCache *lru.ARCCache // task.GetKey() => []byte
}

// NewDelayer creates a new Delayer instance
func NewDelayer(outputSize uint64) *Delayer {
	outputCache, _ := lru.NewARC(8)
	return &Delayer{
		bitSize: outputSize<<2 - 1,
		stopCh:  make(chan struct{}),
		reqCh:   make(chan task),
		resChCh: make(chan (<-chan []byte)),

		outputCache: outputCache,
	}
}

// Verify verifies the given output against the seed and iteration
func (d *Delayer) Verify(seed, output []byte, iteration uint64) bool {
	t := task{
		seed:      seed,
		iteration: iteration,
	}
	if cached, ok := d.outputCache.Get(t.GetKey()); ok {
		return bytes.Equal(output, cached.([]byte))
	}
	return Instance().Verify(seed, output, iteration, d.bitSize)
}

// Get request new delay task and block for output.
//
// The same task request will be ignored; a new task request will interrupt
// and replace the current one.
//
// Requested routine MUST read exactly 1 value from returning chanel, nil value
// indicates that the request is replaced by other request.
func (d *Delayer) Get(seed []byte, iteration uint64) []byte {
	t := task{
		seed:      seed,
		iteration: iteration,
	}
	if output, ok := d.outputCache.Get(t.GetKey()); ok {
		return output.([]byte)
	}
	d.loopOnce.Do(func() {
		go d.loop()
	})
	d.reqCh <- t
	return <-<-d.resChCh
}

// Request requests new delay task and return.
// Return the cached output if available.
func (d *Delayer) Request(seed []byte, iteration uint64) []byte {
	t := task{
		seed:      seed,
		iteration: iteration,
	}
	if output, ok := d.outputCache.Get(t.GetKey()); ok {
		return output.([]byte)
	}
	d.loopOnce.Do(func() {
		go d.loop()
	})
	d.reqCh <- t
	<-d.resChCh // discard the output chan
	return nil
}

// Peek attempts to get the cached result, without requesting
// any delay task.
func (d *Delayer) Peek(seed []byte, iteration uint64) []byte {
	t := task{
		seed:      seed,
		iteration: iteration,
	}
	if output, ok := d.outputCache.Get(t.GetKey()); ok {
		return output.([]byte)
	}
	return nil
}

// Stop stops the current delay task.
func (d *Delayer) Stop() {
	// cancel all currently running routines
	for {
		select {
		case d.stopCh <- struct{}{}:
		default:
			// no more routines to stop
			return
		}
	}
}

func (d *Delayer) loop() {
	var current task
	var resCh chan []byte
	for {
		select {
		case t := <-d.reqCh:
			if t.Equal(current) {
				log.Trace("Delayer: discarding duplicated request", "seed", common.Bytes2Hex(t.seed), "iteration", t.iteration)
				d.resChCh <- resCh
				continue
			}
			d.Stop()
			// new task
			current = t
			resCh = make(chan []byte)
			d.resChCh <- resCh

			// start new worker routine
			go func(t task, resCh chan<- []byte) {
				output, err := Instance().Generate(t.seed, t.iteration, d.bitSize, d.stopCh)
				defer close(resCh)
				if err != nil {
					log.Error("Delayer: VDF worker loop failed", "err", err)
					fmt.Printf("Delayer: VDF worker loop failed, err=%v\n", err)
					return
				}
				if len(output) == 0 {
					log.Info("Delayer: interrupted")
					return
				}
				// cache the result
				d.outputCache.Add(t.GetKey(), output)
				// broadcast to all listening chan
				for {
					select {
					case resCh <- output:
					default:
						// no more listener to broadcast
						break
					}
				}
			}(t, resCh)
			// give the forked worker a chance to start first
			runtime.Gosched()
		}
	}
}

// send empty task{} to stop the calcuation
type task struct {
	seed      []byte
	iteration uint64
}

func (t task) Equal(u task) bool {
	return t.iteration == u.iteration && bytes.Equal(t.seed, u.seed)
}

func (t task) GetKey() string {
	buf := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(buf, t.iteration)
	return string(append(buf, t.seed...))
}
