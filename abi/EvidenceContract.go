// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evidencecontract

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

// EvidencecontractMetaData contains all meta data concerning the Evidencecontract contract.
var EvidencecontractMetaData = &bind.MetaData{
	ABI: "[{\"outputs\":[],\"inputs\":[{\"name\":\"_admin\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"title\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"name\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"content\",\"internalType\":\"string\",\"type\":\"string\"}],\"name\":\"addInfo\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bool\",\"type\":\"bool\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"admins\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_newOwner\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"components\":[{\"name\":\"keyId\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"txTime\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"title\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"name\",\"internalType\":\"string\",\"type\":\"string\"},{\"name\":\"content\",\"internalType\":\"string\",\"type\":\"string\"}],\"name\":\"\",\"internalType\":\"structEvidenceContract.StoreInfo\",\"type\":\"tuple\"}],\"inputs\":[{\"name\":\"_keyId\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"getInfo\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[],\"name\":\"owner\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_admin\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EvidencecontractABI is the input ABI used to generate the binding from.
// Deprecated: Use EvidencecontractMetaData.ABI instead.
var EvidencecontractABI = EvidencecontractMetaData.ABI

// Evidencecontract is an auto generated Go binding around an Ethereum contract.
type Evidencecontract struct {
	EvidencecontractCaller     // Read-only binding to the contract
	EvidencecontractTransactor // Write-only binding to the contract
	EvidencecontractFilterer   // Log filterer for contract events
}

// EvidencecontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EvidencecontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidencecontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EvidencecontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidencecontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EvidencecontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidencecontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EvidencecontractSession struct {
	Contract     *Evidencecontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvidencecontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EvidencecontractCallerSession struct {
	Contract *EvidencecontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EvidencecontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EvidencecontractTransactorSession struct {
	Contract     *EvidencecontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EvidencecontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EvidencecontractRaw struct {
	Contract *Evidencecontract // Generic contract binding to access the raw methods on
}

// EvidencecontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EvidencecontractCallerRaw struct {
	Contract *EvidencecontractCaller // Generic read-only contract binding to access the raw methods on
}

// EvidencecontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EvidencecontractTransactorRaw struct {
	Contract *EvidencecontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvidencecontract creates a new instance of Evidencecontract, bound to a specific deployed contract.
func NewEvidencecontract(address common.Address, backend bind.ContractBackend) (*Evidencecontract, error) {
	contract, err := bindEvidencecontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Evidencecontract{EvidencecontractCaller: EvidencecontractCaller{contract: contract}, EvidencecontractTransactor: EvidencecontractTransactor{contract: contract}, EvidencecontractFilterer: EvidencecontractFilterer{contract: contract}}, nil
}

// NewEvidencecontractCaller creates a new read-only instance of Evidencecontract, bound to a specific deployed contract.
func NewEvidencecontractCaller(address common.Address, caller bind.ContractCaller) (*EvidencecontractCaller, error) {
	contract, err := bindEvidencecontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvidencecontractCaller{contract: contract}, nil
}

// NewEvidencecontractTransactor creates a new write-only instance of Evidencecontract, bound to a specific deployed contract.
func NewEvidencecontractTransactor(address common.Address, transactor bind.ContractTransactor) (*EvidencecontractTransactor, error) {
	contract, err := bindEvidencecontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvidencecontractTransactor{contract: contract}, nil
}

// NewEvidencecontractFilterer creates a new log filterer instance of Evidencecontract, bound to a specific deployed contract.
func NewEvidencecontractFilterer(address common.Address, filterer bind.ContractFilterer) (*EvidencecontractFilterer, error) {
	contract, err := bindEvidencecontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvidencecontractFilterer{contract: contract}, nil
}

// bindEvidencecontract binds a generic wrapper to an already deployed contract.
func bindEvidencecontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EvidencecontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evidencecontract *EvidencecontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evidencecontract.Contract.EvidencecontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evidencecontract *EvidencecontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evidencecontract.Contract.EvidencecontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evidencecontract *EvidencecontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evidencecontract.Contract.EvidencecontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evidencecontract *EvidencecontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evidencecontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evidencecontract *EvidencecontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evidencecontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evidencecontract *EvidencecontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evidencecontract.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Evidencecontract *EvidencecontractCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Evidencecontract.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Evidencecontract *EvidencecontractSession) Admins(arg0 common.Address) (bool, error) {
	return _Evidencecontract.Contract.Admins(&_Evidencecontract.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Evidencecontract *EvidencecontractCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _Evidencecontract.Contract.Admins(&_Evidencecontract.CallOpts, arg0)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _keyId) view returns((uint256,uint256,string,string,string))
func (_Evidencecontract *EvidencecontractCaller) GetInfo(opts *bind.CallOpts, _keyId *big.Int) (EvidenceContractStoreInfo, error) {
	var out []interface{}
	err := _Evidencecontract.contract.Call(opts, &out, "getInfo", _keyId)

	if err != nil {
		return *new(EvidenceContractStoreInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EvidenceContractStoreInfo)).(*EvidenceContractStoreInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _keyId) view returns((uint256,uint256,string,string,string))
func (_Evidencecontract *EvidencecontractSession) GetInfo(_keyId *big.Int) (EvidenceContractStoreInfo, error) {
	return _Evidencecontract.Contract.GetInfo(&_Evidencecontract.CallOpts, _keyId)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _keyId) view returns((uint256,uint256,string,string,string))
func (_Evidencecontract *EvidencecontractCallerSession) GetInfo(_keyId *big.Int) (EvidenceContractStoreInfo, error) {
	return _Evidencecontract.Contract.GetInfo(&_Evidencecontract.CallOpts, _keyId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evidencecontract *EvidencecontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Evidencecontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evidencecontract *EvidencecontractSession) Owner() (common.Address, error) {
	return _Evidencecontract.Contract.Owner(&_Evidencecontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evidencecontract *EvidencecontractCallerSession) Owner() (common.Address, error) {
	return _Evidencecontract.Contract.Owner(&_Evidencecontract.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Evidencecontract *EvidencecontractTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Evidencecontract.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Evidencecontract *EvidencecontractSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecontract.Contract.AddAdmin(&_Evidencecontract.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Evidencecontract *EvidencecontractTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecontract.Contract.AddAdmin(&_Evidencecontract.TransactOpts, _admin)
}

// AddInfo is a paid mutator transaction binding the contract method 0xd7bbffb3.
//
// Solidity: function addInfo(string title, string name, string content) returns()
func (_Evidencecontract *EvidencecontractTransactor) AddInfo(opts *bind.TransactOpts, title string, name string, content string) (*types.Transaction, error) {
	return _Evidencecontract.contract.Transact(opts, "addInfo", title, name, content)
}

// AddInfo is a paid mutator transaction binding the contract method 0xd7bbffb3.
//
// Solidity: function addInfo(string title, string name, string content) returns()
func (_Evidencecontract *EvidencecontractSession) AddInfo(title string, name string, content string) (*types.Transaction, error) {
	return _Evidencecontract.Contract.AddInfo(&_Evidencecontract.TransactOpts, title, name, content)
}

// AddInfo is a paid mutator transaction binding the contract method 0xd7bbffb3.
//
// Solidity: function addInfo(string title, string name, string content) returns()
func (_Evidencecontract *EvidencecontractTransactorSession) AddInfo(title string, name string, content string) (*types.Transaction, error) {
	return _Evidencecontract.Contract.AddInfo(&_Evidencecontract.TransactOpts, title, name, content)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Evidencecontract *EvidencecontractTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Evidencecontract.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Evidencecontract *EvidencecontractSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Evidencecontract.Contract.ChangeOwner(&_Evidencecontract.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Evidencecontract *EvidencecontractTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Evidencecontract.Contract.ChangeOwner(&_Evidencecontract.TransactOpts, _newOwner)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Evidencecontract *EvidencecontractTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Evidencecontract.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Evidencecontract *EvidencecontractSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecontract.Contract.RemoveAdmin(&_Evidencecontract.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Evidencecontract *EvidencecontractTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Evidencecontract.Contract.RemoveAdmin(&_Evidencecontract.TransactOpts, _admin)
}
