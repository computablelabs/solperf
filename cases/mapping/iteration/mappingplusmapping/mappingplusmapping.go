// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mappingplusmapping

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MappingPlusMappingABI is the input ABI used to generate the binding from.
const MappingPlusMappingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"MAX_KOO_KOO_ROOS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"isKooKooRooOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"kooKooRooOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"location\",\"type\":\"string\"}],\"name\":\"becomeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"kooKooRoos\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"listed\",\"type\":\"uint256\"},{\"name\":\"location\",\"type\":\"string\"},{\"name\":\"generalManager\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MappingPlusMappingBin is the compiled bytecode used for deploying new contracts.
const MappingPlusMappingBin = `0x6080604052601460005534801561001557600080fd5b50610696806100256000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663111bf23b811461007157806330b815571461009857806367cadd26146100da57806375490a8514610108578063ee989cd114610128575b600080fd5b34801561007d57600080fd5b50610086610258565b60408051918252519081900360200190f35b3480156100a457600080fd5b506100c673ffffffffffffffffffffffffffffffffffffffff6004351661025e565b604080519115158252519081900360200190f35b3480156100e657600080fd5b5061008673ffffffffffffffffffffffffffffffffffffffff60043516610287565b34801561011457600080fd5b506100c66004803560248101910135610299565b34801561013457600080fd5b506101406004356103dc565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156101b95781810151838201526020016101a1565b50505050905090810190601f1680156101e65780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610219578181015183820152602001610201565b50505050905090810190601f1680156102465780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b60005481565b73ffffffffffffffffffffffffffffffffffffffff166000908152600260205260408120541190565b60026020526000908152604090205481565b6000805433825260026020526040822054829182911061034057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4572726f723a204f6e6c79204469646479206d6179206f776e203231204b6f6f60448201527f204b6f6f20526f6f730000000000000000000000000000000000000000000000606482015290519081900360840190fd5b33600090815260026020908152604091829020805460010190558151601f870182900482028101820190925285825261038e919087908790819084018382808284375061052c945050505050565b6000818152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff19163317815590925090506103cc6002820186866105cf565b5042600191820155949350505050565b6001602081815260009283526040928390208054818401546002808401805488516101009882161598909802600019011691909104601f810186900486028701860190975286865273ffffffffffffffffffffffffffffffffffffffff90921695909492938301828280156104925780601f1061046757610100808354040283529160200191610492565b820191906000526020600020905b81548152906001019060200180831161047557829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529495949350908301828280156105225780601f106104f757610100808354040283529160200191610522565b820191906000526020600020905b81548152906001019060200180831161050557829003601f168201915b5050505050905084565b600060208251111515156105c757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4572726f723a204c6f636174696f6e206d75737420626520333220636861726160448201527f6374657273206f72206c65737300000000000000000000000000000000000000606482015290519081900360840190fd5b506020015190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106105782800160ff1982351617855561063d565b8280016001018555821561063d579182015b8281111561063d578235825591602001919060010190610622565b5061064992915061064d565b5090565b61066791905b808211156106495760008155600101610653565b905600a165627a7a72305820b9f5af4f87e88bf70bd967f41a567796f678462c64340545f9b2f4697a63de5e0029`

// DeployMappingPlusMapping deploys a new Ethereum contract, binding an instance of MappingPlusMapping to it.
func DeployMappingPlusMapping(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MappingPlusMapping, error) {
	parsed, err := abi.JSON(strings.NewReader(MappingPlusMappingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MappingPlusMappingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MappingPlusMapping{MappingPlusMappingCaller: MappingPlusMappingCaller{contract: contract}, MappingPlusMappingTransactor: MappingPlusMappingTransactor{contract: contract}, MappingPlusMappingFilterer: MappingPlusMappingFilterer{contract: contract}}, nil
}

// MappingPlusMapping is an auto generated Go binding around an Ethereum contract.
type MappingPlusMapping struct {
	MappingPlusMappingCaller     // Read-only binding to the contract
	MappingPlusMappingTransactor // Write-only binding to the contract
	MappingPlusMappingFilterer   // Log filterer for contract events
}

// MappingPlusMappingCaller is an auto generated read-only Go binding around an Ethereum contract.
type MappingPlusMappingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingPlusMappingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MappingPlusMappingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingPlusMappingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MappingPlusMappingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingPlusMappingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MappingPlusMappingSession struct {
	Contract     *MappingPlusMapping // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MappingPlusMappingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MappingPlusMappingCallerSession struct {
	Contract *MappingPlusMappingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MappingPlusMappingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MappingPlusMappingTransactorSession struct {
	Contract     *MappingPlusMappingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MappingPlusMappingRaw is an auto generated low-level Go binding around an Ethereum contract.
type MappingPlusMappingRaw struct {
	Contract *MappingPlusMapping // Generic contract binding to access the raw methods on
}

// MappingPlusMappingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MappingPlusMappingCallerRaw struct {
	Contract *MappingPlusMappingCaller // Generic read-only contract binding to access the raw methods on
}

// MappingPlusMappingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MappingPlusMappingTransactorRaw struct {
	Contract *MappingPlusMappingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMappingPlusMapping creates a new instance of MappingPlusMapping, bound to a specific deployed contract.
func NewMappingPlusMapping(address common.Address, backend bind.ContractBackend) (*MappingPlusMapping, error) {
	contract, err := bindMappingPlusMapping(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MappingPlusMapping{MappingPlusMappingCaller: MappingPlusMappingCaller{contract: contract}, MappingPlusMappingTransactor: MappingPlusMappingTransactor{contract: contract}, MappingPlusMappingFilterer: MappingPlusMappingFilterer{contract: contract}}, nil
}

// NewMappingPlusMappingCaller creates a new read-only instance of MappingPlusMapping, bound to a specific deployed contract.
func NewMappingPlusMappingCaller(address common.Address, caller bind.ContractCaller) (*MappingPlusMappingCaller, error) {
	contract, err := bindMappingPlusMapping(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MappingPlusMappingCaller{contract: contract}, nil
}

// NewMappingPlusMappingTransactor creates a new write-only instance of MappingPlusMapping, bound to a specific deployed contract.
func NewMappingPlusMappingTransactor(address common.Address, transactor bind.ContractTransactor) (*MappingPlusMappingTransactor, error) {
	contract, err := bindMappingPlusMapping(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MappingPlusMappingTransactor{contract: contract}, nil
}

// NewMappingPlusMappingFilterer creates a new log filterer instance of MappingPlusMapping, bound to a specific deployed contract.
func NewMappingPlusMappingFilterer(address common.Address, filterer bind.ContractFilterer) (*MappingPlusMappingFilterer, error) {
	contract, err := bindMappingPlusMapping(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MappingPlusMappingFilterer{contract: contract}, nil
}

// bindMappingPlusMapping binds a generic wrapper to an already deployed contract.
func bindMappingPlusMapping(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MappingPlusMappingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MappingPlusMapping *MappingPlusMappingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MappingPlusMapping.Contract.MappingPlusMappingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MappingPlusMapping *MappingPlusMappingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.MappingPlusMappingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MappingPlusMapping *MappingPlusMappingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.MappingPlusMappingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MappingPlusMapping *MappingPlusMappingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MappingPlusMapping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MappingPlusMapping *MappingPlusMappingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MappingPlusMapping *MappingPlusMappingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.contract.Transact(opts, method, params...)
}

// MAXKOOKOOROOS is a free data retrieval call binding the contract method 0x111bf23b.
//
// Solidity: function MAX_KOO_KOO_ROOS() constant returns(uint256)
func (_MappingPlusMapping *MappingPlusMappingCaller) MAXKOOKOOROOS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MappingPlusMapping.contract.Call(opts, out, "MAX_KOO_KOO_ROOS")
	return *ret0, err
}

// MAXKOOKOOROOS is a free data retrieval call binding the contract method 0x111bf23b.
//
// Solidity: function MAX_KOO_KOO_ROOS() constant returns(uint256)
func (_MappingPlusMapping *MappingPlusMappingSession) MAXKOOKOOROOS() (*big.Int, error) {
	return _MappingPlusMapping.Contract.MAXKOOKOOROOS(&_MappingPlusMapping.CallOpts)
}

// MAXKOOKOOROOS is a free data retrieval call binding the contract method 0x111bf23b.
//
// Solidity: function MAX_KOO_KOO_ROOS() constant returns(uint256)
func (_MappingPlusMapping *MappingPlusMappingCallerSession) MAXKOOKOOROOS() (*big.Int, error) {
	return _MappingPlusMapping.Contract.MAXKOOKOOROOS(&_MappingPlusMapping.CallOpts)
}

// KooKooRooOwners is a free data retrieval call binding the contract method 0x67cadd26.
//
// Solidity: function kooKooRooOwners( address) constant returns(uint256)
func (_MappingPlusMapping *MappingPlusMappingCaller) KooKooRooOwners(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MappingPlusMapping.contract.Call(opts, out, "kooKooRooOwners", arg0)
	return *ret0, err
}

// KooKooRooOwners is a free data retrieval call binding the contract method 0x67cadd26.
//
// Solidity: function kooKooRooOwners( address) constant returns(uint256)
func (_MappingPlusMapping *MappingPlusMappingSession) KooKooRooOwners(arg0 common.Address) (*big.Int, error) {
	return _MappingPlusMapping.Contract.KooKooRooOwners(&_MappingPlusMapping.CallOpts, arg0)
}

// KooKooRooOwners is a free data retrieval call binding the contract method 0x67cadd26.
//
// Solidity: function kooKooRooOwners( address) constant returns(uint256)
func (_MappingPlusMapping *MappingPlusMappingCallerSession) KooKooRooOwners(arg0 common.Address) (*big.Int, error) {
	return _MappingPlusMapping.Contract.KooKooRooOwners(&_MappingPlusMapping.CallOpts, arg0)
}

// KooKooRoos is a free data retrieval call binding the contract method 0xee989cd1.
//
// Solidity: function kooKooRoos( bytes32) constant returns(owner address, listed uint256, location string, generalManager string)
func (_MappingPlusMapping *MappingPlusMappingCaller) KooKooRoos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner          common.Address
	Listed         *big.Int
	Location       string
	GeneralManager string
}, error) {
	ret := new(struct {
		Owner          common.Address
		Listed         *big.Int
		Location       string
		GeneralManager string
	})
	out := ret
	err := _MappingPlusMapping.contract.Call(opts, out, "kooKooRoos", arg0)
	return *ret, err
}

// KooKooRoos is a free data retrieval call binding the contract method 0xee989cd1.
//
// Solidity: function kooKooRoos( bytes32) constant returns(owner address, listed uint256, location string, generalManager string)
func (_MappingPlusMapping *MappingPlusMappingSession) KooKooRoos(arg0 [32]byte) (struct {
	Owner          common.Address
	Listed         *big.Int
	Location       string
	GeneralManager string
}, error) {
	return _MappingPlusMapping.Contract.KooKooRoos(&_MappingPlusMapping.CallOpts, arg0)
}

// KooKooRoos is a free data retrieval call binding the contract method 0xee989cd1.
//
// Solidity: function kooKooRoos( bytes32) constant returns(owner address, listed uint256, location string, generalManager string)
func (_MappingPlusMapping *MappingPlusMappingCallerSession) KooKooRoos(arg0 [32]byte) (struct {
	Owner          common.Address
	Listed         *big.Int
	Location       string
	GeneralManager string
}, error) {
	return _MappingPlusMapping.Contract.KooKooRoos(&_MappingPlusMapping.CallOpts, arg0)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0x75490a85.
//
// Solidity: function becomeOwner(location string) returns(bool)
func (_MappingPlusMapping *MappingPlusMappingTransactor) BecomeOwner(opts *bind.TransactOpts, location string) (*types.Transaction, error) {
	return _MappingPlusMapping.contract.Transact(opts, "becomeOwner", location)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0x75490a85.
//
// Solidity: function becomeOwner(location string) returns(bool)
func (_MappingPlusMapping *MappingPlusMappingSession) BecomeOwner(location string) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.BecomeOwner(&_MappingPlusMapping.TransactOpts, location)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0x75490a85.
//
// Solidity: function becomeOwner(location string) returns(bool)
func (_MappingPlusMapping *MappingPlusMappingTransactorSession) BecomeOwner(location string) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.BecomeOwner(&_MappingPlusMapping.TransactOpts, location)
}

// IsKooKooRooOwner is a paid mutator transaction binding the contract method 0x30b81557.
//
// Solidity: function isKooKooRooOwner(candidate address) returns(bool)
func (_MappingPlusMapping *MappingPlusMappingTransactor) IsKooKooRooOwner(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _MappingPlusMapping.contract.Transact(opts, "isKooKooRooOwner", candidate)
}

// IsKooKooRooOwner is a paid mutator transaction binding the contract method 0x30b81557.
//
// Solidity: function isKooKooRooOwner(candidate address) returns(bool)
func (_MappingPlusMapping *MappingPlusMappingSession) IsKooKooRooOwner(candidate common.Address) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.IsKooKooRooOwner(&_MappingPlusMapping.TransactOpts, candidate)
}

// IsKooKooRooOwner is a paid mutator transaction binding the contract method 0x30b81557.
//
// Solidity: function isKooKooRooOwner(candidate address) returns(bool)
func (_MappingPlusMapping *MappingPlusMappingTransactorSession) IsKooKooRooOwner(candidate common.Address) (*types.Transaction, error) {
	return _MappingPlusMapping.Contract.IsKooKooRooOwner(&_MappingPlusMapping.TransactOpts, candidate)
}
