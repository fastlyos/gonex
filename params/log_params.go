// Copyright 2015 The Nexty Authors
// This file is part of the gonex library.
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

package params

const (
	ErrorLogRevertUnknown            = "revert with no reason"
	ErrorLogInvalidOpCode            = "invalid opcode 0x%x"
	ErrorLogStackUnderflow           = "stack underflow (%d <=> %d)"
	ErrorLogStackLimitReached        = "stack limit reached %d (%d)"
	ErrorLogWriteProtection          = "write protection"
	ErrorLogOutOfGas                 = "out of gas"
	ErrorLogGasUintOverflow          = "gas overflow unsigned 64 bit integer"
	ErrorLogReturnDataOutOfBounds    = "return data out of bounds"
	ErrorLogMaxCodeSizeExceeded      = "max code size exceeded"
	ErrorLogInvalidJump              = "invalid jump destination"
	ErrorLogCodeStoreOutOfGas        = "contract creation code storage out of gas"
	ErrorLogDepth                    = "max call depth exceeded"
	ErrorLogInsufficientBalance      = "insufficient balance for transfer"
	ErrorLogContractAddressCollision = "contract address collision"
	ErrorLogTxCodeOverspent          = "tx code value limit overspent"
)

var (
	// GonexErrorSignature is Keccak("Error")
	GonexErrorSignature = []byte{0xe3, 0x42, 0xda, 0xa4}
)
