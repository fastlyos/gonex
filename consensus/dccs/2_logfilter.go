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
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// Keccak256("Joined(address,address)")
	joinedTopic = common.HexToHash("7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea")
	// Keccak256("Left(address,address)")
	leftTopic = common.HexToHash("4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5")
)

type logFilterBackend struct {
	chain consensus.ChainReader
	db    ethdb.Reader
}

func (b *logFilterBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	return b.chain.GetHeaderByNumber(uint64(blockNr.Int64())), nil
}

func (b *logFilterBackend) HeaderByHash(ctx context.Context, blockHash common.Hash) (*types.Header, error) {
	return b.chain.GetHeaderByHash(blockHash), nil
}

func (b *logFilterBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	header := b.chain.GetHeaderByHash(blockHash)
	if header == nil {
		return nil, nil
	}
	receipts := rawdb.ReadReceipts(b.db, blockHash, header.Number.Uint64(), b.chain.Config())
	return receipts, nil
}

func (b *logFilterBackend) GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error) {
	receipts, _ := b.GetReceipts(ctx, blockHash)
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

// This nil assignment ensures compile time that logFilterBackend implements filters.SimpleBackend.
var _ filters.SimpleBackend = (*logFilterBackend)(nil)

// fetchSealerApplications filters the block for any joining or leaving sealer.
// Multiple sealer applications can be confirmed in the same block, the order of
// the requests kept as is.
func (c *Context) fetchSealerApplications(header *types.Header) ([]SealerApplication, error) {
	logs, err := filters.BlockLogs(header,
		[]common.Address{params.GovernanceAddress},
		[][]common.Hash{{joinedTopic, leftTopic}},
		&logFilterBackend{
			chain: c.chain,
			db:    c.engine.db,
		})

	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, nil
	}

	applications := make([]SealerApplication, len(logs))

	for i, l := range logs {
		// len(log.Data) must be 32 * 2 here
		sealer := common.BytesToAddress(l.Data[32:])
		applications[i] = SealerApplication{
			sealer: sealer,
		}
		if l.Topics[0] == joinedTopic {
			staker := common.BytesToAddress(l.Data[:32])
			applications[i].action = ExtendedDataTypeSealerJoin
			log.Info("Sealer application", "join", sealer, "coinbase", staker)
		} else {
			applications[i].action = ExtendedDataTypeSealerLeave
			log.Info("Sealer application", "leave", sealer)
		}
	}
	return applications, nil
}
