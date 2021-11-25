// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// HubMetaData contains all meta data concerning the Hub contract.
var HubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"Client\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"createdMOVROrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"createdRMRKOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"Client\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"filledMOVROrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"filledRMRKOrder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Client\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"addMOVROrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"New\",\"type\":\"address\"}],\"name\":\"changeAddressBank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"New\",\"type\":\"address\"}],\"name\":\"changeAddressReciever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"New\",\"type\":\"uint256\"}],\"name\":\"changeDeployPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"New\",\"type\":\"uint256\"}],\"name\":\"changeServerPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Client\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"createMOVROrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Client\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"createRMRKOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillMOVROrders\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Client\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"fillRMRKOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressBank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressReciever\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployPayments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"UID\",\"type\":\"string\"}],\"name\":\"getID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"getMeta\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getServerPayments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"MyAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myLostMoney\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"New\",\"type\":\"bool\"}],\"name\":\"switchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"takeMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HubABI is the input ABI used to generate the binding from.
// Deprecated: Use HubMetaData.ABI instead.
var HubABI = HubMetaData.ABI

// Hub is an auto generated Go binding around an Ethereum contract.
type Hub struct {
	HubCaller     // Read-only binding to the contract
	HubTransactor // Write-only binding to the contract
	HubFilterer   // Log filterer for contract events
}

// HubCaller is an auto generated read-only Go binding around an Ethereum contract.
type HubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HubSession struct {
	Contract     *Hub              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HubCallerSession struct {
	Contract *HubCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HubTransactorSession struct {
	Contract     *HubTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubRaw is an auto generated low-level Go binding around an Ethereum contract.
type HubRaw struct {
	Contract *Hub // Generic contract binding to access the raw methods on
}

// HubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HubCallerRaw struct {
	Contract *HubCaller // Generic read-only contract binding to access the raw methods on
}

// HubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HubTransactorRaw struct {
	Contract *HubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHub creates a new instance of Hub, bound to a specific deployed contract.
func NewHub(address common.Address, backend bind.ContractBackend) (*Hub, error) {
	contract, err := bindHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hub{HubCaller: HubCaller{contract: contract}, HubTransactor: HubTransactor{contract: contract}, HubFilterer: HubFilterer{contract: contract}}, nil
}

// NewHubCaller creates a new read-only instance of Hub, bound to a specific deployed contract.
func NewHubCaller(address common.Address, caller bind.ContractCaller) (*HubCaller, error) {
	contract, err := bindHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HubCaller{contract: contract}, nil
}

// NewHubTransactor creates a new write-only instance of Hub, bound to a specific deployed contract.
func NewHubTransactor(address common.Address, transactor bind.ContractTransactor) (*HubTransactor, error) {
	contract, err := bindHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HubTransactor{contract: contract}, nil
}

// NewHubFilterer creates a new log filterer instance of Hub, bound to a specific deployed contract.
func NewHubFilterer(address common.Address, filterer bind.ContractFilterer) (*HubFilterer, error) {
	contract, err := bindHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HubFilterer{contract: contract}, nil
}

// bindHub binds a generic wrapper to an already deployed contract.
func bindHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.HubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hub *HubCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hub *HubSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hub.Contract.BalanceOf(&_Hub.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hub *HubCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hub.Contract.BalanceOf(&_Hub.CallOpts, owner)
}

// GetAddressBank is a free data retrieval call binding the contract method 0x8644c571.
//
// Solidity: function getAddressBank() view returns(address)
func (_Hub *HubCaller) GetAddressBank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getAddressBank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressBank is a free data retrieval call binding the contract method 0x8644c571.
//
// Solidity: function getAddressBank() view returns(address)
func (_Hub *HubSession) GetAddressBank() (common.Address, error) {
	return _Hub.Contract.GetAddressBank(&_Hub.CallOpts)
}

// GetAddressBank is a free data retrieval call binding the contract method 0x8644c571.
//
// Solidity: function getAddressBank() view returns(address)
func (_Hub *HubCallerSession) GetAddressBank() (common.Address, error) {
	return _Hub.Contract.GetAddressBank(&_Hub.CallOpts)
}

// GetAddressReciever is a free data retrieval call binding the contract method 0x6fdb3b0d.
//
// Solidity: function getAddressReciever() view returns(address)
func (_Hub *HubCaller) GetAddressReciever(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getAddressReciever")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressReciever is a free data retrieval call binding the contract method 0x6fdb3b0d.
//
// Solidity: function getAddressReciever() view returns(address)
func (_Hub *HubSession) GetAddressReciever() (common.Address, error) {
	return _Hub.Contract.GetAddressReciever(&_Hub.CallOpts)
}

// GetAddressReciever is a free data retrieval call binding the contract method 0x6fdb3b0d.
//
// Solidity: function getAddressReciever() view returns(address)
func (_Hub *HubCallerSession) GetAddressReciever() (common.Address, error) {
	return _Hub.Contract.GetAddressReciever(&_Hub.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hub *HubCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hub *HubSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetApproved(&_Hub.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hub *HubCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetApproved(&_Hub.CallOpts, tokenId)
}

// GetDeployPayments is a free data retrieval call binding the contract method 0xa79199c7.
//
// Solidity: function getDeployPayments() view returns(uint256)
func (_Hub *HubCaller) GetDeployPayments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getDeployPayments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeployPayments is a free data retrieval call binding the contract method 0xa79199c7.
//
// Solidity: function getDeployPayments() view returns(uint256)
func (_Hub *HubSession) GetDeployPayments() (*big.Int, error) {
	return _Hub.Contract.GetDeployPayments(&_Hub.CallOpts)
}

// GetDeployPayments is a free data retrieval call binding the contract method 0xa79199c7.
//
// Solidity: function getDeployPayments() view returns(uint256)
func (_Hub *HubCallerSession) GetDeployPayments() (*big.Int, error) {
	return _Hub.Contract.GetDeployPayments(&_Hub.CallOpts)
}

// GetID is a free data retrieval call binding the contract method 0x70aca6c8.
//
// Solidity: function getID(string UID) view returns(uint256)
func (_Hub *HubCaller) GetID(opts *bind.CallOpts, UID string) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getID", UID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetID is a free data retrieval call binding the contract method 0x70aca6c8.
//
// Solidity: function getID(string UID) view returns(uint256)
func (_Hub *HubSession) GetID(UID string) (*big.Int, error) {
	return _Hub.Contract.GetID(&_Hub.CallOpts, UID)
}

// GetID is a free data retrieval call binding the contract method 0x70aca6c8.
//
// Solidity: function getID(string UID) view returns(uint256)
func (_Hub *HubCallerSession) GetID(UID string) (*big.Int, error) {
	return _Hub.Contract.GetID(&_Hub.CallOpts, UID)
}

// GetMeta is a free data retrieval call binding the contract method 0x36dac2cc.
//
// Solidity: function getMeta(uint256 ID) view returns(string)
func (_Hub *HubCaller) GetMeta(opts *bind.CallOpts, ID *big.Int) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getMeta", ID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMeta is a free data retrieval call binding the contract method 0x36dac2cc.
//
// Solidity: function getMeta(uint256 ID) view returns(string)
func (_Hub *HubSession) GetMeta(ID *big.Int) (string, error) {
	return _Hub.Contract.GetMeta(&_Hub.CallOpts, ID)
}

// GetMeta is a free data retrieval call binding the contract method 0x36dac2cc.
//
// Solidity: function getMeta(uint256 ID) view returns(string)
func (_Hub *HubCallerSession) GetMeta(ID *big.Int) (string, error) {
	return _Hub.Contract.GetMeta(&_Hub.CallOpts, ID)
}

// GetServerPayments is a free data retrieval call binding the contract method 0xd89499cc.
//
// Solidity: function getServerPayments() view returns(uint256)
func (_Hub *HubCaller) GetServerPayments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getServerPayments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetServerPayments is a free data retrieval call binding the contract method 0xd89499cc.
//
// Solidity: function getServerPayments() view returns(uint256)
func (_Hub *HubSession) GetServerPayments() (*big.Int, error) {
	return _Hub.Contract.GetServerPayments(&_Hub.CallOpts)
}

// GetServerPayments is a free data retrieval call binding the contract method 0xd89499cc.
//
// Solidity: function getServerPayments() view returns(uint256)
func (_Hub *HubCallerSession) GetServerPayments() (*big.Int, error) {
	return _Hub.Contract.GetServerPayments(&_Hub.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hub *HubCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hub *HubSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hub.Contract.IsApprovedForAll(&_Hub.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hub *HubCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hub.Contract.IsApprovedForAll(&_Hub.CallOpts, owner, operator)
}

// MyLostMoney is a free data retrieval call binding the contract method 0x3c239b05.
//
// Solidity: function myLostMoney() view returns(uint256)
func (_Hub *HubCaller) MyLostMoney(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "myLostMoney")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyLostMoney is a free data retrieval call binding the contract method 0x3c239b05.
//
// Solidity: function myLostMoney() view returns(uint256)
func (_Hub *HubSession) MyLostMoney() (*big.Int, error) {
	return _Hub.Contract.MyLostMoney(&_Hub.CallOpts)
}

// MyLostMoney is a free data retrieval call binding the contract method 0x3c239b05.
//
// Solidity: function myLostMoney() view returns(uint256)
func (_Hub *HubCallerSession) MyLostMoney() (*big.Int, error) {
	return _Hub.Contract.MyLostMoney(&_Hub.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hub *HubCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hub *HubSession) Name() (string, error) {
	return _Hub.Contract.Name(&_Hub.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hub *HubCallerSession) Name() (string, error) {
	return _Hub.Contract.Name(&_Hub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hub *HubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hub *HubSession) Owner() (common.Address, error) {
	return _Hub.Contract.Owner(&_Hub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hub *HubCallerSession) Owner() (common.Address, error) {
	return _Hub.Contract.Owner(&_Hub.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hub *HubCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hub *HubSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.OwnerOf(&_Hub.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hub *HubCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.OwnerOf(&_Hub.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hub *HubCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hub *HubSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hub.Contract.SupportsInterface(&_Hub.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hub *HubCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hub.Contract.SupportsInterface(&_Hub.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hub *HubCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hub *HubSession) Symbol() (string, error) {
	return _Hub.Contract.Symbol(&_Hub.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hub *HubCallerSession) Symbol() (string, error) {
	return _Hub.Contract.Symbol(&_Hub.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hub *HubCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hub *HubSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hub.Contract.TokenURI(&_Hub.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hub *HubCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hub.Contract.TokenURI(&_Hub.CallOpts, tokenId)
}

// AddMOVROrders is a paid mutator transaction binding the contract method 0x39f4279f.
//
// Solidity: function addMOVROrders(string Client, string UID) returns()
func (_Hub *HubTransactor) AddMOVROrders(opts *bind.TransactOpts, Client string, UID string) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "addMOVROrders", Client, UID)
}

// AddMOVROrders is a paid mutator transaction binding the contract method 0x39f4279f.
//
// Solidity: function addMOVROrders(string Client, string UID) returns()
func (_Hub *HubSession) AddMOVROrders(Client string, UID string) (*types.Transaction, error) {
	return _Hub.Contract.AddMOVROrders(&_Hub.TransactOpts, Client, UID)
}

// AddMOVROrders is a paid mutator transaction binding the contract method 0x39f4279f.
//
// Solidity: function addMOVROrders(string Client, string UID) returns()
func (_Hub *HubTransactorSession) AddMOVROrders(Client string, UID string) (*types.Transaction, error) {
	return _Hub.Contract.AddMOVROrders(&_Hub.TransactOpts, Client, UID)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hub *HubTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hub *HubSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.Approve(&_Hub.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hub *HubTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.Approve(&_Hub.TransactOpts, to, tokenId)
}

// ChangeAddressBank is a paid mutator transaction binding the contract method 0xfb8d0b5f.
//
// Solidity: function changeAddressBank(address New) returns()
func (_Hub *HubTransactor) ChangeAddressBank(opts *bind.TransactOpts, New common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "changeAddressBank", New)
}

// ChangeAddressBank is a paid mutator transaction binding the contract method 0xfb8d0b5f.
//
// Solidity: function changeAddressBank(address New) returns()
func (_Hub *HubSession) ChangeAddressBank(New common.Address) (*types.Transaction, error) {
	return _Hub.Contract.ChangeAddressBank(&_Hub.TransactOpts, New)
}

// ChangeAddressBank is a paid mutator transaction binding the contract method 0xfb8d0b5f.
//
// Solidity: function changeAddressBank(address New) returns()
func (_Hub *HubTransactorSession) ChangeAddressBank(New common.Address) (*types.Transaction, error) {
	return _Hub.Contract.ChangeAddressBank(&_Hub.TransactOpts, New)
}

// ChangeAddressReciever is a paid mutator transaction binding the contract method 0x6226943a.
//
// Solidity: function changeAddressReciever(address New) returns()
func (_Hub *HubTransactor) ChangeAddressReciever(opts *bind.TransactOpts, New common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "changeAddressReciever", New)
}

// ChangeAddressReciever is a paid mutator transaction binding the contract method 0x6226943a.
//
// Solidity: function changeAddressReciever(address New) returns()
func (_Hub *HubSession) ChangeAddressReciever(New common.Address) (*types.Transaction, error) {
	return _Hub.Contract.ChangeAddressReciever(&_Hub.TransactOpts, New)
}

// ChangeAddressReciever is a paid mutator transaction binding the contract method 0x6226943a.
//
// Solidity: function changeAddressReciever(address New) returns()
func (_Hub *HubTransactorSession) ChangeAddressReciever(New common.Address) (*types.Transaction, error) {
	return _Hub.Contract.ChangeAddressReciever(&_Hub.TransactOpts, New)
}

// ChangeDeployPayments is a paid mutator transaction binding the contract method 0xa489bab5.
//
// Solidity: function changeDeployPayments(uint256 New) returns()
func (_Hub *HubTransactor) ChangeDeployPayments(opts *bind.TransactOpts, New *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "changeDeployPayments", New)
}

// ChangeDeployPayments is a paid mutator transaction binding the contract method 0xa489bab5.
//
// Solidity: function changeDeployPayments(uint256 New) returns()
func (_Hub *HubSession) ChangeDeployPayments(New *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.ChangeDeployPayments(&_Hub.TransactOpts, New)
}

// ChangeDeployPayments is a paid mutator transaction binding the contract method 0xa489bab5.
//
// Solidity: function changeDeployPayments(uint256 New) returns()
func (_Hub *HubTransactorSession) ChangeDeployPayments(New *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.ChangeDeployPayments(&_Hub.TransactOpts, New)
}

// ChangeServerPayments is a paid mutator transaction binding the contract method 0xedd43c19.
//
// Solidity: function changeServerPayments(uint256 New) returns()
func (_Hub *HubTransactor) ChangeServerPayments(opts *bind.TransactOpts, New *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "changeServerPayments", New)
}

// ChangeServerPayments is a paid mutator transaction binding the contract method 0xedd43c19.
//
// Solidity: function changeServerPayments(uint256 New) returns()
func (_Hub *HubSession) ChangeServerPayments(New *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.ChangeServerPayments(&_Hub.TransactOpts, New)
}

// ChangeServerPayments is a paid mutator transaction binding the contract method 0xedd43c19.
//
// Solidity: function changeServerPayments(uint256 New) returns()
func (_Hub *HubTransactorSession) ChangeServerPayments(New *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.ChangeServerPayments(&_Hub.TransactOpts, New)
}

// CreateMOVROrder is a paid mutator transaction binding the contract method 0x583ccd94.
//
// Solidity: function createMOVROrder(string Client, string UID) payable returns()
func (_Hub *HubTransactor) CreateMOVROrder(opts *bind.TransactOpts, Client string, UID string) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "createMOVROrder", Client, UID)
}

// CreateMOVROrder is a paid mutator transaction binding the contract method 0x583ccd94.
//
// Solidity: function createMOVROrder(string Client, string UID) payable returns()
func (_Hub *HubSession) CreateMOVROrder(Client string, UID string) (*types.Transaction, error) {
	return _Hub.Contract.CreateMOVROrder(&_Hub.TransactOpts, Client, UID)
}

// CreateMOVROrder is a paid mutator transaction binding the contract method 0x583ccd94.
//
// Solidity: function createMOVROrder(string Client, string UID) payable returns()
func (_Hub *HubTransactorSession) CreateMOVROrder(Client string, UID string) (*types.Transaction, error) {
	return _Hub.Contract.CreateMOVROrder(&_Hub.TransactOpts, Client, UID)
}

// CreateRMRKOrder is a paid mutator transaction binding the contract method 0xa7676976.
//
// Solidity: function createRMRKOrder(address Client, string UID) payable returns()
func (_Hub *HubTransactor) CreateRMRKOrder(opts *bind.TransactOpts, Client common.Address, UID string) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "createRMRKOrder", Client, UID)
}

// CreateRMRKOrder is a paid mutator transaction binding the contract method 0xa7676976.
//
// Solidity: function createRMRKOrder(address Client, string UID) payable returns()
func (_Hub *HubSession) CreateRMRKOrder(Client common.Address, UID string) (*types.Transaction, error) {
	return _Hub.Contract.CreateRMRKOrder(&_Hub.TransactOpts, Client, UID)
}

// CreateRMRKOrder is a paid mutator transaction binding the contract method 0xa7676976.
//
// Solidity: function createRMRKOrder(address Client, string UID) payable returns()
func (_Hub *HubTransactorSession) CreateRMRKOrder(Client common.Address, UID string) (*types.Transaction, error) {
	return _Hub.Contract.CreateRMRKOrder(&_Hub.TransactOpts, Client, UID)
}

// FillMOVROrders is a paid mutator transaction binding the contract method 0xa659e63a.
//
// Solidity: function fillMOVROrders() payable returns()
func (_Hub *HubTransactor) FillMOVROrders(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "fillMOVROrders")
}

// FillMOVROrders is a paid mutator transaction binding the contract method 0xa659e63a.
//
// Solidity: function fillMOVROrders() payable returns()
func (_Hub *HubSession) FillMOVROrders() (*types.Transaction, error) {
	return _Hub.Contract.FillMOVROrders(&_Hub.TransactOpts)
}

// FillMOVROrders is a paid mutator transaction binding the contract method 0xa659e63a.
//
// Solidity: function fillMOVROrders() payable returns()
func (_Hub *HubTransactorSession) FillMOVROrders() (*types.Transaction, error) {
	return _Hub.Contract.FillMOVROrders(&_Hub.TransactOpts)
}

// FillRMRKOrder is a paid mutator transaction binding the contract method 0x5a780f40.
//
// Solidity: function fillRMRKOrder(address Client, string UID, string URI, string data) payable returns()
func (_Hub *HubTransactor) FillRMRKOrder(opts *bind.TransactOpts, Client common.Address, UID string, URI string, data string) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "fillRMRKOrder", Client, UID, URI, data)
}

// FillRMRKOrder is a paid mutator transaction binding the contract method 0x5a780f40.
//
// Solidity: function fillRMRKOrder(address Client, string UID, string URI, string data) payable returns()
func (_Hub *HubSession) FillRMRKOrder(Client common.Address, UID string, URI string, data string) (*types.Transaction, error) {
	return _Hub.Contract.FillRMRKOrder(&_Hub.TransactOpts, Client, UID, URI, data)
}

// FillRMRKOrder is a paid mutator transaction binding the contract method 0x5a780f40.
//
// Solidity: function fillRMRKOrder(address Client, string UID, string URI, string data) payable returns()
func (_Hub *HubTransactorSession) FillRMRKOrder(Client common.Address, UID string, URI string, data string) (*types.Transaction, error) {
	return _Hub.Contract.FillRMRKOrder(&_Hub.TransactOpts, Client, UID, URI, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address MyAddress) returns()
func (_Hub *HubTransactor) Initialize(opts *bind.TransactOpts, MyAddress common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "initialize", MyAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address MyAddress) returns()
func (_Hub *HubSession) Initialize(MyAddress common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Initialize(&_Hub.TransactOpts, MyAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address MyAddress) returns()
func (_Hub *HubTransactorSession) Initialize(MyAddress common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Initialize(&_Hub.TransactOpts, MyAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hub *HubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hub *HubSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hub.Contract.RenounceOwnership(&_Hub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hub *HubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hub.Contract.RenounceOwnership(&_Hub.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hub *HubTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hub *HubSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom0(&_Hub.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hub *HubTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom0(&_Hub.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hub *HubTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hub *HubSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hub.Contract.SetApprovalForAll(&_Hub.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hub *HubTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hub.Contract.SetApprovalForAll(&_Hub.TransactOpts, operator, approved)
}

// SwitchFee is a paid mutator transaction binding the contract method 0x01f048e8.
//
// Solidity: function switchFee(bool New) returns()
func (_Hub *HubTransactor) SwitchFee(opts *bind.TransactOpts, New bool) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "switchFee", New)
}

// SwitchFee is a paid mutator transaction binding the contract method 0x01f048e8.
//
// Solidity: function switchFee(bool New) returns()
func (_Hub *HubSession) SwitchFee(New bool) (*types.Transaction, error) {
	return _Hub.Contract.SwitchFee(&_Hub.TransactOpts, New)
}

// SwitchFee is a paid mutator transaction binding the contract method 0x01f048e8.
//
// Solidity: function switchFee(bool New) returns()
func (_Hub *HubTransactorSession) SwitchFee(New bool) (*types.Transaction, error) {
	return _Hub.Contract.SwitchFee(&_Hub.TransactOpts, New)
}

// TakeMoney is a paid mutator transaction binding the contract method 0xfcae8c06.
//
// Solidity: function takeMoney() returns()
func (_Hub *HubTransactor) TakeMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "takeMoney")
}

// TakeMoney is a paid mutator transaction binding the contract method 0xfcae8c06.
//
// Solidity: function takeMoney() returns()
func (_Hub *HubSession) TakeMoney() (*types.Transaction, error) {
	return _Hub.Contract.TakeMoney(&_Hub.TransactOpts)
}

// TakeMoney is a paid mutator transaction binding the contract method 0xfcae8c06.
//
// Solidity: function takeMoney() returns()
func (_Hub *HubTransactorSession) TakeMoney() (*types.Transaction, error) {
	return _Hub.Contract.TakeMoney(&_Hub.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.TransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.TransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hub *HubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hub *HubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hub.Contract.TransferOwnership(&_Hub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hub *HubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hub.Contract.TransferOwnership(&_Hub.TransactOpts, newOwner)
}

// HubApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hub contract.
type HubApprovalIterator struct {
	Event *HubApproval // Event containing the contract specifics and raw log

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
func (it *HubApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubApproval)
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
		it.Event = new(HubApproval)
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
func (it *HubApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubApproval represents a Approval event raised by the Hub contract.
type HubApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hub *HubFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*HubApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HubApprovalIterator{contract: _Hub.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hub *HubFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HubApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubApproval)
				if err := _Hub.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hub *HubFilterer) ParseApproval(log types.Log) (*HubApproval, error) {
	event := new(HubApproval)
	if err := _Hub.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Hub contract.
type HubApprovalForAllIterator struct {
	Event *HubApprovalForAll // Event containing the contract specifics and raw log

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
func (it *HubApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubApprovalForAll)
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
		it.Event = new(HubApprovalForAll)
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
func (it *HubApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubApprovalForAll represents a ApprovalForAll event raised by the Hub contract.
type HubApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hub *HubFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*HubApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &HubApprovalForAllIterator{contract: _Hub.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hub *HubFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *HubApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubApprovalForAll)
				if err := _Hub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hub *HubFilterer) ParseApprovalForAll(log types.Log) (*HubApprovalForAll, error) {
	event := new(HubApprovalForAll)
	if err := _Hub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hub contract.
type HubOwnershipTransferredIterator struct {
	Event *HubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubOwnershipTransferred)
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
		it.Event = new(HubOwnershipTransferred)
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
func (it *HubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubOwnershipTransferred represents a OwnershipTransferred event raised by the Hub contract.
type HubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hub *HubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HubOwnershipTransferredIterator{contract: _Hub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hub *HubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubOwnershipTransferred)
				if err := _Hub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hub *HubFilterer) ParseOwnershipTransferred(log types.Log) (*HubOwnershipTransferred, error) {
	event := new(HubOwnershipTransferred)
	if err := _Hub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hub contract.
type HubTransferIterator struct {
	Event *HubTransfer // Event containing the contract specifics and raw log

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
func (it *HubTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubTransfer)
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
		it.Event = new(HubTransfer)
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
func (it *HubTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubTransfer represents a Transfer event raised by the Hub contract.
type HubTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hub *HubFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*HubTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HubTransferIterator{contract: _Hub.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hub *HubFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HubTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubTransfer)
				if err := _Hub.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hub *HubFilterer) ParseTransfer(log types.Log) (*HubTransfer, error) {
	event := new(HubTransfer)
	if err := _Hub.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubCreatedMOVROrderIterator is returned from FilterCreatedMOVROrder and is used to iterate over the raw logs and unpacked data for CreatedMOVROrder events raised by the Hub contract.
type HubCreatedMOVROrderIterator struct {
	Event *HubCreatedMOVROrder // Event containing the contract specifics and raw log

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
func (it *HubCreatedMOVROrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubCreatedMOVROrder)
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
		it.Event = new(HubCreatedMOVROrder)
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
func (it *HubCreatedMOVROrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubCreatedMOVROrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubCreatedMOVROrder represents a CreatedMOVROrder event raised by the Hub contract.
type HubCreatedMOVROrder struct {
	Client string
	UID    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreatedMOVROrder is a free log retrieval operation binding the contract event 0x291b181db7bd4d3ff55e53015da42cb7e0a3f808f5d7a3e3de822e327897d423.
//
// Solidity: event createdMOVROrder(string Client, string UID)
func (_Hub *HubFilterer) FilterCreatedMOVROrder(opts *bind.FilterOpts) (*HubCreatedMOVROrderIterator, error) {

	logs, sub, err := _Hub.contract.FilterLogs(opts, "createdMOVROrder")
	if err != nil {
		return nil, err
	}
	return &HubCreatedMOVROrderIterator{contract: _Hub.contract, event: "createdMOVROrder", logs: logs, sub: sub}, nil
}

// WatchCreatedMOVROrder is a free log subscription operation binding the contract event 0x291b181db7bd4d3ff55e53015da42cb7e0a3f808f5d7a3e3de822e327897d423.
//
// Solidity: event createdMOVROrder(string Client, string UID)
func (_Hub *HubFilterer) WatchCreatedMOVROrder(opts *bind.WatchOpts, sink chan<- *HubCreatedMOVROrder) (event.Subscription, error) {

	logs, sub, err := _Hub.contract.WatchLogs(opts, "createdMOVROrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubCreatedMOVROrder)
				if err := _Hub.contract.UnpackLog(event, "createdMOVROrder", log); err != nil {
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

// ParseCreatedMOVROrder is a log parse operation binding the contract event 0x291b181db7bd4d3ff55e53015da42cb7e0a3f808f5d7a3e3de822e327897d423.
//
// Solidity: event createdMOVROrder(string Client, string UID)
func (_Hub *HubFilterer) ParseCreatedMOVROrder(log types.Log) (*HubCreatedMOVROrder, error) {
	event := new(HubCreatedMOVROrder)
	if err := _Hub.contract.UnpackLog(event, "createdMOVROrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubCreatedRMRKOrderIterator is returned from FilterCreatedRMRKOrder and is used to iterate over the raw logs and unpacked data for CreatedRMRKOrder events raised by the Hub contract.
type HubCreatedRMRKOrderIterator struct {
	Event *HubCreatedRMRKOrder // Event containing the contract specifics and raw log

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
func (it *HubCreatedRMRKOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubCreatedRMRKOrder)
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
		it.Event = new(HubCreatedRMRKOrder)
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
func (it *HubCreatedRMRKOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubCreatedRMRKOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubCreatedRMRKOrder represents a CreatedRMRKOrder event raised by the Hub contract.
type HubCreatedRMRKOrder struct {
	Client common.Address
	UID    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreatedRMRKOrder is a free log retrieval operation binding the contract event 0x805ea7ce84de8531f01d31bd45d58f7167a093ae9b87414cc2a0dd484731507e.
//
// Solidity: event createdRMRKOrder(address Client, string UID)
func (_Hub *HubFilterer) FilterCreatedRMRKOrder(opts *bind.FilterOpts) (*HubCreatedRMRKOrderIterator, error) {

	logs, sub, err := _Hub.contract.FilterLogs(opts, "createdRMRKOrder")
	if err != nil {
		return nil, err
	}
	return &HubCreatedRMRKOrderIterator{contract: _Hub.contract, event: "createdRMRKOrder", logs: logs, sub: sub}, nil
}

// WatchCreatedRMRKOrder is a free log subscription operation binding the contract event 0x805ea7ce84de8531f01d31bd45d58f7167a093ae9b87414cc2a0dd484731507e.
//
// Solidity: event createdRMRKOrder(address Client, string UID)
func (_Hub *HubFilterer) WatchCreatedRMRKOrder(opts *bind.WatchOpts, sink chan<- *HubCreatedRMRKOrder) (event.Subscription, error) {

	logs, sub, err := _Hub.contract.WatchLogs(opts, "createdRMRKOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubCreatedRMRKOrder)
				if err := _Hub.contract.UnpackLog(event, "createdRMRKOrder", log); err != nil {
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

// ParseCreatedRMRKOrder is a log parse operation binding the contract event 0x805ea7ce84de8531f01d31bd45d58f7167a093ae9b87414cc2a0dd484731507e.
//
// Solidity: event createdRMRKOrder(address Client, string UID)
func (_Hub *HubFilterer) ParseCreatedRMRKOrder(log types.Log) (*HubCreatedRMRKOrder, error) {
	event := new(HubCreatedRMRKOrder)
	if err := _Hub.contract.UnpackLog(event, "createdRMRKOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubFilledMOVROrderIterator is returned from FilterFilledMOVROrder and is used to iterate over the raw logs and unpacked data for FilledMOVROrder events raised by the Hub contract.
type HubFilledMOVROrderIterator struct {
	Event *HubFilledMOVROrder // Event containing the contract specifics and raw log

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
func (it *HubFilledMOVROrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubFilledMOVROrder)
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
		it.Event = new(HubFilledMOVROrder)
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
func (it *HubFilledMOVROrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubFilledMOVROrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubFilledMOVROrder represents a FilledMOVROrder event raised by the Hub contract.
type HubFilledMOVROrder struct {
	Client string
	UID    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFilledMOVROrder is a free log retrieval operation binding the contract event 0x031d8857fd5e877e23461ea5a8c500c2e5635e3f502c109fcbefba3d0448550e.
//
// Solidity: event filledMOVROrder(string Client, string UID)
func (_Hub *HubFilterer) FilterFilledMOVROrder(opts *bind.FilterOpts) (*HubFilledMOVROrderIterator, error) {

	logs, sub, err := _Hub.contract.FilterLogs(opts, "filledMOVROrder")
	if err != nil {
		return nil, err
	}
	return &HubFilledMOVROrderIterator{contract: _Hub.contract, event: "filledMOVROrder", logs: logs, sub: sub}, nil
}

// WatchFilledMOVROrder is a free log subscription operation binding the contract event 0x031d8857fd5e877e23461ea5a8c500c2e5635e3f502c109fcbefba3d0448550e.
//
// Solidity: event filledMOVROrder(string Client, string UID)
func (_Hub *HubFilterer) WatchFilledMOVROrder(opts *bind.WatchOpts, sink chan<- *HubFilledMOVROrder) (event.Subscription, error) {

	logs, sub, err := _Hub.contract.WatchLogs(opts, "filledMOVROrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubFilledMOVROrder)
				if err := _Hub.contract.UnpackLog(event, "filledMOVROrder", log); err != nil {
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

// ParseFilledMOVROrder is a log parse operation binding the contract event 0x031d8857fd5e877e23461ea5a8c500c2e5635e3f502c109fcbefba3d0448550e.
//
// Solidity: event filledMOVROrder(string Client, string UID)
func (_Hub *HubFilterer) ParseFilledMOVROrder(log types.Log) (*HubFilledMOVROrder, error) {
	event := new(HubFilledMOVROrder)
	if err := _Hub.contract.UnpackLog(event, "filledMOVROrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubFilledRMRKOrderIterator is returned from FilterFilledRMRKOrder and is used to iterate over the raw logs and unpacked data for FilledRMRKOrder events raised by the Hub contract.
type HubFilledRMRKOrderIterator struct {
	Event *HubFilledRMRKOrder // Event containing the contract specifics and raw log

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
func (it *HubFilledRMRKOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubFilledRMRKOrder)
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
		it.Event = new(HubFilledRMRKOrder)
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
func (it *HubFilledRMRKOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubFilledRMRKOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubFilledRMRKOrder represents a FilledRMRKOrder event raised by the Hub contract.
type HubFilledRMRKOrder struct {
	Client common.Address
	UID    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFilledRMRKOrder is a free log retrieval operation binding the contract event 0x2ae423548c2e48daa99c84afe6ceeb49aa42c00a6d70714ed78aadd0800fb727.
//
// Solidity: event filledRMRKOrder(address Client, string UID)
func (_Hub *HubFilterer) FilterFilledRMRKOrder(opts *bind.FilterOpts) (*HubFilledRMRKOrderIterator, error) {

	logs, sub, err := _Hub.contract.FilterLogs(opts, "filledRMRKOrder")
	if err != nil {
		return nil, err
	}
	return &HubFilledRMRKOrderIterator{contract: _Hub.contract, event: "filledRMRKOrder", logs: logs, sub: sub}, nil
}

// WatchFilledRMRKOrder is a free log subscription operation binding the contract event 0x2ae423548c2e48daa99c84afe6ceeb49aa42c00a6d70714ed78aadd0800fb727.
//
// Solidity: event filledRMRKOrder(address Client, string UID)
func (_Hub *HubFilterer) WatchFilledRMRKOrder(opts *bind.WatchOpts, sink chan<- *HubFilledRMRKOrder) (event.Subscription, error) {

	logs, sub, err := _Hub.contract.WatchLogs(opts, "filledRMRKOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubFilledRMRKOrder)
				if err := _Hub.contract.UnpackLog(event, "filledRMRKOrder", log); err != nil {
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

// ParseFilledRMRKOrder is a log parse operation binding the contract event 0x2ae423548c2e48daa99c84afe6ceeb49aa42c00a6d70714ed78aadd0800fb727.
//
// Solidity: event filledRMRKOrder(address Client, string UID)
func (_Hub *HubFilterer) ParseFilledRMRKOrder(log types.Log) (*HubFilledRMRKOrder, error) {
	event := new(HubFilledRMRKOrder)
	if err := _Hub.contract.UnpackLog(event, "filledRMRKOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
