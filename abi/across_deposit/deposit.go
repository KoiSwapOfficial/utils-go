// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package across_deposit

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BridgeV1AcrossParams is an auto generated low-level Go binding around an user-defined struct.
type BridgeV1AcrossParams struct {
	Receiver             common.Address
	Refund               common.Address
	SendingToken         common.Address
	ReceivingToken       common.Address
	OutputAmount         *big.Int
	DestinationChainId   *big.Int
	ExclusiveRelayer     [32]byte
	QuoteTimestamp       uint32
	FillDeadline         uint32
	ExclusivityParameter uint32
	Message              []byte
}

// AcrossDepositMetaData contains all meta data concerning the AcrossDeposit contract.
var AcrossDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAcrossSpokePoolV4\",\"name\":\"_spokePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wrappedNative\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"DepositExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeRescued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRescued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"UpdateFeeCollector\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SPOKEPOOL\",\"outputs\":[{\"internalType\":\"contractIAcrossSpokePoolV4\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"exclusiveRelayer\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"quoteTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fillDeadline\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exclusivityParameter\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structBridgeV1.AcrossParams\",\"name\":\"p\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AcrossDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use AcrossDepositMetaData.ABI instead.
var AcrossDepositABI = AcrossDepositMetaData.ABI

// AcrossDeposit is an auto generated Go binding around an Ethereum contract.
type AcrossDeposit struct {
	AcrossDepositCaller     // Read-only binding to the contract
	AcrossDepositTransactor // Write-only binding to the contract
	AcrossDepositFilterer   // Log filterer for contract events
}

// AcrossDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type AcrossDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcrossDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AcrossDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcrossDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AcrossDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcrossDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AcrossDepositSession struct {
	Contract     *AcrossDeposit    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AcrossDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AcrossDepositCallerSession struct {
	Contract *AcrossDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AcrossDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AcrossDepositTransactorSession struct {
	Contract     *AcrossDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AcrossDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type AcrossDepositRaw struct {
	Contract *AcrossDeposit // Generic contract binding to access the raw methods on
}

// AcrossDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AcrossDepositCallerRaw struct {
	Contract *AcrossDepositCaller // Generic read-only contract binding to access the raw methods on
}

// AcrossDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AcrossDepositTransactorRaw struct {
	Contract *AcrossDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAcrossDeposit creates a new instance of AcrossDeposit, bound to a specific deployed contract.
func NewAcrossDeposit(address common.Address, backend bind.ContractBackend) (*AcrossDeposit, error) {
	contract, err := bindAcrossDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AcrossDeposit{AcrossDepositCaller: AcrossDepositCaller{contract: contract}, AcrossDepositTransactor: AcrossDepositTransactor{contract: contract}, AcrossDepositFilterer: AcrossDepositFilterer{contract: contract}}, nil
}

// NewAcrossDepositCaller creates a new read-only instance of AcrossDeposit, bound to a specific deployed contract.
func NewAcrossDepositCaller(address common.Address, caller bind.ContractCaller) (*AcrossDepositCaller, error) {
	contract, err := bindAcrossDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AcrossDepositCaller{contract: contract}, nil
}

// NewAcrossDepositTransactor creates a new write-only instance of AcrossDeposit, bound to a specific deployed contract.
func NewAcrossDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*AcrossDepositTransactor, error) {
	contract, err := bindAcrossDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AcrossDepositTransactor{contract: contract}, nil
}

// NewAcrossDepositFilterer creates a new log filterer instance of AcrossDeposit, bound to a specific deployed contract.
func NewAcrossDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*AcrossDepositFilterer, error) {
	contract, err := bindAcrossDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AcrossDepositFilterer{contract: contract}, nil
}

// bindAcrossDeposit binds a generic wrapper to an already deployed contract.
func bindAcrossDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AcrossDepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcrossDeposit *AcrossDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AcrossDeposit.Contract.AcrossDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcrossDeposit *AcrossDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.AcrossDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcrossDeposit *AcrossDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.AcrossDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcrossDeposit *AcrossDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AcrossDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcrossDeposit *AcrossDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcrossDeposit *AcrossDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.contract.Transact(opts, method, params...)
}

// SPOKEPOOL is a free data retrieval call binding the contract method 0xf6503992.
//
// Solidity: function SPOKEPOOL() view returns(address)
func (_AcrossDeposit *AcrossDepositCaller) SPOKEPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AcrossDeposit.contract.Call(opts, &out, "SPOKEPOOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPOKEPOOL is a free data retrieval call binding the contract method 0xf6503992.
//
// Solidity: function SPOKEPOOL() view returns(address)
func (_AcrossDeposit *AcrossDepositSession) SPOKEPOOL() (common.Address, error) {
	return _AcrossDeposit.Contract.SPOKEPOOL(&_AcrossDeposit.CallOpts)
}

// SPOKEPOOL is a free data retrieval call binding the contract method 0xf6503992.
//
// Solidity: function SPOKEPOOL() view returns(address)
func (_AcrossDeposit *AcrossDepositCallerSession) SPOKEPOOL() (common.Address, error) {
	return _AcrossDeposit.Contract.SPOKEPOOL(&_AcrossDeposit.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(bytes32)
func (_AcrossDeposit *AcrossDepositCaller) WRAPPEDNATIVE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AcrossDeposit.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(bytes32)
func (_AcrossDeposit *AcrossDepositSession) WRAPPEDNATIVE() ([32]byte, error) {
	return _AcrossDeposit.Contract.WRAPPEDNATIVE(&_AcrossDeposit.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(bytes32)
func (_AcrossDeposit *AcrossDepositCallerSession) WRAPPEDNATIVE() ([32]byte, error) {
	return _AcrossDeposit.Contract.WRAPPEDNATIVE(&_AcrossDeposit.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_AcrossDeposit *AcrossDepositCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AcrossDeposit.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_AcrossDeposit *AcrossDepositSession) FeeCollector() (common.Address, error) {
	return _AcrossDeposit.Contract.FeeCollector(&_AcrossDeposit.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_AcrossDeposit *AcrossDepositCallerSession) FeeCollector() (common.Address, error) {
	return _AcrossDeposit.Contract.FeeCollector(&_AcrossDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AcrossDeposit *AcrossDepositCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AcrossDeposit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AcrossDeposit *AcrossDepositSession) Owner() (common.Address, error) {
	return _AcrossDeposit.Contract.Owner(&_AcrossDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AcrossDeposit *AcrossDepositCallerSession) Owner() (common.Address, error) {
	return _AcrossDeposit.Contract.Owner(&_AcrossDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AcrossDeposit *AcrossDepositCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AcrossDeposit.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AcrossDeposit *AcrossDepositSession) Paused() (bool, error) {
	return _AcrossDeposit.Contract.Paused(&_AcrossDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AcrossDeposit *AcrossDepositCallerSession) Paused() (bool, error) {
	return _AcrossDeposit.Contract.Paused(&_AcrossDeposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x5e3bd9e9.
//
// Solidity: function deposit((address,address,address,address,uint256,uint256,bytes32,uint32,uint32,uint32,bytes) p, uint256 amount, uint256 fee, bool isNative) payable returns()
func (_AcrossDeposit *AcrossDepositTransactor) Deposit(opts *bind.TransactOpts, p BridgeV1AcrossParams, amount *big.Int, fee *big.Int, isNative bool) (*types.Transaction, error) {
	return _AcrossDeposit.contract.Transact(opts, "deposit", p, amount, fee, isNative)
}

// Deposit is a paid mutator transaction binding the contract method 0x5e3bd9e9.
//
// Solidity: function deposit((address,address,address,address,uint256,uint256,bytes32,uint32,uint32,uint32,bytes) p, uint256 amount, uint256 fee, bool isNative) payable returns()
func (_AcrossDeposit *AcrossDepositSession) Deposit(p BridgeV1AcrossParams, amount *big.Int, fee *big.Int, isNative bool) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.Deposit(&_AcrossDeposit.TransactOpts, p, amount, fee, isNative)
}

// Deposit is a paid mutator transaction binding the contract method 0x5e3bd9e9.
//
// Solidity: function deposit((address,address,address,address,uint256,uint256,bytes32,uint32,uint32,uint32,bytes) p, uint256 amount, uint256 fee, bool isNative) payable returns()
func (_AcrossDeposit *AcrossDepositTransactorSession) Deposit(p BridgeV1AcrossParams, amount *big.Int, fee *big.Int, isNative bool) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.Deposit(&_AcrossDeposit.TransactOpts, p, amount, fee, isNative)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AcrossDeposit *AcrossDepositTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcrossDeposit.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AcrossDeposit *AcrossDepositSession) Pause() (*types.Transaction, error) {
	return _AcrossDeposit.Contract.Pause(&_AcrossDeposit.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AcrossDeposit *AcrossDepositTransactorSession) Pause() (*types.Transaction, error) {
	return _AcrossDeposit.Contract.Pause(&_AcrossDeposit.TransactOpts)
}

// RescueNative is a paid mutator transaction binding the contract method 0x454aa669.
//
// Solidity: function rescueNative(uint256 amount) returns()
func (_AcrossDeposit *AcrossDepositTransactor) RescueNative(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AcrossDeposit.contract.Transact(opts, "rescueNative", amount)
}

// RescueNative is a paid mutator transaction binding the contract method 0x454aa669.
//
// Solidity: function rescueNative(uint256 amount) returns()
func (_AcrossDeposit *AcrossDepositSession) RescueNative(amount *big.Int) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.RescueNative(&_AcrossDeposit.TransactOpts, amount)
}

// RescueNative is a paid mutator transaction binding the contract method 0x454aa669.
//
// Solidity: function rescueNative(uint256 amount) returns()
func (_AcrossDeposit *AcrossDepositTransactorSession) RescueNative(amount *big.Int) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.RescueNative(&_AcrossDeposit.TransactOpts, amount)
}

// RescueToken is a paid mutator transaction binding the contract method 0x33f3d628.
//
// Solidity: function rescueToken(address token, uint256 amount) returns()
func (_AcrossDeposit *AcrossDepositTransactor) RescueToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossDeposit.contract.Transact(opts, "rescueToken", token, amount)
}

// RescueToken is a paid mutator transaction binding the contract method 0x33f3d628.
//
// Solidity: function rescueToken(address token, uint256 amount) returns()
func (_AcrossDeposit *AcrossDepositSession) RescueToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.RescueToken(&_AcrossDeposit.TransactOpts, token, amount)
}

// RescueToken is a paid mutator transaction binding the contract method 0x33f3d628.
//
// Solidity: function rescueToken(address token, uint256 amount) returns()
func (_AcrossDeposit *AcrossDepositTransactorSession) RescueToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.RescueToken(&_AcrossDeposit.TransactOpts, token, amount)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_AcrossDeposit *AcrossDepositTransactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _AcrossDeposit.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_AcrossDeposit *AcrossDepositSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.SetFeeCollector(&_AcrossDeposit.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_AcrossDeposit *AcrossDepositTransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _AcrossDeposit.Contract.SetFeeCollector(&_AcrossDeposit.TransactOpts, _feeCollector)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AcrossDeposit *AcrossDepositTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcrossDeposit.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AcrossDeposit *AcrossDepositSession) Unpause() (*types.Transaction, error) {
	return _AcrossDeposit.Contract.Unpause(&_AcrossDeposit.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AcrossDeposit *AcrossDepositTransactorSession) Unpause() (*types.Transaction, error) {
	return _AcrossDeposit.Contract.Unpause(&_AcrossDeposit.TransactOpts)
}

// AcrossDepositDepositExecutedIterator is returned from FilterDepositExecuted and is used to iterate over the raw logs and unpacked data for DepositExecuted events raised by the AcrossDeposit contract.
type AcrossDepositDepositExecutedIterator struct {
	Event *AcrossDepositDepositExecuted // Event containing the contract specifics and raw log

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
func (it *AcrossDepositDepositExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossDepositDepositExecuted)
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
		it.Event = new(AcrossDepositDepositExecuted)
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
func (it *AcrossDepositDepositExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossDepositDepositExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossDepositDepositExecuted represents a DepositExecuted event raised by the AcrossDeposit contract.
type AcrossDepositDepositExecuted struct {
	Caller    common.Address
	Amount    *big.Int
	BridgeFee *big.Int
	IsNative  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositExecuted is a free log retrieval operation binding the contract event 0x1054a5627e951e7fdeda387664ecb4a17707ac049d8d4ea800c87255b9b2e811.
//
// Solidity: event DepositExecuted(address indexed caller, uint256 amount, uint256 bridgeFee, bool isNative)
func (_AcrossDeposit *AcrossDepositFilterer) FilterDepositExecuted(opts *bind.FilterOpts, caller []common.Address) (*AcrossDepositDepositExecutedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _AcrossDeposit.contract.FilterLogs(opts, "DepositExecuted", callerRule)
	if err != nil {
		return nil, err
	}
	return &AcrossDepositDepositExecutedIterator{contract: _AcrossDeposit.contract, event: "DepositExecuted", logs: logs, sub: sub}, nil
}

// WatchDepositExecuted is a free log subscription operation binding the contract event 0x1054a5627e951e7fdeda387664ecb4a17707ac049d8d4ea800c87255b9b2e811.
//
// Solidity: event DepositExecuted(address indexed caller, uint256 amount, uint256 bridgeFee, bool isNative)
func (_AcrossDeposit *AcrossDepositFilterer) WatchDepositExecuted(opts *bind.WatchOpts, sink chan<- *AcrossDepositDepositExecuted, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _AcrossDeposit.contract.WatchLogs(opts, "DepositExecuted", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossDepositDepositExecuted)
				if err := _AcrossDeposit.contract.UnpackLog(event, "DepositExecuted", log); err != nil {
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

// ParseDepositExecuted is a log parse operation binding the contract event 0x1054a5627e951e7fdeda387664ecb4a17707ac049d8d4ea800c87255b9b2e811.
//
// Solidity: event DepositExecuted(address indexed caller, uint256 amount, uint256 bridgeFee, bool isNative)
func (_AcrossDeposit *AcrossDepositFilterer) ParseDepositExecuted(log types.Log) (*AcrossDepositDepositExecuted, error) {
	event := new(AcrossDepositDepositExecuted)
	if err := _AcrossDeposit.contract.UnpackLog(event, "DepositExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcrossDepositNativeRescuedIterator is returned from FilterNativeRescued and is used to iterate over the raw logs and unpacked data for NativeRescued events raised by the AcrossDeposit contract.
type AcrossDepositNativeRescuedIterator struct {
	Event *AcrossDepositNativeRescued // Event containing the contract specifics and raw log

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
func (it *AcrossDepositNativeRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossDepositNativeRescued)
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
		it.Event = new(AcrossDepositNativeRescued)
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
func (it *AcrossDepositNativeRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossDepositNativeRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossDepositNativeRescued represents a NativeRescued event raised by the AcrossDeposit contract.
type AcrossDepositNativeRescued struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeRescued is a free log retrieval operation binding the contract event 0xe3eb98b7fe2a0c1d490b92af73eeae611e9b00ab3c3f70b20bd7bb43f67a0f43.
//
// Solidity: event NativeRescued(address indexed to, uint256 amount)
func (_AcrossDeposit *AcrossDepositFilterer) FilterNativeRescued(opts *bind.FilterOpts, to []common.Address) (*AcrossDepositNativeRescuedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossDeposit.contract.FilterLogs(opts, "NativeRescued", toRule)
	if err != nil {
		return nil, err
	}
	return &AcrossDepositNativeRescuedIterator{contract: _AcrossDeposit.contract, event: "NativeRescued", logs: logs, sub: sub}, nil
}

// WatchNativeRescued is a free log subscription operation binding the contract event 0xe3eb98b7fe2a0c1d490b92af73eeae611e9b00ab3c3f70b20bd7bb43f67a0f43.
//
// Solidity: event NativeRescued(address indexed to, uint256 amount)
func (_AcrossDeposit *AcrossDepositFilterer) WatchNativeRescued(opts *bind.WatchOpts, sink chan<- *AcrossDepositNativeRescued, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossDeposit.contract.WatchLogs(opts, "NativeRescued", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossDepositNativeRescued)
				if err := _AcrossDeposit.contract.UnpackLog(event, "NativeRescued", log); err != nil {
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

// ParseNativeRescued is a log parse operation binding the contract event 0xe3eb98b7fe2a0c1d490b92af73eeae611e9b00ab3c3f70b20bd7bb43f67a0f43.
//
// Solidity: event NativeRescued(address indexed to, uint256 amount)
func (_AcrossDeposit *AcrossDepositFilterer) ParseNativeRescued(log types.Log) (*AcrossDepositNativeRescued, error) {
	event := new(AcrossDepositNativeRescued)
	if err := _AcrossDeposit.contract.UnpackLog(event, "NativeRescued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcrossDepositPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AcrossDeposit contract.
type AcrossDepositPausedIterator struct {
	Event *AcrossDepositPaused // Event containing the contract specifics and raw log

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
func (it *AcrossDepositPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossDepositPaused)
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
		it.Event = new(AcrossDepositPaused)
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
func (it *AcrossDepositPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossDepositPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossDepositPaused represents a Paused event raised by the AcrossDeposit contract.
type AcrossDepositPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AcrossDeposit *AcrossDepositFilterer) FilterPaused(opts *bind.FilterOpts) (*AcrossDepositPausedIterator, error) {

	logs, sub, err := _AcrossDeposit.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AcrossDepositPausedIterator{contract: _AcrossDeposit.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AcrossDeposit *AcrossDepositFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AcrossDepositPaused) (event.Subscription, error) {

	logs, sub, err := _AcrossDeposit.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossDepositPaused)
				if err := _AcrossDeposit.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AcrossDeposit *AcrossDepositFilterer) ParsePaused(log types.Log) (*AcrossDepositPaused, error) {
	event := new(AcrossDepositPaused)
	if err := _AcrossDeposit.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcrossDepositTokenRescuedIterator is returned from FilterTokenRescued and is used to iterate over the raw logs and unpacked data for TokenRescued events raised by the AcrossDeposit contract.
type AcrossDepositTokenRescuedIterator struct {
	Event *AcrossDepositTokenRescued // Event containing the contract specifics and raw log

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
func (it *AcrossDepositTokenRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossDepositTokenRescued)
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
		it.Event = new(AcrossDepositTokenRescued)
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
func (it *AcrossDepositTokenRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossDepositTokenRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossDepositTokenRescued represents a TokenRescued event raised by the AcrossDeposit contract.
type AcrossDepositTokenRescued struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRescued is a free log retrieval operation binding the contract event 0x4143f7b5cb6ea007914c32b8a3e64cebc051d7f493fa0755454da1e47701e125.
//
// Solidity: event TokenRescued(address indexed token, address indexed to, uint256 amount)
func (_AcrossDeposit *AcrossDepositFilterer) FilterTokenRescued(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*AcrossDepositTokenRescuedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossDeposit.contract.FilterLogs(opts, "TokenRescued", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AcrossDepositTokenRescuedIterator{contract: _AcrossDeposit.contract, event: "TokenRescued", logs: logs, sub: sub}, nil
}

// WatchTokenRescued is a free log subscription operation binding the contract event 0x4143f7b5cb6ea007914c32b8a3e64cebc051d7f493fa0755454da1e47701e125.
//
// Solidity: event TokenRescued(address indexed token, address indexed to, uint256 amount)
func (_AcrossDeposit *AcrossDepositFilterer) WatchTokenRescued(opts *bind.WatchOpts, sink chan<- *AcrossDepositTokenRescued, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossDeposit.contract.WatchLogs(opts, "TokenRescued", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossDepositTokenRescued)
				if err := _AcrossDeposit.contract.UnpackLog(event, "TokenRescued", log); err != nil {
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

// ParseTokenRescued is a log parse operation binding the contract event 0x4143f7b5cb6ea007914c32b8a3e64cebc051d7f493fa0755454da1e47701e125.
//
// Solidity: event TokenRescued(address indexed token, address indexed to, uint256 amount)
func (_AcrossDeposit *AcrossDepositFilterer) ParseTokenRescued(log types.Log) (*AcrossDepositTokenRescued, error) {
	event := new(AcrossDepositTokenRescued)
	if err := _AcrossDeposit.contract.UnpackLog(event, "TokenRescued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcrossDepositUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AcrossDeposit contract.
type AcrossDepositUnpausedIterator struct {
	Event *AcrossDepositUnpaused // Event containing the contract specifics and raw log

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
func (it *AcrossDepositUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossDepositUnpaused)
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
		it.Event = new(AcrossDepositUnpaused)
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
func (it *AcrossDepositUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossDepositUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossDepositUnpaused represents a Unpaused event raised by the AcrossDeposit contract.
type AcrossDepositUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AcrossDeposit *AcrossDepositFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AcrossDepositUnpausedIterator, error) {

	logs, sub, err := _AcrossDeposit.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AcrossDepositUnpausedIterator{contract: _AcrossDeposit.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AcrossDeposit *AcrossDepositFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AcrossDepositUnpaused) (event.Subscription, error) {

	logs, sub, err := _AcrossDeposit.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossDepositUnpaused)
				if err := _AcrossDeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AcrossDeposit *AcrossDepositFilterer) ParseUnpaused(log types.Log) (*AcrossDepositUnpaused, error) {
	event := new(AcrossDepositUnpaused)
	if err := _AcrossDeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcrossDepositUpdateFeeCollectorIterator is returned from FilterUpdateFeeCollector and is used to iterate over the raw logs and unpacked data for UpdateFeeCollector events raised by the AcrossDeposit contract.
type AcrossDepositUpdateFeeCollectorIterator struct {
	Event *AcrossDepositUpdateFeeCollector // Event containing the contract specifics and raw log

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
func (it *AcrossDepositUpdateFeeCollectorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossDepositUpdateFeeCollector)
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
		it.Event = new(AcrossDepositUpdateFeeCollector)
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
func (it *AcrossDepositUpdateFeeCollectorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossDepositUpdateFeeCollectorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossDepositUpdateFeeCollector represents a UpdateFeeCollector event raised by the AcrossDeposit contract.
type AcrossDepositUpdateFeeCollector struct {
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeCollector is a free log retrieval operation binding the contract event 0x338e8652f97d92dc615790fbe0fba33e7ce110c00e6e328be604e8119a3b6aa6.
//
// Solidity: event UpdateFeeCollector(address _feeCollector)
func (_AcrossDeposit *AcrossDepositFilterer) FilterUpdateFeeCollector(opts *bind.FilterOpts) (*AcrossDepositUpdateFeeCollectorIterator, error) {

	logs, sub, err := _AcrossDeposit.contract.FilterLogs(opts, "UpdateFeeCollector")
	if err != nil {
		return nil, err
	}
	return &AcrossDepositUpdateFeeCollectorIterator{contract: _AcrossDeposit.contract, event: "UpdateFeeCollector", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeCollector is a free log subscription operation binding the contract event 0x338e8652f97d92dc615790fbe0fba33e7ce110c00e6e328be604e8119a3b6aa6.
//
// Solidity: event UpdateFeeCollector(address _feeCollector)
func (_AcrossDeposit *AcrossDepositFilterer) WatchUpdateFeeCollector(opts *bind.WatchOpts, sink chan<- *AcrossDepositUpdateFeeCollector) (event.Subscription, error) {

	logs, sub, err := _AcrossDeposit.contract.WatchLogs(opts, "UpdateFeeCollector")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossDepositUpdateFeeCollector)
				if err := _AcrossDeposit.contract.UnpackLog(event, "UpdateFeeCollector", log); err != nil {
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

// ParseUpdateFeeCollector is a log parse operation binding the contract event 0x338e8652f97d92dc615790fbe0fba33e7ce110c00e6e328be604e8119a3b6aa6.
//
// Solidity: event UpdateFeeCollector(address _feeCollector)
func (_AcrossDeposit *AcrossDepositFilterer) ParseUpdateFeeCollector(log types.Log) (*AcrossDepositUpdateFeeCollector, error) {
	event := new(AcrossDepositUpdateFeeCollector)
	if err := _AcrossDeposit.contract.UnpackLog(event, "UpdateFeeCollector", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
