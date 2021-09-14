// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goBind

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

// RootChainManagerContractMetaData contains all meta data concerning the RootChainManagerContract contract.
var RootChainManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"MetaTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tokenType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"predicateAddress\",\"type\":\"address\"}],\"name\":\"PredicateRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tokenType\",\"type\":\"bytes32\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC712_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHER_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAPPER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAP_TOKEN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"childChainManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"childToRootToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeperator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processedExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rootToChildToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"typeToPredicate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setupContractId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeEIP712\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStateSender\",\"type\":\"address\"}],\"name\":\"setStateSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateSenderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCheckpointManager\",\"type\":\"address\"}],\"name\":\"setCheckpointManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpointManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newChildChainManager\",\"type\":\"address\"}],\"name\":\"setChildChainManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"predicateAddress\",\"type\":\"address\"}],\"name\":\"registerPredicate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenType\",\"type\":\"bytes32\"}],\"name\":\"mapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"}],\"name\":\"cleanMapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenType\",\"type\":\"bytes32\"}],\"name\":\"remapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"depositEtherFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inputData\",\"type\":\"bytes\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RootChainManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RootChainManagerContractMetaData.ABI instead.
var RootChainManagerContractABI = RootChainManagerContractMetaData.ABI

// RootChainManagerContract is an auto generated Go binding around an Ethereum contract.
type RootChainManagerContract struct {
	RootChainManagerContractCaller     // Read-only binding to the contract
	RootChainManagerContractTransactor // Write-only binding to the contract
	RootChainManagerContractFilterer   // Log filterer for contract events
}

// RootChainManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainManagerContractSession struct {
	Contract     *RootChainManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RootChainManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainManagerContractCallerSession struct {
	Contract *RootChainManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// RootChainManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainManagerContractTransactorSession struct {
	Contract     *RootChainManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// RootChainManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainManagerContractRaw struct {
	Contract *RootChainManagerContract // Generic contract binding to access the raw methods on
}

// RootChainManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainManagerContractCallerRaw struct {
	Contract *RootChainManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainManagerContractTransactorRaw struct {
	Contract *RootChainManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChainManagerContract creates a new instance of RootChainManagerContract, bound to a specific deployed contract.
func NewRootChainManagerContract(address common.Address, backend bind.ContractBackend) (*RootChainManagerContract, error) {
	contract, err := bindRootChainManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContract{RootChainManagerContractCaller: RootChainManagerContractCaller{contract: contract}, RootChainManagerContractTransactor: RootChainManagerContractTransactor{contract: contract}, RootChainManagerContractFilterer: RootChainManagerContractFilterer{contract: contract}}, nil
}

// NewRootChainManagerContractCaller creates a new read-only instance of RootChainManagerContract, bound to a specific deployed contract.
func NewRootChainManagerContractCaller(address common.Address, caller bind.ContractCaller) (*RootChainManagerContractCaller, error) {
	contract, err := bindRootChainManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractCaller{contract: contract}, nil
}

// NewRootChainManagerContractTransactor creates a new write-only instance of RootChainManagerContract, bound to a specific deployed contract.
func NewRootChainManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainManagerContractTransactor, error) {
	contract, err := bindRootChainManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractTransactor{contract: contract}, nil
}

// NewRootChainManagerContractFilterer creates a new log filterer instance of RootChainManagerContract, bound to a specific deployed contract.
func NewRootChainManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainManagerContractFilterer, error) {
	contract, err := bindRootChainManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractFilterer{contract: contract}, nil
}

// bindRootChainManagerContract binds a generic wrapper to an already deployed contract.
func bindRootChainManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainManagerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainManagerContract *RootChainManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootChainManagerContract.Contract.RootChainManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainManagerContract *RootChainManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RootChainManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainManagerContract *RootChainManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RootChainManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainManagerContract *RootChainManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootChainManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainManagerContract *RootChainManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainManagerContract *RootChainManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RootChainManagerContract.Contract.DEFAULTADMINROLE(&_RootChainManagerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RootChainManagerContract.Contract.DEFAULTADMINROLE(&_RootChainManagerContract.CallOpts)
}

// DEPOSIT is a free data retrieval call binding the contract method 0xd81c8e52.
//
// Solidity: function DEPOSIT() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCaller) DEPOSIT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "DEPOSIT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSIT is a free data retrieval call binding the contract method 0xd81c8e52.
//
// Solidity: function DEPOSIT() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractSession) DEPOSIT() ([32]byte, error) {
	return _RootChainManagerContract.Contract.DEPOSIT(&_RootChainManagerContract.CallOpts)
}

// DEPOSIT is a free data retrieval call binding the contract method 0xd81c8e52.
//
// Solidity: function DEPOSIT() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) DEPOSIT() ([32]byte, error) {
	return _RootChainManagerContract.Contract.DEPOSIT(&_RootChainManagerContract.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_RootChainManagerContract *RootChainManagerContractCaller) ERC712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "ERC712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_RootChainManagerContract *RootChainManagerContractSession) ERC712VERSION() (string, error) {
	return _RootChainManagerContract.Contract.ERC712VERSION(&_RootChainManagerContract.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) ERC712VERSION() (string, error) {
	return _RootChainManagerContract.Contract.ERC712VERSION(&_RootChainManagerContract.CallOpts)
}

// ETHERADDRESS is a free data retrieval call binding the contract method 0xcf1d21c0.
//
// Solidity: function ETHER_ADDRESS() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) ETHERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "ETHER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHERADDRESS is a free data retrieval call binding the contract method 0xcf1d21c0.
//
// Solidity: function ETHER_ADDRESS() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) ETHERADDRESS() (common.Address, error) {
	return _RootChainManagerContract.Contract.ETHERADDRESS(&_RootChainManagerContract.CallOpts)
}

// ETHERADDRESS is a free data retrieval call binding the contract method 0xcf1d21c0.
//
// Solidity: function ETHER_ADDRESS() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) ETHERADDRESS() (common.Address, error) {
	return _RootChainManagerContract.Contract.ETHERADDRESS(&_RootChainManagerContract.CallOpts)
}

// MAPPERROLE is a free data retrieval call binding the contract method 0x568b80b5.
//
// Solidity: function MAPPER_ROLE() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCaller) MAPPERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "MAPPER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAPPERROLE is a free data retrieval call binding the contract method 0x568b80b5.
//
// Solidity: function MAPPER_ROLE() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractSession) MAPPERROLE() ([32]byte, error) {
	return _RootChainManagerContract.Contract.MAPPERROLE(&_RootChainManagerContract.CallOpts)
}

// MAPPERROLE is a free data retrieval call binding the contract method 0x568b80b5.
//
// Solidity: function MAPPER_ROLE() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) MAPPERROLE() ([32]byte, error) {
	return _RootChainManagerContract.Contract.MAPPERROLE(&_RootChainManagerContract.CallOpts)
}

// MAPTOKEN is a free data retrieval call binding the contract method 0x886a69ba.
//
// Solidity: function MAP_TOKEN() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCaller) MAPTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "MAP_TOKEN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAPTOKEN is a free data retrieval call binding the contract method 0x886a69ba.
//
// Solidity: function MAP_TOKEN() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractSession) MAPTOKEN() ([32]byte, error) {
	return _RootChainManagerContract.Contract.MAPTOKEN(&_RootChainManagerContract.CallOpts)
}

// MAPTOKEN is a free data retrieval call binding the contract method 0x886a69ba.
//
// Solidity: function MAP_TOKEN() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) MAPTOKEN() ([32]byte, error) {
	return _RootChainManagerContract.Contract.MAPTOKEN(&_RootChainManagerContract.CallOpts)
}

// CheckpointManagerAddress is a free data retrieval call binding the contract method 0x3138b6f1.
//
// Solidity: function checkpointManagerAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) CheckpointManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "checkpointManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckpointManagerAddress is a free data retrieval call binding the contract method 0x3138b6f1.
//
// Solidity: function checkpointManagerAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) CheckpointManagerAddress() (common.Address, error) {
	return _RootChainManagerContract.Contract.CheckpointManagerAddress(&_RootChainManagerContract.CallOpts)
}

// CheckpointManagerAddress is a free data retrieval call binding the contract method 0x3138b6f1.
//
// Solidity: function checkpointManagerAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) CheckpointManagerAddress() (common.Address, error) {
	return _RootChainManagerContract.Contract.CheckpointManagerAddress(&_RootChainManagerContract.CallOpts)
}

// ChildChainManagerAddress is a free data retrieval call binding the contract method 0x04967702.
//
// Solidity: function childChainManagerAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) ChildChainManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "childChainManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildChainManagerAddress is a free data retrieval call binding the contract method 0x04967702.
//
// Solidity: function childChainManagerAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) ChildChainManagerAddress() (common.Address, error) {
	return _RootChainManagerContract.Contract.ChildChainManagerAddress(&_RootChainManagerContract.CallOpts)
}

// ChildChainManagerAddress is a free data retrieval call binding the contract method 0x04967702.
//
// Solidity: function childChainManagerAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) ChildChainManagerAddress() (common.Address, error) {
	return _RootChainManagerContract.Contract.ChildChainManagerAddress(&_RootChainManagerContract.CallOpts)
}

// ChildToRootToken is a free data retrieval call binding the contract method 0x6e86b770.
//
// Solidity: function childToRootToken(address ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) ChildToRootToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "childToRootToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildToRootToken is a free data retrieval call binding the contract method 0x6e86b770.
//
// Solidity: function childToRootToken(address ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) ChildToRootToken(arg0 common.Address) (common.Address, error) {
	return _RootChainManagerContract.Contract.ChildToRootToken(&_RootChainManagerContract.CallOpts, arg0)
}

// ChildToRootToken is a free data retrieval call binding the contract method 0x6e86b770.
//
// Solidity: function childToRootToken(address ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) ChildToRootToken(arg0 common.Address) (common.Address, error) {
	return _RootChainManagerContract.Contract.ChildToRootToken(&_RootChainManagerContract.CallOpts, arg0)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_RootChainManagerContract *RootChainManagerContractCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_RootChainManagerContract *RootChainManagerContractSession) GetChainId() (*big.Int, error) {
	return _RootChainManagerContract.Contract.GetChainId(&_RootChainManagerContract.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) GetChainId() (*big.Int, error) {
	return _RootChainManagerContract.Contract.GetChainId(&_RootChainManagerContract.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCaller) GetDomainSeperator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "getDomainSeperator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractSession) GetDomainSeperator() ([32]byte, error) {
	return _RootChainManagerContract.Contract.GetDomainSeperator(&_RootChainManagerContract.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) GetDomainSeperator() ([32]byte, error) {
	return _RootChainManagerContract.Contract.GetDomainSeperator(&_RootChainManagerContract.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_RootChainManagerContract *RootChainManagerContractCaller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_RootChainManagerContract *RootChainManagerContractSession) GetNonce(user common.Address) (*big.Int, error) {
	return _RootChainManagerContract.Contract.GetNonce(&_RootChainManagerContract.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _RootChainManagerContract.Contract.GetNonce(&_RootChainManagerContract.CallOpts, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RootChainManagerContract.Contract.GetRoleAdmin(&_RootChainManagerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RootChainManagerContract.Contract.GetRoleAdmin(&_RootChainManagerContract.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RootChainManagerContract.Contract.GetRoleMember(&_RootChainManagerContract.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RootChainManagerContract.Contract.GetRoleMember(&_RootChainManagerContract.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RootChainManagerContract *RootChainManagerContractCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RootChainManagerContract *RootChainManagerContractSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RootChainManagerContract.Contract.GetRoleMemberCount(&_RootChainManagerContract.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RootChainManagerContract.Contract.GetRoleMemberCount(&_RootChainManagerContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RootChainManagerContract *RootChainManagerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RootChainManagerContract *RootChainManagerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RootChainManagerContract.Contract.HasRole(&_RootChainManagerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RootChainManagerContract.Contract.HasRole(&_RootChainManagerContract.CallOpts, role, account)
}

// ProcessedExits is a free data retrieval call binding the contract method 0x607f2d42.
//
// Solidity: function processedExits(bytes32 ) view returns(bool)
func (_RootChainManagerContract *RootChainManagerContractCaller) ProcessedExits(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "processedExits", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedExits is a free data retrieval call binding the contract method 0x607f2d42.
//
// Solidity: function processedExits(bytes32 ) view returns(bool)
func (_RootChainManagerContract *RootChainManagerContractSession) ProcessedExits(arg0 [32]byte) (bool, error) {
	return _RootChainManagerContract.Contract.ProcessedExits(&_RootChainManagerContract.CallOpts, arg0)
}

// ProcessedExits is a free data retrieval call binding the contract method 0x607f2d42.
//
// Solidity: function processedExits(bytes32 ) view returns(bool)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) ProcessedExits(arg0 [32]byte) (bool, error) {
	return _RootChainManagerContract.Contract.ProcessedExits(&_RootChainManagerContract.CallOpts, arg0)
}

// RootToChildToken is a free data retrieval call binding the contract method 0xea60c7c4.
//
// Solidity: function rootToChildToken(address ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) RootToChildToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "rootToChildToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootToChildToken is a free data retrieval call binding the contract method 0xea60c7c4.
//
// Solidity: function rootToChildToken(address ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) RootToChildToken(arg0 common.Address) (common.Address, error) {
	return _RootChainManagerContract.Contract.RootToChildToken(&_RootChainManagerContract.CallOpts, arg0)
}

// RootToChildToken is a free data retrieval call binding the contract method 0xea60c7c4.
//
// Solidity: function rootToChildToken(address ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) RootToChildToken(arg0 common.Address) (common.Address, error) {
	return _RootChainManagerContract.Contract.RootToChildToken(&_RootChainManagerContract.CallOpts, arg0)
}

// StateSenderAddress is a free data retrieval call binding the contract method 0xe2c49de1.
//
// Solidity: function stateSenderAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) StateSenderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "stateSenderAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateSenderAddress is a free data retrieval call binding the contract method 0xe2c49de1.
//
// Solidity: function stateSenderAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) StateSenderAddress() (common.Address, error) {
	return _RootChainManagerContract.Contract.StateSenderAddress(&_RootChainManagerContract.CallOpts)
}

// StateSenderAddress is a free data retrieval call binding the contract method 0xe2c49de1.
//
// Solidity: function stateSenderAddress() view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) StateSenderAddress() (common.Address, error) {
	return _RootChainManagerContract.Contract.StateSenderAddress(&_RootChainManagerContract.CallOpts)
}

// TokenToType is a free data retrieval call binding the contract method 0xe43009a6.
//
// Solidity: function tokenToType(address ) view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCaller) TokenToType(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "tokenToType", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenToType is a free data retrieval call binding the contract method 0xe43009a6.
//
// Solidity: function tokenToType(address ) view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractSession) TokenToType(arg0 common.Address) ([32]byte, error) {
	return _RootChainManagerContract.Contract.TokenToType(&_RootChainManagerContract.CallOpts, arg0)
}

// TokenToType is a free data retrieval call binding the contract method 0xe43009a6.
//
// Solidity: function tokenToType(address ) view returns(bytes32)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) TokenToType(arg0 common.Address) ([32]byte, error) {
	return _RootChainManagerContract.Contract.TokenToType(&_RootChainManagerContract.CallOpts, arg0)
}

// TypeToPredicate is a free data retrieval call binding the contract method 0xe66f9603.
//
// Solidity: function typeToPredicate(bytes32 ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCaller) TypeToPredicate(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _RootChainManagerContract.contract.Call(opts, &out, "typeToPredicate", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TypeToPredicate is a free data retrieval call binding the contract method 0xe66f9603.
//
// Solidity: function typeToPredicate(bytes32 ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractSession) TypeToPredicate(arg0 [32]byte) (common.Address, error) {
	return _RootChainManagerContract.Contract.TypeToPredicate(&_RootChainManagerContract.CallOpts, arg0)
}

// TypeToPredicate is a free data retrieval call binding the contract method 0xe66f9603.
//
// Solidity: function typeToPredicate(bytes32 ) view returns(address)
func (_RootChainManagerContract *RootChainManagerContractCallerSession) TypeToPredicate(arg0 [32]byte) (common.Address, error) {
	return _RootChainManagerContract.Contract.TypeToPredicate(&_RootChainManagerContract.CallOpts, arg0)
}

// CleanMapToken is a paid mutator transaction binding the contract method 0x0c3894bb.
//
// Solidity: function cleanMapToken(address rootToken, address childToken) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) CleanMapToken(opts *bind.TransactOpts, rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "cleanMapToken", rootToken, childToken)
}

// CleanMapToken is a paid mutator transaction binding the contract method 0x0c3894bb.
//
// Solidity: function cleanMapToken(address rootToken, address childToken) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) CleanMapToken(rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.CleanMapToken(&_RootChainManagerContract.TransactOpts, rootToken, childToken)
}

// CleanMapToken is a paid mutator transaction binding the contract method 0x0c3894bb.
//
// Solidity: function cleanMapToken(address rootToken, address childToken) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) CleanMapToken(rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.CleanMapToken(&_RootChainManagerContract.TransactOpts, rootToken, childToken)
}

// DepositEtherFor is a paid mutator transaction binding the contract method 0x4faa8a26.
//
// Solidity: function depositEtherFor(address user) payable returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) DepositEtherFor(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "depositEtherFor", user)
}

// DepositEtherFor is a paid mutator transaction binding the contract method 0x4faa8a26.
//
// Solidity: function depositEtherFor(address user) payable returns()
func (_RootChainManagerContract *RootChainManagerContractSession) DepositEtherFor(user common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.DepositEtherFor(&_RootChainManagerContract.TransactOpts, user)
}

// DepositEtherFor is a paid mutator transaction binding the contract method 0x4faa8a26.
//
// Solidity: function depositEtherFor(address user) payable returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) DepositEtherFor(user common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.DepositEtherFor(&_RootChainManagerContract.TransactOpts, user)
}

// DepositFor is a paid mutator transaction binding the contract method 0xe3dec8fb.
//
// Solidity: function depositFor(address user, address rootToken, bytes depositData) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) DepositFor(opts *bind.TransactOpts, user common.Address, rootToken common.Address, depositData []byte) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "depositFor", user, rootToken, depositData)
}

// DepositFor is a paid mutator transaction binding the contract method 0xe3dec8fb.
//
// Solidity: function depositFor(address user, address rootToken, bytes depositData) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) DepositFor(user common.Address, rootToken common.Address, depositData []byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.DepositFor(&_RootChainManagerContract.TransactOpts, user, rootToken, depositData)
}

// DepositFor is a paid mutator transaction binding the contract method 0xe3dec8fb.
//
// Solidity: function depositFor(address user, address rootToken, bytes depositData) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) DepositFor(user common.Address, rootToken common.Address, depositData []byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.DepositFor(&_RootChainManagerContract.TransactOpts, user, rootToken, depositData)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_RootChainManagerContract *RootChainManagerContractTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "executeMetaTransaction", userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_RootChainManagerContract *RootChainManagerContractSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.ExecuteMetaTransaction(&_RootChainManagerContract.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.ExecuteMetaTransaction(&_RootChainManagerContract.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// Exit is a paid mutator transaction binding the contract method 0x3805550f.
//
// Solidity: function exit(bytes inputData) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) Exit(opts *bind.TransactOpts, inputData []byte) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "exit", inputData)
}

// Exit is a paid mutator transaction binding the contract method 0x3805550f.
//
// Solidity: function exit(bytes inputData) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) Exit(inputData []byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.Exit(&_RootChainManagerContract.TransactOpts, inputData)
}

// Exit is a paid mutator transaction binding the contract method 0x3805550f.
//
// Solidity: function exit(bytes inputData) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) Exit(inputData []byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.Exit(&_RootChainManagerContract.TransactOpts, inputData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.GrantRole(&_RootChainManagerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.GrantRole(&_RootChainManagerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.Initialize(&_RootChainManagerContract.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.Initialize(&_RootChainManagerContract.TransactOpts, _owner)
}

// InitializeEIP712 is a paid mutator transaction binding the contract method 0x630fcbfb.
//
// Solidity: function initializeEIP712() returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) InitializeEIP712(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "initializeEIP712")
}

// InitializeEIP712 is a paid mutator transaction binding the contract method 0x630fcbfb.
//
// Solidity: function initializeEIP712() returns()
func (_RootChainManagerContract *RootChainManagerContractSession) InitializeEIP712() (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.InitializeEIP712(&_RootChainManagerContract.TransactOpts)
}

// InitializeEIP712 is a paid mutator transaction binding the contract method 0x630fcbfb.
//
// Solidity: function initializeEIP712() returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) InitializeEIP712() (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.InitializeEIP712(&_RootChainManagerContract.TransactOpts)
}

// MapToken is a paid mutator transaction binding the contract method 0x9173b139.
//
// Solidity: function mapToken(address rootToken, address childToken, bytes32 tokenType) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) MapToken(opts *bind.TransactOpts, rootToken common.Address, childToken common.Address, tokenType [32]byte) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "mapToken", rootToken, childToken, tokenType)
}

// MapToken is a paid mutator transaction binding the contract method 0x9173b139.
//
// Solidity: function mapToken(address rootToken, address childToken, bytes32 tokenType) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) MapToken(rootToken common.Address, childToken common.Address, tokenType [32]byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.MapToken(&_RootChainManagerContract.TransactOpts, rootToken, childToken, tokenType)
}

// MapToken is a paid mutator transaction binding the contract method 0x9173b139.
//
// Solidity: function mapToken(address rootToken, address childToken, bytes32 tokenType) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) MapToken(rootToken common.Address, childToken common.Address, tokenType [32]byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.MapToken(&_RootChainManagerContract.TransactOpts, rootToken, childToken, tokenType)
}

// RegisterPredicate is a paid mutator transaction binding the contract method 0x0c598220.
//
// Solidity: function registerPredicate(bytes32 tokenType, address predicateAddress) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) RegisterPredicate(opts *bind.TransactOpts, tokenType [32]byte, predicateAddress common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "registerPredicate", tokenType, predicateAddress)
}

// RegisterPredicate is a paid mutator transaction binding the contract method 0x0c598220.
//
// Solidity: function registerPredicate(bytes32 tokenType, address predicateAddress) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) RegisterPredicate(tokenType [32]byte, predicateAddress common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RegisterPredicate(&_RootChainManagerContract.TransactOpts, tokenType, predicateAddress)
}

// RegisterPredicate is a paid mutator transaction binding the contract method 0x0c598220.
//
// Solidity: function registerPredicate(bytes32 tokenType, address predicateAddress) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) RegisterPredicate(tokenType [32]byte, predicateAddress common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RegisterPredicate(&_RootChainManagerContract.TransactOpts, tokenType, predicateAddress)
}

// RemapToken is a paid mutator transaction binding the contract method 0xd233a3c7.
//
// Solidity: function remapToken(address rootToken, address childToken, bytes32 tokenType) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) RemapToken(opts *bind.TransactOpts, rootToken common.Address, childToken common.Address, tokenType [32]byte) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "remapToken", rootToken, childToken, tokenType)
}

// RemapToken is a paid mutator transaction binding the contract method 0xd233a3c7.
//
// Solidity: function remapToken(address rootToken, address childToken, bytes32 tokenType) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) RemapToken(rootToken common.Address, childToken common.Address, tokenType [32]byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RemapToken(&_RootChainManagerContract.TransactOpts, rootToken, childToken, tokenType)
}

// RemapToken is a paid mutator transaction binding the contract method 0xd233a3c7.
//
// Solidity: function remapToken(address rootToken, address childToken, bytes32 tokenType) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) RemapToken(rootToken common.Address, childToken common.Address, tokenType [32]byte) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RemapToken(&_RootChainManagerContract.TransactOpts, rootToken, childToken, tokenType)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RenounceRole(&_RootChainManagerContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RenounceRole(&_RootChainManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RevokeRole(&_RootChainManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.RevokeRole(&_RootChainManagerContract.TransactOpts, role, account)
}

// SetCheckpointManager is a paid mutator transaction binding the contract method 0xbc08452b.
//
// Solidity: function setCheckpointManager(address newCheckpointManager) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) SetCheckpointManager(opts *bind.TransactOpts, newCheckpointManager common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "setCheckpointManager", newCheckpointManager)
}

// SetCheckpointManager is a paid mutator transaction binding the contract method 0xbc08452b.
//
// Solidity: function setCheckpointManager(address newCheckpointManager) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) SetCheckpointManager(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetCheckpointManager(&_RootChainManagerContract.TransactOpts, newCheckpointManager)
}

// SetCheckpointManager is a paid mutator transaction binding the contract method 0xbc08452b.
//
// Solidity: function setCheckpointManager(address newCheckpointManager) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) SetCheckpointManager(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetCheckpointManager(&_RootChainManagerContract.TransactOpts, newCheckpointManager)
}

// SetChildChainManagerAddress is a paid mutator transaction binding the contract method 0xdc993a23.
//
// Solidity: function setChildChainManagerAddress(address newChildChainManager) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) SetChildChainManagerAddress(opts *bind.TransactOpts, newChildChainManager common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "setChildChainManagerAddress", newChildChainManager)
}

// SetChildChainManagerAddress is a paid mutator transaction binding the contract method 0xdc993a23.
//
// Solidity: function setChildChainManagerAddress(address newChildChainManager) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) SetChildChainManagerAddress(newChildChainManager common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetChildChainManagerAddress(&_RootChainManagerContract.TransactOpts, newChildChainManager)
}

// SetChildChainManagerAddress is a paid mutator transaction binding the contract method 0xdc993a23.
//
// Solidity: function setChildChainManagerAddress(address newChildChainManager) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) SetChildChainManagerAddress(newChildChainManager common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetChildChainManagerAddress(&_RootChainManagerContract.TransactOpts, newChildChainManager)
}

// SetStateSender is a paid mutator transaction binding the contract method 0x6cb136b0.
//
// Solidity: function setStateSender(address newStateSender) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) SetStateSender(opts *bind.TransactOpts, newStateSender common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "setStateSender", newStateSender)
}

// SetStateSender is a paid mutator transaction binding the contract method 0x6cb136b0.
//
// Solidity: function setStateSender(address newStateSender) returns()
func (_RootChainManagerContract *RootChainManagerContractSession) SetStateSender(newStateSender common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetStateSender(&_RootChainManagerContract.TransactOpts, newStateSender)
}

// SetStateSender is a paid mutator transaction binding the contract method 0x6cb136b0.
//
// Solidity: function setStateSender(address newStateSender) returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) SetStateSender(newStateSender common.Address) (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetStateSender(&_RootChainManagerContract.TransactOpts, newStateSender)
}

// SetupContractId is a paid mutator transaction binding the contract method 0xb4b4f63e.
//
// Solidity: function setupContractId() returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) SetupContractId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.Transact(opts, "setupContractId")
}

// SetupContractId is a paid mutator transaction binding the contract method 0xb4b4f63e.
//
// Solidity: function setupContractId() returns()
func (_RootChainManagerContract *RootChainManagerContractSession) SetupContractId() (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetupContractId(&_RootChainManagerContract.TransactOpts)
}

// SetupContractId is a paid mutator transaction binding the contract method 0xb4b4f63e.
//
// Solidity: function setupContractId() returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) SetupContractId() (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.SetupContractId(&_RootChainManagerContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootChainManagerContract *RootChainManagerContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainManagerContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootChainManagerContract *RootChainManagerContractSession) Receive() (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.Receive(&_RootChainManagerContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootChainManagerContract *RootChainManagerContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RootChainManagerContract.Contract.Receive(&_RootChainManagerContract.TransactOpts)
}

// RootChainManagerContractMetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the RootChainManagerContract contract.
type RootChainManagerContractMetaTransactionExecutedIterator struct {
	Event *RootChainManagerContractMetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *RootChainManagerContractMetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainManagerContractMetaTransactionExecuted)
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
		it.Event = new(RootChainManagerContractMetaTransactionExecuted)
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
func (it *RootChainManagerContractMetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainManagerContractMetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainManagerContractMetaTransactionExecuted represents a MetaTransactionExecuted event raised by the RootChainManagerContract contract.
type RootChainManagerContractMetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_RootChainManagerContract *RootChainManagerContractFilterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*RootChainManagerContractMetaTransactionExecutedIterator, error) {

	logs, sub, err := _RootChainManagerContract.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractMetaTransactionExecutedIterator{contract: _RootChainManagerContract.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_RootChainManagerContract *RootChainManagerContractFilterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *RootChainManagerContractMetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _RootChainManagerContract.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainManagerContractMetaTransactionExecuted)
				if err := _RootChainManagerContract.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_RootChainManagerContract *RootChainManagerContractFilterer) ParseMetaTransactionExecuted(log types.Log) (*RootChainManagerContractMetaTransactionExecuted, error) {
	event := new(RootChainManagerContractMetaTransactionExecuted)
	if err := _RootChainManagerContract.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootChainManagerContractPredicateRegisteredIterator is returned from FilterPredicateRegistered and is used to iterate over the raw logs and unpacked data for PredicateRegistered events raised by the RootChainManagerContract contract.
type RootChainManagerContractPredicateRegisteredIterator struct {
	Event *RootChainManagerContractPredicateRegistered // Event containing the contract specifics and raw log

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
func (it *RootChainManagerContractPredicateRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainManagerContractPredicateRegistered)
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
		it.Event = new(RootChainManagerContractPredicateRegistered)
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
func (it *RootChainManagerContractPredicateRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainManagerContractPredicateRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainManagerContractPredicateRegistered represents a PredicateRegistered event raised by the RootChainManagerContract contract.
type RootChainManagerContractPredicateRegistered struct {
	TokenType        [32]byte
	PredicateAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPredicateRegistered is a free log retrieval operation binding the contract event 0x8643692ae1c12ec91fa18e50b82ed93fa314f580999a236824db6de9ae0d839b.
//
// Solidity: event PredicateRegistered(bytes32 indexed tokenType, address indexed predicateAddress)
func (_RootChainManagerContract *RootChainManagerContractFilterer) FilterPredicateRegistered(opts *bind.FilterOpts, tokenType [][32]byte, predicateAddress []common.Address) (*RootChainManagerContractPredicateRegisteredIterator, error) {

	var tokenTypeRule []interface{}
	for _, tokenTypeItem := range tokenType {
		tokenTypeRule = append(tokenTypeRule, tokenTypeItem)
	}
	var predicateAddressRule []interface{}
	for _, predicateAddressItem := range predicateAddress {
		predicateAddressRule = append(predicateAddressRule, predicateAddressItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.FilterLogs(opts, "PredicateRegistered", tokenTypeRule, predicateAddressRule)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractPredicateRegisteredIterator{contract: _RootChainManagerContract.contract, event: "PredicateRegistered", logs: logs, sub: sub}, nil
}

// WatchPredicateRegistered is a free log subscription operation binding the contract event 0x8643692ae1c12ec91fa18e50b82ed93fa314f580999a236824db6de9ae0d839b.
//
// Solidity: event PredicateRegistered(bytes32 indexed tokenType, address indexed predicateAddress)
func (_RootChainManagerContract *RootChainManagerContractFilterer) WatchPredicateRegistered(opts *bind.WatchOpts, sink chan<- *RootChainManagerContractPredicateRegistered, tokenType [][32]byte, predicateAddress []common.Address) (event.Subscription, error) {

	var tokenTypeRule []interface{}
	for _, tokenTypeItem := range tokenType {
		tokenTypeRule = append(tokenTypeRule, tokenTypeItem)
	}
	var predicateAddressRule []interface{}
	for _, predicateAddressItem := range predicateAddress {
		predicateAddressRule = append(predicateAddressRule, predicateAddressItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.WatchLogs(opts, "PredicateRegistered", tokenTypeRule, predicateAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainManagerContractPredicateRegistered)
				if err := _RootChainManagerContract.contract.UnpackLog(event, "PredicateRegistered", log); err != nil {
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

// ParsePredicateRegistered is a log parse operation binding the contract event 0x8643692ae1c12ec91fa18e50b82ed93fa314f580999a236824db6de9ae0d839b.
//
// Solidity: event PredicateRegistered(bytes32 indexed tokenType, address indexed predicateAddress)
func (_RootChainManagerContract *RootChainManagerContractFilterer) ParsePredicateRegistered(log types.Log) (*RootChainManagerContractPredicateRegistered, error) {
	event := new(RootChainManagerContractPredicateRegistered)
	if err := _RootChainManagerContract.contract.UnpackLog(event, "PredicateRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootChainManagerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RootChainManagerContract contract.
type RootChainManagerContractRoleAdminChangedIterator struct {
	Event *RootChainManagerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RootChainManagerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainManagerContractRoleAdminChanged)
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
		it.Event = new(RootChainManagerContractRoleAdminChanged)
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
func (it *RootChainManagerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainManagerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainManagerContractRoleAdminChanged represents a RoleAdminChanged event raised by the RootChainManagerContract contract.
type RootChainManagerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RootChainManagerContract *RootChainManagerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RootChainManagerContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractRoleAdminChangedIterator{contract: _RootChainManagerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RootChainManagerContract *RootChainManagerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RootChainManagerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainManagerContractRoleAdminChanged)
				if err := _RootChainManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RootChainManagerContract *RootChainManagerContractFilterer) ParseRoleAdminChanged(log types.Log) (*RootChainManagerContractRoleAdminChanged, error) {
	event := new(RootChainManagerContractRoleAdminChanged)
	if err := _RootChainManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootChainManagerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RootChainManagerContract contract.
type RootChainManagerContractRoleGrantedIterator struct {
	Event *RootChainManagerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *RootChainManagerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainManagerContractRoleGranted)
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
		it.Event = new(RootChainManagerContractRoleGranted)
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
func (it *RootChainManagerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainManagerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainManagerContractRoleGranted represents a RoleGranted event raised by the RootChainManagerContract contract.
type RootChainManagerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RootChainManagerContract *RootChainManagerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RootChainManagerContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractRoleGrantedIterator{contract: _RootChainManagerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RootChainManagerContract *RootChainManagerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RootChainManagerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainManagerContractRoleGranted)
				if err := _RootChainManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RootChainManagerContract *RootChainManagerContractFilterer) ParseRoleGranted(log types.Log) (*RootChainManagerContractRoleGranted, error) {
	event := new(RootChainManagerContractRoleGranted)
	if err := _RootChainManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootChainManagerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RootChainManagerContract contract.
type RootChainManagerContractRoleRevokedIterator struct {
	Event *RootChainManagerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RootChainManagerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainManagerContractRoleRevoked)
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
		it.Event = new(RootChainManagerContractRoleRevoked)
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
func (it *RootChainManagerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainManagerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainManagerContractRoleRevoked represents a RoleRevoked event raised by the RootChainManagerContract contract.
type RootChainManagerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RootChainManagerContract *RootChainManagerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RootChainManagerContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractRoleRevokedIterator{contract: _RootChainManagerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RootChainManagerContract *RootChainManagerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RootChainManagerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainManagerContractRoleRevoked)
				if err := _RootChainManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RootChainManagerContract *RootChainManagerContractFilterer) ParseRoleRevoked(log types.Log) (*RootChainManagerContractRoleRevoked, error) {
	event := new(RootChainManagerContractRoleRevoked)
	if err := _RootChainManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootChainManagerContractTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the RootChainManagerContract contract.
type RootChainManagerContractTokenMappedIterator struct {
	Event *RootChainManagerContractTokenMapped // Event containing the contract specifics and raw log

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
func (it *RootChainManagerContractTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainManagerContractTokenMapped)
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
		it.Event = new(RootChainManagerContractTokenMapped)
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
func (it *RootChainManagerContractTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainManagerContractTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainManagerContractTokenMapped represents a TokenMapped event raised by the RootChainManagerContract contract.
type RootChainManagerContractTokenMapped struct {
	RootToken  common.Address
	ChildToken common.Address
	TokenType  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x9e651a8866fbea043e911d816ec254b0e3c992c06fff32d605e72362d6023bd9.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken, bytes32 indexed tokenType)
func (_RootChainManagerContract *RootChainManagerContractFilterer) FilterTokenMapped(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address, tokenType [][32]byte) (*RootChainManagerContractTokenMappedIterator, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}
	var tokenTypeRule []interface{}
	for _, tokenTypeItem := range tokenType {
		tokenTypeRule = append(tokenTypeRule, tokenTypeItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.FilterLogs(opts, "TokenMapped", rootTokenRule, childTokenRule, tokenTypeRule)
	if err != nil {
		return nil, err
	}
	return &RootChainManagerContractTokenMappedIterator{contract: _RootChainManagerContract.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x9e651a8866fbea043e911d816ec254b0e3c992c06fff32d605e72362d6023bd9.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken, bytes32 indexed tokenType)
func (_RootChainManagerContract *RootChainManagerContractFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *RootChainManagerContractTokenMapped, rootToken []common.Address, childToken []common.Address, tokenType [][32]byte) (event.Subscription, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}
	var tokenTypeRule []interface{}
	for _, tokenTypeItem := range tokenType {
		tokenTypeRule = append(tokenTypeRule, tokenTypeItem)
	}

	logs, sub, err := _RootChainManagerContract.contract.WatchLogs(opts, "TokenMapped", rootTokenRule, childTokenRule, tokenTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainManagerContractTokenMapped)
				if err := _RootChainManagerContract.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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

// ParseTokenMapped is a log parse operation binding the contract event 0x9e651a8866fbea043e911d816ec254b0e3c992c06fff32d605e72362d6023bd9.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken, bytes32 indexed tokenType)
func (_RootChainManagerContract *RootChainManagerContractFilterer) ParseTokenMapped(log types.Log) (*RootChainManagerContractTokenMapped, error) {
	event := new(RootChainManagerContractTokenMapped)
	if err := _RootChainManagerContract.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
