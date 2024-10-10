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
)

// BrevisProofData is an auto generated low-level Go binding around an user-defined struct.
type BrevisProofData struct {
	CommitHash    [32]byte
	VkHash        [32]byte
	AppCommitHash [32]byte
	AppVkHash     [32]byte
	SmtRoot       [32]byte
}

// ZkRebateMetaData contains all meta data concerning the ZkRebate contract.
var ZkRebateMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_brv\",\"type\":\"address\",\"internalType\":\"contractIBrevisProof\"},{\"name\":\"_vkHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_uni\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPool\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"brvProof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBrevisProof\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimWithZkProofs\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_proofIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proofDataArray\",\"type\":\"tuple[]\",\"internalType\":\"structBrevis.ProofData[]\",\"components\":[{\"name\":\"commitHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vkHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appCommitHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appVkHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"smtRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_appCircuitOutputs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vkHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"}]",
}

// ZkRebateABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkRebateMetaData.ABI instead.
var ZkRebateABI = ZkRebateMetaData.ABI

// ZkRebate is an auto generated Go binding around an Ethereum contract.
type ZkRebate struct {
	ZkRebateCaller     // Read-only binding to the contract
	ZkRebateTransactor // Write-only binding to the contract
	ZkRebateFilterer   // Log filterer for contract events
}

// ZkRebateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkRebateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkRebateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkRebateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkRebateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkRebateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkRebateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkRebateSession struct {
	Contract     *ZkRebate         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkRebateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkRebateCallerSession struct {
	Contract *ZkRebateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZkRebateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkRebateTransactorSession struct {
	Contract     *ZkRebateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZkRebateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkRebateRaw struct {
	Contract *ZkRebate // Generic contract binding to access the raw methods on
}

// ZkRebateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkRebateCallerRaw struct {
	Contract *ZkRebateCaller // Generic read-only contract binding to access the raw methods on
}

// ZkRebateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkRebateTransactorRaw struct {
	Contract *ZkRebateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkRebate creates a new instance of ZkRebate, bound to a specific deployed contract.
func NewZkRebate(address common.Address, backend bind.ContractBackend) (*ZkRebate, error) {
	contract, err := bindZkRebate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkRebate{ZkRebateCaller: ZkRebateCaller{contract: contract}, ZkRebateTransactor: ZkRebateTransactor{contract: contract}, ZkRebateFilterer: ZkRebateFilterer{contract: contract}}, nil
}

// NewZkRebateCaller creates a new read-only instance of ZkRebate, bound to a specific deployed contract.
func NewZkRebateCaller(address common.Address, caller bind.ContractCaller) (*ZkRebateCaller, error) {
	contract, err := bindZkRebate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkRebateCaller{contract: contract}, nil
}

// NewZkRebateTransactor creates a new write-only instance of ZkRebate, bound to a specific deployed contract.
func NewZkRebateTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkRebateTransactor, error) {
	contract, err := bindZkRebate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkRebateTransactor{contract: contract}, nil
}

// NewZkRebateFilterer creates a new log filterer instance of ZkRebate, bound to a specific deployed contract.
func NewZkRebateFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkRebateFilterer, error) {
	contract, err := bindZkRebate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkRebateFilterer{contract: contract}, nil
}

// bindZkRebate binds a generic wrapper to an already deployed contract.
func bindZkRebate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZkRebateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkRebate *ZkRebateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkRebate.Contract.ZkRebateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkRebate *ZkRebateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkRebate.Contract.ZkRebateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkRebate *ZkRebateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkRebate.Contract.ZkRebateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkRebate *ZkRebateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkRebate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkRebate *ZkRebateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkRebate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkRebate *ZkRebateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkRebate.Contract.contract.Transact(opts, method, params...)
}

// BrvProof is a free data retrieval call binding the contract method 0x80d91692.
//
// Solidity: function brvProof() view returns(address)
func (_ZkRebate *ZkRebateCaller) BrvProof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkRebate.contract.Call(opts, &out, "brvProof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BrvProof is a free data retrieval call binding the contract method 0x80d91692.
//
// Solidity: function brvProof() view returns(address)
func (_ZkRebate *ZkRebateSession) BrvProof() (common.Address, error) {
	return _ZkRebate.Contract.BrvProof(&_ZkRebate.CallOpts)
}

// BrvProof is a free data retrieval call binding the contract method 0x80d91692.
//
// Solidity: function brvProof() view returns(address)
func (_ZkRebate *ZkRebateCallerSession) BrvProof() (common.Address, error) {
	return _ZkRebate.Contract.BrvProof(&_ZkRebate.CallOpts)
}

// Uni is a free data retrieval call binding the contract method 0xedc9af95.
//
// Solidity: function uni() view returns(address)
func (_ZkRebate *ZkRebateCaller) Uni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkRebate.contract.Call(opts, &out, "uni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uni is a free data retrieval call binding the contract method 0xedc9af95.
//
// Solidity: function uni() view returns(address)
func (_ZkRebate *ZkRebateSession) Uni() (common.Address, error) {
	return _ZkRebate.Contract.Uni(&_ZkRebate.CallOpts)
}

// Uni is a free data retrieval call binding the contract method 0xedc9af95.
//
// Solidity: function uni() view returns(address)
func (_ZkRebate *ZkRebateCallerSession) Uni() (common.Address, error) {
	return _ZkRebate.Contract.Uni(&_ZkRebate.CallOpts)
}

// VkHash is a free data retrieval call binding the contract method 0x4fe840f5.
//
// Solidity: function vkHash() view returns(bytes32)
func (_ZkRebate *ZkRebateCaller) VkHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkRebate.contract.Call(opts, &out, "vkHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VkHash is a free data retrieval call binding the contract method 0x4fe840f5.
//
// Solidity: function vkHash() view returns(bytes32)
func (_ZkRebate *ZkRebateSession) VkHash() ([32]byte, error) {
	return _ZkRebate.Contract.VkHash(&_ZkRebate.CallOpts)
}

// VkHash is a free data retrieval call binding the contract method 0x4fe840f5.
//
// Solidity: function vkHash() view returns(bytes32)
func (_ZkRebate *ZkRebateCallerSession) VkHash() ([32]byte, error) {
	return _ZkRebate.Contract.VkHash(&_ZkRebate.CallOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0x5bef1040.
//
// Solidity: function addPool((address,address,uint24,int24,address) key) returns()
func (_ZkRebate *ZkRebateTransactor) AddPool(opts *bind.TransactOpts, key PoolKey) (*types.Transaction, error) {
	return _ZkRebate.contract.Transact(opts, "addPool", key)
}

// AddPool is a paid mutator transaction binding the contract method 0x5bef1040.
//
// Solidity: function addPool((address,address,uint24,int24,address) key) returns()
func (_ZkRebate *ZkRebateSession) AddPool(key PoolKey) (*types.Transaction, error) {
	return _ZkRebate.Contract.AddPool(&_ZkRebate.TransactOpts, key)
}

// AddPool is a paid mutator transaction binding the contract method 0x5bef1040.
//
// Solidity: function addPool((address,address,uint24,int24,address) key) returns()
func (_ZkRebate *ZkRebateTransactorSession) AddPool(key PoolKey) (*types.Transaction, error) {
	return _ZkRebate.Contract.AddPool(&_ZkRebate.TransactOpts, key)
}

// ClaimWithZkProofs is a paid mutator transaction binding the contract method 0x27c08d01.
//
// Solidity: function claimWithZkProofs(address receiver, bytes32[] _proofIds, bytes _proof, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs) returns()
func (_ZkRebate *ZkRebateTransactor) ClaimWithZkProofs(opts *bind.TransactOpts, receiver common.Address, _proofIds [][32]byte, _proof []byte, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte) (*types.Transaction, error) {
	return _ZkRebate.contract.Transact(opts, "claimWithZkProofs", receiver, _proofIds, _proof, _proofDataArray, _appCircuitOutputs)
}

// ClaimWithZkProofs is a paid mutator transaction binding the contract method 0x27c08d01.
//
// Solidity: function claimWithZkProofs(address receiver, bytes32[] _proofIds, bytes _proof, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs) returns()
func (_ZkRebate *ZkRebateSession) ClaimWithZkProofs(receiver common.Address, _proofIds [][32]byte, _proof []byte, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte) (*types.Transaction, error) {
	return _ZkRebate.Contract.ClaimWithZkProofs(&_ZkRebate.TransactOpts, receiver, _proofIds, _proof, _proofDataArray, _appCircuitOutputs)
}

// ClaimWithZkProofs is a paid mutator transaction binding the contract method 0x27c08d01.
//
// Solidity: function claimWithZkProofs(address receiver, bytes32[] _proofIds, bytes _proof, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs) returns()
func (_ZkRebate *ZkRebateTransactorSession) ClaimWithZkProofs(receiver common.Address, _proofIds [][32]byte, _proof []byte, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte) (*types.Transaction, error) {
	return _ZkRebate.Contract.ClaimWithZkProofs(&_ZkRebate.TransactOpts, receiver, _proofIds, _proof, _proofDataArray, _appCircuitOutputs)
}
