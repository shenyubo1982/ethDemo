// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evidencecountract

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

// EvidenceContractStoreInfo is an auto generated low-level Go binding around an user-defined struct.
type EvidenceContractStoreInfo struct {
	KeyId   *big.Int
	TxTime  *big.Int
	Title   string
	Name    string
	Content string
}

// EvidencecountractMetaData contains all meta data concerning the Evidencecountract contract.
var EvidencecountractMetaData = &bind.MetaData{
	ABI: "[{\"outputs\":[],\"inputs\":[{\"name\":\"_admin\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"title\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"name\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"content\",\"internalType\":\"string\",\"type\":\"string\"}],\"name\":\"addInfo\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bool\",\"type\":\"bool\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"admins\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_newOwner\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"components\":[{\"name\":\"keyId\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"txTime\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"title\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"name\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"content\",\"internalType\":\"string\",\"type\":\"string\"}],\"name\":\"\",\"internalType\":\"structEvidenceContract.StoreInfo\",\"type\":\"tuple\"}],\"inputs\":[{\"name\":\"_keyId\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"getInfo\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[],\"name\":\"owner\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_admin\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EvidencecountractABI is the input ABI used to generate the binding from.
// Deprecated: Use EvidencecountractMetaData.ABI instead.
var EvidencecountractABI = EvidencecountractMetaData.ABI

// Evidencecountract is an auto generated Go binding around an Ethereum contract.
type Evidencecountract struct {
	EvidencecountractCaller     // Read-only binding to the contract
	EvidencecountractTransactor // Write-only binding to the contract
	EvidencecountractFilterer   // Log filterer for contract events
}

// EvidencecountractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EvidencecountractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidencecountractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EvidencecountractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidencecountractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EvidencecountractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidencecountractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EvidencecountractSession struct {
	Contract     *Evidencecountract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EvidencecountractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EvidencecountractCallerSession struct {
	Contract *EvidencecountractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EvidencecountractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EvidencecountractTransactorSession struct {
	Contract     *EvidencecountractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EvidencecountractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EvidencecountractRaw struct {
	Contract *Evidencecountract // Generic contract binding to access the raw methods on
}

// EvidencecountractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EvidencecountractCallerRaw struct {
	Contract *EvidencecountractCaller // Generic read-only contract binding to access the raw methods on
}

// EvidencecountractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EvidencecountractTransactorRaw struct {
	Contract *EvidencecountractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvidencecountract creates a new instance of Evidencecountract, bound to a specific deployed contract.
func NewEvidencecountract(address common.Address, backend bind.ContractBackend) (*Evidencecountract, error) {
	contract, err := bindEvidencecountract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Evidencecountract{EvidencecountractCaller: EvidencecountractCaller{contract: contract}, EvidencecountractTransactor: EvidencecountractTransactor{contract: contract}, EvidencecountractFilterer: EvidencecountractFilterer{contract: contract}}, nil
}

// NewEvidencecountractCaller creates a new read-only instance of Evidencecountract, bound to a specific deployed contract.
func NewEvidencecountractCaller(address common.Address, caller bind.ContractCaller) (*EvidencecountractCaller, error) {
	contract, err := bindEvidencecountract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvidencecountractCaller{contract: contract}, nil
}

// NewEvidencecountractTransactor creates a new write-only instance of Evidencecountract, bound to a specific deployed contract.
func NewEvidencecountractTransactor(address common.Address, transactor bind.ContractTransactor) (*EvidencecountractTransactor, error) {
	contract, err := bindEvidencecountract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvidencecountractTransactor{contract: contract}, nil
}

// NewEvidencecountractFilterer creates a new log filterer instance of Evidencecountract, bound to a specific deployed contract.
func NewEvidencecountractFilterer(address common.Address, filterer bind.ContractFilterer) (*EvidencecountractFilterer, error) {
	contract, err := bindEvidencecountract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvidencecountractFilterer{contract: contract}, nil
}

// bindEvidencecountract binds a generic wrapper to an already deployed contract.
func bindEvidencecountract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EvidencecountractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evidencecountract *EvidencecountractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evidencecountract.Contract.EvidencecountractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evidencecountract *EvidencecountractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evidencecountract.Contract.EvidencecountractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evidencecountract *EvidencecountractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evidencecountract.Contract.EvidencecountractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evidencecountract *EvidencecountractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evidencecountract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evidencecountract *EvidencecountractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evidencecountract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evidencecountract *EvidencecountractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evidencecountract.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Evidencecountract *EvidencecountractCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Evidencecountract.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call nbinding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Evidencecountract *EvidencecountractSession) Admins(arg0 common.Address) (bool, error) {
	return _Evidencecountract.Contract.Admins(&_Evidencecountract.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Evidencecountract *EvidencecountractCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _Evidencecountract.Contract.Admins(&_Evidencecountract.CallOpts, arg0)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _keyId) view returns((uint256,uint256,string,string,string))
func (_Evidencecountract *EvidencecountractCaller) GetInfo(opts *bind.CallOpts, _keyId *big.Int) (EvidenceContractStoreInfo, error) {
	var out []interface{}
	err := _Evidencecountract.contract.Call(opts, &out, "getInfo", _keyId)

	if err != nil {
		return *new(EvidenceContractStoreInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EvidenceContractStoreInfo)).(*EvidenceContractStoreInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _keyId) view returns((uint256,uint256,string,string,string))
func (_Evidencecountract *EvidencecountractSession) GetInfo(_keyId *big.Int) (EvidenceContractStoreInfo, error) {
	return _Evidencecountract.Contract.GetInfo(&_Evidencecountract.CallOpts, _keyId)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _keyId) view returns((uint256,uint256,string,string,string))
func (_Evidencecountract *EvidencecountractCallerSession) GetInfo(_keyId *big.Int) (EvidenceContractStoreInfo, error) {
	return _Evidencecountract.Contract.GetInfo(&_Evidencecountract.CallOpts, _keyId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evidencecountract *EvidencecountractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Evidencecountract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evidencecountract *EvidencecountractSession) Owner() (common.Address, error) {
	return _Evidencecountract.Contract.Owner(&_Evidencecountract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evidencecountract *EvidencecountractCallerSession) Owner() (common.Address, error) {
	return _Evidencecountract.Contract.Owner(&_Evidencecountract.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Evidencecountract *EvidencecountractTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Evidencecountract.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Evidencecountract *EvidencecountractSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecountract.Contract.AddAdmin(&_Evidencecountract.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Evidencecountract *EvidencecountractTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecountract.Contract.AddAdmin(&_Evidencecountract.TransactOpts, _admin)
}

// AddInfo is a paid mutator transaction binding the contract method 0xd7bbffb3.
//
// Solidity: function addInfo(string title, string name, string content) returns()
func (_Evidencecountract *EvidencecountractTransactor) AddInfo(opts *bind.TransactOpts, title string, name string, content string) (*types.Transaction, error) {
	return _Evidencecountract.contract.Transact(opts, "addInfo", title, name, content)
}

// AddInfo is a paid mutator transaction binding the contract method 0xd7bbffb3.
//
// Solidity: function addInfo(string title, string name, string content) returns()
func (_Evidencecountract *EvidencecountractSession) AddInfo(title string, name string, content string) (*types.Transaction, error) {
	return _Evidencecountract.Contract.AddInfo(&_Evidencecountract.TransactOpts, title, name, content)
}

// AddInfo is a paid mutator transaction binding the contract method 0xd7bbffb3.
//
// Solidity: function addInfo(string title, string name, string content) returns()
func (_Evidencecountract *EvidencecountractTransactorSession) AddInfo(title string, name string, content string) (*types.Transaction, error) {
	return _Evidencecountract.Contract.AddInfo(&_Evidencecountract.TransactOpts, title, name, content)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Evidencecountract *EvidencecountractTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Evidencecountract.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Evidencecountract *EvidencecountractSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Evidencecountract.Contract.ChangeOwner(&_Evidencecountract.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Evidencecountract *EvidencecountractTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Evidencecountract.Contract.ChangeOwner(&_Evidencecountract.TransactOpts, _newOwner)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Evidencecountract *EvidencecountractTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Evidencecountract.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Evidencecountract *EvidencecountractSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecountract.Contract.RemoveAdmin(&_Evidencecountract.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Evidencecountract *EvidencecountractTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecountract.Contract.RemoveAdmin(&_Evidencecountract.TransactOpts, _admin)
}
