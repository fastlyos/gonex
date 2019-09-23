// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endurio

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbsorbableABI is the input ABI used to generate the binding from.
const AbsorbableABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbsorbableFuncSigs maps the 4-byte function signature to its string representation.
var AbsorbableFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"8aa3f897": "top(bool)",
}

// AbsorbableBin is the compiled bytecode used for deploying new contracts.
var AbsorbableBin = "0x608060405262049d4060035560026003548161001757fe5b0460045534801561002757600080fd5b50604051611dfc380380611dfc8339818101604052604081101561004a57600080fd5b508051602090910151801561005f5760038190555b600082116100705760028104610072565b815b6004555050611d76806100866000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638aa3f897116100715780638aa3f897146101ba578063aa1c259c146101d9578063be91d72914610207578063ced4aac814610224578063ee1a68c614610264578063f318722b14610287576100b4565b806307c399a3146100b95780630d90b10a1461011357806343271d791461014a5780634ea097971461017157806369c07d31146101965780636e6452cb146101b2575b600080fd5b6100de600480360360408110156100cf57600080fd5b508035151590602001356102b3565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101386004803603604081101561012957600080fd5b508035151590602001356102fa565b60408051918252519081900360200190f35b61016f6004803603604081101561016057600080fd5b50803515159060200135610322565b005b6101386004803603604081101561018757600080fd5b508035151590602001356103a8565b61019e6103cc565b604080519115158252519081900360200190f35b61019e6103d1565b610138600480360360208110156101d057600080fd5b503515156103d6565b61016f600480360360408110156101ef57600080fd5b506001600160a01b03813581169160200135166103f9565b61016f6004803603602081101561021d57600080fd5b5035610574565b610138600480360360a081101561023a57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610711565b61026c610779565b60408051921515835260208301919091528051918290030190f35b6101386004803603604081101561029d57600080fd5b506001600160a01b038135169060200135610823565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610392576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b6103a2828463ffffffff61083616565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b80151560009081526020819052604081206103f081610930565b9150505b919050565b6001546001600160a01b031615610457576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b0316156104b5576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556104ef8282610949565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561053457600080fd5b505afa158015610548573d6000803e3d6000fd5b505050506040513d602081101561055e57600080fd5b5051905061056f81806000806109bd565b505050565b33156105b8576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6105c26005610a5c565b156105cf576105cf610a75565b6105d96009610ab9565b156105e6576105e6610ad6565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561062b57600080fd5b505afa15801561063f573d6000803e3d6000fd5b505050506040513d602081101561065557600080fd5b5051905081156106dd57610667610bd9565b1561067e5761067982826000806109bd565b61069b565b6106888183610beb565b1561069b5761069b8282600160006109bd565b6106a56009610c99565b156106dd5760006106b68383610cb5565b90506106c181610ccf565b80156106cf575060085460ff165b156106db57505061070e565b505b6106ee60058263ffffffff610ded16565b1561070c5760006106fd610e22565b905061070881610ed7565b5050505b505b50565b841515600090815260208190526040812061072a611ced565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261076d82828663ffffffff610f7e16565b98975050505050505050565b60075460009081906107905750600090508061081f565b600161081a600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107e957600080fd5b505afa1580156107fd573d6000803e3d6000fd5b505050506040513d602081101561081357600080fd5b5051610cb5565b915091505b9091565b600061082f8383610ff0565b9392505050565b61083e611ced565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b0316815260018201549281018390529281015493830193909352600383015460608301526004909201546080820152901561092057825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b1580156108f357600080fd5b505af1158015610907573d6000803e3d6000fd5b505050506040513d602081101561091d57600080fd5b50505b61056f838363ffffffff6110bd16565b6000808052600282016020526040902060040154919050565b60008080526020526109827fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff61111d16565b6001600090815260205261070c7fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff61111d16565b604080516080810182526003544301808252602082018690529181018690528215156060909101819052600591909155600684905560078590556008805460ff191690911790556000610a108585610cb5565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6000610a67826112ca565b801561031c57505054431190565b60006005819055600681905560078190556008805460ff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b6000610ac4826112d0565b801561031c5750506004015443101590565b610ae060096112d0565b610ae957610bd7565b600b5415610b7a57600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b505050506040513d6020811015610b7757600080fd5b50505b6009546040516001600160a01b03909116907f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d31826057290600090a2600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b6000610be56005610a5c565b90505b90565b600082821415610bfd5750600061031c565b6006546007541415610c115750600161031c565b82821115610c6557600654600754848403911015610c4757600654600754036002818381610c3b57fe5b0410159250505061031c565b600754600654036002828281610c5957fe5b0411159250505061031c565b600654600754838503911115610c8757600754600654036002818381610c3b57fe5b600654600754036002828281610c5957fe5b6000610ca4826112d0565b801561031c57505060040154431090565b6000818311610cc95782820360000361082f565b50900390565b6000610ce26009600101546000846112df565b610cee575060006103f4565b600c54600a546000919084830381610d0257fe5b0581610d0a57fe5b04905080610d16575060015b600b54811115610d345750600b54610d2c610a75565b610d34610ad6565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015610d8a57600080fd5b505af1158015610d9e573d6000803e3d6000fd5b50506009546040805185815290516001600160a01b0390921693507fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26925081900360200190a250600192915050565b6000610df8836112ca565b8015610e08575082600201548214155b801561082f575061082f836001015483856002015461130f565b600080610e39600560020154600560010154610cb5565b600754600254604080516318160ddd60e01b81529051939450600093610e8993926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156107e957600080fd5b9050610e97600082846112df565b610ea657600092505050610be8565b60006004548381610eb357fe5b059050610ec2600082846112df565b610ed057509150610be89050565b9250505090565b6000806000806000808613610eed576001610ef0565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff1615610f5757610f4c81610f318761133a565b600954859291906001600160a01b031663ffffffff61135016565b935093505050610f79565b610f7281610f648761133a565b84919063ffffffff61170316565b9350935050505b915091565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610fb6848263ffffffff6119d016565b15610f8f575b6003015460008181526002860160205260409020909250610fe3848263ffffffff6119d016565b610fbc5750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106110605780518252601f199092019160209182019101611041565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561109f573d6000803e3d6000fd5b5050506040513d60208110156110b457600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b54151590565b546001600160a01b0316151590565b60008284131580156112f15750818313155b8061130757508284121580156113075750818312155b949350505050565b60008284111580156113215750818311155b8061130757508284101580156113075750501115919050565b600080821361134c578160000361031c565b5090565b60008061136486868663ffffffff61170316565b90925090506000808661137857838361137b565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b1580156113d157600080fd5b505afa1580156113e5573d6000803e3d6000fd5b505050506040513d60208110156113fb57600080fd5b5051821180159061148357508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b15801561145357600080fd5b505afa158015611467573d6000803e3d6000fd5b505050506040513d602081101561147d57600080fd5b50518211155b156116f8578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b1580156114d757600080fd5b505af11580156114eb573d6000803e3d6000fd5b505050506040513d602081101561150157600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b15801561155a57600080fd5b505af115801561156e573d6000803e3d6000fd5b505050506040513d602081101561158457600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156115d257600080fd5b505af11580156115e6573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561163857600080fd5b505af115801561164c573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156116a857600080fd5b505af11580156116bc573d6000803e3d6000fd5b505050506040513d60208110156116d257600080fd5b508790506116e057806116e2565b815b83019250866116f157816116f3565b805b840193505b505094509492505050565b600080600061171186610930565b90505b600019811480159061172557508382105b156119c75760008181526002870160205260408120908661174a578160020154611750565b81600101545b9050600087611763578260010154611769565b82600201545b90508661177c868463ffffffff6119ee16565b1161186b57885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156117d257600080fd5b505af11580156117e6573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561183d57600080fd5b505af1158015611851573d6000803e3d6000fd5b5050505060048301546118648a86611a48565b935061199b565b600061187d888763ffffffff611b0b16565b9050828183028161188a57fe5b04915080925060008961189d578261189f565b835b905060008a6118ae57846118b0565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b1580156118ff57600080fd5b505af1158015611913573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561196557600080fd5b505af1158015611979573d6000803e3d6000fd5b5061199292508e9150899050848463ffffffff611b6216565b50600019955050505b6119ab858363ffffffff6119ee16565b94506119bd868263ffffffff6119ee16565b9550505050611714565b50935093915050565b60408201516001820154600290920154602090930151910291021190565b60008282018381101561082f576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b611a50611ced565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b031681526001820154928101929092529182015492810183905260038201546060820152600490910154608082015290156109205760018301548151604080840151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b1580156108f357600080fd5b600082821115610cc9576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b60008381526002850160205260409020611b7b816112d0565b611bbe576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b8060010154831115611c17576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b600181018054849003905560028101548210611c4357611c3d858563ffffffff611a4816565b506103a2565b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015611c9d57600080fd5b505af1158015611cb1573d6000803e3d6000fd5b505050506040513d6020811015611cc757600080fd5b50506002810180548390039055600181015461070857610708858563ffffffff6110bd16565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a72315820f6726f19173b0dfbceb27ec7a7219c26308696686f87d028b096ff896f841c5564736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployAbsorbable deploys a new Ethereum contract, binding an instance of Absorbable to it.
func DeployAbsorbable(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int) (common.Address, *types.Transaction, *Absorbable, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsorbableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbsorbableBin), backend, absorptionDuration, absorptionExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Absorbable{AbsorbableCaller: AbsorbableCaller{contract: contract}, AbsorbableTransactor: AbsorbableTransactor{contract: contract}, AbsorbableFilterer: AbsorbableFilterer{contract: contract}}, nil
}

// Absorbable is an auto generated Go binding around an Ethereum contract.
type Absorbable struct {
	AbsorbableCaller     // Read-only binding to the contract
	AbsorbableTransactor // Write-only binding to the contract
	AbsorbableFilterer   // Log filterer for contract events
}

// AbsorbableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsorbableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsorbableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsorbableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsorbableSession struct {
	Contract     *Absorbable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsorbableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsorbableCallerSession struct {
	Contract *AbsorbableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AbsorbableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsorbableTransactorSession struct {
	Contract     *AbsorbableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AbsorbableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsorbableRaw struct {
	Contract *Absorbable // Generic contract binding to access the raw methods on
}

// AbsorbableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsorbableCallerRaw struct {
	Contract *AbsorbableCaller // Generic read-only contract binding to access the raw methods on
}

// AbsorbableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsorbableTransactorRaw struct {
	Contract *AbsorbableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsorbable creates a new instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbable(address common.Address, backend bind.ContractBackend) (*Absorbable, error) {
	contract, err := bindAbsorbable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Absorbable{AbsorbableCaller: AbsorbableCaller{contract: contract}, AbsorbableTransactor: AbsorbableTransactor{contract: contract}, AbsorbableFilterer: AbsorbableFilterer{contract: contract}}, nil
}

// NewAbsorbableCaller creates a new read-only instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableCaller(address common.Address, caller bind.ContractCaller) (*AbsorbableCaller, error) {
	contract, err := bindAbsorbable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsorbableCaller{contract: contract}, nil
}

// NewAbsorbableTransactor creates a new write-only instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsorbableTransactor, error) {
	contract, err := bindAbsorbable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsorbableTransactor{contract: contract}, nil
}

// NewAbsorbableFilterer creates a new log filterer instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsorbableFilterer, error) {
	contract, err := bindAbsorbable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsorbableFilterer{contract: contract}, nil
}

// bindAbsorbable binds a generic wrapper to an already deployed contract.
func bindAbsorbable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsorbableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absorbable *AbsorbableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absorbable.Contract.AbsorbableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absorbable *AbsorbableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absorbable.Contract.AbsorbableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absorbable *AbsorbableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absorbable.Contract.AbsorbableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absorbable *AbsorbableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absorbable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absorbable *AbsorbableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absorbable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absorbable *AbsorbableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absorbable.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Absorbable *AbsorbableCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Absorbable *AbsorbableSession) Ask() (bool, error) {
	return _Absorbable.Contract.Ask(&_Absorbable.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Absorbable *AbsorbableCallerSession) Ask() (bool, error) {
	return _Absorbable.Contract.Ask(&_Absorbable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Absorbable *AbsorbableCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Absorbable *AbsorbableSession) Bid() (bool, error) {
	return _Absorbable.Contract.Bid(&_Absorbable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Absorbable *AbsorbableCallerSession) Bid() (bool, error) {
	return _Absorbable.Contract.Bid(&_Absorbable.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.CalcOrderID(&_Absorbable.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.CalcOrderID(&_Absorbable.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.FindAssistingID(&_Absorbable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.FindAssistingID(&_Absorbable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Absorbable *AbsorbableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Absorbable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Absorbable *AbsorbableSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Absorbable.Contract.GetOrder(&_Absorbable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Absorbable *AbsorbableCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Absorbable.Contract.GetOrder(&_Absorbable.CallOpts, _orderType, _id)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Absorbable *AbsorbableCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Absorbable.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Absorbable *AbsorbableSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Absorbable.Contract.GetRemainToAbsorb(&_Absorbable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Absorbable *AbsorbableCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Absorbable.Contract.GetRemainToAbsorb(&_Absorbable.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Next(&_Absorbable.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Next(&_Absorbable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Prev(&_Absorbable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Prev(&_Absorbable.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) Top(orderType bool) ([32]byte, error) {
	return _Absorbable.Contract.Top(&_Absorbable.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Absorbable.Contract.Top(&_Absorbable.CallOpts, orderType)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.Contract.Cancel(&_Absorbable.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.Contract.Cancel(&_Absorbable.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Absorbable.Contract.OnBlockInitialized(&_Absorbable.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Absorbable.Contract.OnBlockInitialized(&_Absorbable.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.Contract.RegisterTokens(&_Absorbable.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.Contract.RegisterTokens(&_Absorbable.TransactOpts, volatileToken, stablizeToken)
}

// AbsorbableAbsorptionIterator is returned from FilterAbsorption and is used to iterate over the raw logs and unpacked data for Absorption events raised by the Absorbable contract.
type AbsorbableAbsorptionIterator struct {
	Event *AbsorbableAbsorption // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbsorbableAbsorptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsorbableAbsorption)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbsorbableAbsorption)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbsorbableAbsorptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsorbableAbsorptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsorbableAbsorption represents a Absorption event raised by the Absorbable contract.
type AbsorbableAbsorption struct {
	Amount  *big.Int
	Supply  *big.Int
	Emptive bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbsorption is a free log retrieval operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Absorbable *AbsorbableFilterer) FilterAbsorption(opts *bind.FilterOpts) (*AbsorbableAbsorptionIterator, error) {

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return &AbsorbableAbsorptionIterator{contract: _Absorbable.contract, event: "Absorption", logs: logs, sub: sub}, nil
}

// WatchAbsorption is a free log subscription operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Absorbable *AbsorbableFilterer) WatchAbsorption(opts *bind.WatchOpts, sink chan<- *AbsorbableAbsorption) (event.Subscription, error) {

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsorbableAbsorption)
				if err := _Absorbable.contract.UnpackLog(event, "Absorption", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbsorption is a log parse operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Absorbable *AbsorbableFilterer) ParseAbsorption(log types.Log) (*AbsorbableAbsorption, error) {
	event := new(AbsorbableAbsorption)
	if err := _Absorbable.contract.UnpackLog(event, "Absorption", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbsorbableSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the Absorbable contract.
type AbsorbableSlashIterator struct {
	Event *AbsorbableSlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbsorbableSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsorbableSlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbsorbableSlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbsorbableSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsorbableSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsorbableSlash represents a Slash event raised by the Absorbable contract.
type AbsorbableSlash struct {
	Maker  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Absorbable *AbsorbableFilterer) FilterSlash(opts *bind.FilterOpts, maker []common.Address) (*AbsorbableSlashIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return &AbsorbableSlashIterator{contract: _Absorbable.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Absorbable *AbsorbableFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *AbsorbableSlash, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsorbableSlash)
				if err := _Absorbable.contract.UnpackLog(event, "Slash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlash is a log parse operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Absorbable *AbsorbableFilterer) ParseSlash(log types.Log) (*AbsorbableSlash, error) {
	event := new(AbsorbableSlash)
	if err := _Absorbable.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbsorbableStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the Absorbable contract.
type AbsorbableStopIterator struct {
	Event *AbsorbableStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbsorbableStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsorbableStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbsorbableStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbsorbableStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsorbableStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsorbableStop represents a Stop event raised by the Absorbable contract.
type AbsorbableStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Absorbable *AbsorbableFilterer) FilterStop(opts *bind.FilterOpts) (*AbsorbableStopIterator, error) {

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &AbsorbableStopIterator{contract: _Absorbable.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Absorbable *AbsorbableFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *AbsorbableStop) (event.Subscription, error) {

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsorbableStop)
				if err := _Absorbable.contract.UnpackLog(event, "Stop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Absorbable *AbsorbableFilterer) ParseStop(log types.Log) (*AbsorbableStop, error) {
	event := new(AbsorbableStop)
	if err := _Absorbable.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbsorbableUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the Absorbable contract.
type AbsorbableUnlockIterator struct {
	Event *AbsorbableUnlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbsorbableUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsorbableUnlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbsorbableUnlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbsorbableUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsorbableUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsorbableUnlock represents a Unlock event raised by the Absorbable contract.
type AbsorbableUnlock struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Absorbable *AbsorbableFilterer) FilterUnlock(opts *bind.FilterOpts, maker []common.Address) (*AbsorbableUnlockIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return &AbsorbableUnlockIterator{contract: _Absorbable.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Absorbable *AbsorbableFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *AbsorbableUnlock, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsorbableUnlock)
				if err := _Absorbable.contract.UnpackLog(event, "Unlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlock is a log parse operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Absorbable *AbsorbableFilterer) ParseUnlock(log types.Log) (*AbsorbableUnlock, error) {
	event := new(AbsorbableUnlock)
	if err := _Absorbable.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ITokenABI is the input ABI used to generate the binding from.
const ITokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"dexBurn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"dexMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITokenFuncSigs maps the 4-byte function signature to its string representation.
var ITokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"692058c2": "dex()",
	"117f5a55": "dexBurn(uint256)",
	"bdfde911": "dexMint(uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IToken is an auto generated Go binding around an Ethereum contract.
type IToken struct {
	ITokenCaller     // Read-only binding to the contract
	ITokenTransactor // Write-only binding to the contract
	ITokenFilterer   // Log filterer for contract events
}

// ITokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSession struct {
	Contract     *IToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenCallerSession struct {
	Contract *ITokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransactorSession struct {
	Contract     *ITokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenRaw struct {
	Contract *IToken // Generic contract binding to access the raw methods on
}

// ITokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenCallerRaw struct {
	Contract *ITokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransactorRaw struct {
	Contract *ITokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIToken creates a new instance of IToken, bound to a specific deployed contract.
func NewIToken(address common.Address, backend bind.ContractBackend) (*IToken, error) {
	contract, err := bindIToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IToken{ITokenCaller: ITokenCaller{contract: contract}, ITokenTransactor: ITokenTransactor{contract: contract}, ITokenFilterer: ITokenFilterer{contract: contract}}, nil
}

// NewITokenCaller creates a new read-only instance of IToken, bound to a specific deployed contract.
func NewITokenCaller(address common.Address, caller bind.ContractCaller) (*ITokenCaller, error) {
	contract, err := bindIToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenCaller{contract: contract}, nil
}

// NewITokenTransactor creates a new write-only instance of IToken, bound to a specific deployed contract.
func NewITokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransactor, error) {
	contract, err := bindIToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransactor{contract: contract}, nil
}

// NewITokenFilterer creates a new log filterer instance of IToken, bound to a specific deployed contract.
func NewITokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFilterer, error) {
	contract, err := bindIToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFilterer{contract: contract}, nil
}

// bindIToken binds a generic wrapper to an already deployed contract.
func bindIToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.ITokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IToken *ITokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IToken *ITokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IToken *ITokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IToken *ITokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IToken *ITokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IToken *ITokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IToken *ITokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IToken *ITokenSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IToken *ITokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, value)
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenTransactor) Dex(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dex")
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenSession) Dex() (*types.Transaction, error) {
	return _IToken.Contract.Dex(&_IToken.TransactOpts)
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenTransactorSession) Dex() (*types.Transaction, error) {
	return _IToken.Contract.Dex(&_IToken.TransactOpts)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 _amount) returns()
func (_IToken *ITokenTransactor) DexBurn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dexBurn", _amount)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 _amount) returns()
func (_IToken *ITokenSession) DexBurn(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexBurn(&_IToken.TransactOpts, _amount)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 _amount) returns()
func (_IToken *ITokenTransactorSession) DexBurn(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexBurn(&_IToken.TransactOpts, _amount)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 _amount) returns()
func (_IToken *ITokenTransactor) DexMint(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dexMint", _amount)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 _amount) returns()
func (_IToken *ITokenSession) DexMint(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexMint(&_IToken.TransactOpts, _amount)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 _amount) returns()
func (_IToken *ITokenTransactorSession) DexMint(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexMint(&_IToken.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, from, to, value)
}

// ITokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IToken contract.
type ITokenApprovalIterator struct {
	Event *ITokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ITokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ITokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ITokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenApproval represents a Approval event raised by the IToken contract.
type ITokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ITokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITokenApprovalIterator{contract: _IToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenApproval)
				if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) ParseApproval(log types.Log) (*ITokenApproval, error) {
	event := new(ITokenApproval)
	if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ITokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IToken contract.
type ITokenTransferIterator struct {
	Event *ITokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ITokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ITokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ITokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenTransfer represents a Transfer event raised by the IToken contract.
type ITokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ITokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferIterator{contract: _IToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ITokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenTransfer)
				if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) ParseTransfer(log types.Log) (*ITokenTransfer, error) {
	event := new(ITokenTransfer)
	if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582066e022996ecb52b4dc7ecb9074bceffaa3550425aabecc8970c25026f2db944a64736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OrderbookABI is the input ABI used to generate the binding from.
const OrderbookABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OrderbookFuncSigs maps the 4-byte function signature to its string representation.
var OrderbookFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"4ea09797": "next(bool,bytes32)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"8aa3f897": "top(bool)",
}

// OrderbookBin is the compiled bytecode used for deploying new contracts.
var OrderbookBin = "0x608060405234801561001057600080fd5b50610994806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636e6452cb116100665780636e6452cb1461019c5780638aa3f897146101a4578063aa1c259c146101c3578063ced4aac8146101f1578063f318722b146102315761009e565b806307c399a3146100a35780630d90b10a146100fd57806343271d79146101345780634ea097971461015b57806369c07d3114610180575b600080fd5b6100c8600480360360408110156100b957600080fd5b5080351515906020013561025d565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101226004803603604081101561011357600080fd5b508035151590602001356102a4565b60408051918252519081900360200190f35b6101596004803603604081101561014a57600080fd5b508035151590602001356102c8565b005b6101226004803603604081101561017157600080fd5b5080351515906020013561034e565b610188610372565b604080519115158252519081900360200190f35b610188610377565b610122600480360360208110156101ba57600080fd5b5035151561037c565b610159600480360360408110156101d957600080fd5b506001600160a01b038135811691602001351661039d565b610122600480360360a081101561020757600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610415565b6101226004803603604081101561024757600080fd5b506001600160a01b03813516906020013561047d565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b90151560009081526020818152604080832093835260029093019052206003015490565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610338576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610348828463ffffffff61048916565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b801515600090815260208190526040812061039681610588565b9392505050565b60008080526020526103d67fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff6105a116565b600160009081526020526104117fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff6105a116565b5050565b841515600090815260208190526040812061042e61090b565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261047182828663ffffffff61074e16565b98975050505050505050565b600061039683836107c0565b61049161090b565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b0316815260018201549281018390529281015493830193909352600383015460608301526004909201546080820152901561057357825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b15801561054657600080fd5b505af115801561055a573d6000803e3d6000fd5b505050506040513d602081101561057057600080fd5b50505b610583838363ffffffff61088d16565b505050565b6000808052600282016020526040902060040154919050565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610786848263ffffffff6108ed16565b1561075f575b60030154600081815260028601602052604090209092506107b3848263ffffffff6108ed16565b61078c5750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106108305780518252601f199092019160209182019101610811565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561086f573d6000803e3d6000fd5b5050506040513d602081101561088457600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60408201516001820154600290920154602090930151910291021190565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a72315820105d72feaf4e7b2a5823076ad8fb02917dcedd004db61197c828e4aaee22415a64736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployOrderbook deploys a new Ethereum contract, binding an instance of Orderbook to it.
func DeployOrderbook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Orderbook, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderbookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// Orderbook is an auto generated Go binding around an Ethereum contract.
type Orderbook struct {
	OrderbookCaller     // Read-only binding to the contract
	OrderbookTransactor // Write-only binding to the contract
	OrderbookFilterer   // Log filterer for contract events
}

// OrderbookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderbookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderbookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderbookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderbookSession struct {
	Contract     *Orderbook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderbookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderbookCallerSession struct {
	Contract *OrderbookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderbookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderbookTransactorSession struct {
	Contract     *OrderbookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderbookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderbookRaw struct {
	Contract *Orderbook // Generic contract binding to access the raw methods on
}

// OrderbookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderbookCallerRaw struct {
	Contract *OrderbookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderbookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderbookTransactorRaw struct {
	Contract *OrderbookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderbook creates a new instance of Orderbook, bound to a specific deployed contract.
func NewOrderbook(address common.Address, backend bind.ContractBackend) (*Orderbook, error) {
	contract, err := bindOrderbook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// NewOrderbookCaller creates a new read-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookCaller(address common.Address, caller bind.ContractCaller) (*OrderbookCaller, error) {
	contract, err := bindOrderbook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookCaller{contract: contract}, nil
}

// NewOrderbookTransactor creates a new write-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderbookTransactor, error) {
	contract, err := bindOrderbook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookTransactor{contract: contract}, nil
}

// NewOrderbookFilterer creates a new log filterer instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderbookFilterer, error) {
	contract, err := bindOrderbook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderbookFilterer{contract: contract}, nil
}

// bindOrderbook binds a generic wrapper to an already deployed contract.
func bindOrderbook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.OrderbookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Orderbook *OrderbookCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Orderbook *OrderbookSession) Ask() (bool, error) {
	return _Orderbook.Contract.Ask(&_Orderbook.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Orderbook *OrderbookCallerSession) Ask() (bool, error) {
	return _Orderbook.Contract.Ask(&_Orderbook.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Orderbook *OrderbookCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Orderbook *OrderbookSession) Bid() (bool, error) {
	return _Orderbook.Contract.Bid(&_Orderbook.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Orderbook *OrderbookCallerSession) Bid() (bool, error) {
	return _Orderbook.Contract.Bid(&_Orderbook.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Orderbook *OrderbookSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.CalcOrderID(&_Orderbook.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.CalcOrderID(&_Orderbook.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Orderbook *OrderbookSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.FindAssistingID(&_Orderbook.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.FindAssistingID(&_Orderbook.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Orderbook *OrderbookCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Orderbook.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Orderbook *OrderbookSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Orderbook *OrderbookCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _orderType, _id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Next(&_Orderbook.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Next(&_Orderbook.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Prev(&_Orderbook.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Prev(&_Orderbook.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Orderbook *OrderbookSession) Top(orderType bool) ([32]byte, error) {
	return _Orderbook.Contract.Top(&_Orderbook.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Orderbook.Contract.Top(&_Orderbook.CallOpts, orderType)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.Cancel(&_Orderbook.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.Cancel(&_Orderbook.TransactOpts, orderType, id)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.RegisterTokens(&_Orderbook.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.RegisterTokens(&_Orderbook.TransactOpts, volatileToken, stablizeToken)
}

// PreemptivableABI is the input ABI used to generate the binding from.
const PreemptivableABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"name\":\"Preemptive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PreemptivableFuncSigs maps the 4-byte function signature to its string representation.
var PreemptivableFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"c7f758a8": "getProposal(uint256)",
	"c08cc02d": "getProposalCount()",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"74a8f103": "revoke(address)",
	"c0ee0b8a": "tokenFallback(address,uint256,bytes)",
	"8aa3f897": "top(bool)",
	"4def5645": "totalVote(address)",
	"bd041c4d": "vote(address,bool)",
}

// PreemptivableBin is the compiled bytecode used for deploying new contracts.
var PreemptivableBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562127500600e556002600e54816200003057fe5b04600f55600060105560006011553480156200004b57600080fd5b5060405162003bc538038062003bc5833981810160405260808110156200007157600080fd5b508051602082015160408301516060909301519192909183838015620000975760038190555b60008211620000aa5760028104620000ac565b815b60045550508015620000be57600e8190555b60008211620000db576002600e5481620000d457fe5b04620000dd565b815b600f5550505050613ad180620000f46000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063aa1c259c116100a2578063c0ee0b8a11610071578063c0ee0b8a14610308578063c7f758a81461038d578063ced4aac8146103e7578063ee1a68c614610427578063f318722b1461044a57610116565b8063aa1c259c14610287578063bd041c4d146102b5578063be91d729146102e3578063c08cc02d1461030057610116565b80634ea09797116100e95780634ea09797146101f957806369c07d311461021e5780636e6452cb1461023a57806374a8f103146102425780638aa3f8971461026857610116565b806307c399a31461011b5780630d90b10a1461017557806343271d79146101ac5780634def5645146101d3575b600080fd5b6101406004803603604081101561013157600080fd5b50803515159060200135610476565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61019a6004803603604081101561018b57600080fd5b508035151590602001356104bd565b60408051918252519081900360200190f35b6101d1600480360360408110156101c257600080fd5b508035151590602001356104e5565b005b61019a600480360360208110156101e957600080fd5b50356001600160a01b031661056b565b61019a6004803603604081101561020f57600080fd5b50803515159060200135610593565b6102266105b7565b604080519115158252519081900360200190f35b6102266105bc565b6101d16004803603602081101561025857600080fd5b50356001600160a01b03166105c1565b61019a6004803603602081101561027e57600080fd5b5035151561075c565b6101d16004803603604081101561029d57600080fd5b506001600160a01b0381358116916020013516610776565b6101d1600480360360408110156102cb57600080fd5b506001600160a01b03813516906020013515156108f1565b6101d1600480360360208110156102f957600080fd5b503561096b565b61019a610a09565b6101d16004803603606081101561031e57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561034e57600080fd5b82018360208201111561036057600080fd5b8035906020019184600183028401116401000000008311171561038257600080fd5b509092509050610a1b565b6103aa600480360360208110156103a357600080fd5b5035610b3e565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61019a600480360360a08110156103fd57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610b92565b61042f610bfa565b60408051921515835260208301919091528051918290030190f35b61019a6004803603604081101561046057600080fd5b506001600160a01b038135169060200135610ca4565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610555576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610565828463ffffffff610cb716565b50505050565b60008061057f60128463ffffffff610db116565b905061058a81610df3565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b60006105d460128363ffffffff610db116565b80549091506001600160a01b03838116911614610638576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b601480546001810180835560009290925260068301805490916003027fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec019061068490829084906137c1565b5050600154835460028501546040805163a9059cbb60e01b81526001600160a01b03938416600482015260248101929092525191909216935063a9059cbb925060448083019260209291908290030181600087803b1580156106e557600080fd5b505af11580156106f9573d6000803e3d6000fd5b505050506040513d602081101561070f57600080fd5b50610723905060128363ffffffff610ed916565b506040516001600160a01b038316907f9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c90600090a25050565b801515600090815260208190526040812061058a81610f21565b6001546001600160a01b0316156107d4576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b031615610832576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b031992831617909255600280549284169290911691909117905561086c8282610f3a565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156108b157600080fd5b505afa1580156108c5573d6000803e3d6000fd5b505050506040513d60208110156108db57600080fd5b505190506108ec8180600080610f8e565b505050565b61090260128363ffffffff61102d16565b610946576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b600061095960128463ffffffff610db116565b90506108ec818363ffffffff61104e16565b33156109af576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b6014548110156109e7576109df601482815481106109cc57fe5b9060005260206000209060030201611062565b6001016109b2565b506109f46014600061380d565b6109fc6110cf565b50610a0681611143565b50565b6000610a1560126112de565b90505b90565b608081148015610a3557506001546001600160a01b031633145b15610ad857610a4b60128563ffffffff61102d16565b15610a96576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610aac57600080fd5b50803594506020810135935060400135915060009050610acf88858986866112e2565b50505050610565565b600080806060841415610b0c5784846060811015610af557600080fd5b508035935060208101359250604001359050610b28565b84846040811015610b1c57600080fd5b50803593506020013591505b610b3587848885856114bd565b50505050505050565b6000808080808080610b5760128963ffffffff61153e16565b805460028201546001830154600384015460048501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610bab61382e565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610bee82828663ffffffff61154f16565b98975050505050505050565b6007546000908190610c1157506000905080610ca0565b6001610c9b600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c6a57600080fd5b505afa158015610c7e573d6000803e3d6000fd5b505050506040513d6020811015610c9457600080fd5b50516115c1565b915091505b9091565b6000610cb083836115db565b9392505050565b610cbf61382e565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810183905292810154938301939093526003830154606083015260049092015460808201529015610da157825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015610d7457600080fd5b505af1158015610d88573d6000803e3d6000fd5b505050506040513d6020811015610d9e57600080fd5b50505b6108ec838363ffffffff6116a816565b6001600160a01b03811660009081526001830160205260408120548354849160001901908110610ddd57fe5b9060005260206000209060090201905092915050565b600080805b610e04846006016112de565b811015610ed257600080610e21600687018463ffffffff61170816565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610e7557600080fd5b505afa158015610e89573d6000803e3d6000fd5b505050506040513d6020811015610e9f57600080fd5b50516001600160a01b038416310190508115610ebe5793840193610ec4565b80850394505b505050806001019050610df8565b5092915050565b6001600160a01b038116600090815260018301602052604081205480610f035760009150506104df565b610f1784600019830163ffffffff61174716565b5060019392505050565b6000808052600282016020526040902060040154919050565b6000808052602052610f61600080516020613a57833981519152838363ffffffff6118fd16565b60016000908152602052610f8a600080516020613a37833981519152828463ffffffff6118fd16565b5050565b604080516080810182526003544301808252602082018690529181018690528215156060909101819052600591909155600684905560078590556008805460ff191690911790556000610fe185856115c1565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6001600160a01b031660009081526001919091016020526040902054151590565b6108ec60068301338363ffffffff611aaa16565b60005b81548110156110c357600082600001828154811061107f57fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff19169055919091019050611065565b50610a0681600061385c565b60006110db6009611bc6565b156110e857506000610a18565b6000806110f3611be2565b90925090506001600160a01b03821661111157600092505050610a18565b600061112460128463ffffffff610db116565b905061112f81611c93565b6111398183611e16565b6001935050505090565b3315611187576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6111916005611e59565b1561119e5761119e611e72565b6111a86009611eb6565b156111b5576111b5611ed3565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156111fa57600080fd5b505afa15801561120e573d6000803e3d6000fd5b505050506040513d602081101561122457600080fd5b5051905081156112ac57611236611fd6565b1561124d576112488282600080610f8e565b61126a565b6112578183611fe2565b1561126a5761126a828260016000610f8e565b6112746009611bc6565b156112ac57600061128583836115c1565b905061129081612090565b801561129e575060085460ff165b156112aa575050610a06565b505b6112bd60058263ffffffff6121ae16565b15610f8a5760006112cc6121e3565b90506112d781612298565b5050505050565b5490565b6003601154816112ee57fe5b0460115403831015611337576040805162461bcd60e51b815260206004820152600d60248201526c7374616b6520746f6f206c6f7760981b604482015290519081900360640190fd5b61133f61387a565b82156113b6576003600f548161135157fe5b04600f54018311156113aa576040805162461bcd60e51b815260206004820181905260248201527f736c617368696e67206475726174696f6e20706172616d20746f6f206c6f6e67604482015290519081900360640190fd5b606081018390526113bf565b600f5460608201525b8115611420576003600e54816113d157fe5b04600e54038210156114145760405162461bcd60e51b8152600401808060200182810382526021815260200180613a166021913960400191505060405180910390fd5b60808101829052611429565b600e5460808201525b6001600160a01b038616815260208101859052604081018490524360a082015261145a60128263ffffffff61233f16565b50606080820151608080840151604080518a8152602081018a9052808201949094529383015291516001600160a01b038916927f56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca928290030190a2505050505050565b60006114c761382e565b6114d38787878761248c565b9150915060006114e233612584565b90506114ff826114f133612667565b83919063ffffffff61272116565b61150882612ad2565b156115225761151d818363ffffffff612aeb16565b611534565b6115348184848763ffffffff612b7816565b5050505050505050565b6000826000018281548110610ddd57fe5b600081815260028401602052604081205b6004015460008181526002860160205260409020909250611587848263ffffffff612c4e16565b15611560575b60030154600081815260028601602052604090209092506115b4848263ffffffff612c4e16565b61158d5750909392505050565b60008183116115d557828203600003610cb0565b50900390565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831061164b5780518252601f19909201916020918201910161162c565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561168a573d6000803e3d6000fd5b5050506040513d602081101561169f57600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b6000808061171c858563ffffffff612c6c16565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b81600101600083600001838154811061175c57fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040018120558154600019018114611885578154829060001981019081106117a457fe5b90600052602060002090600902018260000182815481106117c157fe5b60009182526020909120825460099092020180546001600160a01b0319166001600160a01b03909216919091178155600180830154908201556002808301549082015560038083015490820155600480830154908201556005808301549082015560068083018054909183019061183b90829084906137c1565b5050509050508060010182600101600084600001848154811061185a57fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040019020555b81548290600019810190811061189757fe5b60009182526020822060099091020180546001600160a01b031916815560018101829055600281018290556003810182905560048101829055600581018290559060068201816118e7828261385c565b5050835491506108ec90508360001983016138c5565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff19168515151790556001860190915281205480611b2d575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb0565b8454811180611b6b5750836001600160a01b0316856000016001830381548110611b5357fe5b6000918252602090912001546001600160a01b031614155b15611bbb575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb0565b506000949350505050565b6000611bd182612c99565b80156104df57505060040154431090565b6000806000600360105481611bf357fe5b60105491900490039050600080805b611c0c60126112de565b811015611c89576000611c2660128363ffffffff61153e16565b90506004600e5481611c3457fe5b04816005015443031015611c485750611c81565b6000611c5382612ca8565b905085811215611c64575050611c81565b84811315611c7e5781549094506001600160a01b03169250835b50505b600101611c02565b5093509150509091565b611c9f81600601611062565b6040805160a08101825282546001600160a01b039081168083526001850154602084018190526002860154948401859052600386015460608501819052600487015443016080909501859052600980546001600160a01b031916909317909255600a55600b93909355600c92909255600d558154611d26916012911663ffffffff610ed916565b50600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015611d6c57600080fd5b505afa158015611d80573d6000803e3d6000fd5b505050506040513d6020811015611d9657600080fd5b5051600a54909150600090611dac908390612cd6565b9050611dbb8183600180610f8e565b600954600b54600c54600d5460408051938452602084019290925282820152516001600160a01b03909216917f8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d089181900360600190a2505050565b60108054909101600190811c909155600282015460118054909101821c90556004820154600e8054909101821c9055600390910154600f805490910190911c9055565b6000611e6482612cf2565b80156104df57505054431190565b60006005819055600681905560078190556008805460ff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b6000611ec182612c99565b80156104df5750506004015443101590565b611edd6009612c99565b611ee657611fd4565b600b5415611f7757600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015611f4a57600080fd5b505af1158015611f5e573d6000803e3d6000fd5b505050506040513d6020811015611f7457600080fd5b50505b6009546040516001600160a01b03909116907f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d31826057290600090a2600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b6000610a156005611e59565b600082821415611ff4575060006104df565b6006546007541415612008575060016104df565b8282111561205c5760065460075484840391101561203e5760065460075403600281838161203257fe5b041015925050506104df565b60075460065403600282828161205057fe5b041115925050506104df565b60065460075483850391111561207e5760075460065403600281838161203257fe5b60065460075403600282828161205057fe5b60006120a3600960010154600084612cf8565b6120af5750600061058e565b600c54600a5460009190848303816120c357fe5b05816120cb57fe5b049050806120d7575060015b600b548111156120f55750600b546120ed611e72565b6120f5611ed3565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b15801561214b57600080fd5b505af115801561215f573d6000803e3d6000fd5b50506009546040805185815290516001600160a01b0390921693507fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26925081900360200190a250600192915050565b60006121b983612cf2565b80156121c9575082600201548214155b8015610cb05750610cb08360010154838560020154612d28565b6000806121fa6005600201546005600101546115c1565b600754600254604080516318160ddd60e01b8152905193945060009361224a93926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610c6a57600080fd5b905061225860008284612cf8565b61226757600092505050610a18565b6000600454838161227457fe5b05905061228360008284612cf8565b61229157509150610a189050565b9250505090565b60008060008060008086136122ae5760016122b1565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff16156123185761230d816122f287612d53565b600954859291906001600160a01b031663ffffffff612d6916565b93509350505061233a565b6123338161232587612d53565b84919063ffffffff61311c16565b9350935050505b915091565b80516001600160a01b0316600090815260018301602052604081205480156123ae576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b835460018082018087556000878152602090819020875160099095020180546001600160a01b0319166001600160a01b03909516949094178455868101519284019290925560408601516002840155606086015160038401556080860151600484015560a0860151600584015560c086015180518051929488949093600685019261243d9284929101906138f1565b505086516001600160a01b03166000908152600189016020526040902084905550508554919250859160001984019150811061247557fe5b906000526020600020906009020191505092915050565b600061249661382e565b6000841180156124a65750600083115b6124e4576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b841080156124fa5750600160801b83105b61253e576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b61254886866115db565b6040805160a0810182526001600160a01b0398909816885260208801959095529386019290925250506000606084018190526080840152929050565b60008080526020819052600080516020613a57833981519152546001600160a01b03838116911614156125ce57506000808052602052600080516020613a5783398151915261058e565b60016000908152602052600080516020613a37833981519152546001600160a01b038381169116141561261a575060016000908152602052600080516020613a3783398151915261058e565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b03838116911614156126c357506000808052602052600080516020613a5783398151915261058e565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b038381169116141561261a575060016000908152602052600080516020613a3783398151915261058e565b600061272c82610f21565b90505b60001981146105655761274183612ad2565b1561274b57610565565b6000818152600283016020526040902061276b848263ffffffff6133e916565b6127755750610565565b806001015484604001511015612986576000816001015482600201548660400151028161279e57fe5b04905081600201548111156127f2576040805162461bcd60e51b8152602060048201526015602482015274199a5b1b18589b19480f881c99591c9bc81dd85b9d605a1b604482015290519081900360640190fd5b8460200151811115612843576040805162461bcd60e51b815260206004820152601560248201527466696c6c61626c65203e206f72646572206861766560581b604482015290519081900360640190fd5b600282018054829003905560408086015160018085018054929092039091558501548354825163a9059cbb60e01b81526001600160a01b03918216600482015260248101859052925191169163a9059cbb9160448083019260209291908290030181600087803b1580156128b657600080fd5b505af11580156128ca573d6000803e3d6000fd5b505050506040513d60208110156128e057600080fd5b505060018601548551604080880151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561294157600080fd5b505af1158015612955573d6000803e3d6000fd5b505050506040513d602081101561296b57600080fd5b50506000604086015260208501805191909103905250610565565b60018082015460408087018051839003905260028401546020808901805192909203909152928801548751825163a9059cbb60e01b81526001600160a01b0391821660048201526024810194909452915191169263a9059cbb92604480820193918290030181600087803b1580156129fd57600080fd5b505af1158015612a11573d6000803e3d6000fd5b505050506040513d6020811015612a2757600080fd5b50506001830154815460028301546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015612a8857600080fd5b505af1158015612a9c573d6000803e3d6000fd5b505050506040513d6020811015612ab257600080fd5b50506004810154612ac9848463ffffffff6116a816565b915061272f9050565b60008160200151600014806104df575050604001511590565b602081015115610f8a57815481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015612b4e57600080fd5b505af1158015612b62573d6000803e3d6000fd5b505050506040513d602081101561056557600080fd5b60008381526002850160205260409020612b9190612c99565b15612bd8576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b6000838152600285810160209081526040808420865181546001600160a01b0319166001600160a01b0390911617815591860151600183015585015191810191909155606084015160038201556080840151600490910155612c3b85848461154f565b90506112d785858363ffffffff61340816565b60408201516001820154600290920154602090930151910291021190565b6000826000018281548110612c7d57fe5b6000918252602090912001546001600160a01b03169392505050565b546001600160a01b0316151590565b600080612cb483610df3565b905060008113612cc857600091505061058e565b61058a818460020154613443565b600080821215612ced5781600003830390506104df565b500190565b54151590565b6000828413158015612d0a5750818313155b80612d205750828412158015612d205750818312155b949350505050565b6000828411158015612d3a5750818311155b80612d205750828410158015612d205750501115919050565b6000808213612d6557816000036104df565b5090565b600080612d7d86868663ffffffff61311c16565b909250905060008086612d91578383612d94565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b158015612dea57600080fd5b505afa158015612dfe573d6000803e3d6000fd5b505050506040513d6020811015612e1457600080fd5b50518211801590612e9c57508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b158015612e6c57600080fd5b505afa158015612e80573d6000803e3d6000fd5b505050506040513d6020811015612e9657600080fd5b50518211155b15613111578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b158015612ef057600080fd5b505af1158015612f04573d6000803e3d6000fd5b505050506040513d6020811015612f1a57600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b158015612f7357600080fd5b505af1158015612f87573d6000803e3d6000fd5b505050506040513d6020811015612f9d57600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612feb57600080fd5b505af1158015612fff573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561305157600080fd5b505af1158015613065573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156130c157600080fd5b505af11580156130d5573d6000803e3d6000fd5b505050506040513d60208110156130eb57600080fd5b508790506130f957806130fb565b815b830192508661310a578161310c565b805b840193505b505094509492505050565b600080600061312a86610f21565b90505b600019811480159061313e57508382105b156133e057600081815260028701602052604081209086613163578160020154613169565b81600101545b905060008761317c578260010154613182565b82600201545b905086613195868463ffffffff61349b16565b1161328457885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156131eb57600080fd5b505af11580156131ff573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561325657600080fd5b505af115801561326a573d6000803e3d6000fd5b50505050600483015461327d8a866134f5565b93506133b4565b6000613296888763ffffffff6135b816565b905082818302816132a357fe5b0491508092506000896132b657826132b8565b835b905060008a6132c757846132c9565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561331857600080fd5b505af115801561332c573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561337e57600080fd5b505af1158015613392573d6000803e3d6000fd5b506133ab92508e9150899050848463ffffffff61360f16565b50600019955050505b6133c4858363ffffffff61349b16565b94506133d6868263ffffffff61349b16565b955050505061312d565b50935093915050565b6040820151600282015460019092015460209093015191029102101590565b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b600082613452575060006104df565b8282028284828161345f57fe5b05141561346d5790506104df565b6134798460008561379a565b1561348b5750600160ff1b90506104df565b506001600160ff1b039392505050565b600082820183811015610cb0576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6134fd61382e565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810192909252918201549281018390526003820154606082015260049091015460808201529015610da15760018301548151604080840151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015610d7457600080fd5b6000828211156115d5576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b6000838152600285016020526040902061362881612c99565b61366b576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b80600101548311156136c4576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b6001810180548490039055600281015482106136f0576136ea858563ffffffff6134f516565b50610565565b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561374a57600080fd5b505af115801561375e573d6000803e3d6000fd5b505050506040513d602081101561377457600080fd5b5050600281018054839003905560018101546112d7576112d7858563ffffffff6116a816565b600082841280156137aa57508183125b80612d2057508284138015612d2057505012919050565b8280548282559060005260206000209081019282156138015760005260206000209182015b828111156138015782548255916001019190600101906137e6565b50612d65929150613946565b5080546000825560030290600052602060002090810190610a06919061396a565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b5080546000825590600052602060002090810190610a06919061398d565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016138c06139a7565b905290565b8154818355818111156108ec576009028160090283600052602060002091820191016108ec91906139ba565b828054828255906000526020600020908101928215613801579160200282015b8281111561380157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613911565b610a1891905b80821115612d655780546001600160a01b031916815560010161394c565b610a1891905b80821115612d65576000613984828261385c565b50600301613970565b610a1891905b80821115612d655760008155600101613993565b6040518060200160405280606081525090565b610a1891905b80821115612d655780546001600160a01b0319168155600060018201819055600282018190556003820181905560048201819055600582018190556006820181613a0a828261385c565b5050506009016139c056fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a72315820df209652742d78c9e1c951a001b9eb70bd437e7e5c5e54b3b19ee4eb29131eb164736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployPreemptivable deploys a new Ethereum contract, binding an instance of Preemptivable to it.
func DeployPreemptivable(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int, initialSlashingDuration *big.Int, initialLockdownExpiration *big.Int) (common.Address, *types.Transaction, *Preemptivable, error) {
	parsed, err := abi.JSON(strings.NewReader(PreemptivableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PreemptivableBin), backend, absorptionDuration, absorptionExpiration, initialSlashingDuration, initialLockdownExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Preemptivable{PreemptivableCaller: PreemptivableCaller{contract: contract}, PreemptivableTransactor: PreemptivableTransactor{contract: contract}, PreemptivableFilterer: PreemptivableFilterer{contract: contract}}, nil
}

// Preemptivable is an auto generated Go binding around an Ethereum contract.
type Preemptivable struct {
	PreemptivableCaller     // Read-only binding to the contract
	PreemptivableTransactor // Write-only binding to the contract
	PreemptivableFilterer   // Log filterer for contract events
}

// PreemptivableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreemptivableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreemptivableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreemptivableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreemptivableSession struct {
	Contract     *Preemptivable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreemptivableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreemptivableCallerSession struct {
	Contract *PreemptivableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PreemptivableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreemptivableTransactorSession struct {
	Contract     *PreemptivableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PreemptivableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreemptivableRaw struct {
	Contract *Preemptivable // Generic contract binding to access the raw methods on
}

// PreemptivableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreemptivableCallerRaw struct {
	Contract *PreemptivableCaller // Generic read-only contract binding to access the raw methods on
}

// PreemptivableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreemptivableTransactorRaw struct {
	Contract *PreemptivableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreemptivable creates a new instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivable(address common.Address, backend bind.ContractBackend) (*Preemptivable, error) {
	contract, err := bindPreemptivable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Preemptivable{PreemptivableCaller: PreemptivableCaller{contract: contract}, PreemptivableTransactor: PreemptivableTransactor{contract: contract}, PreemptivableFilterer: PreemptivableFilterer{contract: contract}}, nil
}

// NewPreemptivableCaller creates a new read-only instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableCaller(address common.Address, caller bind.ContractCaller) (*PreemptivableCaller, error) {
	contract, err := bindPreemptivable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreemptivableCaller{contract: contract}, nil
}

// NewPreemptivableTransactor creates a new write-only instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableTransactor(address common.Address, transactor bind.ContractTransactor) (*PreemptivableTransactor, error) {
	contract, err := bindPreemptivable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreemptivableTransactor{contract: contract}, nil
}

// NewPreemptivableFilterer creates a new log filterer instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableFilterer(address common.Address, filterer bind.ContractFilterer) (*PreemptivableFilterer, error) {
	contract, err := bindPreemptivable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreemptivableFilterer{contract: contract}, nil
}

// bindPreemptivable binds a generic wrapper to an already deployed contract.
func bindPreemptivable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PreemptivableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preemptivable *PreemptivableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Preemptivable.Contract.PreemptivableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preemptivable *PreemptivableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preemptivable.Contract.PreemptivableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preemptivable *PreemptivableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preemptivable.Contract.PreemptivableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preemptivable *PreemptivableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Preemptivable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preemptivable *PreemptivableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preemptivable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preemptivable *PreemptivableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preemptivable.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Preemptivable *PreemptivableCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Preemptivable *PreemptivableSession) Ask() (bool, error) {
	return _Preemptivable.Contract.Ask(&_Preemptivable.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Preemptivable *PreemptivableCallerSession) Ask() (bool, error) {
	return _Preemptivable.Contract.Ask(&_Preemptivable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Preemptivable *PreemptivableCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Preemptivable *PreemptivableSession) Bid() (bool, error) {
	return _Preemptivable.Contract.Bid(&_Preemptivable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Preemptivable *PreemptivableCallerSession) Bid() (bool, error) {
	return _Preemptivable.Contract.Bid(&_Preemptivable.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.CalcOrderID(&_Preemptivable.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.CalcOrderID(&_Preemptivable.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.FindAssistingID(&_Preemptivable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.FindAssistingID(&_Preemptivable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Preemptivable *PreemptivableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Preemptivable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Preemptivable *PreemptivableSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Preemptivable.Contract.GetOrder(&_Preemptivable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Preemptivable *PreemptivableCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Preemptivable.Contract.GetOrder(&_Preemptivable.CallOpts, _orderType, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableCaller) GetProposal(opts *bind.CallOpts, idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	ret := new(struct {
		Maker              common.Address
		Stake              *big.Int
		Amount             *big.Int
		SlashingDuration   *big.Int
		LockdownExpiration *big.Int
		Number             *big.Int
	})
	out := ret
	err := _Preemptivable.contract.Call(opts, out, "getProposal", idx)
	return *ret, err
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Preemptivable.Contract.GetProposal(&_Preemptivable.CallOpts, idx)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableCallerSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Preemptivable.Contract.GetProposal(&_Preemptivable.CallOpts, idx)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Preemptivable *PreemptivableCaller) GetProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "getProposalCount")
	return *ret0, err
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Preemptivable *PreemptivableSession) GetProposalCount() (*big.Int, error) {
	return _Preemptivable.Contract.GetProposalCount(&_Preemptivable.CallOpts)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Preemptivable *PreemptivableCallerSession) GetProposalCount() (*big.Int, error) {
	return _Preemptivable.Contract.GetProposalCount(&_Preemptivable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Preemptivable *PreemptivableCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Preemptivable.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Preemptivable *PreemptivableSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Preemptivable.Contract.GetRemainToAbsorb(&_Preemptivable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Preemptivable *PreemptivableCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Preemptivable.Contract.GetRemainToAbsorb(&_Preemptivable.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Next(&_Preemptivable.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Next(&_Preemptivable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Prev(&_Preemptivable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Prev(&_Preemptivable.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) Top(orderType bool) ([32]byte, error) {
	return _Preemptivable.Contract.Top(&_Preemptivable.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Preemptivable.Contract.Top(&_Preemptivable.CallOpts, orderType)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Preemptivable *PreemptivableCaller) TotalVote(opts *bind.CallOpts, maker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "totalVote", maker)
	return *ret0, err
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Preemptivable *PreemptivableSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Preemptivable.Contract.TotalVote(&_Preemptivable.CallOpts, maker)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Preemptivable *PreemptivableCallerSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Preemptivable.Contract.TotalVote(&_Preemptivable.CallOpts, maker)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.Cancel(&_Preemptivable.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.Cancel(&_Preemptivable.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.Contract.OnBlockInitialized(&_Preemptivable.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.Contract.OnBlockInitialized(&_Preemptivable.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.RegisterTokens(&_Preemptivable.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.RegisterTokens(&_Preemptivable.TransactOpts, volatileToken, stablizeToken)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableTransactor) Revoke(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "revoke", maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.Revoke(&_Preemptivable.TransactOpts, maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableTransactorSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.Revoke(&_Preemptivable.TransactOpts, maker)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableTransactor) TokenFallback(opts *bind.TransactOpts, maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "tokenFallback", maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.TokenFallback(&_Preemptivable.TransactOpts, maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableTransactorSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.TokenFallback(&_Preemptivable.TransactOpts, maker, value, data)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableTransactor) Vote(opts *bind.TransactOpts, maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "vote", maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.Contract.Vote(&_Preemptivable.TransactOpts, maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableTransactorSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.Contract.Vote(&_Preemptivable.TransactOpts, maker, up)
}

// PreemptivableAbsorptionIterator is returned from FilterAbsorption and is used to iterate over the raw logs and unpacked data for Absorption events raised by the Preemptivable contract.
type PreemptivableAbsorptionIterator struct {
	Event *PreemptivableAbsorption // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableAbsorptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableAbsorption)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableAbsorption)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableAbsorptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableAbsorptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableAbsorption represents a Absorption event raised by the Preemptivable contract.
type PreemptivableAbsorption struct {
	Amount  *big.Int
	Supply  *big.Int
	Emptive bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbsorption is a free log retrieval operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Preemptivable *PreemptivableFilterer) FilterAbsorption(opts *bind.FilterOpts) (*PreemptivableAbsorptionIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return &PreemptivableAbsorptionIterator{contract: _Preemptivable.contract, event: "Absorption", logs: logs, sub: sub}, nil
}

// WatchAbsorption is a free log subscription operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Preemptivable *PreemptivableFilterer) WatchAbsorption(opts *bind.WatchOpts, sink chan<- *PreemptivableAbsorption) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableAbsorption)
				if err := _Preemptivable.contract.UnpackLog(event, "Absorption", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbsorption is a log parse operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Preemptivable *PreemptivableFilterer) ParseAbsorption(log types.Log) (*PreemptivableAbsorption, error) {
	event := new(PreemptivableAbsorption)
	if err := _Preemptivable.contract.UnpackLog(event, "Absorption", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivablePreemptiveIterator is returned from FilterPreemptive and is used to iterate over the raw logs and unpacked data for Preemptive events raised by the Preemptivable contract.
type PreemptivablePreemptiveIterator struct {
	Event *PreemptivablePreemptive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivablePreemptiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivablePreemptive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivablePreemptive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivablePreemptiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivablePreemptiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivablePreemptive represents a Preemptive event raised by the Preemptivable contract.
type PreemptivablePreemptive struct {
	Maker              common.Address
	Stake              *big.Int
	LockdownExpiration *big.Int
	UnlockNumber       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPreemptive is a free log retrieval operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) FilterPreemptive(opts *bind.FilterOpts, maker []common.Address) (*PreemptivablePreemptiveIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivablePreemptiveIterator{contract: _Preemptivable.contract, event: "Preemptive", logs: logs, sub: sub}, nil
}

// WatchPreemptive is a free log subscription operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) WatchPreemptive(opts *bind.WatchOpts, sink chan<- *PreemptivablePreemptive, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivablePreemptive)
				if err := _Preemptivable.contract.UnpackLog(event, "Preemptive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreemptive is a log parse operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) ParsePreemptive(log types.Log) (*PreemptivablePreemptive, error) {
	event := new(PreemptivablePreemptive)
	if err := _Preemptivable.contract.UnpackLog(event, "Preemptive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the Preemptivable contract.
type PreemptivableProposeIterator struct {
	Event *PreemptivablePropose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivablePropose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivablePropose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivablePropose represents a Propose event raised by the Preemptivable contract.
type PreemptivablePropose struct {
	Maker              common.Address
	Amount             *big.Int
	Stake              *big.Int
	LockdownExpiration *big.Int
	SlashingDuration   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Preemptivable *PreemptivableFilterer) FilterPropose(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableProposeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableProposeIterator{contract: _Preemptivable.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Preemptivable *PreemptivableFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *PreemptivablePropose, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivablePropose)
				if err := _Preemptivable.contract.UnpackLog(event, "Propose", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePropose is a log parse operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Preemptivable *PreemptivableFilterer) ParsePropose(log types.Log) (*PreemptivablePropose, error) {
	event := new(PreemptivablePropose)
	if err := _Preemptivable.contract.UnpackLog(event, "Propose", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Preemptivable contract.
type PreemptivableRevokeIterator struct {
	Event *PreemptivableRevoke // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableRevoke)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableRevoke)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableRevoke represents a Revoke event raised by the Preemptivable contract.
type PreemptivableRevoke struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) FilterRevoke(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableRevokeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableRevokeIterator{contract: _Preemptivable.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *PreemptivableRevoke, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableRevoke)
				if err := _Preemptivable.contract.UnpackLog(event, "Revoke", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevoke is a log parse operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) ParseRevoke(log types.Log) (*PreemptivableRevoke, error) {
	event := new(PreemptivableRevoke)
	if err := _Preemptivable.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the Preemptivable contract.
type PreemptivableSlashIterator struct {
	Event *PreemptivableSlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableSlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableSlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableSlash represents a Slash event raised by the Preemptivable contract.
type PreemptivableSlash struct {
	Maker  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) FilterSlash(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableSlashIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableSlashIterator{contract: _Preemptivable.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *PreemptivableSlash, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableSlash)
				if err := _Preemptivable.contract.UnpackLog(event, "Slash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlash is a log parse operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) ParseSlash(log types.Log) (*PreemptivableSlash, error) {
	event := new(PreemptivableSlash)
	if err := _Preemptivable.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the Preemptivable contract.
type PreemptivableStopIterator struct {
	Event *PreemptivableStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableStop represents a Stop event raised by the Preemptivable contract.
type PreemptivableStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Preemptivable *PreemptivableFilterer) FilterStop(opts *bind.FilterOpts) (*PreemptivableStopIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &PreemptivableStopIterator{contract: _Preemptivable.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Preemptivable *PreemptivableFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *PreemptivableStop) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableStop)
				if err := _Preemptivable.contract.UnpackLog(event, "Stop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Preemptivable *PreemptivableFilterer) ParseStop(log types.Log) (*PreemptivableStop, error) {
	event := new(PreemptivableStop)
	if err := _Preemptivable.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the Preemptivable contract.
type PreemptivableUnlockIterator struct {
	Event *PreemptivableUnlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableUnlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableUnlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableUnlock represents a Unlock event raised by the Preemptivable contract.
type PreemptivableUnlock struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) FilterUnlock(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableUnlockIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableUnlockIterator{contract: _Preemptivable.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *PreemptivableUnlock, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableUnlock)
				if err := _Preemptivable.contract.UnpackLog(event, "Unlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlock is a log parse operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) ParseUnlock(log types.Log) (*PreemptivableUnlock, error) {
	event := new(PreemptivableUnlock)
	if err := _Preemptivable.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820106b135b11dd6fc9d42dc3a3a7a9d3c68052b8d556f2d6b7a6105d4dd14182ac64736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SeigniorageABI is the input ABI used to generate the binding from.
const SeigniorageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"name\":\"Preemptive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SeigniorageFuncSigs maps the 4-byte function signature to its string representation.
var SeigniorageFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"c7f758a8": "getProposal(uint256)",
	"c08cc02d": "getProposalCount()",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"74a8f103": "revoke(address)",
	"c0ee0b8a": "tokenFallback(address,uint256,bytes)",
	"8aa3f897": "top(bool)",
	"4def5645": "totalVote(address)",
	"bd041c4d": "vote(address,bool)",
}

// SeigniorageBin is the compiled bytecode used for deploying new contracts.
var SeigniorageBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562127500600e556002600e54816200003057fe5b04600f55600060105560006011553480156200004b57600080fd5b5060405162003bcd38038062003bcd833981810160405260808110156200007157600080fd5b508051602082015160408301516060909301519192909183838383838380156200009b5760038190555b60008211620000ae5760028104620000b0565b815b60045550508015620000c257600e8190555b60008211620000df576002600e5481620000d857fe5b04620000e1565b815b600f555050505050505050613ad180620000fc6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063aa1c259c116100a2578063c0ee0b8a11610071578063c0ee0b8a14610308578063c7f758a81461038d578063ced4aac8146103e7578063ee1a68c614610427578063f318722b1461044a57610116565b8063aa1c259c14610287578063bd041c4d146102b5578063be91d729146102e3578063c08cc02d1461030057610116565b80634ea09797116100e95780634ea09797146101f957806369c07d311461021e5780636e6452cb1461023a57806374a8f103146102425780638aa3f8971461026857610116565b806307c399a31461011b5780630d90b10a1461017557806343271d79146101ac5780634def5645146101d3575b600080fd5b6101406004803603604081101561013157600080fd5b50803515159060200135610476565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61019a6004803603604081101561018b57600080fd5b508035151590602001356104bd565b60408051918252519081900360200190f35b6101d1600480360360408110156101c257600080fd5b508035151590602001356104e5565b005b61019a600480360360208110156101e957600080fd5b50356001600160a01b031661056b565b61019a6004803603604081101561020f57600080fd5b50803515159060200135610593565b6102266105b7565b604080519115158252519081900360200190f35b6102266105bc565b6101d16004803603602081101561025857600080fd5b50356001600160a01b03166105c1565b61019a6004803603602081101561027e57600080fd5b5035151561075c565b6101d16004803603604081101561029d57600080fd5b506001600160a01b0381358116916020013516610776565b6101d1600480360360408110156102cb57600080fd5b506001600160a01b03813516906020013515156108f1565b6101d1600480360360208110156102f957600080fd5b503561096b565b61019a610a09565b6101d16004803603606081101561031e57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561034e57600080fd5b82018360208201111561036057600080fd5b8035906020019184600183028401116401000000008311171561038257600080fd5b509092509050610a1b565b6103aa600480360360208110156103a357600080fd5b5035610b3e565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61019a600480360360a08110156103fd57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610b92565b61042f610bfa565b60408051921515835260208301919091528051918290030190f35b61019a6004803603604081101561046057600080fd5b506001600160a01b038135169060200135610ca4565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610555576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610565828463ffffffff610cb716565b50505050565b60008061057f60128463ffffffff610db116565b905061058a81610df3565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b60006105d460128363ffffffff610db116565b80549091506001600160a01b03838116911614610638576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b601480546001810180835560009290925260068301805490916003027fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec019061068490829084906137c1565b5050600154835460028501546040805163a9059cbb60e01b81526001600160a01b03938416600482015260248101929092525191909216935063a9059cbb925060448083019260209291908290030181600087803b1580156106e557600080fd5b505af11580156106f9573d6000803e3d6000fd5b505050506040513d602081101561070f57600080fd5b50610723905060128363ffffffff610ed916565b506040516001600160a01b038316907f9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c90600090a25050565b801515600090815260208190526040812061058a81610f21565b6001546001600160a01b0316156107d4576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b031615610832576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b031992831617909255600280549284169290911691909117905561086c8282610f3a565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156108b157600080fd5b505afa1580156108c5573d6000803e3d6000fd5b505050506040513d60208110156108db57600080fd5b505190506108ec8180600080610f8e565b505050565b61090260128363ffffffff61102d16565b610946576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b600061095960128463ffffffff610db116565b90506108ec818363ffffffff61104e16565b33156109af576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b6014548110156109e7576109df601482815481106109cc57fe5b9060005260206000209060030201611062565b6001016109b2565b506109f46014600061380d565b6109fc6110cf565b50610a0681611143565b50565b6000610a1560126112de565b90505b90565b608081148015610a3557506001546001600160a01b031633145b15610ad857610a4b60128563ffffffff61102d16565b15610a96576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610aac57600080fd5b50803594506020810135935060400135915060009050610acf88858986866112e2565b50505050610565565b600080806060841415610b0c5784846060811015610af557600080fd5b508035935060208101359250604001359050610b28565b84846040811015610b1c57600080fd5b50803593506020013591505b610b3587848885856114bd565b50505050505050565b6000808080808080610b5760128963ffffffff61153e16565b805460028201546001830154600384015460048501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610bab61382e565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610bee82828663ffffffff61154f16565b98975050505050505050565b6007546000908190610c1157506000905080610ca0565b6001610c9b600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c6a57600080fd5b505afa158015610c7e573d6000803e3d6000fd5b505050506040513d6020811015610c9457600080fd5b50516115c1565b915091505b9091565b6000610cb083836115db565b9392505050565b610cbf61382e565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810183905292810154938301939093526003830154606083015260049092015460808201529015610da157825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015610d7457600080fd5b505af1158015610d88573d6000803e3d6000fd5b505050506040513d6020811015610d9e57600080fd5b50505b6108ec838363ffffffff6116a816565b6001600160a01b03811660009081526001830160205260408120548354849160001901908110610ddd57fe5b9060005260206000209060090201905092915050565b600080805b610e04846006016112de565b811015610ed257600080610e21600687018463ffffffff61170816565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610e7557600080fd5b505afa158015610e89573d6000803e3d6000fd5b505050506040513d6020811015610e9f57600080fd5b50516001600160a01b038416310190508115610ebe5793840193610ec4565b80850394505b505050806001019050610df8565b5092915050565b6001600160a01b038116600090815260018301602052604081205480610f035760009150506104df565b610f1784600019830163ffffffff61174716565b5060019392505050565b6000808052600282016020526040902060040154919050565b6000808052602052610f61600080516020613a57833981519152838363ffffffff6118fd16565b60016000908152602052610f8a600080516020613a37833981519152828463ffffffff6118fd16565b5050565b604080516080810182526003544301808252602082018690529181018690528215156060909101819052600591909155600684905560078590556008805460ff191690911790556000610fe185856115c1565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6001600160a01b031660009081526001919091016020526040902054151590565b6108ec60068301338363ffffffff611aaa16565b60005b81548110156110c357600082600001828154811061107f57fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff19169055919091019050611065565b50610a0681600061385c565b60006110db6009611bc6565b156110e857506000610a18565b6000806110f3611be2565b90925090506001600160a01b03821661111157600092505050610a18565b600061112460128463ffffffff610db116565b905061112f81611c93565b6111398183611e16565b6001935050505090565b3315611187576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6111916005611e59565b1561119e5761119e611e72565b6111a86009611eb6565b156111b5576111b5611ed3565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156111fa57600080fd5b505afa15801561120e573d6000803e3d6000fd5b505050506040513d602081101561122457600080fd5b5051905081156112ac57611236611fd6565b1561124d576112488282600080610f8e565b61126a565b6112578183611fe2565b1561126a5761126a828260016000610f8e565b6112746009611bc6565b156112ac57600061128583836115c1565b905061129081612090565b801561129e575060085460ff165b156112aa575050610a06565b505b6112bd60058263ffffffff6121ae16565b15610f8a5760006112cc6121e3565b90506112d781612298565b5050505050565b5490565b6003601154816112ee57fe5b0460115403831015611337576040805162461bcd60e51b815260206004820152600d60248201526c7374616b6520746f6f206c6f7760981b604482015290519081900360640190fd5b61133f61387a565b82156113b6576003600f548161135157fe5b04600f54018311156113aa576040805162461bcd60e51b815260206004820181905260248201527f736c617368696e67206475726174696f6e20706172616d20746f6f206c6f6e67604482015290519081900360640190fd5b606081018390526113bf565b600f5460608201525b8115611420576003600e54816113d157fe5b04600e54038210156114145760405162461bcd60e51b8152600401808060200182810382526021815260200180613a166021913960400191505060405180910390fd5b60808101829052611429565b600e5460808201525b6001600160a01b038616815260208101859052604081018490524360a082015261145a60128263ffffffff61233f16565b50606080820151608080840151604080518a8152602081018a9052808201949094529383015291516001600160a01b038916927f56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca928290030190a2505050505050565b60006114c761382e565b6114d38787878761248c565b9150915060006114e233612584565b90506114ff826114f133612667565b83919063ffffffff61272116565b61150882612ad2565b156115225761151d818363ffffffff612aeb16565b611534565b6115348184848763ffffffff612b7816565b5050505050505050565b6000826000018281548110610ddd57fe5b600081815260028401602052604081205b6004015460008181526002860160205260409020909250611587848263ffffffff612c4e16565b15611560575b60030154600081815260028601602052604090209092506115b4848263ffffffff612c4e16565b61158d5750909392505050565b60008183116115d557828203600003610cb0565b50900390565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831061164b5780518252601f19909201916020918201910161162c565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561168a573d6000803e3d6000fd5b5050506040513d602081101561169f57600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b6000808061171c858563ffffffff612c6c16565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b81600101600083600001838154811061175c57fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040018120558154600019018114611885578154829060001981019081106117a457fe5b90600052602060002090600902018260000182815481106117c157fe5b60009182526020909120825460099092020180546001600160a01b0319166001600160a01b03909216919091178155600180830154908201556002808301549082015560038083015490820155600480830154908201556005808301549082015560068083018054909183019061183b90829084906137c1565b5050509050508060010182600101600084600001848154811061185a57fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040019020555b81548290600019810190811061189757fe5b60009182526020822060099091020180546001600160a01b031916815560018101829055600281018290556003810182905560048101829055600581018290559060068201816118e7828261385c565b5050835491506108ec90508360001983016138c5565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff19168515151790556001860190915281205480611b2d575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb0565b8454811180611b6b5750836001600160a01b0316856000016001830381548110611b5357fe5b6000918252602090912001546001600160a01b031614155b15611bbb575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb0565b506000949350505050565b6000611bd182612c99565b80156104df57505060040154431090565b6000806000600360105481611bf357fe5b60105491900490039050600080805b611c0c60126112de565b811015611c89576000611c2660128363ffffffff61153e16565b90506004600e5481611c3457fe5b04816005015443031015611c485750611c81565b6000611c5382612ca8565b905085811215611c64575050611c81565b84811315611c7e5781549094506001600160a01b03169250835b50505b600101611c02565b5093509150509091565b611c9f81600601611062565b6040805160a08101825282546001600160a01b039081168083526001850154602084018190526002860154948401859052600386015460608501819052600487015443016080909501859052600980546001600160a01b031916909317909255600a55600b93909355600c92909255600d558154611d26916012911663ffffffff610ed916565b50600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015611d6c57600080fd5b505afa158015611d80573d6000803e3d6000fd5b505050506040513d6020811015611d9657600080fd5b5051600a54909150600090611dac908390612cd6565b9050611dbb8183600180610f8e565b600954600b54600c54600d5460408051938452602084019290925282820152516001600160a01b03909216917f8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d089181900360600190a2505050565b60108054909101600190811c909155600282015460118054909101821c90556004820154600e8054909101821c9055600390910154600f805490910190911c9055565b6000611e6482612cf2565b80156104df57505054431190565b60006005819055600681905560078190556008805460ff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b6000611ec182612c99565b80156104df5750506004015443101590565b611edd6009612c99565b611ee657611fd4565b600b5415611f7757600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015611f4a57600080fd5b505af1158015611f5e573d6000803e3d6000fd5b505050506040513d6020811015611f7457600080fd5b50505b6009546040516001600160a01b03909116907f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d31826057290600090a2600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b6000610a156005611e59565b600082821415611ff4575060006104df565b6006546007541415612008575060016104df565b8282111561205c5760065460075484840391101561203e5760065460075403600281838161203257fe5b041015925050506104df565b60075460065403600282828161205057fe5b041115925050506104df565b60065460075483850391111561207e5760075460065403600281838161203257fe5b60065460075403600282828161205057fe5b60006120a3600960010154600084612cf8565b6120af5750600061058e565b600c54600a5460009190848303816120c357fe5b05816120cb57fe5b049050806120d7575060015b600b548111156120f55750600b546120ed611e72565b6120f5611ed3565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b15801561214b57600080fd5b505af115801561215f573d6000803e3d6000fd5b50506009546040805185815290516001600160a01b0390921693507fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26925081900360200190a250600192915050565b60006121b983612cf2565b80156121c9575082600201548214155b8015610cb05750610cb08360010154838560020154612d28565b6000806121fa6005600201546005600101546115c1565b600754600254604080516318160ddd60e01b8152905193945060009361224a93926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610c6a57600080fd5b905061225860008284612cf8565b61226757600092505050610a18565b6000600454838161227457fe5b05905061228360008284612cf8565b61229157509150610a189050565b9250505090565b60008060008060008086136122ae5760016122b1565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff16156123185761230d816122f287612d53565b600954859291906001600160a01b031663ffffffff612d6916565b93509350505061233a565b6123338161232587612d53565b84919063ffffffff61311c16565b9350935050505b915091565b80516001600160a01b0316600090815260018301602052604081205480156123ae576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b835460018082018087556000878152602090819020875160099095020180546001600160a01b0319166001600160a01b03909516949094178455868101519284019290925560408601516002840155606086015160038401556080860151600484015560a0860151600584015560c086015180518051929488949093600685019261243d9284929101906138f1565b505086516001600160a01b03166000908152600189016020526040902084905550508554919250859160001984019150811061247557fe5b906000526020600020906009020191505092915050565b600061249661382e565b6000841180156124a65750600083115b6124e4576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b841080156124fa5750600160801b83105b61253e576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b61254886866115db565b6040805160a0810182526001600160a01b0398909816885260208801959095529386019290925250506000606084018190526080840152929050565b60008080526020819052600080516020613a57833981519152546001600160a01b03838116911614156125ce57506000808052602052600080516020613a5783398151915261058e565b60016000908152602052600080516020613a37833981519152546001600160a01b038381169116141561261a575060016000908152602052600080516020613a3783398151915261058e565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b03838116911614156126c357506000808052602052600080516020613a5783398151915261058e565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b038381169116141561261a575060016000908152602052600080516020613a3783398151915261058e565b600061272c82610f21565b90505b60001981146105655761274183612ad2565b1561274b57610565565b6000818152600283016020526040902061276b848263ffffffff6133e916565b6127755750610565565b806001015484604001511015612986576000816001015482600201548660400151028161279e57fe5b04905081600201548111156127f2576040805162461bcd60e51b8152602060048201526015602482015274199a5b1b18589b19480f881c99591c9bc81dd85b9d605a1b604482015290519081900360640190fd5b8460200151811115612843576040805162461bcd60e51b815260206004820152601560248201527466696c6c61626c65203e206f72646572206861766560581b604482015290519081900360640190fd5b600282018054829003905560408086015160018085018054929092039091558501548354825163a9059cbb60e01b81526001600160a01b03918216600482015260248101859052925191169163a9059cbb9160448083019260209291908290030181600087803b1580156128b657600080fd5b505af11580156128ca573d6000803e3d6000fd5b505050506040513d60208110156128e057600080fd5b505060018601548551604080880151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561294157600080fd5b505af1158015612955573d6000803e3d6000fd5b505050506040513d602081101561296b57600080fd5b50506000604086015260208501805191909103905250610565565b60018082015460408087018051839003905260028401546020808901805192909203909152928801548751825163a9059cbb60e01b81526001600160a01b0391821660048201526024810194909452915191169263a9059cbb92604480820193918290030181600087803b1580156129fd57600080fd5b505af1158015612a11573d6000803e3d6000fd5b505050506040513d6020811015612a2757600080fd5b50506001830154815460028301546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015612a8857600080fd5b505af1158015612a9c573d6000803e3d6000fd5b505050506040513d6020811015612ab257600080fd5b50506004810154612ac9848463ffffffff6116a816565b915061272f9050565b60008160200151600014806104df575050604001511590565b602081015115610f8a57815481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015612b4e57600080fd5b505af1158015612b62573d6000803e3d6000fd5b505050506040513d602081101561056557600080fd5b60008381526002850160205260409020612b9190612c99565b15612bd8576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b6000838152600285810160209081526040808420865181546001600160a01b0319166001600160a01b0390911617815591860151600183015585015191810191909155606084015160038201556080840151600490910155612c3b85848461154f565b90506112d785858363ffffffff61340816565b60408201516001820154600290920154602090930151910291021190565b6000826000018281548110612c7d57fe5b6000918252602090912001546001600160a01b03169392505050565b546001600160a01b0316151590565b600080612cb483610df3565b905060008113612cc857600091505061058e565b61058a818460020154613443565b600080821215612ced5781600003830390506104df565b500190565b54151590565b6000828413158015612d0a5750818313155b80612d205750828412158015612d205750818312155b949350505050565b6000828411158015612d3a5750818311155b80612d205750828410158015612d205750501115919050565b6000808213612d6557816000036104df565b5090565b600080612d7d86868663ffffffff61311c16565b909250905060008086612d91578383612d94565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b158015612dea57600080fd5b505afa158015612dfe573d6000803e3d6000fd5b505050506040513d6020811015612e1457600080fd5b50518211801590612e9c57508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b158015612e6c57600080fd5b505afa158015612e80573d6000803e3d6000fd5b505050506040513d6020811015612e9657600080fd5b50518211155b15613111578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b158015612ef057600080fd5b505af1158015612f04573d6000803e3d6000fd5b505050506040513d6020811015612f1a57600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b158015612f7357600080fd5b505af1158015612f87573d6000803e3d6000fd5b505050506040513d6020811015612f9d57600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612feb57600080fd5b505af1158015612fff573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561305157600080fd5b505af1158015613065573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156130c157600080fd5b505af11580156130d5573d6000803e3d6000fd5b505050506040513d60208110156130eb57600080fd5b508790506130f957806130fb565b815b830192508661310a578161310c565b805b840193505b505094509492505050565b600080600061312a86610f21565b90505b600019811480159061313e57508382105b156133e057600081815260028701602052604081209086613163578160020154613169565b81600101545b905060008761317c578260010154613182565b82600201545b905086613195868463ffffffff61349b16565b1161328457885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156131eb57600080fd5b505af11580156131ff573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561325657600080fd5b505af115801561326a573d6000803e3d6000fd5b50505050600483015461327d8a866134f5565b93506133b4565b6000613296888763ffffffff6135b816565b905082818302816132a357fe5b0491508092506000896132b657826132b8565b835b905060008a6132c757846132c9565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561331857600080fd5b505af115801561332c573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561337e57600080fd5b505af1158015613392573d6000803e3d6000fd5b506133ab92508e9150899050848463ffffffff61360f16565b50600019955050505b6133c4858363ffffffff61349b16565b94506133d6868263ffffffff61349b16565b955050505061312d565b50935093915050565b6040820151600282015460019092015460209093015191029102101590565b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b600082613452575060006104df565b8282028284828161345f57fe5b05141561346d5790506104df565b6134798460008561379a565b1561348b5750600160ff1b90506104df565b506001600160ff1b039392505050565b600082820183811015610cb0576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6134fd61382e565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810192909252918201549281018390526003820154606082015260049091015460808201529015610da15760018301548151604080840151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015610d7457600080fd5b6000828211156115d5576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b6000838152600285016020526040902061362881612c99565b61366b576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b80600101548311156136c4576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b6001810180548490039055600281015482106136f0576136ea858563ffffffff6134f516565b50610565565b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561374a57600080fd5b505af115801561375e573d6000803e3d6000fd5b505050506040513d602081101561377457600080fd5b5050600281018054839003905560018101546112d7576112d7858563ffffffff6116a816565b600082841280156137aa57508183125b80612d2057508284138015612d2057505012919050565b8280548282559060005260206000209081019282156138015760005260206000209182015b828111156138015782548255916001019190600101906137e6565b50612d65929150613946565b5080546000825560030290600052602060002090810190610a06919061396a565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b5080546000825590600052602060002090810190610a06919061398d565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016138c06139a7565b905290565b8154818355818111156108ec576009028160090283600052602060002091820191016108ec91906139ba565b828054828255906000526020600020908101928215613801579160200282015b8281111561380157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613911565b610a1891905b80821115612d655780546001600160a01b031916815560010161394c565b610a1891905b80821115612d65576000613984828261385c565b50600301613970565b610a1891905b80821115612d655760008155600101613993565b6040518060200160405280606081525090565b610a1891905b80821115612d655780546001600160a01b0319168155600060018201819055600282018190556003820181905560048201819055600582018190556006820181613a0a828261385c565b5050506009016139c056fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a72315820e4b786db9e16ff769c9394309c26b3832b5f80c18bf89ed3ff4e3d084d2d151064736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeploySeigniorage deploys a new Ethereum contract, binding an instance of Seigniorage to it.
func DeploySeigniorage(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int, initialSlashingDuration *big.Int, initialLockdownExpiration *big.Int) (common.Address, *types.Transaction, *Seigniorage, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigniorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SeigniorageBin), backend, absorptionDuration, absorptionExpiration, initialSlashingDuration, initialLockdownExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Seigniorage{SeigniorageCaller: SeigniorageCaller{contract: contract}, SeigniorageTransactor: SeigniorageTransactor{contract: contract}, SeigniorageFilterer: SeigniorageFilterer{contract: contract}}, nil
}

// Seigniorage is an auto generated Go binding around an Ethereum contract.
type Seigniorage struct {
	SeigniorageCaller     // Read-only binding to the contract
	SeigniorageTransactor // Write-only binding to the contract
	SeigniorageFilterer   // Log filterer for contract events
}

// SeigniorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeigniorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeigniorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeigniorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeigniorageSession struct {
	Contract     *Seigniorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeigniorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeigniorageCallerSession struct {
	Contract *SeigniorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SeigniorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeigniorageTransactorSession struct {
	Contract     *SeigniorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SeigniorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeigniorageRaw struct {
	Contract *Seigniorage // Generic contract binding to access the raw methods on
}

// SeigniorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeigniorageCallerRaw struct {
	Contract *SeigniorageCaller // Generic read-only contract binding to access the raw methods on
}

// SeigniorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeigniorageTransactorRaw struct {
	Contract *SeigniorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeigniorage creates a new instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorage(address common.Address, backend bind.ContractBackend) (*Seigniorage, error) {
	contract, err := bindSeigniorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Seigniorage{SeigniorageCaller: SeigniorageCaller{contract: contract}, SeigniorageTransactor: SeigniorageTransactor{contract: contract}, SeigniorageFilterer: SeigniorageFilterer{contract: contract}}, nil
}

// NewSeigniorageCaller creates a new read-only instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageCaller(address common.Address, caller bind.ContractCaller) (*SeigniorageCaller, error) {
	contract, err := bindSeigniorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeigniorageCaller{contract: contract}, nil
}

// NewSeigniorageTransactor creates a new write-only instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SeigniorageTransactor, error) {
	contract, err := bindSeigniorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeigniorageTransactor{contract: contract}, nil
}

// NewSeigniorageFilterer creates a new log filterer instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SeigniorageFilterer, error) {
	contract, err := bindSeigniorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeigniorageFilterer{contract: contract}, nil
}

// bindSeigniorage binds a generic wrapper to an already deployed contract.
func bindSeigniorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigniorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seigniorage *SeigniorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Seigniorage.Contract.SeigniorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seigniorage *SeigniorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seigniorage.Contract.SeigniorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seigniorage *SeigniorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seigniorage.Contract.SeigniorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seigniorage *SeigniorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Seigniorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seigniorage *SeigniorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seigniorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seigniorage *SeigniorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seigniorage.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Seigniorage *SeigniorageCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Seigniorage *SeigniorageSession) Ask() (bool, error) {
	return _Seigniorage.Contract.Ask(&_Seigniorage.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Seigniorage *SeigniorageCallerSession) Ask() (bool, error) {
	return _Seigniorage.Contract.Ask(&_Seigniorage.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Seigniorage *SeigniorageCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Seigniorage *SeigniorageSession) Bid() (bool, error) {
	return _Seigniorage.Contract.Bid(&_Seigniorage.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Seigniorage *SeigniorageCallerSession) Bid() (bool, error) {
	return _Seigniorage.Contract.Bid(&_Seigniorage.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.CalcOrderID(&_Seigniorage.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.CalcOrderID(&_Seigniorage.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.FindAssistingID(&_Seigniorage.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.FindAssistingID(&_Seigniorage.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Seigniorage *SeigniorageCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Seigniorage.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Seigniorage *SeigniorageSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Seigniorage.Contract.GetOrder(&_Seigniorage.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Seigniorage *SeigniorageCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Seigniorage.Contract.GetOrder(&_Seigniorage.CallOpts, _orderType, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageCaller) GetProposal(opts *bind.CallOpts, idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	ret := new(struct {
		Maker              common.Address
		Stake              *big.Int
		Amount             *big.Int
		SlashingDuration   *big.Int
		LockdownExpiration *big.Int
		Number             *big.Int
	})
	out := ret
	err := _Seigniorage.contract.Call(opts, out, "getProposal", idx)
	return *ret, err
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Seigniorage.Contract.GetProposal(&_Seigniorage.CallOpts, idx)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageCallerSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Seigniorage.Contract.GetProposal(&_Seigniorage.CallOpts, idx)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Seigniorage *SeigniorageCaller) GetProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "getProposalCount")
	return *ret0, err
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Seigniorage *SeigniorageSession) GetProposalCount() (*big.Int, error) {
	return _Seigniorage.Contract.GetProposalCount(&_Seigniorage.CallOpts)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Seigniorage *SeigniorageCallerSession) GetProposalCount() (*big.Int, error) {
	return _Seigniorage.Contract.GetProposalCount(&_Seigniorage.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Seigniorage *SeigniorageCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Seigniorage.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Seigniorage *SeigniorageSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Seigniorage.Contract.GetRemainToAbsorb(&_Seigniorage.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Seigniorage *SeigniorageCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Seigniorage.Contract.GetRemainToAbsorb(&_Seigniorage.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Next(&_Seigniorage.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Next(&_Seigniorage.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Prev(&_Seigniorage.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Prev(&_Seigniorage.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) Top(orderType bool) ([32]byte, error) {
	return _Seigniorage.Contract.Top(&_Seigniorage.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Seigniorage.Contract.Top(&_Seigniorage.CallOpts, orderType)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Seigniorage *SeigniorageCaller) TotalVote(opts *bind.CallOpts, maker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "totalVote", maker)
	return *ret0, err
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Seigniorage *SeigniorageSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Seigniorage.Contract.TotalVote(&_Seigniorage.CallOpts, maker)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Seigniorage *SeigniorageCallerSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Seigniorage.Contract.TotalVote(&_Seigniorage.CallOpts, maker)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.Cancel(&_Seigniorage.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.Cancel(&_Seigniorage.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.Contract.OnBlockInitialized(&_Seigniorage.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.Contract.OnBlockInitialized(&_Seigniorage.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.RegisterTokens(&_Seigniorage.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.RegisterTokens(&_Seigniorage.TransactOpts, volatileToken, stablizeToken)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageTransactor) Revoke(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "revoke", maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.Revoke(&_Seigniorage.TransactOpts, maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageTransactorSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.Revoke(&_Seigniorage.TransactOpts, maker)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageTransactor) TokenFallback(opts *bind.TransactOpts, maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "tokenFallback", maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.TokenFallback(&_Seigniorage.TransactOpts, maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageTransactorSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.TokenFallback(&_Seigniorage.TransactOpts, maker, value, data)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageTransactor) Vote(opts *bind.TransactOpts, maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "vote", maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.Contract.Vote(&_Seigniorage.TransactOpts, maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageTransactorSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.Contract.Vote(&_Seigniorage.TransactOpts, maker, up)
}

// SeigniorageAbsorptionIterator is returned from FilterAbsorption and is used to iterate over the raw logs and unpacked data for Absorption events raised by the Seigniorage contract.
type SeigniorageAbsorptionIterator struct {
	Event *SeigniorageAbsorption // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageAbsorptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageAbsorption)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageAbsorption)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageAbsorptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageAbsorptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageAbsorption represents a Absorption event raised by the Seigniorage contract.
type SeigniorageAbsorption struct {
	Amount  *big.Int
	Supply  *big.Int
	Emptive bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbsorption is a free log retrieval operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Seigniorage *SeigniorageFilterer) FilterAbsorption(opts *bind.FilterOpts) (*SeigniorageAbsorptionIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return &SeigniorageAbsorptionIterator{contract: _Seigniorage.contract, event: "Absorption", logs: logs, sub: sub}, nil
}

// WatchAbsorption is a free log subscription operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Seigniorage *SeigniorageFilterer) WatchAbsorption(opts *bind.WatchOpts, sink chan<- *SeigniorageAbsorption) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageAbsorption)
				if err := _Seigniorage.contract.UnpackLog(event, "Absorption", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbsorption is a log parse operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Seigniorage *SeigniorageFilterer) ParseAbsorption(log types.Log) (*SeigniorageAbsorption, error) {
	event := new(SeigniorageAbsorption)
	if err := _Seigniorage.contract.UnpackLog(event, "Absorption", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeignioragePreemptiveIterator is returned from FilterPreemptive and is used to iterate over the raw logs and unpacked data for Preemptive events raised by the Seigniorage contract.
type SeignioragePreemptiveIterator struct {
	Event *SeignioragePreemptive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeignioragePreemptiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeignioragePreemptive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeignioragePreemptive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeignioragePreemptiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeignioragePreemptiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeignioragePreemptive represents a Preemptive event raised by the Seigniorage contract.
type SeignioragePreemptive struct {
	Maker              common.Address
	Stake              *big.Int
	LockdownExpiration *big.Int
	UnlockNumber       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPreemptive is a free log retrieval operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) FilterPreemptive(opts *bind.FilterOpts, maker []common.Address) (*SeignioragePreemptiveIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeignioragePreemptiveIterator{contract: _Seigniorage.contract, event: "Preemptive", logs: logs, sub: sub}, nil
}

// WatchPreemptive is a free log subscription operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) WatchPreemptive(opts *bind.WatchOpts, sink chan<- *SeignioragePreemptive, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeignioragePreemptive)
				if err := _Seigniorage.contract.UnpackLog(event, "Preemptive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreemptive is a log parse operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) ParsePreemptive(log types.Log) (*SeignioragePreemptive, error) {
	event := new(SeignioragePreemptive)
	if err := _Seigniorage.contract.UnpackLog(event, "Preemptive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the Seigniorage contract.
type SeigniorageProposeIterator struct {
	Event *SeignioragePropose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeignioragePropose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeignioragePropose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeignioragePropose represents a Propose event raised by the Seigniorage contract.
type SeignioragePropose struct {
	Maker              common.Address
	Amount             *big.Int
	Stake              *big.Int
	LockdownExpiration *big.Int
	SlashingDuration   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Seigniorage *SeigniorageFilterer) FilterPropose(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageProposeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageProposeIterator{contract: _Seigniorage.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Seigniorage *SeigniorageFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *SeignioragePropose, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeignioragePropose)
				if err := _Seigniorage.contract.UnpackLog(event, "Propose", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePropose is a log parse operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Seigniorage *SeigniorageFilterer) ParsePropose(log types.Log) (*SeignioragePropose, error) {
	event := new(SeignioragePropose)
	if err := _Seigniorage.contract.UnpackLog(event, "Propose", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Seigniorage contract.
type SeigniorageRevokeIterator struct {
	Event *SeigniorageRevoke // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageRevoke)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageRevoke)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageRevoke represents a Revoke event raised by the Seigniorage contract.
type SeigniorageRevoke struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) FilterRevoke(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageRevokeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageRevokeIterator{contract: _Seigniorage.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *SeigniorageRevoke, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageRevoke)
				if err := _Seigniorage.contract.UnpackLog(event, "Revoke", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevoke is a log parse operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) ParseRevoke(log types.Log) (*SeigniorageRevoke, error) {
	event := new(SeigniorageRevoke)
	if err := _Seigniorage.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the Seigniorage contract.
type SeigniorageSlashIterator struct {
	Event *SeigniorageSlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageSlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageSlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageSlash represents a Slash event raised by the Seigniorage contract.
type SeigniorageSlash struct {
	Maker  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) FilterSlash(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageSlashIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageSlashIterator{contract: _Seigniorage.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *SeigniorageSlash, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageSlash)
				if err := _Seigniorage.contract.UnpackLog(event, "Slash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlash is a log parse operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) ParseSlash(log types.Log) (*SeigniorageSlash, error) {
	event := new(SeigniorageSlash)
	if err := _Seigniorage.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the Seigniorage contract.
type SeigniorageStopIterator struct {
	Event *SeigniorageStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageStop represents a Stop event raised by the Seigniorage contract.
type SeigniorageStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Seigniorage *SeigniorageFilterer) FilterStop(opts *bind.FilterOpts) (*SeigniorageStopIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &SeigniorageStopIterator{contract: _Seigniorage.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Seigniorage *SeigniorageFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *SeigniorageStop) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageStop)
				if err := _Seigniorage.contract.UnpackLog(event, "Stop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Seigniorage *SeigniorageFilterer) ParseStop(log types.Log) (*SeigniorageStop, error) {
	event := new(SeigniorageStop)
	if err := _Seigniorage.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the Seigniorage contract.
type SeigniorageUnlockIterator struct {
	Event *SeigniorageUnlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageUnlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageUnlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageUnlock represents a Unlock event raised by the Seigniorage contract.
type SeigniorageUnlock struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) FilterUnlock(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageUnlockIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageUnlockIterator{contract: _Seigniorage.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *SeigniorageUnlock, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageUnlock)
				if err := _Seigniorage.contract.UnpackLog(event, "Unlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlock is a log parse operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) ParseUnlock(log types.Log) (*SeigniorageUnlock, error) {
	event := new(SeigniorageUnlock)
	if err := _Seigniorage.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbsnABI is the input ABI used to generate the binding from.
const AbsnABI = "[]"

// AbsnBin is the compiled bytecode used for deploying new contracts.
var AbsnBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208437123d003f48edfc97833e0cd31d70d899e7882c5f4cfd90f1b2d190e6d00c64736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployAbsn deploys a new Ethereum contract, binding an instance of Absn to it.
func DeployAbsn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Absn, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbsnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Absn{AbsnCaller: AbsnCaller{contract: contract}, AbsnTransactor: AbsnTransactor{contract: contract}, AbsnFilterer: AbsnFilterer{contract: contract}}, nil
}

// Absn is an auto generated Go binding around an Ethereum contract.
type Absn struct {
	AbsnCaller     // Read-only binding to the contract
	AbsnTransactor // Write-only binding to the contract
	AbsnFilterer   // Log filterer for contract events
}

// AbsnCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsnSession struct {
	Contract     *Absn             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsnCallerSession struct {
	Contract *AbsnCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbsnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsnTransactorSession struct {
	Contract     *AbsnTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsnRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsnRaw struct {
	Contract *Absn // Generic contract binding to access the raw methods on
}

// AbsnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsnCallerRaw struct {
	Contract *AbsnCaller // Generic read-only contract binding to access the raw methods on
}

// AbsnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsnTransactorRaw struct {
	Contract *AbsnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsn creates a new instance of Absn, bound to a specific deployed contract.
func NewAbsn(address common.Address, backend bind.ContractBackend) (*Absn, error) {
	contract, err := bindAbsn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Absn{AbsnCaller: AbsnCaller{contract: contract}, AbsnTransactor: AbsnTransactor{contract: contract}, AbsnFilterer: AbsnFilterer{contract: contract}}, nil
}

// NewAbsnCaller creates a new read-only instance of Absn, bound to a specific deployed contract.
func NewAbsnCaller(address common.Address, caller bind.ContractCaller) (*AbsnCaller, error) {
	contract, err := bindAbsn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsnCaller{contract: contract}, nil
}

// NewAbsnTransactor creates a new write-only instance of Absn, bound to a specific deployed contract.
func NewAbsnTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsnTransactor, error) {
	contract, err := bindAbsn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsnTransactor{contract: contract}, nil
}

// NewAbsnFilterer creates a new log filterer instance of Absn, bound to a specific deployed contract.
func NewAbsnFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsnFilterer, error) {
	contract, err := bindAbsn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsnFilterer{contract: contract}, nil
}

// bindAbsn binds a generic wrapper to an already deployed contract.
func bindAbsn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absn *AbsnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absn.Contract.AbsnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absn *AbsnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absn.Contract.AbsnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absn *AbsnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absn.Contract.AbsnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absn *AbsnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absn *AbsnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absn *AbsnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absn.Contract.contract.Transact(opts, method, params...)
}

// DexABI is the input ABI used to generate the binding from.
const DexABI = "[]"

// DexBin is the compiled bytecode used for deploying new contracts.
var DexBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200ec25d1218ed721ae0695ed10839100f3fd5320afc7a449ea05a95f6e1c3631764736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployDex deploys a new Ethereum contract, binding an instance of Dex to it.
func DeployDex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dex, error) {
	parsed, err := abi.JSON(strings.NewReader(DexABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// Dex is an auto generated Go binding around an Ethereum contract.
type Dex struct {
	DexCaller     // Read-only binding to the contract
	DexTransactor // Write-only binding to the contract
	DexFilterer   // Log filterer for contract events
}

// DexCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexSession struct {
	Contract     *Dex              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexCallerSession struct {
	Contract *DexCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexTransactorSession struct {
	Contract     *DexTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexRaw struct {
	Contract *Dex // Generic contract binding to access the raw methods on
}

// DexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexCallerRaw struct {
	Contract *DexCaller // Generic read-only contract binding to access the raw methods on
}

// DexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexTransactorRaw struct {
	Contract *DexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDex creates a new instance of Dex, bound to a specific deployed contract.
func NewDex(address common.Address, backend bind.ContractBackend) (*Dex, error) {
	contract, err := bindDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// NewDexCaller creates a new read-only instance of Dex, bound to a specific deployed contract.
func NewDexCaller(address common.Address, caller bind.ContractCaller) (*DexCaller, error) {
	contract, err := bindDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexCaller{contract: contract}, nil
}

// NewDexTransactor creates a new write-only instance of Dex, bound to a specific deployed contract.
func NewDexTransactor(address common.Address, transactor bind.ContractTransactor) (*DexTransactor, error) {
	contract, err := bindDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexTransactor{contract: contract}, nil
}

// NewDexFilterer creates a new log filterer instance of Dex, bound to a specific deployed contract.
func NewDexFilterer(address common.Address, filterer bind.ContractFilterer) (*DexFilterer, error) {
	contract, err := bindDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexFilterer{contract: contract}, nil
}

// bindDex binds a generic wrapper to an already deployed contract.
func bindDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.DexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transact(opts, method, params...)
}

// MapABI is the input ABI used to generate the binding from.
const MapABI = "[]"

// MapBin is the compiled bytecode used for deploying new contracts.
var MapBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158206ffb8a90208ce61a9de53239079ca4f697256bd94ba98d12ba290f9e9e5445f564736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployMap deploys a new Ethereum contract, binding an instance of Map to it.
func DeployMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Map, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// Map is an auto generated Go binding around an Ethereum contract.
type Map struct {
	MapCaller     // Read-only binding to the contract
	MapTransactor // Write-only binding to the contract
	MapFilterer   // Log filterer for contract events
}

// MapCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapSession struct {
	Contract     *Map              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapCallerSession struct {
	Contract *MapCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapTransactorSession struct {
	Contract     *MapTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapRaw struct {
	Contract *Map // Generic contract binding to access the raw methods on
}

// MapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapCallerRaw struct {
	Contract *MapCaller // Generic read-only contract binding to access the raw methods on
}

// MapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapTransactorRaw struct {
	Contract *MapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMap creates a new instance of Map, bound to a specific deployed contract.
func NewMap(address common.Address, backend bind.ContractBackend) (*Map, error) {
	contract, err := bindMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// NewMapCaller creates a new read-only instance of Map, bound to a specific deployed contract.
func NewMapCaller(address common.Address, caller bind.ContractCaller) (*MapCaller, error) {
	contract, err := bindMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapCaller{contract: contract}, nil
}

// NewMapTransactor creates a new write-only instance of Map, bound to a specific deployed contract.
func NewMapTransactor(address common.Address, transactor bind.ContractTransactor) (*MapTransactor, error) {
	contract, err := bindMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapTransactor{contract: contract}, nil
}

// NewMapFilterer creates a new log filterer instance of Map, bound to a specific deployed contract.
func NewMapFilterer(address common.Address, filterer bind.ContractFilterer) (*MapFilterer, error) {
	contract, err := bindMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapFilterer{contract: contract}, nil
}

// bindMap binds a generic wrapper to an already deployed contract.
func bindMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.MapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.contract.Transact(opts, method, params...)
}

// UtilABI is the input ABI used to generate the binding from.
const UtilABI = "[]"

// UtilBin is the compiled bytecode used for deploying new contracts.
var UtilBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820749ca204754b91c51046675cd7b8c995cc59b945312a6dd8e83a4da798e45b4d64736f6c637828302e352e31322d646576656c6f702e323031392e392e31392b636f6d6d69742e30343738656231650058"

// DeployUtil deploys a new Ethereum contract, binding an instance of Util to it.
func DeployUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Util, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Util{UtilCaller: UtilCaller{contract: contract}, UtilTransactor: UtilTransactor{contract: contract}, UtilFilterer: UtilFilterer{contract: contract}}, nil
}

// Util is an auto generated Go binding around an Ethereum contract.
type Util struct {
	UtilCaller     // Read-only binding to the contract
	UtilTransactor // Write-only binding to the contract
	UtilFilterer   // Log filterer for contract events
}

// UtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilSession struct {
	Contract     *Util             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilCallerSession struct {
	Contract *UtilCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilTransactorSession struct {
	Contract     *UtilTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilRaw struct {
	Contract *Util // Generic contract binding to access the raw methods on
}

// UtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilCallerRaw struct {
	Contract *UtilCaller // Generic read-only contract binding to access the raw methods on
}

// UtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilTransactorRaw struct {
	Contract *UtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtil creates a new instance of Util, bound to a specific deployed contract.
func NewUtil(address common.Address, backend bind.ContractBackend) (*Util, error) {
	contract, err := bindUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Util{UtilCaller: UtilCaller{contract: contract}, UtilTransactor: UtilTransactor{contract: contract}, UtilFilterer: UtilFilterer{contract: contract}}, nil
}

// NewUtilCaller creates a new read-only instance of Util, bound to a specific deployed contract.
func NewUtilCaller(address common.Address, caller bind.ContractCaller) (*UtilCaller, error) {
	contract, err := bindUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilCaller{contract: contract}, nil
}

// NewUtilTransactor creates a new write-only instance of Util, bound to a specific deployed contract.
func NewUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilTransactor, error) {
	contract, err := bindUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilTransactor{contract: contract}, nil
}

// NewUtilFilterer creates a new log filterer instance of Util, bound to a specific deployed contract.
func NewUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilFilterer, error) {
	contract, err := bindUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilFilterer{contract: contract}, nil
}

// bindUtil binds a generic wrapper to an already deployed contract.
func bindUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Util *UtilRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Util.Contract.UtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Util *UtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Util.Contract.UtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Util *UtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Util.Contract.UtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Util *UtilCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Util.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Util *UtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Util.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Util *UtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Util.Contract.contract.Transact(opts, method, params...)
}
