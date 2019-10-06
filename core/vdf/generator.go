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
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/log"
	"github.com/harmony-one/vdf/src/vdf_go"
)

const vdfGenInternal = "internal"
const vdfGenDisable = "disable"

type generator struct {
	cli string
}

var goGenerator = generator{}
var generators sync.Map

func Generator(cliName string) *generator {
	if len(cliName) == 0 || cliName == vdfGenDisable {
		log.Warn("VDF generator: disabled")
		return nil
	}
	if cliName == vdfGenInternal {
		log.Warn("VDF generator: internal")
		return &goGenerator
	}
	if val, ok := generators.Load(cliName); ok {
		return val.(*generator)
	}
	cli, err := exec.LookPath(cliName)
	if err != nil {
		log.Error("VDF generator: disabled", cliName, "not found")
		return nil
	}
	gen := generator{cli}
	ret, _ := generators.LoadOrStore(cliName, &gen)
	log.Info("VDF generator", "cli", cli)
	return ret.(*generator)
}

// Generate generates the vdf output = (y, proof)
func (g *generator) Generate(seed []byte, iteration uint64, bitSize uint64, stop <-chan struct{}) (output []byte, err error) {
	if g == nil {
		return nil, errors.New("VDF generator disabled")
	}
	if len(g.cli) > 0 {
		return g.generateCLI(seed, iteration, bitSize, stop)
	}
	defer func() {
		if x := recover(); x != nil {
			log.Error("vdf.Generate: generation process panic", "reason", x)
			err = fmt.Errorf("%v", x)
		}
	}()
	blockingStop := make(chan struct{})
	var done chan struct{}
	if stop != nil {
		// always-listen adapter for blocking stop chan
		done = make(chan struct{})
		go func() {
			select {
			case <-stop:
				log.Trace("vdf.Generate: vdf-go interrupted")
				blockingStop <- struct{}{}
				return
			case <-done:
				log.Trace("vdf.Generate: vdf-go done")
				return
			}
		}()
		// give channel listening routine a chance to run first
		runtime.Gosched()
	}
	y, proof := vdf_go.GenerateVDFWithStopChan(seed, int(iteration), int(bitSize), blockingStop)

	if stop != nil && done != nil {
		// release the stopping goroutine above
		select {
		case done <- struct{}{}:
		default:
		}
	}

	if y == nil || proof == nil {
		return nil, nil
	}
	return append(y, proof...), nil
}

func (g *generator) generateCLI(seed []byte, iteration uint64, bitSize uint64, stop <-chan struct{}) (output []byte, err error) {
	cmd := exec.Command(g.cli,
		"-l"+strconv.Itoa(int(bitSize)),
		common.Bytes2Hex(seed),
		strconv.Itoa(int(iteration)))

	log.Trace(g.cli + " -l" + strconv.Itoa(int(bitSize)) + " " + common.Bytes2Hex(seed) + " " + strconv.Itoa(int(iteration)))

	var done chan struct{}
	if stop != nil {
		done = make(chan struct{})
		go func() {
			select {
			case <-stop:
				if cmd == nil {
					return
				}
				if cmd.Process == nil {
					// process is signaled to kill before it's even started
					if !func() bool {
						for i := 0; i < 1000000; i++ {
							// yeild and wait for process to start first
							runtime.Gosched()
							if cmd.Process != nil {
								return true
							}
						}
						return false
					}() {
						log.Info("vdf.Generate: non-exist vdf-cli interrupted")
						return
					}
				}
				log.Trace("vdf.Generate: vdf-cli interrupted")
				if err := cmd.Process.Kill(); err != nil {
					log.Error("vdf.Generate: failed to kill vdf-cli process", "err", err)
				}
				return
			case <-done:
				log.Trace("vdf.Generate: vdf-cli done")
				return
			}
		}()
		// give channel listening routine a chance to run first
		runtime.Gosched()
	}

	log.Trace("vdf.Generate", "seed", common.Bytes2Hex(seed), "iteration", iteration, "output", common.Bytes2Hex(output))
	output, err = cmd.Output()

	if stop != nil && done != nil {
		// release the stopping goroutine above
		select {
		case done <- struct{}{}:
		default:
		}
	}

	if err != nil {
		if cmd.ProcessState != nil {
			status := cmd.ProcessState.Sys().(syscall.WaitStatus)
			if status.Signaled() && status.Signal() == syscall.SIGKILL {
				// interrupted, nuffin wong
				log.Error("vdf.Generate: interrupted")
				return nil, nil
			}
		}
		if err, ok := err.(*exec.ExitError); ok {
			// verification failed
			log.Trace("vdf.Generate", "error code", err.Error())
			return nil, err
		}
		// sum ting wong
		log.Error("vdf.Generate", "error", err)
		return nil, err
	}

	strOutput := strings.TrimSpace(string(output))

	log.Trace("vdf.Generate", "output", strOutput)
	return common.Hex2Bytes(strOutput), nil
}
