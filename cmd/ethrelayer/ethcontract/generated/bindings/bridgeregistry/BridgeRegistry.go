// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeRegistry

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeRegistryABI is the input ABI used to generate the binding from.
const BridgeRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"evrnetBridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_evrnetBridge\",\"type\":\"address\"},{\"name\":\"_bridgeBank\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_valset\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_evrnetBridge\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_bridgeBank\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_oracle\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_valset\",\"type\":\"address\"}],\"name\":\"LogContractsRegistered\",\"type\":\"event\"}]"

// BridgeRegistryBin is the compiled bytecode used for deploying new contracts.
var BridgeRegistryBin = "0x608060405234801561001057600080fd5b5060405160808061057b8339810180604052608081101561003057600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f039b733f31259b106f1d278c726870d5b28c7db22957d63df8dbaa70bd3a032a6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390a150505050610280806102fb6000396000f3fe60806040526004361061005c576000357c0100000000000000000000000000000000000000000000000000000000900480630e41f373146100615780637dc0d1d0146100b85780637f54af0c1461010f578063e48e227d14610166575b600080fd5b34801561006d57600080fd5b506100766101bd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100c457600080fd5b506100cd6101e3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561011b57600080fd5b50610124610209565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017257600080fd5b5061017b61022f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a72305820a8d4f345d755d5195a20144d9a364f67b3c566f76ee9f92e48699ad417a7ba880029"

// DeployBridgeRegistry deploys a new Ethereum contract, binding an instance of BridgeRegistry to it.
func DeployBridgeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _evrnetBridge common.Address, _bridgeBank common.Address, _oracle common.Address, _valset common.Address) (common.Address, *types.Transaction, *BridgeRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeRegistryBin), backend, _evrnetBridge, _bridgeBank, _oracle, _valset)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeRegistry{BridgeRegistryCaller: BridgeRegistryCaller{contract: contract}, BridgeRegistryTransactor: BridgeRegistryTransactor{contract: contract}, BridgeRegistryFilterer: BridgeRegistryFilterer{contract: contract}}, nil
}

// BridgeRegistry is an auto generated Go binding around an Ethereum contract.
type BridgeRegistry struct {
	BridgeRegistryCaller     // Read-only binding to the contract
	BridgeRegistryTransactor // Write-only binding to the contract
	BridgeRegistryFilterer   // Log filterer for contract events
}

// BridgeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeRegistrySession struct {
	Contract     *BridgeRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeRegistryCallerSession struct {
	Contract *BridgeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeRegistryTransactorSession struct {
	Contract     *BridgeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRegistryRaw struct {
	Contract *BridgeRegistry // Generic contract binding to access the raw methods on
}

// BridgeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeRegistryCallerRaw struct {
	Contract *BridgeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeRegistryTransactorRaw struct {
	Contract *BridgeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeRegistry creates a new instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistry(address common.Address, backend bind.ContractBackend) (*BridgeRegistry, error) {
	contract, err := bindBridgeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistry{BridgeRegistryCaller: BridgeRegistryCaller{contract: contract}, BridgeRegistryTransactor: BridgeRegistryTransactor{contract: contract}, BridgeRegistryFilterer: BridgeRegistryFilterer{contract: contract}}, nil
}

// NewBridgeRegistryCaller creates a new read-only instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryCaller(address common.Address, caller bind.ContractCaller) (*BridgeRegistryCaller, error) {
	contract, err := bindBridgeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryCaller{contract: contract}, nil
}

// NewBridgeRegistryTransactor creates a new write-only instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeRegistryTransactor, error) {
	contract, err := bindBridgeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryTransactor{contract: contract}, nil
}

// NewBridgeRegistryFilterer creates a new log filterer instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeRegistryFilterer, error) {
	contract, err := bindBridgeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryFilterer{contract: contract}, nil
}

// bindBridgeRegistry binds a generic wrapper to an already deployed contract.
func bindBridgeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRegistry *BridgeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeRegistry.Contract.BridgeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRegistry *BridgeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.BridgeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRegistry *BridgeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.BridgeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRegistry *BridgeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRegistry *BridgeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRegistry *BridgeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.contract.Transact(opts, method, params...)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "bridgeBank")
	return *ret0, err
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) BridgeBank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeBank(&_BridgeRegistry.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) BridgeBank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeBank(&_BridgeRegistry.CallOpts)
}

// EvrnetBridge is a free data retrieval call binding the contract method 0xe48e227d.
//
// Solidity: function evrnetBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) EvrnetBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "evrnetBridge")
	return *ret0, err
}

// EvrnetBridge is a free data retrieval call binding the contract method 0xe48e227d.
//
// Solidity: function evrnetBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) EvrnetBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.EvrnetBridge(&_BridgeRegistry.CallOpts)
}

// EvrnetBridge is a free data retrieval call binding the contract method 0xe48e227d.
//
// Solidity: function evrnetBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) EvrnetBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.EvrnetBridge(&_BridgeRegistry.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) Oracle() (common.Address, error) {
	return _BridgeRegistry.Contract.Oracle(&_BridgeRegistry.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) Oracle() (common.Address, error) {
	return _BridgeRegistry.Contract.Oracle(&_BridgeRegistry.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) Valset() (common.Address, error) {
	return _BridgeRegistry.Contract.Valset(&_BridgeRegistry.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) Valset() (common.Address, error) {
	return _BridgeRegistry.Contract.Valset(&_BridgeRegistry.CallOpts)
}

// BridgeRegistryLogContractsRegisteredIterator is returned from FilterLogContractsRegistered and is used to iterate over the raw logs and unpacked data for LogContractsRegistered events raised by the BridgeRegistry contract.
type BridgeRegistryLogContractsRegisteredIterator struct {
	Event *BridgeRegistryLogContractsRegistered // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryLogContractsRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryLogContractsRegistered)
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
		it.Event = new(BridgeRegistryLogContractsRegistered)
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
func (it *BridgeRegistryLogContractsRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryLogContractsRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryLogContractsRegistered represents a LogContractsRegistered event raised by the BridgeRegistry contract.
type BridgeRegistryLogContractsRegistered struct {
	EvrnetBridge common.Address
	BridgeBank   common.Address
	Oracle       common.Address
	Valset       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogContractsRegistered is a free log retrieval operation binding the contract event 0x039b733f31259b106f1d278c726870d5b28c7db22957d63df8dbaa70bd3a032a.
//
// Solidity: event LogContractsRegistered(address _evrnetBridge, address _bridgeBank, address _oracle, address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterLogContractsRegistered(opts *bind.FilterOpts) (*BridgeRegistryLogContractsRegisteredIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "LogContractsRegistered")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryLogContractsRegisteredIterator{contract: _BridgeRegistry.contract, event: "LogContractsRegistered", logs: logs, sub: sub}, nil
}

// WatchLogContractsRegistered is a free log subscription operation binding the contract event 0x039b733f31259b106f1d278c726870d5b28c7db22957d63df8dbaa70bd3a032a.
//
// Solidity: event LogContractsRegistered(address _evrnetBridge, address _bridgeBank, address _oracle, address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchLogContractsRegistered(opts *bind.WatchOpts, sink chan<- *BridgeRegistryLogContractsRegistered) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "LogContractsRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryLogContractsRegistered)
				if err := _BridgeRegistry.contract.UnpackLog(event, "LogContractsRegistered", log); err != nil {
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

// ParseLogContractsRegistered is a log parse operation binding the contract event 0x039b733f31259b106f1d278c726870d5b28c7db22957d63df8dbaa70bd3a032a.
//
// Solidity: event LogContractsRegistered(address _evrnetBridge, address _bridgeBank, address _oracle, address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseLogContractsRegistered(log types.Log) (*BridgeRegistryLogContractsRegistered, error) {
	event := new(BridgeRegistryLogContractsRegistered)
	if err := _BridgeRegistry.contract.UnpackLog(event, "LogContractsRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
