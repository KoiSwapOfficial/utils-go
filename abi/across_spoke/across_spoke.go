// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package across_spoke

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

// AcrossAcrossParams is an auto generated low-level Go binding around an user-defined struct.
type AcrossAcrossParams struct {
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

// AcrossSpokeMetaData contains all meta data concerning the AcrossSpoke contract.
var AcrossSpokeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAcrossSpokePoolV4\",\"name\":\"_spokePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wrappedNative\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"DepositExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeRescued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRescued\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SPOKEPOOL\",\"outputs\":[{\"internalType\":\"contractIAcrossSpokePoolV4\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"exclusiveRelayer\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"quoteTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fillDeadline\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exclusivityParameter\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structAcross.AcrossParams\",\"name\":\"p\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AcrossSpokeABI is the input ABI used to generate the binding from.
// Deprecated: Use AcrossSpokeMetaData.ABI instead.
var AcrossSpokeABI = AcrossSpokeMetaData.ABI

// AcrossSpoke is an auto generated Go binding around an Ethereum contract.
type AcrossSpoke struct {
	AcrossSpokeCaller     // Read-only binding to the contract
	AcrossSpokeTransactor // Write-only binding to the contract
	AcrossSpokeFilterer   // Log filterer for contract events
}

// AcrossSpokeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AcrossSpokeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcrossSpokeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AcrossSpokeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcrossSpokeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AcrossSpokeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcrossSpokeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AcrossSpokeSession struct {
	Contract     *AcrossSpoke      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AcrossSpokeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AcrossSpokeCallerSession struct {
	Contract *AcrossSpokeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AcrossSpokeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AcrossSpokeTransactorSession struct {
	Contract     *AcrossSpokeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AcrossSpokeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AcrossSpokeRaw struct {
	Contract *AcrossSpoke // Generic contract binding to access the raw methods on
}

// AcrossSpokeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AcrossSpokeCallerRaw struct {
	Contract *AcrossSpokeCaller // Generic read-only contract binding to access the raw methods on
}

// AcrossSpokeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AcrossSpokeTransactorRaw struct {
	Contract *AcrossSpokeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAcrossSpoke creates a new instance of AcrossSpoke, bound to a specific deployed contract.
func NewAcrossSpoke(address common.Address, backend bind.ContractBackend) (*AcrossSpoke, error) {
	contract, err := bindAcrossSpoke(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AcrossSpoke{AcrossSpokeCaller: AcrossSpokeCaller{contract: contract}, AcrossSpokeTransactor: AcrossSpokeTransactor{contract: contract}, AcrossSpokeFilterer: AcrossSpokeFilterer{contract: contract}}, nil
}

// NewAcrossSpokeCaller creates a new read-only instance of AcrossSpoke, bound to a specific deployed contract.
func NewAcrossSpokeCaller(address common.Address, caller bind.ContractCaller) (*AcrossSpokeCaller, error) {
	contract, err := bindAcrossSpoke(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AcrossSpokeCaller{contract: contract}, nil
}

// NewAcrossSpokeTransactor creates a new write-only instance of AcrossSpoke, bound to a specific deployed contract.
func NewAcrossSpokeTransactor(address common.Address, transactor bind.ContractTransactor) (*AcrossSpokeTransactor, error) {
	contract, err := bindAcrossSpoke(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AcrossSpokeTransactor{contract: contract}, nil
}

// NewAcrossSpokeFilterer creates a new log filterer instance of AcrossSpoke, bound to a specific deployed contract.
func NewAcrossSpokeFilterer(address common.Address, filterer bind.ContractFilterer) (*AcrossSpokeFilterer, error) {
	contract, err := bindAcrossSpoke(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AcrossSpokeFilterer{contract: contract}, nil
}

// bindAcrossSpoke binds a generic wrapper to an already deployed contract.
func bindAcrossSpoke(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AcrossSpokeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcrossSpoke *AcrossSpokeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AcrossSpoke.Contract.AcrossSpokeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcrossSpoke *AcrossSpokeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.AcrossSpokeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcrossSpoke *AcrossSpokeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.AcrossSpokeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcrossSpoke *AcrossSpokeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AcrossSpoke.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcrossSpoke *AcrossSpokeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcrossSpoke *AcrossSpokeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.contract.Transact(opts, method, params...)
}

// SPOKEPOOL is a free data retrieval call binding the contract method 0xf6503992.
//
// Solidity: function SPOKEPOOL() view returns(address)
func (_AcrossSpoke *AcrossSpokeCaller) SPOKEPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AcrossSpoke.contract.Call(opts, &out, "SPOKEPOOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPOKEPOOL is a free data retrieval call binding the contract method 0xf6503992.
//
// Solidity: function SPOKEPOOL() view returns(address)
func (_AcrossSpoke *AcrossSpokeSession) SPOKEPOOL() (common.Address, error) {
	return _AcrossSpoke.Contract.SPOKEPOOL(&_AcrossSpoke.CallOpts)
}

// SPOKEPOOL is a free data retrieval call binding the contract method 0xf6503992.
//
// Solidity: function SPOKEPOOL() view returns(address)
func (_AcrossSpoke *AcrossSpokeCallerSession) SPOKEPOOL() (common.Address, error) {
	return _AcrossSpoke.Contract.SPOKEPOOL(&_AcrossSpoke.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(bytes32)
func (_AcrossSpoke *AcrossSpokeCaller) WRAPPEDNATIVE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AcrossSpoke.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(bytes32)
func (_AcrossSpoke *AcrossSpokeSession) WRAPPEDNATIVE() ([32]byte, error) {
	return _AcrossSpoke.Contract.WRAPPEDNATIVE(&_AcrossSpoke.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(bytes32)
func (_AcrossSpoke *AcrossSpokeCallerSession) WRAPPEDNATIVE() ([32]byte, error) {
	return _AcrossSpoke.Contract.WRAPPEDNATIVE(&_AcrossSpoke.CallOpts)
}

// Maker is a free data retrieval call binding the contract method 0x50655d8c.
//
// Solidity: function maker() view returns(address)
func (_AcrossSpoke *AcrossSpokeCaller) Maker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AcrossSpoke.contract.Call(opts, &out, "maker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Maker is a free data retrieval call binding the contract method 0x50655d8c.
//
// Solidity: function maker() view returns(address)
func (_AcrossSpoke *AcrossSpokeSession) Maker() (common.Address, error) {
	return _AcrossSpoke.Contract.Maker(&_AcrossSpoke.CallOpts)
}

// Maker is a free data retrieval call binding the contract method 0x50655d8c.
//
// Solidity: function maker() view returns(address)
func (_AcrossSpoke *AcrossSpokeCallerSession) Maker() (common.Address, error) {
	return _AcrossSpoke.Contract.Maker(&_AcrossSpoke.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x5e3bd9e9.
//
// Solidity: function deposit((address,address,address,address,uint256,uint256,bytes32,uint32,uint32,uint32,bytes) p, uint256 bridgeFee, uint256 amount, bool isNative) payable returns()
func (_AcrossSpoke *AcrossSpokeTransactor) Deposit(opts *bind.TransactOpts, p AcrossAcrossParams, bridgeFee *big.Int, amount *big.Int, isNative bool) (*types.Transaction, error) {
	return _AcrossSpoke.contract.Transact(opts, "deposit", p, bridgeFee, amount, isNative)
}

// Deposit is a paid mutator transaction binding the contract method 0x5e3bd9e9.
//
// Solidity: function deposit((address,address,address,address,uint256,uint256,bytes32,uint32,uint32,uint32,bytes) p, uint256 bridgeFee, uint256 amount, bool isNative) payable returns()
func (_AcrossSpoke *AcrossSpokeSession) Deposit(p AcrossAcrossParams, bridgeFee *big.Int, amount *big.Int, isNative bool) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.Deposit(&_AcrossSpoke.TransactOpts, p, bridgeFee, amount, isNative)
}

// Deposit is a paid mutator transaction binding the contract method 0x5e3bd9e9.
//
// Solidity: function deposit((address,address,address,address,uint256,uint256,bytes32,uint32,uint32,uint32,bytes) p, uint256 bridgeFee, uint256 amount, bool isNative) payable returns()
func (_AcrossSpoke *AcrossSpokeTransactorSession) Deposit(p AcrossAcrossParams, bridgeFee *big.Int, amount *big.Int, isNative bool) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.Deposit(&_AcrossSpoke.TransactOpts, p, bridgeFee, amount, isNative)
}

// RescueNative is a paid mutator transaction binding the contract method 0x1291f79d.
//
// Solidity: function rescueNative(address to, uint256 amount) returns()
func (_AcrossSpoke *AcrossSpokeTransactor) RescueNative(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossSpoke.contract.Transact(opts, "rescueNative", to, amount)
}

// RescueNative is a paid mutator transaction binding the contract method 0x1291f79d.
//
// Solidity: function rescueNative(address to, uint256 amount) returns()
func (_AcrossSpoke *AcrossSpokeSession) RescueNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.RescueNative(&_AcrossSpoke.TransactOpts, to, amount)
}

// RescueNative is a paid mutator transaction binding the contract method 0x1291f79d.
//
// Solidity: function rescueNative(address to, uint256 amount) returns()
func (_AcrossSpoke *AcrossSpokeTransactorSession) RescueNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.RescueNative(&_AcrossSpoke.TransactOpts, to, amount)
}

// RescueToken is a paid mutator transaction binding the contract method 0xe5711e8b.
//
// Solidity: function rescueToken(address token, address to, uint256 amount) returns()
func (_AcrossSpoke *AcrossSpokeTransactor) RescueToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossSpoke.contract.Transact(opts, "rescueToken", token, to, amount)
}

// RescueToken is a paid mutator transaction binding the contract method 0xe5711e8b.
//
// Solidity: function rescueToken(address token, address to, uint256 amount) returns()
func (_AcrossSpoke *AcrossSpokeSession) RescueToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.RescueToken(&_AcrossSpoke.TransactOpts, token, to, amount)
}

// RescueToken is a paid mutator transaction binding the contract method 0xe5711e8b.
//
// Solidity: function rescueToken(address token, address to, uint256 amount) returns()
func (_AcrossSpoke *AcrossSpokeTransactorSession) RescueToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AcrossSpoke.Contract.RescueToken(&_AcrossSpoke.TransactOpts, token, to, amount)
}

// AcrossSpokeDepositExecutedIterator is returned from FilterDepositExecuted and is used to iterate over the raw logs and unpacked data for DepositExecuted events raised by the AcrossSpoke contract.
type AcrossSpokeDepositExecutedIterator struct {
	Event *AcrossSpokeDepositExecuted // Event containing the contract specifics and raw log

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
func (it *AcrossSpokeDepositExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossSpokeDepositExecuted)
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
		it.Event = new(AcrossSpokeDepositExecuted)
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
func (it *AcrossSpokeDepositExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossSpokeDepositExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossSpokeDepositExecuted represents a DepositExecuted event raised by the AcrossSpoke contract.
type AcrossSpokeDepositExecuted struct {
	Caller    common.Address
	Amount    *big.Int
	BridgeFee *big.Int
	IsNative  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositExecuted is a free log retrieval operation binding the contract event 0x1054a5627e951e7fdeda387664ecb4a17707ac049d8d4ea800c87255b9b2e811.
//
// Solidity: event DepositExecuted(address indexed caller, uint256 amount, uint256 bridgeFee, bool isNative)
func (_AcrossSpoke *AcrossSpokeFilterer) FilterDepositExecuted(opts *bind.FilterOpts, caller []common.Address) (*AcrossSpokeDepositExecutedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _AcrossSpoke.contract.FilterLogs(opts, "DepositExecuted", callerRule)
	if err != nil {
		return nil, err
	}
	return &AcrossSpokeDepositExecutedIterator{contract: _AcrossSpoke.contract, event: "DepositExecuted", logs: logs, sub: sub}, nil
}

// WatchDepositExecuted is a free log subscription operation binding the contract event 0x1054a5627e951e7fdeda387664ecb4a17707ac049d8d4ea800c87255b9b2e811.
//
// Solidity: event DepositExecuted(address indexed caller, uint256 amount, uint256 bridgeFee, bool isNative)
func (_AcrossSpoke *AcrossSpokeFilterer) WatchDepositExecuted(opts *bind.WatchOpts, sink chan<- *AcrossSpokeDepositExecuted, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _AcrossSpoke.contract.WatchLogs(opts, "DepositExecuted", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossSpokeDepositExecuted)
				if err := _AcrossSpoke.contract.UnpackLog(event, "DepositExecuted", log); err != nil {
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
func (_AcrossSpoke *AcrossSpokeFilterer) ParseDepositExecuted(log types.Log) (*AcrossSpokeDepositExecuted, error) {
	event := new(AcrossSpokeDepositExecuted)
	if err := _AcrossSpoke.contract.UnpackLog(event, "DepositExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcrossSpokeNativeRescuedIterator is returned from FilterNativeRescued and is used to iterate over the raw logs and unpacked data for NativeRescued events raised by the AcrossSpoke contract.
type AcrossSpokeNativeRescuedIterator struct {
	Event *AcrossSpokeNativeRescued // Event containing the contract specifics and raw log

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
func (it *AcrossSpokeNativeRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossSpokeNativeRescued)
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
		it.Event = new(AcrossSpokeNativeRescued)
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
func (it *AcrossSpokeNativeRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossSpokeNativeRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossSpokeNativeRescued represents a NativeRescued event raised by the AcrossSpoke contract.
type AcrossSpokeNativeRescued struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeRescued is a free log retrieval operation binding the contract event 0xe3eb98b7fe2a0c1d490b92af73eeae611e9b00ab3c3f70b20bd7bb43f67a0f43.
//
// Solidity: event NativeRescued(address indexed to, uint256 amount)
func (_AcrossSpoke *AcrossSpokeFilterer) FilterNativeRescued(opts *bind.FilterOpts, to []common.Address) (*AcrossSpokeNativeRescuedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossSpoke.contract.FilterLogs(opts, "NativeRescued", toRule)
	if err != nil {
		return nil, err
	}
	return &AcrossSpokeNativeRescuedIterator{contract: _AcrossSpoke.contract, event: "NativeRescued", logs: logs, sub: sub}, nil
}

// WatchNativeRescued is a free log subscription operation binding the contract event 0xe3eb98b7fe2a0c1d490b92af73eeae611e9b00ab3c3f70b20bd7bb43f67a0f43.
//
// Solidity: event NativeRescued(address indexed to, uint256 amount)
func (_AcrossSpoke *AcrossSpokeFilterer) WatchNativeRescued(opts *bind.WatchOpts, sink chan<- *AcrossSpokeNativeRescued, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossSpoke.contract.WatchLogs(opts, "NativeRescued", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossSpokeNativeRescued)
				if err := _AcrossSpoke.contract.UnpackLog(event, "NativeRescued", log); err != nil {
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
func (_AcrossSpoke *AcrossSpokeFilterer) ParseNativeRescued(log types.Log) (*AcrossSpokeNativeRescued, error) {
	event := new(AcrossSpokeNativeRescued)
	if err := _AcrossSpoke.contract.UnpackLog(event, "NativeRescued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcrossSpokeTokenRescuedIterator is returned from FilterTokenRescued and is used to iterate over the raw logs and unpacked data for TokenRescued events raised by the AcrossSpoke contract.
type AcrossSpokeTokenRescuedIterator struct {
	Event *AcrossSpokeTokenRescued // Event containing the contract specifics and raw log

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
func (it *AcrossSpokeTokenRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcrossSpokeTokenRescued)
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
		it.Event = new(AcrossSpokeTokenRescued)
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
func (it *AcrossSpokeTokenRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcrossSpokeTokenRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcrossSpokeTokenRescued represents a TokenRescued event raised by the AcrossSpoke contract.
type AcrossSpokeTokenRescued struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRescued is a free log retrieval operation binding the contract event 0x4143f7b5cb6ea007914c32b8a3e64cebc051d7f493fa0755454da1e47701e125.
//
// Solidity: event TokenRescued(address indexed token, address indexed to, uint256 amount)
func (_AcrossSpoke *AcrossSpokeFilterer) FilterTokenRescued(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*AcrossSpokeTokenRescuedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossSpoke.contract.FilterLogs(opts, "TokenRescued", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AcrossSpokeTokenRescuedIterator{contract: _AcrossSpoke.contract, event: "TokenRescued", logs: logs, sub: sub}, nil
}

// WatchTokenRescued is a free log subscription operation binding the contract event 0x4143f7b5cb6ea007914c32b8a3e64cebc051d7f493fa0755454da1e47701e125.
//
// Solidity: event TokenRescued(address indexed token, address indexed to, uint256 amount)
func (_AcrossSpoke *AcrossSpokeFilterer) WatchTokenRescued(opts *bind.WatchOpts, sink chan<- *AcrossSpokeTokenRescued, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AcrossSpoke.contract.WatchLogs(opts, "TokenRescued", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcrossSpokeTokenRescued)
				if err := _AcrossSpoke.contract.UnpackLog(event, "TokenRescued", log); err != nil {
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
func (_AcrossSpoke *AcrossSpokeFilterer) ParseTokenRescued(log types.Log) (*AcrossSpokeTokenRescued, error) {
	event := new(AcrossSpokeTokenRescued)
	if err := _AcrossSpoke.contract.UnpackLog(event, "TokenRescued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
