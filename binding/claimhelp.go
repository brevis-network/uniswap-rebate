// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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

// ClaimHelpMetaData contains all meta data concerning the ClaimHelp contract.
var ClaimHelpMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Claimer\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ClaimHelpABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimHelpMetaData.ABI instead.
var ClaimHelpABI = ClaimHelpMetaData.ABI

// ClaimHelp is an auto generated Go binding around an Ethereum contract.
type ClaimHelp struct {
	ClaimHelpCaller     // Read-only binding to the contract
	ClaimHelpTransactor // Write-only binding to the contract
	ClaimHelpFilterer   // Log filterer for contract events
}

// ClaimHelpCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimHelpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHelpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimHelpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHelpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimHelpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHelpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimHelpSession struct {
	Contract     *ClaimHelp        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimHelpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimHelpCallerSession struct {
	Contract *ClaimHelpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ClaimHelpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimHelpTransactorSession struct {
	Contract     *ClaimHelpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ClaimHelpRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimHelpRaw struct {
	Contract *ClaimHelp // Generic contract binding to access the raw methods on
}

// ClaimHelpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimHelpCallerRaw struct {
	Contract *ClaimHelpCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimHelpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimHelpTransactorRaw struct {
	Contract *ClaimHelpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimHelp creates a new instance of ClaimHelp, bound to a specific deployed contract.
func NewClaimHelp(address common.Address, backend bind.ContractBackend) (*ClaimHelp, error) {
	contract, err := bindClaimHelp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimHelp{ClaimHelpCaller: ClaimHelpCaller{contract: contract}, ClaimHelpTransactor: ClaimHelpTransactor{contract: contract}, ClaimHelpFilterer: ClaimHelpFilterer{contract: contract}}, nil
}

// NewClaimHelpCaller creates a new read-only instance of ClaimHelp, bound to a specific deployed contract.
func NewClaimHelpCaller(address common.Address, caller bind.ContractCaller) (*ClaimHelpCaller, error) {
	contract, err := bindClaimHelp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimHelpCaller{contract: contract}, nil
}

// NewClaimHelpTransactor creates a new write-only instance of ClaimHelp, bound to a specific deployed contract.
func NewClaimHelpTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimHelpTransactor, error) {
	contract, err := bindClaimHelp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimHelpTransactor{contract: contract}, nil
}

// NewClaimHelpFilterer creates a new log filterer instance of ClaimHelp, bound to a specific deployed contract.
func NewClaimHelpFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimHelpFilterer, error) {
	contract, err := bindClaimHelp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimHelpFilterer{contract: contract}, nil
}

// bindClaimHelp binds a generic wrapper to an already deployed contract.
func bindClaimHelp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimHelpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimHelp *ClaimHelpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimHelp.Contract.ClaimHelpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimHelp *ClaimHelpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimHelp.Contract.ClaimHelpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimHelp *ClaimHelpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimHelp.Contract.ClaimHelpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimHelp *ClaimHelpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimHelp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimHelp *ClaimHelpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimHelp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimHelp *ClaimHelpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimHelp.Contract.contract.Transact(opts, method, params...)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address router) returns()
func (_ClaimHelp *ClaimHelpTransactor) Claim(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _ClaimHelp.contract.Transact(opts, "claim", router)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address router) returns()
func (_ClaimHelp *ClaimHelpSession) Claim(router common.Address) (*types.Transaction, error) {
	return _ClaimHelp.Contract.Claim(&_ClaimHelp.TransactOpts, router)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address router) returns()
func (_ClaimHelp *ClaimHelpTransactorSession) Claim(router common.Address) (*types.Transaction, error) {
	return _ClaimHelp.Contract.Claim(&_ClaimHelp.TransactOpts, router)
}

// ClaimHelpClaimerIterator is returned from FilterClaimer and is used to iterate over the raw logs and unpacked data for Claimer events raised by the ClaimHelp contract.
type ClaimHelpClaimerIterator struct {
	Event *ClaimHelpClaimer // Event containing the contract specifics and raw log

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
func (it *ClaimHelpClaimerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHelpClaimer)
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
		it.Event = new(ClaimHelpClaimer)
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
func (it *ClaimHelpClaimerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHelpClaimerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHelpClaimer represents a Claimer event raised by the ClaimHelp contract.
type ClaimHelpClaimer struct {
	Router  common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimer is a free log retrieval operation binding the contract event 0xf0d796bb38c321bf748f9334d1b7b16ba5fb79e2112396aa77c47cd5d21a8b2f.
//
// Solidity: event Claimer(address indexed router, address indexed claimer)
func (_ClaimHelp *ClaimHelpFilterer) FilterClaimer(opts *bind.FilterOpts, router []common.Address, claimer []common.Address) (*ClaimHelpClaimerIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ClaimHelp.contract.FilterLogs(opts, "Claimer", routerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHelpClaimerIterator{contract: _ClaimHelp.contract, event: "Claimer", logs: logs, sub: sub}, nil
}

// WatchClaimer is a free log subscription operation binding the contract event 0xf0d796bb38c321bf748f9334d1b7b16ba5fb79e2112396aa77c47cd5d21a8b2f.
//
// Solidity: event Claimer(address indexed router, address indexed claimer)
func (_ClaimHelp *ClaimHelpFilterer) WatchClaimer(opts *bind.WatchOpts, sink chan<- *ClaimHelpClaimer, router []common.Address, claimer []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ClaimHelp.contract.WatchLogs(opts, "Claimer", routerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHelpClaimer)
				if err := _ClaimHelp.contract.UnpackLog(event, "Claimer", log); err != nil {
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

// ParseClaimer is a log parse operation binding the contract event 0xf0d796bb38c321bf748f9334d1b7b16ba5fb79e2112396aa77c47cd5d21a8b2f.
//
// Solidity: event Claimer(address indexed router, address indexed claimer)
func (_ClaimHelp *ClaimHelpFilterer) ParseClaimer(log types.Log) (*ClaimHelpClaimer, error) {
	event := new(ClaimHelpClaimer)
	if err := _ClaimHelp.contract.UnpackLog(event, "Claimer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
