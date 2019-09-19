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
const AbsorbableABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
var AbsorbableBin = "0x608060405262049d4060035560026003548161001757fe5b0460045534801561002757600080fd5b50604051611db4380380611db48339818101604052604081101561004a57600080fd5b508051602090910151801561005f5760038190555b600082116100705760028104610072565b815b6004555050611d2e806100866000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638aa3f897116100715780638aa3f897146101ba578063aa1c259c146101d9578063be91d72914610207578063ced4aac814610224578063ee1a68c614610264578063f318722b14610287576100b4565b806307c399a3146100b95780630d90b10a1461011357806343271d791461014a5780634ea097971461017157806369c07d31146101965780636e6452cb146101b2575b600080fd5b6100de600480360360408110156100cf57600080fd5b508035151590602001356102b3565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101386004803603604081101561012957600080fd5b508035151590602001356102fa565b60408051918252519081900360200190f35b61016f6004803603604081101561016057600080fd5b50803515159060200135610322565b005b6101386004803603604081101561018757600080fd5b508035151590602001356103a8565b61019e6103cc565b604080519115158252519081900360200190f35b61019e6103d1565b610138600480360360208110156101d057600080fd5b503515156103d6565b61016f600480360360408110156101ef57600080fd5b506001600160a01b03813581169160200135166103f9565b61016f6004803603602081101561021d57600080fd5b5035610574565b610138600480360360a081101561023a57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610711565b61026c610779565b60408051921515835260208301919091528051918290030190f35b6101386004803603604081101561029d57600080fd5b506001600160a01b038135169060200135610823565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610392576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b6103a2828463ffffffff61083616565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b80151560009081526020819052604081206103f0816108f5565b9150505b919050565b6001546001600160a01b031615610457576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b0316156104b5576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556104ef828261090e565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561053457600080fd5b505afa158015610548573d6000803e3d6000fd5b505050506040513d602081101561055e57600080fd5b5051905061056f8180600080610982565b505050565b33156105b8576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6105c26005610a21565b156105cf576105cf610a3a565b6105d96009610a7e565b156105e6576105e6610a9b565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561062b57600080fd5b505afa15801561063f573d6000803e3d6000fd5b505050506040513d602081101561065557600080fd5b5051905081156106dd57610667610ba5565b1561067e576106798282600080610982565b61069b565b6106888183610bb7565b1561069b5761069b828260016000610982565b6106a56009610c65565b156106dd5760006106b68383610c81565b90506106c181610c9b565b80156106cf575060085460ff165b156106db57505061070e565b505b6106ee60058263ffffffff610dbd16565b1561070c5760006106fd610df2565b905061070881610ea7565b5050505b505b50565b841515600090815260208190526040812061072a611ca5565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261076d82828663ffffffff610f4e16565b98975050505050505050565b60075460009081906107905750600090508061081f565b600161081a600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107e957600080fd5b505afa1580156107fd573d6000803e3d6000fd5b505050506040513d602081101561081357600080fd5b5051610c81565b915091505b9091565b600061082f8383610fc0565b9392505050565b6000818152600283016020526040902060010154156108e5578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b1580156108b857600080fd5b505af11580156108cc573d6000803e3d6000fd5b505050506040513d60208110156108e257600080fd5b50505b61070c828263ffffffff61108d16565b6000808052600282016020526040902060040154919050565b60008080526020526109477fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff6110ed16565b6001600090815260205261070c7fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff6110ed16565b604080516080810182526003544301808252602082018690529181018690528215156060909101819052600591909155600684905560078590556008805460ff1916909117905560006109d58585610c81565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6000610a2c8261129a565b801561031c57505054431190565b60006005819055600681905560078190556008805460ff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b6000610a89826112a0565b801561031c5750506004015443101590565b610aa560096112a0565b610aae57610ba3565b600b5415610b3f57600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015610b1257600080fd5b505af1158015610b26573d6000803e3d6000fd5b505050506040513d6020811015610b3c57600080fd5b50505b600954604080516001600160a01b039092168252517f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d3182605729181900360200190a1600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b6000610bb16005610a21565b90505b90565b600082821415610bc95750600061031c565b6006546007541415610bdd5750600161031c565b82821115610c3157600654600754848403911015610c1357600654600754036002818381610c0757fe5b0410159250505061031c565b600754600654036002828281610c2557fe5b0411159250505061031c565b600654600754838503911115610c5357600754600654036002818381610c0757fe5b600654600754036002828281610c2557fe5b6000610c70826112a0565b801561031c57505060040154431090565b6000818311610c955782820360000361082f565b50900390565b6000610cae6009600101546000846112af565b610cba575060006103f4565b600c54600a546000919084830381610cce57fe5b0581610cd657fe5b04905080610ce2575060015b600b54811115610d005750600b54610cf8610a3a565b610d00610a9b565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015610d5657600080fd5b505af1158015610d6a573d6000803e3d6000fd5b5050600954604080516001600160a01b0390921682526020820185905280517fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26945091829003019150a150600192915050565b6000610dc88361129a565b8015610dd8575082600201548214155b801561082f575061082f83600101548385600201546112df565b600080610e09600560020154600560010154610c81565b600754600254604080516318160ddd60e01b81529051939450600093610e5993926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156107e957600080fd5b9050610e67600082846112af565b610e7657600092505050610bb4565b60006004548381610e8357fe5b059050610e92600082846112af565b610ea057509150610bb49050565b9250505090565b6000806000806000808613610ebd576001610ec0565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff1615610f2757610f1c81610f018761130a565b600954859291906001600160a01b031663ffffffff61132016565b935093505050610f49565b610f4281610f348761130a565b84919063ffffffff6116d316565b9350935050505b915091565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610f86848263ffffffff6119a016565b15610f5f575b6003015460008181526002860160205260409020909250610fb3848263ffffffff6119a016565b610f8c5750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106110305780518252601f199092019160209182019101611011565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561106f573d6000803e3d6000fd5b5050506040513d602081101561108457600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b54151590565b546001600160a01b0316151590565b60008284131580156112c15750818313155b806112d757508284121580156112d75750818312155b949350505050565b60008284111580156112f15750818311155b806112d757508284101580156112d75750501115919050565b600080821361131c578160000361031c565b5090565b60008061133486868663ffffffff6116d316565b90925090506000808661134857838361134b565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b1580156113a157600080fd5b505afa1580156113b5573d6000803e3d6000fd5b505050506040513d60208110156113cb57600080fd5b5051821180159061145357508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b15801561142357600080fd5b505afa158015611437573d6000803e3d6000fd5b505050506040513d602081101561144d57600080fd5b50518211155b156116c8578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b1580156114a757600080fd5b505af11580156114bb573d6000803e3d6000fd5b505050506040513d60208110156114d157600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b15801561152a57600080fd5b505af115801561153e573d6000803e3d6000fd5b505050506040513d602081101561155457600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156115a257600080fd5b505af11580156115b6573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561160857600080fd5b505af115801561161c573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561167857600080fd5b505af115801561168c573d6000803e3d6000fd5b505050506040513d60208110156116a257600080fd5b508790506116b057806116b2565b815b83019250866116c157816116c3565b805b840193505b505094509492505050565b60008060006116e1866108f5565b90505b60001981148015906116f557508382105b156119975760008181526002870160205260408120908661171a578160020154611720565b81600101545b9050600087611733578260010154611739565b82600201545b90508661174c868463ffffffff6119be16565b1161183b57885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156117a257600080fd5b505af11580156117b6573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561180d57600080fd5b505af1158015611821573d6000803e3d6000fd5b5050505060048301546118348a86611a18565b935061196b565b600061184d888763ffffffff611a9d16565b9050828183028161185a57fe5b04915080925060008961186d578261186f565b835b905060008a61187e5784611880565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b1580156118cf57600080fd5b505af11580156118e3573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561193557600080fd5b505af1158015611949573d6000803e3d6000fd5b5061196292508e9150899050848463ffffffff611af416565b50600019955050505b61197b858363ffffffff6119be16565b945061198d868263ffffffff6119be16565b95505050506116e4565b50935093915050565b60408201516001820154600290920154602090930151910291021190565b60008282018381101561082f576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000818152600280840160205260409091200154156108e557600182015460008281526002808501602090815260408084208054930154815163a9059cbb60e01b81526001600160a01b03948516600482015260248101919091529051929094169363a9059cbb93604480830194928390030190829087803b1580156108b857600080fd5b600082821115610c95576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b60008381526002850160205260409020611b0d816112a0565b611b50576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b8060010154831115611ba9576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b6001810154611bbe908463ffffffff611a9d16565b60018201556002810154821015611bdf576002810180548390039055611be7565b600060028201555b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015611c4157600080fd5b505af1158015611c55573d6000803e3d6000fd5b505050506040513d6020811015611c6b57600080fd5b50611c77905081611c8c565b1561070857610708858563ffffffff61083616565b600081600101546000148061031c575050600201541590565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a7231582086d9ec963eb7264b7addce9a7e5adef697d0aa0de1b629163079ff6f4d3a16c364736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Absorbable *AbsorbableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Absorbable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Absorbable *AbsorbableSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Absorbable.Contract.GetOrder(&_Absorbable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Absorbable *AbsorbableCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
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
// Solidity: event Slash(address maker, uint256 amount)
func (_Absorbable *AbsorbableFilterer) FilterSlash(opts *bind.FilterOpts) (*AbsorbableSlashIterator, error) {

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Slash")
	if err != nil {
		return nil, err
	}
	return &AbsorbableSlashIterator{contract: _Absorbable.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address maker, uint256 amount)
func (_Absorbable *AbsorbableFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *AbsorbableSlash) (event.Subscription, error) {

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Slash")
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
// Solidity: event Slash(address maker, uint256 amount)
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
// Solidity: event Unlock(address maker)
func (_Absorbable *AbsorbableFilterer) FilterUnlock(opts *bind.FilterOpts) (*AbsorbableUnlockIterator, error) {

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return &AbsorbableUnlockIterator{contract: _Absorbable.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address maker)
func (_Absorbable *AbsorbableFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *AbsorbableUnlock) (event.Subscription, error) {

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Unlock")
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
// Solidity: event Unlock(address maker)
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
var MathBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b2bbfb8914c8057a34fa7a1e0ec1d6ab0eb4014fe34e9ec9050f44de0da584cc64736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
const OrderbookABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
var OrderbookBin = "0x608060405234801561001057600080fd5b50610954806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636e6452cb116100665780636e6452cb1461019c5780638aa3f897146101a4578063aa1c259c146101c3578063ced4aac8146101f1578063f318722b146102315761009e565b806307c399a3146100a35780630d90b10a146100fd57806343271d79146101345780634ea097971461015b57806369c07d3114610180575b600080fd5b6100c8600480360360408110156100b957600080fd5b5080351515906020013561025d565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101226004803603604081101561011357600080fd5b508035151590602001356102a4565b60408051918252519081900360200190f35b6101596004803603604081101561014a57600080fd5b508035151590602001356102c8565b005b6101226004803603604081101561017157600080fd5b5080351515906020013561034e565b610188610372565b604080519115158252519081900360200190f35b610188610377565b610122600480360360208110156101ba57600080fd5b5035151561037c565b610159600480360360408110156101d957600080fd5b506001600160a01b038135811691602001351661039d565b610122600480360360a081101561020757600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610415565b6101226004803603604081101561024757600080fd5b506001600160a01b03813516906020013561047d565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b90151560009081526020818152604080832093835260029093019052206003015490565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610338576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610348828463ffffffff61048916565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b801515600090815260208190526040812061039681610548565b9392505050565b60008080526020526103d67fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff61056116565b600160009081526020526104117fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff61056116565b5050565b841515600090815260208190526040812061042e6108cb565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261047182828663ffffffff61070e16565b98975050505050505050565b60006103968383610780565b600081815260028301602052604090206001015415610538578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b15801561050b57600080fd5b505af115801561051f573d6000803e3d6000fd5b505050506040513d602081101561053557600080fd5b50505b610411828263ffffffff61084d16565b6000808052600282016020526040902060040154919050565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610746848263ffffffff6108ad16565b1561071f575b6003015460008181526002860160205260409020909250610773848263ffffffff6108ad16565b61074c5750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106107f05780518252601f1990920191602091820191016107d1565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561082f573d6000803e3d6000fd5b5050506040513d602081101561084457600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60408201516001820154600290920154602090930151910291021190565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a72315820e3aa3d3b2dc883d9915905a6f6e1bd1ea9b08ee7aa234fdb8a7c308927914b1d64736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Orderbook *OrderbookCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Orderbook.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Orderbook *OrderbookSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Orderbook *OrderbookCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
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
const PreemptivableABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"name\":\"Preemptive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
var PreemptivableBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562127500600e556002600e54816200003057fe5b04600f55600060105560006011553480156200004b57600080fd5b50604051620038ce380380620038ce833981810160405260808110156200007157600080fd5b508051602082015160408301516060909301519192909183838015620000975760038190555b60008211620000aa5760028104620000ac565b815b60045550508015620000be57600e8190555b60008211620000db576002600e5481620000d457fe5b04620000dd565b815b600f55505050506137da80620000f46000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063aa1c259c116100a2578063c0ee0b8a11610071578063c0ee0b8a14610308578063c7f758a81461038d578063ced4aac8146103e7578063ee1a68c614610427578063f318722b1461044a57610116565b8063aa1c259c14610287578063bd041c4d146102b5578063be91d729146102e3578063c08cc02d1461030057610116565b80634ea09797116100e95780634ea09797146101f957806369c07d311461021e5780636e6452cb1461023a57806374a8f103146102425780638aa3f8971461026857610116565b806307c399a31461011b5780630d90b10a1461017557806343271d79146101ac5780634def5645146101d3575b600080fd5b6101406004803603604081101561013157600080fd5b50803515159060200135610476565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61019a6004803603604081101561018b57600080fd5b508035151590602001356104bd565b60408051918252519081900360200190f35b6101d1600480360360408110156101c257600080fd5b508035151590602001356104e5565b005b61019a600480360360208110156101e957600080fd5b50356001600160a01b031661056b565b61019a6004803603604081101561020f57600080fd5b50803515159060200135610593565b6102266105b7565b604080519115158252519081900360200190f35b6102266105bc565b6101d16004803603602081101561025857600080fd5b50356001600160a01b03166105c1565b61019a6004803603602081101561027e57600080fd5b50351515610764565b6101d16004803603604081101561029d57600080fd5b506001600160a01b038135811691602001351661077e565b6101d1600480360360408110156102cb57600080fd5b506001600160a01b03813516906020013515156108f9565b6101d1600480360360208110156102f957600080fd5b5035610973565b61019a610a11565b6101d16004803603606081101561031e57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561034e57600080fd5b82018360208201111561036057600080fd5b8035906020019184600183028401116401000000008311171561038257600080fd5b509092509050610a23565b6103aa600480360360208110156103a357600080fd5b5035610b46565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61019a600480360360a08110156103fd57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610b9a565b61042f610c02565b60408051921515835260208301919091528051918290030190f35b61019a6004803603604081101561046057600080fd5b506001600160a01b038135169060200135610cac565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610555576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610565828463ffffffff610cbf16565b50505050565b60008061057f60128463ffffffff610d8216565b905061058a81610dc4565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b60006105d460128363ffffffff610d8216565b80549091506001600160a01b03838116911614610638576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b601480546001810180835560009290925260068301805490916003027fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec019061068490829084906134ca565b5050600154835460028501546040805163a9059cbb60e01b81526001600160a01b03938416600482015260248101929092525191909216935063a9059cbb925060448083019260209291908290030181600087803b1580156106e557600080fd5b505af11580156106f9573d6000803e3d6000fd5b505050506040513d602081101561070f57600080fd5b50610723905060128363ffffffff610eaa16565b50604080516001600160a01b038416815290517f9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c9181900360200190a15050565b801515600090815260208190526040812061058a81610ef2565b6001546001600160a01b0316156107dc576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b03161561083a576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556108748282610f0b565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156108b957600080fd5b505afa1580156108cd573d6000803e3d6000fd5b505050506040513d60208110156108e357600080fd5b505190506108f48180600080610f5b565b505050565b61090a60128363ffffffff610ffa16565b61094e576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b600061096160128463ffffffff610d8216565b90506108f4818363ffffffff61101b16565b33156109b7576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b6014548110156109ef576109e7601482815481106109d457fe5b906000526020600020906003020161102f565b6001016109ba565b506109fc60146000613516565b610a0461109c565b50610a0e81611110565b50565b6000610a1d60126112ab565b90505b90565b608081148015610a3d57506001546001600160a01b031633145b15610ae057610a5360128563ffffffff610ffa16565b15610a9e576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610ab457600080fd5b50803594506020810135935060400135915060009050610ad788858986866112af565b50505050610565565b600080806060841415610b145784846060811015610afd57600080fd5b508035935060208101359250604001359050610b30565b84846040811015610b2457600080fd5b50803593506020013591505b610b3d878488858561148f565b50505050505050565b6000808080808080610b5f60128963ffffffff6114ea16565b805460028201546001830154600384015460048501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610bb3613537565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610bf682828663ffffffff6114fb16565b98975050505050505050565b6007546000908190610c1957506000905080610ca8565b6001610ca3600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c7257600080fd5b505afa158015610c86573d6000803e3d6000fd5b505050506040513d6020811015610c9c57600080fd5b505161156d565b915091505b9091565b6000610cb88383611587565b9392505050565b600081815260028301602052604090206001015415610d6e578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b158015610d4157600080fd5b505af1158015610d55573d6000803e3d6000fd5b505050506040513d6020811015610d6b57600080fd5b50505b610d7e828263ffffffff61165416565b5050565b6001600160a01b03811660009081526001830160205260408120548354849160001901908110610dae57fe5b9060005260206000209060090201905092915050565b600080805b610dd5846006016112ab565b811015610ea357600080610df2600687018463ffffffff6116b416565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610e4657600080fd5b505afa158015610e5a573d6000803e3d6000fd5b505050506040513d6020811015610e7057600080fd5b50516001600160a01b038416310190508115610e8f5793840193610e95565b80850394505b505050806001019050610dc9565b5092915050565b6001600160a01b038116600090815260018301602052604081205480610ed45760009150506104df565b610ee884600019830163ffffffff6116f316565b5060019392505050565b6000808052600282016020526040902060040154919050565b6000808052602052610f32600080516020613760833981519152838363ffffffff6118a916565b60016000908152602052610d7e600080516020613740833981519152828463ffffffff6118a916565b604080516080810182526003544301808252602082018690529181018690528215156060909101819052600591909155600684905560078590556008805460ff191690911790556000610fae858561156d565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6001600160a01b031660009081526001919091016020526040902054151590565b6108f460068301338363ffffffff611a5616565b60005b815481101561109057600082600001828154811061104c57fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff19169055919091019050611032565b50610a0e816000613565565b60006110a86009611b72565b156110b557506000610a20565b6000806110c0611b8e565b90925090506001600160a01b0382166110de57600092505050610a20565b60006110f160128463ffffffff610d8216565b90506110fc81611c3f565b6111068183611dc8565b6001935050505090565b3315611154576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b61115e6005611e0b565b1561116b5761116b611e24565b6111756009611e68565b1561118257611182611e85565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156111c757600080fd5b505afa1580156111db573d6000803e3d6000fd5b505050506040513d60208110156111f157600080fd5b50519050811561127957611203611f8f565b1561121a576112158282600080610f5b565b611237565b6112248183611f9b565b1561123757611237828260016000610f5b565b6112416009611b72565b15611279576000611252838361156d565b905061125d81612049565b801561126b575060085460ff165b15611277575050610a0e565b505b61128a60058263ffffffff61216b16565b15610d7e5760006112996121a0565b90506112a481612255565b5050505050565b5490565b6003601154816112bb57fe5b0460115403831015611304576040805162461bcd60e51b815260206004820152600d60248201526c7374616b6520746f6f206c6f7760981b604482015290519081900360640190fd5b61130c613583565b8215611383576003600f548161131e57fe5b04600f5401831115611377576040805162461bcd60e51b815260206004820181905260248201527f736c617368696e67206475726174696f6e20706172616d20746f6f206c6f6e67604482015290519081900360640190fd5b6060810183905261138c565b600f5460608201525b81156113ed576003600e548161139e57fe5b04600e54038210156113e15760405162461bcd60e51b815260040180806020018281038252602181526020018061371f6021913960400191505060405180910390fd5b608081018290526113f6565b600e5460808201525b6001600160a01b038616815260208101859052604081018490524360a082015261142760128263ffffffff6122fc16565b50606080820151608080840151604080516001600160a01b038c168152602081018b90528082018a90529485019390935290830152517f56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca9181900360a00190a1505050505050565b600061149a33612449565b905060006114b1828888888863ffffffff61252c16565b90506114c482828563ffffffff6126c416565b506114e0816114d233612702565b84919063ffffffff6127bc16565b5050505050505050565b6000826000018281548110610dae57fe5b600081815260028401602052604081205b6004015460008181526002860160205260409020909250611533848263ffffffff61295216565b1561150c575b6003015460008181526002860160205260409020909250611560848263ffffffff61295216565b6115395750909392505050565b600081831161158157828203600003610cb8565b50900390565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106115f75780518252601f1990920191602091820191016115d8565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611636573d6000803e3d6000fd5b5050506040513d602081101561164b57600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b600080806116c8858563ffffffff61297016565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b81600101600083600001838154811061170857fe5b600091825260208083206009909202909101546001600160a01b0316835282019290925260400181205581546000190181146118315781548290600019810190811061175057fe5b906000526020600020906009020182600001828154811061176d57fe5b60009182526020909120825460099092020180546001600160a01b0319166001600160a01b0390921691909117815560018083015490820155600280830154908201556003808301549082015560048083015490820155600580830154908201556006808301805490918301906117e790829084906134ca565b5050509050508060010182600101600084600001848154811061180657fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040019020555b81548290600019810190811061184357fe5b60009182526020822060099091020180546001600160a01b031916815560018101829055600281018290556003810182905560048101829055600581018290559060068201816118938282613565565b5050835491506108f490508360001983016135ce565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff19168515151790556001860190915281205480611ad9575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb8565b8454811180611b175750836001600160a01b0316856000016001830381548110611aff57fe5b6000918252602090912001546001600160a01b031614155b15611b67575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb8565b506000949350505050565b6000611b7d8261299d565b80156104df57505060040154431090565b6000806000600360105481611b9f57fe5b60105491900490039050600080805b611bb860126112ab565b811015611c35576000611bd260128363ffffffff6114ea16565b90506004600e5481611be057fe5b04816005015443031015611bf45750611c2d565b6000611bff826129ac565b905085811215611c10575050611c2d565b84811315611c2a5781549094506001600160a01b03169250835b50505b600101611bae565b5093509150509091565b611c4b8160060161102f565b6040805160a08101825282546001600160a01b039081168083526001850154602084018190526002860154948401859052600386015460608501819052600487015443016080909501859052600980546001600160a01b031916909317909255600a55600b93909355600c92909255600d558154611cd2916012911663ffffffff610eaa16565b50600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015611d1857600080fd5b505afa158015611d2c573d6000803e3d6000fd5b505050506040513d6020811015611d4257600080fd5b5051600a54909150600090611d589083906129e3565b9050611d678183600180610f5b565b600954600b54600c54600d54604080516001600160a01b0390951685526020850193909352838301919091526060830152517f8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d089181900360800190a1505050565b60108054909101600190811c909155600282015460118054909101821c90556004820154600e8054909101821c9055600390910154600f805490910190911c9055565b6000611e16826129ff565b80156104df57505054431190565b60006005819055600681905560078190556008805460ff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b6000611e738261299d565b80156104df5750506004015443101590565b611e8f600961299d565b611e9857611f8d565b600b5415611f2957600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015611efc57600080fd5b505af1158015611f10573d6000803e3d6000fd5b505050506040513d6020811015611f2657600080fd5b50505b600954604080516001600160a01b039092168252517f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d3182605729181900360200190a1600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b6000610a1d6005611e0b565b600082821415611fad575060006104df565b6006546007541415611fc1575060016104df565b8282111561201557600654600754848403911015611ff757600654600754036002818381611feb57fe5b041015925050506104df565b60075460065403600282828161200957fe5b041115925050506104df565b60065460075483850391111561203757600754600654036002818381611feb57fe5b60065460075403600282828161200957fe5b600061205c600960010154600084612a05565b6120685750600061058e565b600c54600a54600091908483038161207c57fe5b058161208457fe5b04905080612090575060015b600b548111156120ae5750600b546120a6611e24565b6120ae611e85565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b15801561210457600080fd5b505af1158015612118573d6000803e3d6000fd5b5050600954604080516001600160a01b0390921682526020820185905280517fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26945091829003019150a150600192915050565b6000612176836129ff565b8015612186575082600201548214155b8015610cb85750610cb88360010154838560020154612a35565b6000806121b760056002015460056001015461156d565b600754600254604080516318160ddd60e01b8152905193945060009361220793926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610c7257600080fd5b905061221560008284612a05565b61222457600092505050610a20565b6000600454838161223157fe5b05905061224060008284612a05565b61224e57509150610a209050565b9250505090565b600080600080600080861361226b57600161226e565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff16156122d5576122ca816122af87612a60565b600954859291906001600160a01b031663ffffffff612a7616565b9350935050506122f7565b6122f0816122e287612a60565b84919063ffffffff612e2916565b9350935050505b915091565b80516001600160a01b03166000908152600183016020526040812054801561236b576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b835460018082018087556000878152602090819020875160099095020180546001600160a01b0319166001600160a01b03909516949094178455868101519284019290925560408601516002840155606086015160038401556080860151600484015560a0860151600584015560c08601518051805192948894909360068501926123fa9284929101906135fa565b505086516001600160a01b03166000908152600189016020526040902084905550508554919250859160001984019150811061243257fe5b906000526020600020906009020191505092915050565b60008080526020819052600080516020613760833981519152546001600160a01b03838116911614156124935750600080805260205260008051602061376083398151915261058e565b60016000908152602052600080516020613740833981519152546001600160a01b03838116911614156124df57506001600090815260205260008051602061374083398151915261058e565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b6000808311801561253d5750600082115b61257b576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b831080156125915750600160801b82105b6125d5576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b60006125e18686611587565b600081815260028901602052604090209091506125fd8161299d565b15612644576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b506040805160a0810182526001600160a01b03978816815260208082019687528183019586526000606083018181526080840182815286835260029c8d01909352939020915182546001600160a01b0319169916989098178155945160018601559251968401969096555051600382015591516004909201919091555090565b60008281526002840160205260408120816126e686838663ffffffff6130f616565b90506126f986868363ffffffff61316816565b95945050505050565b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b038381169116141561275e5750600080805260205260008051602061376083398151915261058e565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b03838116911614156124df57506001600090815260205260008051602061374083398151915261058e565b60008281526002840160205260408120816127d684610ef2565b90505b60001981146126f9576127eb826131a3565b156127f5576126f9565b600081815260028086016020908152604092839020835160a08101855286546001600160a01b03168152600187015492810192909252918501549281019290925260038401546060830152600484015460808301529061285b908263ffffffff6131bc16565b61286557506126f9565b806002015483600101541015612916576000816002015482600101548560010154028161288e57fe5b04905081600101548111156128dc576040805162461bcd60e51b815260206004820152600f60248201526e66696c6c61626c65203e206861766560881b604482015290519081900360640190fd5b60018401546128f69087908590849063ffffffff6131db16565b600184015461290f90899089908463ffffffff6131db16565b50506126f9565b6002810154600182015461293491899189919063ffffffff6131db16565b6004810154612949868463ffffffff61337316565b91506127d99050565b60408201516001820154600290920154602090930151910291021190565b600082600001828154811061298157fe5b6000918252602090912001546001600160a01b03169392505050565b546001600160a01b0316151590565b6000806129b883610dc4565b9050600081136129cc57600091505061058e565b6129d583610dc4565b836002015402915050919050565b6000808212156129fa5781600003830390506104df565b500190565b54151590565b6000828413158015612a175750818313155b80612a2d5750828412158015612a2d5750818312155b949350505050565b6000828411158015612a475750818311155b80612a2d5750828410158015612a2d5750501115919050565b6000808213612a7257816000036104df565b5090565b600080612a8a86868663ffffffff612e2916565b909250905060008086612a9e578383612aa1565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b158015612af757600080fd5b505afa158015612b0b573d6000803e3d6000fd5b505050506040513d6020811015612b2157600080fd5b50518211801590612ba957508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b158015612b7957600080fd5b505afa158015612b8d573d6000803e3d6000fd5b505050506040513d6020811015612ba357600080fd5b50518211155b15612e1e578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b158015612bfd57600080fd5b505af1158015612c11573d6000803e3d6000fd5b505050506040513d6020811015612c2757600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b158015612c8057600080fd5b505af1158015612c94573d6000803e3d6000fd5b505050506040513d6020811015612caa57600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612cf857600080fd5b505af1158015612d0c573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612d5e57600080fd5b505af1158015612d72573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015612dce57600080fd5b505af1158015612de2573d6000803e3d6000fd5b505050506040513d6020811015612df857600080fd5b50879050612e065780612e08565b815b8301925086612e175781612e19565b805b840193505b505094509492505050565b6000806000612e3786610ef2565b90505b6000198114801590612e4b57508382105b156130ed57600081815260028701602052604081209086612e70578160020154612e76565b81600101545b9050600087612e89578260010154612e8f565b82600201545b905086612ea2868463ffffffff6133f816565b11612f9157885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612ef857600080fd5b505af1158015612f0c573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612f6357600080fd5b505af1158015612f77573d6000803e3d6000fd5b505050506004830154612f8a8a86613373565b93506130c1565b6000612fa3888763ffffffff61345216565b90508281830281612fb057fe5b049150809250600089612fc35782612fc5565b835b905060008a612fd45784612fd6565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561302557600080fd5b505af1158015613039573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561308b57600080fd5b505af115801561309f573d6000803e3d6000fd5b506130b892508e9150899050848463ffffffff6131db16565b50600019955050505b6130d1858363ffffffff6133f816565b94506130e3868263ffffffff6133f816565b9550505050612e3a565b50935093915050565b600081815260028401602052604081205b600401546000818152600286016020526040902090925061312e848263ffffffff6134a916565b15613107575b600301546000818152600286016020526040902090925061315b848263ffffffff6134a916565b6131345750909392505050565b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b60008160010154600014806104df575050600201541590565b6040820151600282015460019092015460209093015191029102101590565b600083815260028501602052604090206131f48161299d565b613237576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b8060010154831115613290576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b60018101546132a5908463ffffffff61345216565b600182015560028101548210156132c65760028101805483900390556132ce565b600060028201555b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561332857600080fd5b505af115801561333c573d6000803e3d6000fd5b505050506040513d602081101561335257600080fd5b5061335e9050816131a3565b156112a4576112a4858563ffffffff610cbf16565b600081815260028084016020526040909120015415610d6e57600182015460008281526002808501602090815260408084208054930154815163a9059cbb60e01b81526001600160a01b03948516600482015260248101919091529051929094169363a9059cbb93604480830194928390030190829087803b158015610d4157600080fd5b600082820183811015610cb8576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600082821115611581576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b60008260020154826001015402826002015484600101540211905092915050565b82805482825590600052602060002090810192821561350a5760005260206000209182015b8281111561350a5782548255916001019190600101906134ef565b50612a7292915061364f565b5080546000825560030290600052602060002090810190610a0e9190613673565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b5080546000825590600052602060002090810190610a0e9190613696565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016135c96136b0565b905290565b8154818355818111156108f4576009028160090283600052602060002091820191016108f491906136c3565b82805482825590600052602060002090810192821561350a579160200282015b8281111561350a57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061361a565b610a2091905b80821115612a725780546001600160a01b0319168155600101613655565b610a2091905b80821115612a7257600061368d8282613565565b50600301613679565b610a2091905b80821115612a72576000815560010161369c565b6040518060200160405280606081525090565b610a2091905b80821115612a725780546001600160a01b03191681556000600182018190556002820181905560038201819055600482018190556005820181905560068201816137138282613565565b5050506009016136c956fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a72315820062f083646abb708f3a60faf9cef70629b2e2549940dcba8ee4f2773320628b764736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Preemptivable *PreemptivableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Preemptivable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Preemptivable *PreemptivableSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Preemptivable.Contract.GetOrder(&_Preemptivable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Preemptivable *PreemptivableCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
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
// Solidity: event Preemptive(address maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) FilterPreemptive(opts *bind.FilterOpts) (*PreemptivablePreemptiveIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Preemptive")
	if err != nil {
		return nil, err
	}
	return &PreemptivablePreemptiveIterator{contract: _Preemptivable.contract, event: "Preemptive", logs: logs, sub: sub}, nil
}

// WatchPreemptive is a free log subscription operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) WatchPreemptive(opts *bind.WatchOpts, sink chan<- *PreemptivablePreemptive) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Preemptive")
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
// Solidity: event Preemptive(address maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
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
// Solidity: event Propose(address maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Preemptivable *PreemptivableFilterer) FilterPropose(opts *bind.FilterOpts) (*PreemptivableProposeIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &PreemptivableProposeIterator{contract: _Preemptivable.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Preemptivable *PreemptivableFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *PreemptivablePropose) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Propose")
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
// Solidity: event Propose(address maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
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
// Solidity: event Revoke(address maker)
func (_Preemptivable *PreemptivableFilterer) FilterRevoke(opts *bind.FilterOpts) (*PreemptivableRevokeIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return &PreemptivableRevokeIterator{contract: _Preemptivable.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address maker)
func (_Preemptivable *PreemptivableFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *PreemptivableRevoke) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Revoke")
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
// Solidity: event Revoke(address maker)
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
// Solidity: event Slash(address maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) FilterSlash(opts *bind.FilterOpts) (*PreemptivableSlashIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Slash")
	if err != nil {
		return nil, err
	}
	return &PreemptivableSlashIterator{contract: _Preemptivable.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *PreemptivableSlash) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Slash")
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
// Solidity: event Slash(address maker, uint256 amount)
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
// Solidity: event Unlock(address maker)
func (_Preemptivable *PreemptivableFilterer) FilterUnlock(opts *bind.FilterOpts) (*PreemptivableUnlockIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return &PreemptivableUnlockIterator{contract: _Preemptivable.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address maker)
func (_Preemptivable *PreemptivableFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *PreemptivableUnlock) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Unlock")
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
// Solidity: event Unlock(address maker)
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
var SafeMathBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205378d85ffc165fefe0f3429dc6265fde1cbc56ca647062588e9de51382169c3764736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
const SeigniorageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"name\":\"Preemptive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
var SeigniorageBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562127500600e556002600e54816200003057fe5b04600f55600060105560006011553480156200004b57600080fd5b50604051620038d6380380620038d6833981810160405260808110156200007157600080fd5b508051602082015160408301516060909301519192909183838383838380156200009b5760038190555b60008211620000ae5760028104620000b0565b815b60045550508015620000c257600e8190555b60008211620000df576002600e5481620000d857fe5b04620000e1565b815b600f5550505050505050506137da80620000fc6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063aa1c259c116100a2578063c0ee0b8a11610071578063c0ee0b8a14610308578063c7f758a81461038d578063ced4aac8146103e7578063ee1a68c614610427578063f318722b1461044a57610116565b8063aa1c259c14610287578063bd041c4d146102b5578063be91d729146102e3578063c08cc02d1461030057610116565b80634ea09797116100e95780634ea09797146101f957806369c07d311461021e5780636e6452cb1461023a57806374a8f103146102425780638aa3f8971461026857610116565b806307c399a31461011b5780630d90b10a1461017557806343271d79146101ac5780634def5645146101d3575b600080fd5b6101406004803603604081101561013157600080fd5b50803515159060200135610476565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61019a6004803603604081101561018b57600080fd5b508035151590602001356104bd565b60408051918252519081900360200190f35b6101d1600480360360408110156101c257600080fd5b508035151590602001356104e5565b005b61019a600480360360208110156101e957600080fd5b50356001600160a01b031661056b565b61019a6004803603604081101561020f57600080fd5b50803515159060200135610593565b6102266105b7565b604080519115158252519081900360200190f35b6102266105bc565b6101d16004803603602081101561025857600080fd5b50356001600160a01b03166105c1565b61019a6004803603602081101561027e57600080fd5b50351515610764565b6101d16004803603604081101561029d57600080fd5b506001600160a01b038135811691602001351661077e565b6101d1600480360360408110156102cb57600080fd5b506001600160a01b03813516906020013515156108f9565b6101d1600480360360208110156102f957600080fd5b5035610973565b61019a610a11565b6101d16004803603606081101561031e57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561034e57600080fd5b82018360208201111561036057600080fd5b8035906020019184600183028401116401000000008311171561038257600080fd5b509092509050610a23565b6103aa600480360360208110156103a357600080fd5b5035610b46565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61019a600480360360a08110156103fd57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610b9a565b61042f610c02565b60408051921515835260208301919091528051918290030190f35b61019a6004803603604081101561046057600080fd5b506001600160a01b038135169060200135610cac565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610555576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610565828463ffffffff610cbf16565b50505050565b60008061057f60128463ffffffff610d8216565b905061058a81610dc4565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b60006105d460128363ffffffff610d8216565b80549091506001600160a01b03838116911614610638576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b601480546001810180835560009290925260068301805490916003027fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec019061068490829084906134ca565b5050600154835460028501546040805163a9059cbb60e01b81526001600160a01b03938416600482015260248101929092525191909216935063a9059cbb925060448083019260209291908290030181600087803b1580156106e557600080fd5b505af11580156106f9573d6000803e3d6000fd5b505050506040513d602081101561070f57600080fd5b50610723905060128363ffffffff610eaa16565b50604080516001600160a01b038416815290517f9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c9181900360200190a15050565b801515600090815260208190526040812061058a81610ef2565b6001546001600160a01b0316156107dc576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b03161561083a576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556108748282610f0b565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156108b957600080fd5b505afa1580156108cd573d6000803e3d6000fd5b505050506040513d60208110156108e357600080fd5b505190506108f48180600080610f5b565b505050565b61090a60128363ffffffff610ffa16565b61094e576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b600061096160128463ffffffff610d8216565b90506108f4818363ffffffff61101b16565b33156109b7576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b6014548110156109ef576109e7601482815481106109d457fe5b906000526020600020906003020161102f565b6001016109ba565b506109fc60146000613516565b610a0461109c565b50610a0e81611110565b50565b6000610a1d60126112ab565b90505b90565b608081148015610a3d57506001546001600160a01b031633145b15610ae057610a5360128563ffffffff610ffa16565b15610a9e576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610ab457600080fd5b50803594506020810135935060400135915060009050610ad788858986866112af565b50505050610565565b600080806060841415610b145784846060811015610afd57600080fd5b508035935060208101359250604001359050610b30565b84846040811015610b2457600080fd5b50803593506020013591505b610b3d878488858561148f565b50505050505050565b6000808080808080610b5f60128963ffffffff6114ea16565b805460028201546001830154600384015460048501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610bb3613537565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610bf682828663ffffffff6114fb16565b98975050505050505050565b6007546000908190610c1957506000905080610ca8565b6001610ca3600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c7257600080fd5b505afa158015610c86573d6000803e3d6000fd5b505050506040513d6020811015610c9c57600080fd5b505161156d565b915091505b9091565b6000610cb88383611587565b9392505050565b600081815260028301602052604090206001015415610d6e578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b158015610d4157600080fd5b505af1158015610d55573d6000803e3d6000fd5b505050506040513d6020811015610d6b57600080fd5b50505b610d7e828263ffffffff61165416565b5050565b6001600160a01b03811660009081526001830160205260408120548354849160001901908110610dae57fe5b9060005260206000209060090201905092915050565b600080805b610dd5846006016112ab565b811015610ea357600080610df2600687018463ffffffff6116b416565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610e4657600080fd5b505afa158015610e5a573d6000803e3d6000fd5b505050506040513d6020811015610e7057600080fd5b50516001600160a01b038416310190508115610e8f5793840193610e95565b80850394505b505050806001019050610dc9565b5092915050565b6001600160a01b038116600090815260018301602052604081205480610ed45760009150506104df565b610ee884600019830163ffffffff6116f316565b5060019392505050565b6000808052600282016020526040902060040154919050565b6000808052602052610f32600080516020613760833981519152838363ffffffff6118a916565b60016000908152602052610d7e600080516020613740833981519152828463ffffffff6118a916565b604080516080810182526003544301808252602082018690529181018690528215156060909101819052600591909155600684905560078590556008805460ff191690911790556000610fae858561156d565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6001600160a01b031660009081526001919091016020526040902054151590565b6108f460068301338363ffffffff611a5616565b60005b815481101561109057600082600001828154811061104c57fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff19169055919091019050611032565b50610a0e816000613565565b60006110a86009611b72565b156110b557506000610a20565b6000806110c0611b8e565b90925090506001600160a01b0382166110de57600092505050610a20565b60006110f160128463ffffffff610d8216565b90506110fc81611c3f565b6111068183611dc8565b6001935050505090565b3315611154576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b61115e6005611e0b565b1561116b5761116b611e24565b6111756009611e68565b1561118257611182611e85565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156111c757600080fd5b505afa1580156111db573d6000803e3d6000fd5b505050506040513d60208110156111f157600080fd5b50519050811561127957611203611f8f565b1561121a576112158282600080610f5b565b611237565b6112248183611f9b565b1561123757611237828260016000610f5b565b6112416009611b72565b15611279576000611252838361156d565b905061125d81612049565b801561126b575060085460ff165b15611277575050610a0e565b505b61128a60058263ffffffff61216b16565b15610d7e5760006112996121a0565b90506112a481612255565b5050505050565b5490565b6003601154816112bb57fe5b0460115403831015611304576040805162461bcd60e51b815260206004820152600d60248201526c7374616b6520746f6f206c6f7760981b604482015290519081900360640190fd5b61130c613583565b8215611383576003600f548161131e57fe5b04600f5401831115611377576040805162461bcd60e51b815260206004820181905260248201527f736c617368696e67206475726174696f6e20706172616d20746f6f206c6f6e67604482015290519081900360640190fd5b6060810183905261138c565b600f5460608201525b81156113ed576003600e548161139e57fe5b04600e54038210156113e15760405162461bcd60e51b815260040180806020018281038252602181526020018061371f6021913960400191505060405180910390fd5b608081018290526113f6565b600e5460808201525b6001600160a01b038616815260208101859052604081018490524360a082015261142760128263ffffffff6122fc16565b50606080820151608080840151604080516001600160a01b038c168152602081018b90528082018a90529485019390935290830152517f56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca9181900360a00190a1505050505050565b600061149a33612449565b905060006114b1828888888863ffffffff61252c16565b90506114c482828563ffffffff6126c416565b506114e0816114d233612702565b84919063ffffffff6127bc16565b5050505050505050565b6000826000018281548110610dae57fe5b600081815260028401602052604081205b6004015460008181526002860160205260409020909250611533848263ffffffff61295216565b1561150c575b6003015460008181526002860160205260409020909250611560848263ffffffff61295216565b6115395750909392505050565b600081831161158157828203600003610cb8565b50900390565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106115f75780518252601f1990920191602091820191016115d8565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611636573d6000803e3d6000fd5b5050506040513d602081101561164b57600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b600080806116c8858563ffffffff61297016565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b81600101600083600001838154811061170857fe5b600091825260208083206009909202909101546001600160a01b0316835282019290925260400181205581546000190181146118315781548290600019810190811061175057fe5b906000526020600020906009020182600001828154811061176d57fe5b60009182526020909120825460099092020180546001600160a01b0319166001600160a01b0390921691909117815560018083015490820155600280830154908201556003808301549082015560048083015490820155600580830154908201556006808301805490918301906117e790829084906134ca565b5050509050508060010182600101600084600001848154811061180657fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040019020555b81548290600019810190811061184357fe5b60009182526020822060099091020180546001600160a01b031916815560018101829055600281018290556003810182905560048101829055600581018290559060068201816118938282613565565b5050835491506108f490508360001983016135ce565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff19168515151790556001860190915281205480611ad9575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb8565b8454811180611b175750836001600160a01b0316856000016001830381548110611aff57fe5b6000918252602090912001546001600160a01b031614155b15611b67575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610cb8565b506000949350505050565b6000611b7d8261299d565b80156104df57505060040154431090565b6000806000600360105481611b9f57fe5b60105491900490039050600080805b611bb860126112ab565b811015611c35576000611bd260128363ffffffff6114ea16565b90506004600e5481611be057fe5b04816005015443031015611bf45750611c2d565b6000611bff826129ac565b905085811215611c10575050611c2d565b84811315611c2a5781549094506001600160a01b03169250835b50505b600101611bae565b5093509150509091565b611c4b8160060161102f565b6040805160a08101825282546001600160a01b039081168083526001850154602084018190526002860154948401859052600386015460608501819052600487015443016080909501859052600980546001600160a01b031916909317909255600a55600b93909355600c92909255600d558154611cd2916012911663ffffffff610eaa16565b50600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015611d1857600080fd5b505afa158015611d2c573d6000803e3d6000fd5b505050506040513d6020811015611d4257600080fd5b5051600a54909150600090611d589083906129e3565b9050611d678183600180610f5b565b600954600b54600c54600d54604080516001600160a01b0390951685526020850193909352838301919091526060830152517f8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d089181900360800190a1505050565b60108054909101600190811c909155600282015460118054909101821c90556004820154600e8054909101821c9055600390910154600f805490910190911c9055565b6000611e16826129ff565b80156104df57505054431190565b60006005819055600681905560078190556008805460ff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b6000611e738261299d565b80156104df5750506004015443101590565b611e8f600961299d565b611e9857611f8d565b600b5415611f2957600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015611efc57600080fd5b505af1158015611f10573d6000803e3d6000fd5b505050506040513d6020811015611f2657600080fd5b50505b600954604080516001600160a01b039092168252517f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d3182605729181900360200190a1600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b6000610a1d6005611e0b565b600082821415611fad575060006104df565b6006546007541415611fc1575060016104df565b8282111561201557600654600754848403911015611ff757600654600754036002818381611feb57fe5b041015925050506104df565b60075460065403600282828161200957fe5b041115925050506104df565b60065460075483850391111561203757600754600654036002818381611feb57fe5b60065460075403600282828161200957fe5b600061205c600960010154600084612a05565b6120685750600061058e565b600c54600a54600091908483038161207c57fe5b058161208457fe5b04905080612090575060015b600b548111156120ae5750600b546120a6611e24565b6120ae611e85565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b15801561210457600080fd5b505af1158015612118573d6000803e3d6000fd5b5050600954604080516001600160a01b0390921682526020820185905280517fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26945091829003019150a150600192915050565b6000612176836129ff565b8015612186575082600201548214155b8015610cb85750610cb88360010154838560020154612a35565b6000806121b760056002015460056001015461156d565b600754600254604080516318160ddd60e01b8152905193945060009361220793926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610c7257600080fd5b905061221560008284612a05565b61222457600092505050610a20565b6000600454838161223157fe5b05905061224060008284612a05565b61224e57509150610a209050565b9250505090565b600080600080600080861361226b57600161226e565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff16156122d5576122ca816122af87612a60565b600954859291906001600160a01b031663ffffffff612a7616565b9350935050506122f7565b6122f0816122e287612a60565b84919063ffffffff612e2916565b9350935050505b915091565b80516001600160a01b03166000908152600183016020526040812054801561236b576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b835460018082018087556000878152602090819020875160099095020180546001600160a01b0319166001600160a01b03909516949094178455868101519284019290925560408601516002840155606086015160038401556080860151600484015560a0860151600584015560c08601518051805192948894909360068501926123fa9284929101906135fa565b505086516001600160a01b03166000908152600189016020526040902084905550508554919250859160001984019150811061243257fe5b906000526020600020906009020191505092915050565b60008080526020819052600080516020613760833981519152546001600160a01b03838116911614156124935750600080805260205260008051602061376083398151915261058e565b60016000908152602052600080516020613740833981519152546001600160a01b03838116911614156124df57506001600090815260205260008051602061374083398151915261058e565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b6000808311801561253d5750600082115b61257b576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b831080156125915750600160801b82105b6125d5576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b60006125e18686611587565b600081815260028901602052604090209091506125fd8161299d565b15612644576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b506040805160a0810182526001600160a01b03978816815260208082019687528183019586526000606083018181526080840182815286835260029c8d01909352939020915182546001600160a01b0319169916989098178155945160018601559251968401969096555051600382015591516004909201919091555090565b60008281526002840160205260408120816126e686838663ffffffff6130f616565b90506126f986868363ffffffff61316816565b95945050505050565b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b038381169116141561275e5750600080805260205260008051602061376083398151915261058e565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b03838116911614156124df57506001600090815260205260008051602061374083398151915261058e565b60008281526002840160205260408120816127d684610ef2565b90505b60001981146126f9576127eb826131a3565b156127f5576126f9565b600081815260028086016020908152604092839020835160a08101855286546001600160a01b03168152600187015492810192909252918501549281019290925260038401546060830152600484015460808301529061285b908263ffffffff6131bc16565b61286557506126f9565b806002015483600101541015612916576000816002015482600101548560010154028161288e57fe5b04905081600101548111156128dc576040805162461bcd60e51b815260206004820152600f60248201526e66696c6c61626c65203e206861766560881b604482015290519081900360640190fd5b60018401546128f69087908590849063ffffffff6131db16565b600184015461290f90899089908463ffffffff6131db16565b50506126f9565b6002810154600182015461293491899189919063ffffffff6131db16565b6004810154612949868463ffffffff61337316565b91506127d99050565b60408201516001820154600290920154602090930151910291021190565b600082600001828154811061298157fe5b6000918252602090912001546001600160a01b03169392505050565b546001600160a01b0316151590565b6000806129b883610dc4565b9050600081136129cc57600091505061058e565b6129d583610dc4565b836002015402915050919050565b6000808212156129fa5781600003830390506104df565b500190565b54151590565b6000828413158015612a175750818313155b80612a2d5750828412158015612a2d5750818312155b949350505050565b6000828411158015612a475750818311155b80612a2d5750828410158015612a2d5750501115919050565b6000808213612a7257816000036104df565b5090565b600080612a8a86868663ffffffff612e2916565b909250905060008086612a9e578383612aa1565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b158015612af757600080fd5b505afa158015612b0b573d6000803e3d6000fd5b505050506040513d6020811015612b2157600080fd5b50518211801590612ba957508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b158015612b7957600080fd5b505afa158015612b8d573d6000803e3d6000fd5b505050506040513d6020811015612ba357600080fd5b50518211155b15612e1e578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b158015612bfd57600080fd5b505af1158015612c11573d6000803e3d6000fd5b505050506040513d6020811015612c2757600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b158015612c8057600080fd5b505af1158015612c94573d6000803e3d6000fd5b505050506040513d6020811015612caa57600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612cf857600080fd5b505af1158015612d0c573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612d5e57600080fd5b505af1158015612d72573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015612dce57600080fd5b505af1158015612de2573d6000803e3d6000fd5b505050506040513d6020811015612df857600080fd5b50879050612e065780612e08565b815b8301925086612e175781612e19565b805b840193505b505094509492505050565b6000806000612e3786610ef2565b90505b6000198114801590612e4b57508382105b156130ed57600081815260028701602052604081209086612e70578160020154612e76565b81600101545b9050600087612e89578260010154612e8f565b82600201545b905086612ea2868463ffffffff6133f816565b11612f9157885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612ef857600080fd5b505af1158015612f0c573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612f6357600080fd5b505af1158015612f77573d6000803e3d6000fd5b505050506004830154612f8a8a86613373565b93506130c1565b6000612fa3888763ffffffff61345216565b90508281830281612fb057fe5b049150809250600089612fc35782612fc5565b835b905060008a612fd45784612fd6565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561302557600080fd5b505af1158015613039573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561308b57600080fd5b505af115801561309f573d6000803e3d6000fd5b506130b892508e9150899050848463ffffffff6131db16565b50600019955050505b6130d1858363ffffffff6133f816565b94506130e3868263ffffffff6133f816565b9550505050612e3a565b50935093915050565b600081815260028401602052604081205b600401546000818152600286016020526040902090925061312e848263ffffffff6134a916565b15613107575b600301546000818152600286016020526040902090925061315b848263ffffffff6134a916565b6131345750909392505050565b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b60008160010154600014806104df575050600201541590565b6040820151600282015460019092015460209093015191029102101590565b600083815260028501602052604090206131f48161299d565b613237576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b8060010154831115613290576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b60018101546132a5908463ffffffff61345216565b600182015560028101548210156132c65760028101805483900390556132ce565b600060028201555b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561332857600080fd5b505af115801561333c573d6000803e3d6000fd5b505050506040513d602081101561335257600080fd5b5061335e9050816131a3565b156112a4576112a4858563ffffffff610cbf16565b600081815260028084016020526040909120015415610d6e57600182015460008281526002808501602090815260408084208054930154815163a9059cbb60e01b81526001600160a01b03948516600482015260248101919091529051929094169363a9059cbb93604480830194928390030190829087803b158015610d4157600080fd5b600082820183811015610cb8576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600082821115611581576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b60008260020154826001015402826002015484600101540211905092915050565b82805482825590600052602060002090810192821561350a5760005260206000209182015b8281111561350a5782548255916001019190600101906134ef565b50612a7292915061364f565b5080546000825560030290600052602060002090810190610a0e9190613673565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b5080546000825590600052602060002090810190610a0e9190613696565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016135c96136b0565b905290565b8154818355818111156108f4576009028160090283600052602060002091820191016108f491906136c3565b82805482825590600052602060002090810192821561350a579160200282015b8281111561350a57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061361a565b610a2091905b80821115612a725780546001600160a01b0319168155600101613655565b610a2091905b80821115612a7257600061368d8282613565565b50600301613679565b610a2091905b80821115612a72576000815560010161369c565b6040518060200160405280606081525090565b610a2091905b80821115612a725780546001600160a01b03191681556000600182018190556002820181905560038201819055600482018190556005820181905560068201816137138282613565565b5050506009016136c956fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a72315820fe56f2738ab6bdc82b576b8aebe2691fadfb19603f72861976d1801714b3cdc564736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Seigniorage *SeigniorageCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Seigniorage.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Seigniorage *SeigniorageSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Seigniorage.Contract.GetOrder(&_Seigniorage.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Seigniorage *SeigniorageCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
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
// Solidity: event Preemptive(address maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) FilterPreemptive(opts *bind.FilterOpts) (*SeignioragePreemptiveIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Preemptive")
	if err != nil {
		return nil, err
	}
	return &SeignioragePreemptiveIterator{contract: _Seigniorage.contract, event: "Preemptive", logs: logs, sub: sub}, nil
}

// WatchPreemptive is a free log subscription operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) WatchPreemptive(opts *bind.WatchOpts, sink chan<- *SeignioragePreemptive) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Preemptive")
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
// Solidity: event Preemptive(address maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
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
// Solidity: event Propose(address maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Seigniorage *SeigniorageFilterer) FilterPropose(opts *bind.FilterOpts) (*SeigniorageProposeIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &SeigniorageProposeIterator{contract: _Seigniorage.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
func (_Seigniorage *SeigniorageFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *SeignioragePropose) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Propose")
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
// Solidity: event Propose(address maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingDuration)
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
// Solidity: event Revoke(address maker)
func (_Seigniorage *SeigniorageFilterer) FilterRevoke(opts *bind.FilterOpts) (*SeigniorageRevokeIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return &SeigniorageRevokeIterator{contract: _Seigniorage.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address maker)
func (_Seigniorage *SeigniorageFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *SeigniorageRevoke) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Revoke")
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
// Solidity: event Revoke(address maker)
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
// Solidity: event Slash(address maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) FilterSlash(opts *bind.FilterOpts) (*SeigniorageSlashIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Slash")
	if err != nil {
		return nil, err
	}
	return &SeigniorageSlashIterator{contract: _Seigniorage.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *SeigniorageSlash) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Slash")
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
// Solidity: event Slash(address maker, uint256 amount)
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
// Solidity: event Unlock(address maker)
func (_Seigniorage *SeigniorageFilterer) FilterUnlock(opts *bind.FilterOpts) (*SeigniorageUnlockIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return &SeigniorageUnlockIterator{contract: _Seigniorage.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address maker)
func (_Seigniorage *SeigniorageFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *SeigniorageUnlock) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Unlock")
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
// Solidity: event Unlock(address maker)
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
var AbsnBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d55f66724a91ea8ff5ee01d4bee15b211f80a6b3decc1de52e44c77db816b6bc64736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
var DexBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820cf636214d56f49da0b01901abfe6f7ea9697992b93e5d7390a8df17c9e6ec18a64736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
var MapBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d42b0dca1f35cadb0d925aa8b4b0136bd092bfd9ccb934cc00400da50da323b264736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
var UtilBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158203a14a12bbe230d50e75680e38a1c1fc0ede823afd6d208022c5d441217fdc89664736f6c637828302e352e31322d646576656c6f702e323031392e392e31372b636f6d6d69742e35613935303930380058"

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
