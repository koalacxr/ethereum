// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unionpay

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

// UnionPayABI is the input ABI used to generate the binding from.
const UnionPayABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_transId\",\"type\":\"uint256\"}],\"name\":\"receiptUsed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_transId\",\"type\":\"uint256\"},{\"name\":\"_feePercentage\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"safePay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platform\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_checker\",\"type\":\"address\"}],\"name\":\"setPlatform\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"plainPay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountIndeed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transId\",\"type\":\"uint256\"}],\"name\":\"UserPay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BareUserPay\",\"type\":\"event\"}]"

// UnionPayBin is the compiled bytecode used for deploying new contracts.
const UnionPayBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a0319918216811790925560018054909116909117905561087b8061004b6000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630c433edf811461010d57806312065fe01461014b57806333c91611146101725780633ccfd60b146101d15780634bde38c8146101e85780636945c5ea146102195780638da5cb5b1461023a578063eb3f24271461024f578063f2fde38b14610257575b7f39a90d638bd2f864aeb4ed151f1668f52e4383137a90ee960743ad694d259bf333346000366040518085600160a060020a0316600160a060020a03168152602001848152602001806020018281038252848482818152602001925080828437604051920182900397509095505050505050a1005b34801561011957600080fd5b50610137600160a060020a0360043581169060243516604435610278565b604080519115158252519081900360200190f35b34801561015757600080fd5b506101606102a9565b60408051918252519081900360200190f35b604080516020601f606435600481810135928301849004840285018401909552818452610137948035946024803595600160a060020a036044351695369560849493019181908401838280828437509497506102b79650505050505050565b3480156101dd57600080fd5b506101e6610550565b005b3480156101f457600080fd5b506101fd6105c0565b60408051600160a060020a039092168252519081900360200190f35b34801561022557600080fd5b506101e6600160a060020a03600435166105cf565b34801561024657600080fd5b506101fd61062e565b61013761063d565b34801561026357600080fd5b506101e6600160a060020a03600435166106b8565b60006002600061028986868661070f565b815260208101919091526040016000205460ff1660011490509392505050565b600160a060020a0330163190565b60008060008086101580156102cd575060648611155b15156102d857600080fd5b600160a060020a03851615156102ed57600080fd5b600260006102fc33888b61070f565b815260208101919091526040016000205460ff161561031a57600080fd5b600154600160a060020a0316151561033157600080fd5b604080516c01000000000000000000000000600160a060020a033381168202835288160260148201523460288201526048810188905260688101899052905190819003608801902061038290610751565b600154909250600160a060020a031661039b838661078f565b600160a060020a0316146103ae57600080fd5b6001600260006103bf33898c61070f565b81526020810191909152604001600020805460ff191660ff9290921691909117905585151561048857600034111561042857604051600160a060020a038616903480156108fc02916000818181858888f19350505050158015610426573d6000803e3d6000fd5b505b60408051600160a060020a033381168252871660208201523481830181905260608201526080810189905290517f72d0b2856e77baaa53d185402480e6844c2d105276b126fdd46c34d14f4ac5f79181900360a00190a160019250610546565b348602905034868281151561049957fe5b04146104a157fe5b60649004348110156104e757604051600160a060020a038616903483900380156108fc02916000818181858888f193505050501580156104e5573d6000803e3d6000fd5b505b60408051600160a060020a033381168252871660208201523481830181905283900360608201526080810189905290517f72d0b2856e77baaa53d185402480e6844c2d105276b126fdd46c34d14f4ac5f79181900360a00190a1600192505b5050949350505050565b60005433600160a060020a0390811691161461056b57600080fd5b600154600160a060020a0316151561058257600080fd5b600154604051600160a060020a039182169130163180156108fc02916000818181858888f193505050501580156105bd573d6000803e3d6000fd5b50565b600154600160a060020a031681565b60005433600160a060020a039081169116146105ea57600080fd5b600160a060020a03811615156105ff57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b60007f39a90d638bd2f864aeb4ed151f1668f52e4383137a90ee960743ad694d259bf333346000366040518085600160a060020a0316600160a060020a03168152602001848152602001806020018281038252848482818152602001925080828437604051920182900397509095505050505050a150600190565b60005433600160a060020a039081169116146106d357600080fd5b600160a060020a038116156105bd5760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b604080516c01000000000000000000000000600160a060020a038087168202835285160260148201526028810183905290519081900360480190209392505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101839052905190819003603c019020919050565b60008060008061079e85610816565b60408051600080825260208083018085528d905260ff8716838501526060830186905260808301859052925195985093965091945060019360a0808401949293601f19830193908390039091019190865af1158015610801573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000806000806000808651604114151561082f57600080fd5b505050506020830151604084015160609094015160001a949093925090505600a165627a7a72305820258023a84a29d72d746c435fe478e4c603edd1c9acfb91a931158a4da69e71ca0029`

// DeployUnionPay deploys a new Ethereum contract, binding an instance of UnionPay to it.
func DeployUnionPay(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnionPay, error) {
	parsed, err := abi.JSON(strings.NewReader(UnionPayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnionPayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnionPay{UnionPayCaller: UnionPayCaller{contract: contract}, UnionPayTransactor: UnionPayTransactor{contract: contract}, UnionPayFilterer: UnionPayFilterer{contract: contract}}, nil
}

// UnionPay is an auto generated Go binding around an Ethereum contract.
type UnionPay struct {
	UnionPayCaller     // Read-only binding to the contract
	UnionPayTransactor // Write-only binding to the contract
	UnionPayFilterer   // Log filterer for contract events
}

// UnionPayCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnionPayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionPayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnionPayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionPayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnionPayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionPaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnionPaySession struct {
	Contract     *UnionPay         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnionPayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnionPayCallerSession struct {
	Contract *UnionPayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UnionPayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnionPayTransactorSession struct {
	Contract     *UnionPayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnionPayRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnionPayRaw struct {
	Contract *UnionPay // Generic contract binding to access the raw methods on
}

// UnionPayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnionPayCallerRaw struct {
	Contract *UnionPayCaller // Generic read-only contract binding to access the raw methods on
}

// UnionPayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnionPayTransactorRaw struct {
	Contract *UnionPayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnionPay creates a new instance of UnionPay, bound to a specific deployed contract.
func NewUnionPay(address common.Address, backend bind.ContractBackend) (*UnionPay, error) {
	contract, err := bindUnionPay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnionPay{UnionPayCaller: UnionPayCaller{contract: contract}, UnionPayTransactor: UnionPayTransactor{contract: contract}, UnionPayFilterer: UnionPayFilterer{contract: contract}}, nil
}

// NewUnionPayCaller creates a new read-only instance of UnionPay, bound to a specific deployed contract.
func NewUnionPayCaller(address common.Address, caller bind.ContractCaller) (*UnionPayCaller, error) {
	contract, err := bindUnionPay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnionPayCaller{contract: contract}, nil
}

// NewUnionPayTransactor creates a new write-only instance of UnionPay, bound to a specific deployed contract.
func NewUnionPayTransactor(address common.Address, transactor bind.ContractTransactor) (*UnionPayTransactor, error) {
	contract, err := bindUnionPay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnionPayTransactor{contract: contract}, nil
}

// NewUnionPayFilterer creates a new log filterer instance of UnionPay, bound to a specific deployed contract.
func NewUnionPayFilterer(address common.Address, filterer bind.ContractFilterer) (*UnionPayFilterer, error) {
	contract, err := bindUnionPay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnionPayFilterer{contract: contract}, nil
}

// bindUnionPay binds a generic wrapper to an already deployed contract.
func bindUnionPay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnionPayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnionPay *UnionPayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnionPay.Contract.UnionPayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnionPay *UnionPayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.Contract.UnionPayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnionPay *UnionPayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnionPay.Contract.UnionPayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnionPay *UnionPayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnionPay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnionPay *UnionPayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnionPay *UnionPayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnionPay.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_UnionPay *UnionPayCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_UnionPay *UnionPaySession) GetBalance() (*big.Int, error) {
	return _UnionPay.Contract.GetBalance(&_UnionPay.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_UnionPay *UnionPayCallerSession) GetBalance() (*big.Int, error) {
	return _UnionPay.Contract.GetBalance(&_UnionPay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnionPay *UnionPayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnionPay *UnionPaySession) Owner() (common.Address, error) {
	return _UnionPay.Contract.Owner(&_UnionPay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnionPay *UnionPayCallerSession) Owner() (common.Address, error) {
	return _UnionPay.Contract.Owner(&_UnionPay.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_UnionPay *UnionPayCaller) Platform(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "platform")
	return *ret0, err
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_UnionPay *UnionPaySession) Platform() (common.Address, error) {
	return _UnionPay.Contract.Platform(&_UnionPay.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_UnionPay *UnionPayCallerSession) Platform() (common.Address, error) {
	return _UnionPay.Contract.Platform(&_UnionPay.CallOpts)
}

// ReceiptUsed is a free data retrieval call binding the contract method 0x0c433edf.
//
// Solidity: function receiptUsed(_from address, _to address, _transId uint256) constant returns(bool)
func (_UnionPay *UnionPayCaller) ReceiptUsed(opts *bind.CallOpts, _from common.Address, _to common.Address, _transId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "receiptUsed", _from, _to, _transId)
	return *ret0, err
}

// ReceiptUsed is a free data retrieval call binding the contract method 0x0c433edf.
//
// Solidity: function receiptUsed(_from address, _to address, _transId uint256) constant returns(bool)
func (_UnionPay *UnionPaySession) ReceiptUsed(_from common.Address, _to common.Address, _transId *big.Int) (bool, error) {
	return _UnionPay.Contract.ReceiptUsed(&_UnionPay.CallOpts, _from, _to, _transId)
}

// ReceiptUsed is a free data retrieval call binding the contract method 0x0c433edf.
//
// Solidity: function receiptUsed(_from address, _to address, _transId uint256) constant returns(bool)
func (_UnionPay *UnionPayCallerSession) ReceiptUsed(_from common.Address, _to common.Address, _transId *big.Int) (bool, error) {
	return _UnionPay.Contract.ReceiptUsed(&_UnionPay.CallOpts, _from, _to, _transId)
}

// PlainPay is a paid mutator transaction binding the contract method 0xeb3f2427.
//
// Solidity: function plainPay() returns(bool)
func (_UnionPay *UnionPayTransactor) PlainPay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "plainPay")
}

// PlainPay is a paid mutator transaction binding the contract method 0xeb3f2427.
//
// Solidity: function plainPay() returns(bool)
func (_UnionPay *UnionPaySession) PlainPay() (*types.Transaction, error) {
	return _UnionPay.Contract.PlainPay(&_UnionPay.TransactOpts)
}

// PlainPay is a paid mutator transaction binding the contract method 0xeb3f2427.
//
// Solidity: function plainPay() returns(bool)
func (_UnionPay *UnionPayTransactorSession) PlainPay() (*types.Transaction, error) {
	return _UnionPay.Contract.PlainPay(&_UnionPay.TransactOpts)
}

// SafePay is a paid mutator transaction binding the contract method 0x33c91611.
//
// Solidity: function safePay(_transId uint256, _feePercentage uint256, _to address, _sig bytes) returns(bool)
func (_UnionPay *UnionPayTransactor) SafePay(opts *bind.TransactOpts, _transId *big.Int, _feePercentage *big.Int, _to common.Address, _sig []byte) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "safePay", _transId, _feePercentage, _to, _sig)
}

// SafePay is a paid mutator transaction binding the contract method 0x33c91611.
//
// Solidity: function safePay(_transId uint256, _feePercentage uint256, _to address, _sig bytes) returns(bool)
func (_UnionPay *UnionPaySession) SafePay(_transId *big.Int, _feePercentage *big.Int, _to common.Address, _sig []byte) (*types.Transaction, error) {
	return _UnionPay.Contract.SafePay(&_UnionPay.TransactOpts, _transId, _feePercentage, _to, _sig)
}

// SafePay is a paid mutator transaction binding the contract method 0x33c91611.
//
// Solidity: function safePay(_transId uint256, _feePercentage uint256, _to address, _sig bytes) returns(bool)
func (_UnionPay *UnionPayTransactorSession) SafePay(_transId *big.Int, _feePercentage *big.Int, _to common.Address, _sig []byte) (*types.Transaction, error) {
	return _UnionPay.Contract.SafePay(&_UnionPay.TransactOpts, _transId, _feePercentage, _to, _sig)
}

// SetPlatform is a paid mutator transaction binding the contract method 0x6945c5ea.
//
// Solidity: function setPlatform(_checker address) returns()
func (_UnionPay *UnionPayTransactor) SetPlatform(opts *bind.TransactOpts, _checker common.Address) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "setPlatform", _checker)
}

// SetPlatform is a paid mutator transaction binding the contract method 0x6945c5ea.
//
// Solidity: function setPlatform(_checker address) returns()
func (_UnionPay *UnionPaySession) SetPlatform(_checker common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.SetPlatform(&_UnionPay.TransactOpts, _checker)
}

// SetPlatform is a paid mutator transaction binding the contract method 0x6945c5ea.
//
// Solidity: function setPlatform(_checker address) returns()
func (_UnionPay *UnionPayTransactorSession) SetPlatform(_checker common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.SetPlatform(&_UnionPay.TransactOpts, _checker)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnionPay *UnionPayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnionPay *UnionPaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.TransferOwnership(&_UnionPay.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnionPay *UnionPayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.TransferOwnership(&_UnionPay.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_UnionPay *UnionPayTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_UnionPay *UnionPaySession) Withdraw() (*types.Transaction, error) {
	return _UnionPay.Contract.Withdraw(&_UnionPay.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_UnionPay *UnionPayTransactorSession) Withdraw() (*types.Transaction, error) {
	return _UnionPay.Contract.Withdraw(&_UnionPay.TransactOpts)
}

// UnionPayBareUserPayIterator is returned from FilterBareUserPay and is used to iterate over the raw logs and unpacked data for BareUserPay events raised by the UnionPay contract.
type UnionPayBareUserPayIterator struct {
	Event *UnionPayBareUserPay // Event containing the contract specifics and raw log

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
func (it *UnionPayBareUserPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnionPayBareUserPay)
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
		it.Event = new(UnionPayBareUserPay)
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
func (it *UnionPayBareUserPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnionPayBareUserPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnionPayBareUserPay represents a BareUserPay event raised by the UnionPay contract.
type UnionPayBareUserPay struct {
	From   common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBareUserPay is a free log retrieval operation binding the contract event 0x39a90d638bd2f864aeb4ed151f1668f52e4383137a90ee960743ad694d259bf3.
//
// Solidity: e BareUserPay(from address, amount uint256, data bytes)
func (_UnionPay *UnionPayFilterer) FilterBareUserPay(opts *bind.FilterOpts) (*UnionPayBareUserPayIterator, error) {

	logs, sub, err := _UnionPay.contract.FilterLogs(opts, "BareUserPay")
	if err != nil {
		return nil, err
	}
	return &UnionPayBareUserPayIterator{contract: _UnionPay.contract, event: "BareUserPay", logs: logs, sub: sub}, nil
}

// WatchBareUserPay is a free log subscription operation binding the contract event 0x39a90d638bd2f864aeb4ed151f1668f52e4383137a90ee960743ad694d259bf3.
//
// Solidity: e BareUserPay(from address, amount uint256, data bytes)
func (_UnionPay *UnionPayFilterer) WatchBareUserPay(opts *bind.WatchOpts, sink chan<- *UnionPayBareUserPay) (event.Subscription, error) {

	logs, sub, err := _UnionPay.contract.WatchLogs(opts, "BareUserPay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnionPayBareUserPay)
				if err := _UnionPay.contract.UnpackLog(event, "BareUserPay", log); err != nil {
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

// UnionPayUserPayIterator is returned from FilterUserPay and is used to iterate over the raw logs and unpacked data for UserPay events raised by the UnionPay contract.
type UnionPayUserPayIterator struct {
	Event *UnionPayUserPay // Event containing the contract specifics and raw log

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
func (it *UnionPayUserPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnionPayUserPay)
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
		it.Event = new(UnionPayUserPay)
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
func (it *UnionPayUserPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnionPayUserPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnionPayUserPay represents a UserPay event raised by the UnionPay contract.
type UnionPayUserPay struct {
	From         common.Address
	To           common.Address
	Amount       *big.Int
	AmountIndeed *big.Int
	TransId      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserPay is a free log retrieval operation binding the contract event 0x72d0b2856e77baaa53d185402480e6844c2d105276b126fdd46c34d14f4ac5f7.
//
// Solidity: e UserPay(from address, to address, amount uint256, amountIndeed uint256, transId uint256)
func (_UnionPay *UnionPayFilterer) FilterUserPay(opts *bind.FilterOpts) (*UnionPayUserPayIterator, error) {

	logs, sub, err := _UnionPay.contract.FilterLogs(opts, "UserPay")
	if err != nil {
		return nil, err
	}
	return &UnionPayUserPayIterator{contract: _UnionPay.contract, event: "UserPay", logs: logs, sub: sub}, nil
}

// WatchUserPay is a free log subscription operation binding the contract event 0x72d0b2856e77baaa53d185402480e6844c2d105276b126fdd46c34d14f4ac5f7.
//
// Solidity: e UserPay(from address, to address, amount uint256, amountIndeed uint256, transId uint256)
func (_UnionPay *UnionPayFilterer) WatchUserPay(opts *bind.WatchOpts, sink chan<- *UnionPayUserPay) (event.Subscription, error) {

	logs, sub, err := _UnionPay.contract.WatchLogs(opts, "UserPay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnionPayUserPay)
				if err := _UnionPay.contract.UnpackLog(event, "UserPay", log); err != nil {
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
