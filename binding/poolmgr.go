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

// IPoolManagerModifyLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerModifyLiquidityParams struct {
	TickLower      *big.Int
	TickUpper      *big.Int
	LiquidityDelta *big.Int
	Salt           [32]byte
}

// IPoolManagerSwapParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerSwapParams struct {
	ZeroForOne        bool
	AmountSpecified   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// PoolMgrMetaData contains all meta data concerning the PoolMgr contract.
var PoolMgrMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clear\",\"inputs\":[{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collectProtocolFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountCollected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"donate\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extsload\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extsload\",\"inputs\":[{\"name\":\"startSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nSlots\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"values\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extsload\",\"inputs\":[{\"name\":\"slots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"values\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exttload\",\"inputs\":[{\"name\":\"slots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"values\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exttload\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyLiquidity\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"callerDelta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"feesAccrued\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"protocolFeeController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProtocolFeeController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFeesAccrued\",\"inputs\":[{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"Currency\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFee\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"newProtocolFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFeeController\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"contractIProtocolFeeController\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[],\"outputs\":[{\"name\":\"paid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"settleFor\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"paid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"swapDelta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sync\",\"inputs\":[{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"Currency\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"take\",\"inputs\":[{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlock\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDynamicLPFee\",\"inputs\":[{\"name\":\"key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"newDynamicLPFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Donate\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialize\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"currency0\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIHooks\"},{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModifyLiquidity\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSet\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolFeeControllerUpdated\",\"inputs\":[{\"name\":\"protocolFeeController\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolFeeUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"protocolFee\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"int128\",\"indexed\":false,\"internalType\":\"int128\"},{\"name\":\"amount1\",\"type\":\"int128\",\"indexed\":false,\"internalType\":\"int128\"},{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"liquidity\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"},{\"name\":\"fee\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyUnlocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContractUnlocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrenciesOutOfOrderOrEqual\",\"inputs\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CurrencyNotSettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManagerLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustClearExactPositiveDelta\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonzeroNativeValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolNotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProtocolFeeCannotBeFetched\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProtocolFeeTooLarge\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}]},{\"type\":\"error\",\"name\":\"SwapAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TickSpacingTooLarge\",\"inputs\":[{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"}]},{\"type\":\"error\",\"name\":\"TickSpacingTooSmall\",\"inputs\":[{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedDynamicLPFeeUpdate\",\"inputs\":[]}]",
}

// PoolMgrABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMgrMetaData.ABI instead.
var PoolMgrABI = PoolMgrMetaData.ABI

// PoolMgr is an auto generated Go binding around an Ethereum contract.
type PoolMgr struct {
	PoolMgrCaller     // Read-only binding to the contract
	PoolMgrTransactor // Write-only binding to the contract
	PoolMgrFilterer   // Log filterer for contract events
}

// PoolMgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolMgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolMgrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolMgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolMgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolMgrSession struct {
	Contract     *PoolMgr          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolMgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolMgrCallerSession struct {
	Contract *PoolMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PoolMgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolMgrTransactorSession struct {
	Contract     *PoolMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PoolMgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolMgrRaw struct {
	Contract *PoolMgr // Generic contract binding to access the raw methods on
}

// PoolMgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolMgrCallerRaw struct {
	Contract *PoolMgrCaller // Generic read-only contract binding to access the raw methods on
}

// PoolMgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolMgrTransactorRaw struct {
	Contract *PoolMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolMgr creates a new instance of PoolMgr, bound to a specific deployed contract.
func NewPoolMgr(address common.Address, backend bind.ContractBackend) (*PoolMgr, error) {
	contract, err := bindPoolMgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolMgr{PoolMgrCaller: PoolMgrCaller{contract: contract}, PoolMgrTransactor: PoolMgrTransactor{contract: contract}, PoolMgrFilterer: PoolMgrFilterer{contract: contract}}, nil
}

// NewPoolMgrCaller creates a new read-only instance of PoolMgr, bound to a specific deployed contract.
func NewPoolMgrCaller(address common.Address, caller bind.ContractCaller) (*PoolMgrCaller, error) {
	contract, err := bindPoolMgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolMgrCaller{contract: contract}, nil
}

// NewPoolMgrTransactor creates a new write-only instance of PoolMgr, bound to a specific deployed contract.
func NewPoolMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolMgrTransactor, error) {
	contract, err := bindPoolMgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolMgrTransactor{contract: contract}, nil
}

// NewPoolMgrFilterer creates a new log filterer instance of PoolMgr, bound to a specific deployed contract.
func NewPoolMgrFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolMgrFilterer, error) {
	contract, err := bindPoolMgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolMgrFilterer{contract: contract}, nil
}

// bindPoolMgr binds a generic wrapper to an already deployed contract.
func bindPoolMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolMgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolMgr *PoolMgrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolMgr.Contract.PoolMgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolMgr *PoolMgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolMgr.Contract.PoolMgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolMgr *PoolMgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolMgr.Contract.PoolMgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolMgr *PoolMgrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolMgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolMgr *PoolMgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolMgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolMgr *PoolMgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolMgr.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256 amount)
func (_PoolMgr *PoolMgrCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "allowance", owner, spender, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256 amount)
func (_PoolMgr *PoolMgrSession) Allowance(owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	return _PoolMgr.Contract.Allowance(&_PoolMgr.CallOpts, owner, spender, id)
}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256 amount)
func (_PoolMgr *PoolMgrCallerSession) Allowance(owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	return _PoolMgr.Contract.Allowance(&_PoolMgr.CallOpts, owner, spender, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 amount)
func (_PoolMgr *PoolMgrCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 amount)
func (_PoolMgr *PoolMgrSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _PoolMgr.Contract.BalanceOf(&_PoolMgr.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 amount)
func (_PoolMgr *PoolMgrCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _PoolMgr.Contract.BalanceOf(&_PoolMgr.CallOpts, owner, id)
}

// Extsload is a free data retrieval call binding the contract method 0x1e2eaeaf.
//
// Solidity: function extsload(bytes32 slot) view returns(bytes32 value)
func (_PoolMgr *PoolMgrCaller) Extsload(opts *bind.CallOpts, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "extsload", slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Extsload is a free data retrieval call binding the contract method 0x1e2eaeaf.
//
// Solidity: function extsload(bytes32 slot) view returns(bytes32 value)
func (_PoolMgr *PoolMgrSession) Extsload(slot [32]byte) ([32]byte, error) {
	return _PoolMgr.Contract.Extsload(&_PoolMgr.CallOpts, slot)
}

// Extsload is a free data retrieval call binding the contract method 0x1e2eaeaf.
//
// Solidity: function extsload(bytes32 slot) view returns(bytes32 value)
func (_PoolMgr *PoolMgrCallerSession) Extsload(slot [32]byte) ([32]byte, error) {
	return _PoolMgr.Contract.Extsload(&_PoolMgr.CallOpts, slot)
}

// Extsload0 is a free data retrieval call binding the contract method 0x35fd631a.
//
// Solidity: function extsload(bytes32 startSlot, uint256 nSlots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrCaller) Extsload0(opts *bind.CallOpts, startSlot [32]byte, nSlots *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "extsload0", startSlot, nSlots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Extsload0 is a free data retrieval call binding the contract method 0x35fd631a.
//
// Solidity: function extsload(bytes32 startSlot, uint256 nSlots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrSession) Extsload0(startSlot [32]byte, nSlots *big.Int) ([][32]byte, error) {
	return _PoolMgr.Contract.Extsload0(&_PoolMgr.CallOpts, startSlot, nSlots)
}

// Extsload0 is a free data retrieval call binding the contract method 0x35fd631a.
//
// Solidity: function extsload(bytes32 startSlot, uint256 nSlots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrCallerSession) Extsload0(startSlot [32]byte, nSlots *big.Int) ([][32]byte, error) {
	return _PoolMgr.Contract.Extsload0(&_PoolMgr.CallOpts, startSlot, nSlots)
}

// Extsload1 is a free data retrieval call binding the contract method 0xdbd035ff.
//
// Solidity: function extsload(bytes32[] slots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrCaller) Extsload1(opts *bind.CallOpts, slots [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "extsload1", slots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Extsload1 is a free data retrieval call binding the contract method 0xdbd035ff.
//
// Solidity: function extsload(bytes32[] slots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrSession) Extsload1(slots [][32]byte) ([][32]byte, error) {
	return _PoolMgr.Contract.Extsload1(&_PoolMgr.CallOpts, slots)
}

// Extsload1 is a free data retrieval call binding the contract method 0xdbd035ff.
//
// Solidity: function extsload(bytes32[] slots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrCallerSession) Extsload1(slots [][32]byte) ([][32]byte, error) {
	return _PoolMgr.Contract.Extsload1(&_PoolMgr.CallOpts, slots)
}

// Exttload is a free data retrieval call binding the contract method 0x9bf6645f.
//
// Solidity: function exttload(bytes32[] slots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrCaller) Exttload(opts *bind.CallOpts, slots [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "exttload", slots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Exttload is a free data retrieval call binding the contract method 0x9bf6645f.
//
// Solidity: function exttload(bytes32[] slots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrSession) Exttload(slots [][32]byte) ([][32]byte, error) {
	return _PoolMgr.Contract.Exttload(&_PoolMgr.CallOpts, slots)
}

// Exttload is a free data retrieval call binding the contract method 0x9bf6645f.
//
// Solidity: function exttload(bytes32[] slots) view returns(bytes32[] values)
func (_PoolMgr *PoolMgrCallerSession) Exttload(slots [][32]byte) ([][32]byte, error) {
	return _PoolMgr.Contract.Exttload(&_PoolMgr.CallOpts, slots)
}

// Exttload0 is a free data retrieval call binding the contract method 0xf135baaa.
//
// Solidity: function exttload(bytes32 slot) view returns(bytes32 value)
func (_PoolMgr *PoolMgrCaller) Exttload0(opts *bind.CallOpts, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "exttload0", slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Exttload0 is a free data retrieval call binding the contract method 0xf135baaa.
//
// Solidity: function exttload(bytes32 slot) view returns(bytes32 value)
func (_PoolMgr *PoolMgrSession) Exttload0(slot [32]byte) ([32]byte, error) {
	return _PoolMgr.Contract.Exttload0(&_PoolMgr.CallOpts, slot)
}

// Exttload0 is a free data retrieval call binding the contract method 0xf135baaa.
//
// Solidity: function exttload(bytes32 slot) view returns(bytes32 value)
func (_PoolMgr *PoolMgrCallerSession) Exttload0(slot [32]byte) ([32]byte, error) {
	return _PoolMgr.Contract.Exttload0(&_PoolMgr.CallOpts, slot)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address spender) view returns(bool approved)
func (_PoolMgr *PoolMgrCaller) IsOperator(opts *bind.CallOpts, owner common.Address, spender common.Address) (bool, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "isOperator", owner, spender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address spender) view returns(bool approved)
func (_PoolMgr *PoolMgrSession) IsOperator(owner common.Address, spender common.Address) (bool, error) {
	return _PoolMgr.Contract.IsOperator(&_PoolMgr.CallOpts, owner, spender)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address spender) view returns(bool approved)
func (_PoolMgr *PoolMgrCallerSession) IsOperator(owner common.Address, spender common.Address) (bool, error) {
	return _PoolMgr.Contract.IsOperator(&_PoolMgr.CallOpts, owner, spender)
}

// ProtocolFeeController is a free data retrieval call binding the contract method 0xf02de3b2.
//
// Solidity: function protocolFeeController() view returns(address)
func (_PoolMgr *PoolMgrCaller) ProtocolFeeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "protocolFeeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeController is a free data retrieval call binding the contract method 0xf02de3b2.
//
// Solidity: function protocolFeeController() view returns(address)
func (_PoolMgr *PoolMgrSession) ProtocolFeeController() (common.Address, error) {
	return _PoolMgr.Contract.ProtocolFeeController(&_PoolMgr.CallOpts)
}

// ProtocolFeeController is a free data retrieval call binding the contract method 0xf02de3b2.
//
// Solidity: function protocolFeeController() view returns(address)
func (_PoolMgr *PoolMgrCallerSession) ProtocolFeeController() (common.Address, error) {
	return _PoolMgr.Contract.ProtocolFeeController(&_PoolMgr.CallOpts)
}

// ProtocolFeesAccrued is a free data retrieval call binding the contract method 0x97e8cd4e.
//
// Solidity: function protocolFeesAccrued(address currency) view returns(uint256 amount)
func (_PoolMgr *PoolMgrCaller) ProtocolFeesAccrued(opts *bind.CallOpts, currency common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolMgr.contract.Call(opts, &out, "protocolFeesAccrued", currency)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeesAccrued is a free data retrieval call binding the contract method 0x97e8cd4e.
//
// Solidity: function protocolFeesAccrued(address currency) view returns(uint256 amount)
func (_PoolMgr *PoolMgrSession) ProtocolFeesAccrued(currency common.Address) (*big.Int, error) {
	return _PoolMgr.Contract.ProtocolFeesAccrued(&_PoolMgr.CallOpts, currency)
}

// ProtocolFeesAccrued is a free data retrieval call binding the contract method 0x97e8cd4e.
//
// Solidity: function protocolFeesAccrued(address currency) view returns(uint256 amount)
func (_PoolMgr *PoolMgrCallerSession) ProtocolFeesAccrued(currency common.Address) (*big.Int, error) {
	return _PoolMgr.Contract.ProtocolFeesAccrued(&_PoolMgr.CallOpts, currency)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "approve", spender, id, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrSession) Approve(spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Approve(&_PoolMgr.TransactOpts, spender, id, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrTransactorSession) Approve(spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Approve(&_PoolMgr.TransactOpts, spender, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactor) Burn(opts *bind.TransactOpts, from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "burn", from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_PoolMgr *PoolMgrSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Burn(&_PoolMgr.TransactOpts, from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactorSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Burn(&_PoolMgr.TransactOpts, from, id, amount)
}

// Clear is a paid mutator transaction binding the contract method 0x80f0b44c.
//
// Solidity: function clear(address currency, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactor) Clear(opts *bind.TransactOpts, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "clear", currency, amount)
}

// Clear is a paid mutator transaction binding the contract method 0x80f0b44c.
//
// Solidity: function clear(address currency, uint256 amount) returns()
func (_PoolMgr *PoolMgrSession) Clear(currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Clear(&_PoolMgr.TransactOpts, currency, amount)
}

// Clear is a paid mutator transaction binding the contract method 0x80f0b44c.
//
// Solidity: function clear(address currency, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactorSession) Clear(currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Clear(&_PoolMgr.TransactOpts, currency, amount)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0x8161b874.
//
// Solidity: function collectProtocolFees(address recipient, address currency, uint256 amount) returns(uint256 amountCollected)
func (_PoolMgr *PoolMgrTransactor) CollectProtocolFees(opts *bind.TransactOpts, recipient common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "collectProtocolFees", recipient, currency, amount)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0x8161b874.
//
// Solidity: function collectProtocolFees(address recipient, address currency, uint256 amount) returns(uint256 amountCollected)
func (_PoolMgr *PoolMgrSession) CollectProtocolFees(recipient common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.CollectProtocolFees(&_PoolMgr.TransactOpts, recipient, currency, amount)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0x8161b874.
//
// Solidity: function collectProtocolFees(address recipient, address currency, uint256 amount) returns(uint256 amountCollected)
func (_PoolMgr *PoolMgrTransactorSession) CollectProtocolFees(recipient common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.CollectProtocolFees(&_PoolMgr.TransactOpts, recipient, currency, amount)
}

// Donate is a paid mutator transaction binding the contract method 0x234266d7.
//
// Solidity: function donate((address,address,uint24,int24,address) key, uint256 amount0, uint256 amount1, bytes hookData) returns(int256)
func (_PoolMgr *PoolMgrTransactor) Donate(opts *bind.TransactOpts, key PoolKey, amount0 *big.Int, amount1 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "donate", key, amount0, amount1, hookData)
}

// Donate is a paid mutator transaction binding the contract method 0x234266d7.
//
// Solidity: function donate((address,address,uint24,int24,address) key, uint256 amount0, uint256 amount1, bytes hookData) returns(int256)
func (_PoolMgr *PoolMgrSession) Donate(key PoolKey, amount0 *big.Int, amount1 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Donate(&_PoolMgr.TransactOpts, key, amount0, amount1, hookData)
}

// Donate is a paid mutator transaction binding the contract method 0x234266d7.
//
// Solidity: function donate((address,address,uint24,int24,address) key, uint256 amount0, uint256 amount1, bytes hookData) returns(int256)
func (_PoolMgr *PoolMgrTransactorSession) Donate(key PoolKey, amount0 *big.Int, amount1 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Donate(&_PoolMgr.TransactOpts, key, amount0, amount1, hookData)
}

// Initialize is a paid mutator transaction binding the contract method 0x695c5bf5.
//
// Solidity: function initialize((address,address,uint24,int24,address) key, uint160 sqrtPriceX96, bytes hookData) returns(int24 tick)
func (_PoolMgr *PoolMgrTransactor) Initialize(opts *bind.TransactOpts, key PoolKey, sqrtPriceX96 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "initialize", key, sqrtPriceX96, hookData)
}

// Initialize is a paid mutator transaction binding the contract method 0x695c5bf5.
//
// Solidity: function initialize((address,address,uint24,int24,address) key, uint160 sqrtPriceX96, bytes hookData) returns(int24 tick)
func (_PoolMgr *PoolMgrSession) Initialize(key PoolKey, sqrtPriceX96 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Initialize(&_PoolMgr.TransactOpts, key, sqrtPriceX96, hookData)
}

// Initialize is a paid mutator transaction binding the contract method 0x695c5bf5.
//
// Solidity: function initialize((address,address,uint24,int24,address) key, uint160 sqrtPriceX96, bytes hookData) returns(int24 tick)
func (_PoolMgr *PoolMgrTransactorSession) Initialize(key PoolKey, sqrtPriceX96 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Initialize(&_PoolMgr.TransactOpts, key, sqrtPriceX96, hookData)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "mint", to, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_PoolMgr *PoolMgrSession) Mint(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Mint(&_PoolMgr.TransactOpts, to, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Mint(&_PoolMgr.TransactOpts, to, id, amount)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) returns(int256 callerDelta, int256 feesAccrued)
func (_PoolMgr *PoolMgrTransactor) ModifyLiquidity(opts *bind.TransactOpts, key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "modifyLiquidity", key, params, hookData)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) returns(int256 callerDelta, int256 feesAccrued)
func (_PoolMgr *PoolMgrSession) ModifyLiquidity(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.ModifyLiquidity(&_PoolMgr.TransactOpts, key, params, hookData)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) returns(int256 callerDelta, int256 feesAccrued)
func (_PoolMgr *PoolMgrTransactorSession) ModifyLiquidity(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.ModifyLiquidity(&_PoolMgr.TransactOpts, key, params, hookData)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator, bool approved) returns(bool)
func (_PoolMgr *PoolMgrTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "setOperator", operator, approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator, bool approved) returns(bool)
func (_PoolMgr *PoolMgrSession) SetOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PoolMgr.Contract.SetOperator(&_PoolMgr.TransactOpts, operator, approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator, bool approved) returns(bool)
func (_PoolMgr *PoolMgrTransactorSession) SetOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PoolMgr.Contract.SetOperator(&_PoolMgr.TransactOpts, operator, approved)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7e87ce7d.
//
// Solidity: function setProtocolFee((address,address,uint24,int24,address) key, uint24 newProtocolFee) returns()
func (_PoolMgr *PoolMgrTransactor) SetProtocolFee(opts *bind.TransactOpts, key PoolKey, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "setProtocolFee", key, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7e87ce7d.
//
// Solidity: function setProtocolFee((address,address,uint24,int24,address) key, uint24 newProtocolFee) returns()
func (_PoolMgr *PoolMgrSession) SetProtocolFee(key PoolKey, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.SetProtocolFee(&_PoolMgr.TransactOpts, key, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7e87ce7d.
//
// Solidity: function setProtocolFee((address,address,uint24,int24,address) key, uint24 newProtocolFee) returns()
func (_PoolMgr *PoolMgrTransactorSession) SetProtocolFee(key PoolKey, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.SetProtocolFee(&_PoolMgr.TransactOpts, key, newProtocolFee)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address controller) returns()
func (_PoolMgr *PoolMgrTransactor) SetProtocolFeeController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "setProtocolFeeController", controller)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address controller) returns()
func (_PoolMgr *PoolMgrSession) SetProtocolFeeController(controller common.Address) (*types.Transaction, error) {
	return _PoolMgr.Contract.SetProtocolFeeController(&_PoolMgr.TransactOpts, controller)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address controller) returns()
func (_PoolMgr *PoolMgrTransactorSession) SetProtocolFeeController(controller common.Address) (*types.Transaction, error) {
	return _PoolMgr.Contract.SetProtocolFeeController(&_PoolMgr.TransactOpts, controller)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() payable returns(uint256 paid)
func (_PoolMgr *PoolMgrTransactor) Settle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "settle")
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() payable returns(uint256 paid)
func (_PoolMgr *PoolMgrSession) Settle() (*types.Transaction, error) {
	return _PoolMgr.Contract.Settle(&_PoolMgr.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() payable returns(uint256 paid)
func (_PoolMgr *PoolMgrTransactorSession) Settle() (*types.Transaction, error) {
	return _PoolMgr.Contract.Settle(&_PoolMgr.TransactOpts)
}

// SettleFor is a paid mutator transaction binding the contract method 0x3dd45adb.
//
// Solidity: function settleFor(address recipient) payable returns(uint256 paid)
func (_PoolMgr *PoolMgrTransactor) SettleFor(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "settleFor", recipient)
}

// SettleFor is a paid mutator transaction binding the contract method 0x3dd45adb.
//
// Solidity: function settleFor(address recipient) payable returns(uint256 paid)
func (_PoolMgr *PoolMgrSession) SettleFor(recipient common.Address) (*types.Transaction, error) {
	return _PoolMgr.Contract.SettleFor(&_PoolMgr.TransactOpts, recipient)
}

// SettleFor is a paid mutator transaction binding the contract method 0x3dd45adb.
//
// Solidity: function settleFor(address recipient) payable returns(uint256 paid)
func (_PoolMgr *PoolMgrTransactorSession) SettleFor(recipient common.Address) (*types.Transaction, error) {
	return _PoolMgr.Contract.SettleFor(&_PoolMgr.TransactOpts, recipient)
}

// Swap is a paid mutator transaction binding the contract method 0xf3cd914c.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes hookData) returns(int256 swapDelta)
func (_PoolMgr *PoolMgrTransactor) Swap(opts *bind.TransactOpts, key PoolKey, params IPoolManagerSwapParams, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "swap", key, params, hookData)
}

// Swap is a paid mutator transaction binding the contract method 0xf3cd914c.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes hookData) returns(int256 swapDelta)
func (_PoolMgr *PoolMgrSession) Swap(key PoolKey, params IPoolManagerSwapParams, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Swap(&_PoolMgr.TransactOpts, key, params, hookData)
}

// Swap is a paid mutator transaction binding the contract method 0xf3cd914c.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes hookData) returns(int256 swapDelta)
func (_PoolMgr *PoolMgrTransactorSession) Swap(key PoolKey, params IPoolManagerSwapParams, hookData []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Swap(&_PoolMgr.TransactOpts, key, params, hookData)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address currency) returns()
func (_PoolMgr *PoolMgrTransactor) Sync(opts *bind.TransactOpts, currency common.Address) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "sync", currency)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address currency) returns()
func (_PoolMgr *PoolMgrSession) Sync(currency common.Address) (*types.Transaction, error) {
	return _PoolMgr.Contract.Sync(&_PoolMgr.TransactOpts, currency)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address currency) returns()
func (_PoolMgr *PoolMgrTransactorSession) Sync(currency common.Address) (*types.Transaction, error) {
	return _PoolMgr.Contract.Sync(&_PoolMgr.TransactOpts, currency)
}

// Take is a paid mutator transaction binding the contract method 0x0b0d9c09.
//
// Solidity: function take(address currency, address to, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactor) Take(opts *bind.TransactOpts, currency common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "take", currency, to, amount)
}

// Take is a paid mutator transaction binding the contract method 0x0b0d9c09.
//
// Solidity: function take(address currency, address to, uint256 amount) returns()
func (_PoolMgr *PoolMgrSession) Take(currency common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Take(&_PoolMgr.TransactOpts, currency, to, amount)
}

// Take is a paid mutator transaction binding the contract method 0x0b0d9c09.
//
// Solidity: function take(address currency, address to, uint256 amount) returns()
func (_PoolMgr *PoolMgrTransactorSession) Take(currency common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Take(&_PoolMgr.TransactOpts, currency, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "transfer", receiver, id, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrSession) Transfer(receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Transfer(&_PoolMgr.TransactOpts, receiver, id, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrTransactorSession) Transfer(receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.Transfer(&_PoolMgr.TransactOpts, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "transferFrom", sender, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrSession) TransferFrom(sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.TransferFrom(&_PoolMgr.TransactOpts, sender, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_PoolMgr *PoolMgrTransactorSession) TransferFrom(sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.TransferFrom(&_PoolMgr.TransactOpts, sender, receiver, id, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes)
func (_PoolMgr *PoolMgrTransactor) Unlock(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "unlock", data)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes)
func (_PoolMgr *PoolMgrSession) Unlock(data []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Unlock(&_PoolMgr.TransactOpts, data)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes)
func (_PoolMgr *PoolMgrTransactorSession) Unlock(data []byte) (*types.Transaction, error) {
	return _PoolMgr.Contract.Unlock(&_PoolMgr.TransactOpts, data)
}

// UpdateDynamicLPFee is a paid mutator transaction binding the contract method 0x52759651.
//
// Solidity: function updateDynamicLPFee((address,address,uint24,int24,address) key, uint24 newDynamicLPFee) returns()
func (_PoolMgr *PoolMgrTransactor) UpdateDynamicLPFee(opts *bind.TransactOpts, key PoolKey, newDynamicLPFee *big.Int) (*types.Transaction, error) {
	return _PoolMgr.contract.Transact(opts, "updateDynamicLPFee", key, newDynamicLPFee)
}

// UpdateDynamicLPFee is a paid mutator transaction binding the contract method 0x52759651.
//
// Solidity: function updateDynamicLPFee((address,address,uint24,int24,address) key, uint24 newDynamicLPFee) returns()
func (_PoolMgr *PoolMgrSession) UpdateDynamicLPFee(key PoolKey, newDynamicLPFee *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.UpdateDynamicLPFee(&_PoolMgr.TransactOpts, key, newDynamicLPFee)
}

// UpdateDynamicLPFee is a paid mutator transaction binding the contract method 0x52759651.
//
// Solidity: function updateDynamicLPFee((address,address,uint24,int24,address) key, uint24 newDynamicLPFee) returns()
func (_PoolMgr *PoolMgrTransactorSession) UpdateDynamicLPFee(key PoolKey, newDynamicLPFee *big.Int) (*types.Transaction, error) {
	return _PoolMgr.Contract.UpdateDynamicLPFee(&_PoolMgr.TransactOpts, key, newDynamicLPFee)
}

// PoolMgrApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PoolMgr contract.
type PoolMgrApprovalIterator struct {
	Event *PoolMgrApproval // Event containing the contract specifics and raw log

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
func (it *PoolMgrApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrApproval)
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
		it.Event = new(PoolMgrApproval)
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
func (it *PoolMgrApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrApproval represents a Approval event raised by the PoolMgr contract.
type PoolMgrApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_PoolMgr *PoolMgrFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*PoolMgrApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrApprovalIterator{contract: _PoolMgr.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_PoolMgr *PoolMgrFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolMgrApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrApproval)
				if err := _PoolMgr.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_PoolMgr *PoolMgrFilterer) ParseApproval(log types.Log) (*PoolMgrApproval, error) {
	event := new(PoolMgrApproval)
	if err := _PoolMgr.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the PoolMgr contract.
type PoolMgrDonateIterator struct {
	Event *PoolMgrDonate // Event containing the contract specifics and raw log

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
func (it *PoolMgrDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrDonate)
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
		it.Event = new(PoolMgrDonate)
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
func (it *PoolMgrDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrDonate represents a Donate event raised by the PoolMgr contract.
type PoolMgrDonate struct {
	Id      [32]byte
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0x29ef05caaff9404b7cb6d1c0e9bbae9eaa7ab2541feba1a9c4248594c08156cb.
//
// Solidity: event Donate(bytes32 indexed id, address indexed sender, uint256 amount0, uint256 amount1)
func (_PoolMgr *PoolMgrFilterer) FilterDonate(opts *bind.FilterOpts, id [][32]byte, sender []common.Address) (*PoolMgrDonateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "Donate", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrDonateIterator{contract: _PoolMgr.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0x29ef05caaff9404b7cb6d1c0e9bbae9eaa7ab2541feba1a9c4248594c08156cb.
//
// Solidity: event Donate(bytes32 indexed id, address indexed sender, uint256 amount0, uint256 amount1)
func (_PoolMgr *PoolMgrFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *PoolMgrDonate, id [][32]byte, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "Donate", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrDonate)
				if err := _PoolMgr.contract.UnpackLog(event, "Donate", log); err != nil {
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

// ParseDonate is a log parse operation binding the contract event 0x29ef05caaff9404b7cb6d1c0e9bbae9eaa7ab2541feba1a9c4248594c08156cb.
//
// Solidity: event Donate(bytes32 indexed id, address indexed sender, uint256 amount0, uint256 amount1)
func (_PoolMgr *PoolMgrFilterer) ParseDonate(log types.Log) (*PoolMgrDonate, error) {
	event := new(PoolMgrDonate)
	if err := _PoolMgr.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the PoolMgr contract.
type PoolMgrInitializeIterator struct {
	Event *PoolMgrInitialize // Event containing the contract specifics and raw log

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
func (it *PoolMgrInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrInitialize)
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
		it.Event = new(PoolMgrInitialize)
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
func (it *PoolMgrInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrInitialize represents a Initialize event raised by the PoolMgr contract.
type PoolMgrInitialize struct {
	Id           [32]byte
	Currency0    common.Address
	Currency1    common.Address
	Fee          *big.Int
	TickSpacing  *big.Int
	Hooks        common.Address
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0xdd466e674ea557f56295e2d0218a125ea4b4f0f6f3307b95f85e6110838d6438.
//
// Solidity: event Initialize(bytes32 indexed id, address indexed currency0, address indexed currency1, uint24 fee, int24 tickSpacing, address hooks, uint160 sqrtPriceX96, int24 tick)
func (_PoolMgr *PoolMgrFilterer) FilterInitialize(opts *bind.FilterOpts, id [][32]byte, currency0 []common.Address, currency1 []common.Address) (*PoolMgrInitializeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var currency0Rule []interface{}
	for _, currency0Item := range currency0 {
		currency0Rule = append(currency0Rule, currency0Item)
	}
	var currency1Rule []interface{}
	for _, currency1Item := range currency1 {
		currency1Rule = append(currency1Rule, currency1Item)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "Initialize", idRule, currency0Rule, currency1Rule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrInitializeIterator{contract: _PoolMgr.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0xdd466e674ea557f56295e2d0218a125ea4b4f0f6f3307b95f85e6110838d6438.
//
// Solidity: event Initialize(bytes32 indexed id, address indexed currency0, address indexed currency1, uint24 fee, int24 tickSpacing, address hooks, uint160 sqrtPriceX96, int24 tick)
func (_PoolMgr *PoolMgrFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *PoolMgrInitialize, id [][32]byte, currency0 []common.Address, currency1 []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var currency0Rule []interface{}
	for _, currency0Item := range currency0 {
		currency0Rule = append(currency0Rule, currency0Item)
	}
	var currency1Rule []interface{}
	for _, currency1Item := range currency1 {
		currency1Rule = append(currency1Rule, currency1Item)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "Initialize", idRule, currency0Rule, currency1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrInitialize)
				if err := _PoolMgr.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0xdd466e674ea557f56295e2d0218a125ea4b4f0f6f3307b95f85e6110838d6438.
//
// Solidity: event Initialize(bytes32 indexed id, address indexed currency0, address indexed currency1, uint24 fee, int24 tickSpacing, address hooks, uint160 sqrtPriceX96, int24 tick)
func (_PoolMgr *PoolMgrFilterer) ParseInitialize(log types.Log) (*PoolMgrInitialize, error) {
	event := new(PoolMgrInitialize)
	if err := _PoolMgr.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrModifyLiquidityIterator is returned from FilterModifyLiquidity and is used to iterate over the raw logs and unpacked data for ModifyLiquidity events raised by the PoolMgr contract.
type PoolMgrModifyLiquidityIterator struct {
	Event *PoolMgrModifyLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolMgrModifyLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrModifyLiquidity)
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
		it.Event = new(PoolMgrModifyLiquidity)
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
func (it *PoolMgrModifyLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrModifyLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrModifyLiquidity represents a ModifyLiquidity event raised by the PoolMgr contract.
type PoolMgrModifyLiquidity struct {
	Id             [32]byte
	Sender         common.Address
	TickLower      *big.Int
	TickUpper      *big.Int
	LiquidityDelta *big.Int
	Salt           [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterModifyLiquidity is a free log retrieval operation binding the contract event 0xf208f4912782fd25c7f114ca3723a2d5dd6f3bcc3ac8db5af63baa85f711d5ec.
//
// Solidity: event ModifyLiquidity(bytes32 indexed id, address indexed sender, int24 tickLower, int24 tickUpper, int256 liquidityDelta, bytes32 salt)
func (_PoolMgr *PoolMgrFilterer) FilterModifyLiquidity(opts *bind.FilterOpts, id [][32]byte, sender []common.Address) (*PoolMgrModifyLiquidityIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "ModifyLiquidity", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrModifyLiquidityIterator{contract: _PoolMgr.contract, event: "ModifyLiquidity", logs: logs, sub: sub}, nil
}

// WatchModifyLiquidity is a free log subscription operation binding the contract event 0xf208f4912782fd25c7f114ca3723a2d5dd6f3bcc3ac8db5af63baa85f711d5ec.
//
// Solidity: event ModifyLiquidity(bytes32 indexed id, address indexed sender, int24 tickLower, int24 tickUpper, int256 liquidityDelta, bytes32 salt)
func (_PoolMgr *PoolMgrFilterer) WatchModifyLiquidity(opts *bind.WatchOpts, sink chan<- *PoolMgrModifyLiquidity, id [][32]byte, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "ModifyLiquidity", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrModifyLiquidity)
				if err := _PoolMgr.contract.UnpackLog(event, "ModifyLiquidity", log); err != nil {
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

// ParseModifyLiquidity is a log parse operation binding the contract event 0xf208f4912782fd25c7f114ca3723a2d5dd6f3bcc3ac8db5af63baa85f711d5ec.
//
// Solidity: event ModifyLiquidity(bytes32 indexed id, address indexed sender, int24 tickLower, int24 tickUpper, int256 liquidityDelta, bytes32 salt)
func (_PoolMgr *PoolMgrFilterer) ParseModifyLiquidity(log types.Log) (*PoolMgrModifyLiquidity, error) {
	event := new(PoolMgrModifyLiquidity)
	if err := _PoolMgr.contract.UnpackLog(event, "ModifyLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrOperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the PoolMgr contract.
type PoolMgrOperatorSetIterator struct {
	Event *PoolMgrOperatorSet // Event containing the contract specifics and raw log

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
func (it *PoolMgrOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrOperatorSet)
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
		it.Event = new(PoolMgrOperatorSet)
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
func (it *PoolMgrOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrOperatorSet represents a OperatorSet event raised by the PoolMgr contract.
type PoolMgrOperatorSet struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed operator, bool approved)
func (_PoolMgr *PoolMgrFilterer) FilterOperatorSet(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PoolMgrOperatorSetIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "OperatorSet", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrOperatorSetIterator{contract: _PoolMgr.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed operator, bool approved)
func (_PoolMgr *PoolMgrFilterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *PoolMgrOperatorSet, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "OperatorSet", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrOperatorSet)
				if err := _PoolMgr.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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

// ParseOperatorSet is a log parse operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed operator, bool approved)
func (_PoolMgr *PoolMgrFilterer) ParseOperatorSet(log types.Log) (*PoolMgrOperatorSet, error) {
	event := new(PoolMgrOperatorSet)
	if err := _PoolMgr.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrProtocolFeeControllerUpdatedIterator is returned from FilterProtocolFeeControllerUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeControllerUpdated events raised by the PoolMgr contract.
type PoolMgrProtocolFeeControllerUpdatedIterator struct {
	Event *PoolMgrProtocolFeeControllerUpdated // Event containing the contract specifics and raw log

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
func (it *PoolMgrProtocolFeeControllerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrProtocolFeeControllerUpdated)
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
		it.Event = new(PoolMgrProtocolFeeControllerUpdated)
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
func (it *PoolMgrProtocolFeeControllerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrProtocolFeeControllerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrProtocolFeeControllerUpdated represents a ProtocolFeeControllerUpdated event raised by the PoolMgr contract.
type PoolMgrProtocolFeeControllerUpdated struct {
	ProtocolFeeController common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeControllerUpdated is a free log retrieval operation binding the contract event 0xb4bd8ef53df690b9943d3318996006dbb82a25f54719d8c8035b516a2a5b8acc.
//
// Solidity: event ProtocolFeeControllerUpdated(address indexed protocolFeeController)
func (_PoolMgr *PoolMgrFilterer) FilterProtocolFeeControllerUpdated(opts *bind.FilterOpts, protocolFeeController []common.Address) (*PoolMgrProtocolFeeControllerUpdatedIterator, error) {

	var protocolFeeControllerRule []interface{}
	for _, protocolFeeControllerItem := range protocolFeeController {
		protocolFeeControllerRule = append(protocolFeeControllerRule, protocolFeeControllerItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "ProtocolFeeControllerUpdated", protocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrProtocolFeeControllerUpdatedIterator{contract: _PoolMgr.contract, event: "ProtocolFeeControllerUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeControllerUpdated is a free log subscription operation binding the contract event 0xb4bd8ef53df690b9943d3318996006dbb82a25f54719d8c8035b516a2a5b8acc.
//
// Solidity: event ProtocolFeeControllerUpdated(address indexed protocolFeeController)
func (_PoolMgr *PoolMgrFilterer) WatchProtocolFeeControllerUpdated(opts *bind.WatchOpts, sink chan<- *PoolMgrProtocolFeeControllerUpdated, protocolFeeController []common.Address) (event.Subscription, error) {

	var protocolFeeControllerRule []interface{}
	for _, protocolFeeControllerItem := range protocolFeeController {
		protocolFeeControllerRule = append(protocolFeeControllerRule, protocolFeeControllerItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "ProtocolFeeControllerUpdated", protocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrProtocolFeeControllerUpdated)
				if err := _PoolMgr.contract.UnpackLog(event, "ProtocolFeeControllerUpdated", log); err != nil {
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

// ParseProtocolFeeControllerUpdated is a log parse operation binding the contract event 0xb4bd8ef53df690b9943d3318996006dbb82a25f54719d8c8035b516a2a5b8acc.
//
// Solidity: event ProtocolFeeControllerUpdated(address indexed protocolFeeController)
func (_PoolMgr *PoolMgrFilterer) ParseProtocolFeeControllerUpdated(log types.Log) (*PoolMgrProtocolFeeControllerUpdated, error) {
	event := new(PoolMgrProtocolFeeControllerUpdated)
	if err := _PoolMgr.contract.UnpackLog(event, "ProtocolFeeControllerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrProtocolFeeUpdatedIterator is returned from FilterProtocolFeeUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeUpdated events raised by the PoolMgr contract.
type PoolMgrProtocolFeeUpdatedIterator struct {
	Event *PoolMgrProtocolFeeUpdated // Event containing the contract specifics and raw log

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
func (it *PoolMgrProtocolFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrProtocolFeeUpdated)
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
		it.Event = new(PoolMgrProtocolFeeUpdated)
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
func (it *PoolMgrProtocolFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrProtocolFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrProtocolFeeUpdated represents a ProtocolFeeUpdated event raised by the PoolMgr contract.
type PoolMgrProtocolFeeUpdated struct {
	Id          [32]byte
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeUpdated is a free log retrieval operation binding the contract event 0xe9c42593e71f84403b84352cd168d693e2c9fcd1fdbcc3feb21d92b43e6696f9.
//
// Solidity: event ProtocolFeeUpdated(bytes32 indexed id, uint24 protocolFee)
func (_PoolMgr *PoolMgrFilterer) FilterProtocolFeeUpdated(opts *bind.FilterOpts, id [][32]byte) (*PoolMgrProtocolFeeUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "ProtocolFeeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrProtocolFeeUpdatedIterator{contract: _PoolMgr.contract, event: "ProtocolFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeUpdated is a free log subscription operation binding the contract event 0xe9c42593e71f84403b84352cd168d693e2c9fcd1fdbcc3feb21d92b43e6696f9.
//
// Solidity: event ProtocolFeeUpdated(bytes32 indexed id, uint24 protocolFee)
func (_PoolMgr *PoolMgrFilterer) WatchProtocolFeeUpdated(opts *bind.WatchOpts, sink chan<- *PoolMgrProtocolFeeUpdated, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "ProtocolFeeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrProtocolFeeUpdated)
				if err := _PoolMgr.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
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

// ParseProtocolFeeUpdated is a log parse operation binding the contract event 0xe9c42593e71f84403b84352cd168d693e2c9fcd1fdbcc3feb21d92b43e6696f9.
//
// Solidity: event ProtocolFeeUpdated(bytes32 indexed id, uint24 protocolFee)
func (_PoolMgr *PoolMgrFilterer) ParseProtocolFeeUpdated(log types.Log) (*PoolMgrProtocolFeeUpdated, error) {
	event := new(PoolMgrProtocolFeeUpdated)
	if err := _PoolMgr.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the PoolMgr contract.
type PoolMgrSwapIterator struct {
	Event *PoolMgrSwap // Event containing the contract specifics and raw log

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
func (it *PoolMgrSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrSwap)
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
		it.Event = new(PoolMgrSwap)
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
func (it *PoolMgrSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrSwap represents a Swap event raised by the PoolMgr contract.
type PoolMgrSwap struct {
	Id           [32]byte
	Sender       common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Fee          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f.
//
// Solidity: event Swap(bytes32 indexed id, address indexed sender, int128 amount0, int128 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint24 fee)
func (_PoolMgr *PoolMgrFilterer) FilterSwap(opts *bind.FilterOpts, id [][32]byte, sender []common.Address) (*PoolMgrSwapIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "Swap", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrSwapIterator{contract: _PoolMgr.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f.
//
// Solidity: event Swap(bytes32 indexed id, address indexed sender, int128 amount0, int128 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint24 fee)
func (_PoolMgr *PoolMgrFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PoolMgrSwap, id [][32]byte, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "Swap", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrSwap)
				if err := _PoolMgr.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f.
//
// Solidity: event Swap(bytes32 indexed id, address indexed sender, int128 amount0, int128 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint24 fee)
func (_PoolMgr *PoolMgrFilterer) ParseSwap(log types.Log) (*PoolMgrSwap, error) {
	event := new(PoolMgrSwap)
	if err := _PoolMgr.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMgrTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PoolMgr contract.
type PoolMgrTransferIterator struct {
	Event *PoolMgrTransfer // Event containing the contract specifics and raw log

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
func (it *PoolMgrTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMgrTransfer)
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
		it.Event = new(PoolMgrTransfer)
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
func (it *PoolMgrTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMgrTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMgrTransfer represents a Transfer event raised by the PoolMgr contract.
type PoolMgrTransfer struct {
	Caller common.Address
	From   common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed from, address indexed to, uint256 indexed id, uint256 amount)
func (_PoolMgr *PoolMgrFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*PoolMgrTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PoolMgr.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PoolMgrTransferIterator{contract: _PoolMgr.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed from, address indexed to, uint256 indexed id, uint256 amount)
func (_PoolMgr *PoolMgrFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolMgrTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PoolMgr.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMgrTransfer)
				if err := _PoolMgr.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed from, address indexed to, uint256 indexed id, uint256 amount)
func (_PoolMgr *PoolMgrFilterer) ParseTransfer(log types.Log) (*PoolMgrTransfer, error) {
	event := new(PoolMgrTransfer)
	if err := _PoolMgr.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
