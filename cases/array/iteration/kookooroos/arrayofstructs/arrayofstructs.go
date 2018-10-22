// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arrayofstructs

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ArrayOfStructsABI is the input ABI used to generate the binding from.
const ArrayOfStructsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"MAX_KOO_KOO_ROOS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"isKooKooRooOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kooKooRoos\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"listed\",\"type\":\"uint256\"},{\"name\":\"location\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"location\",\"type\":\"string\"}],\"name\":\"becomeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"string\"}],\"name\":\"stringToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArrayOfStructsBin is the compiled bytecode used for deploying new contracts.
const ArrayOfStructsBin = `0x6080604052601460005534801561001557600080fd5b506103e7806100256000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663111bf23b811461007157806330b8155714610098578063550b9f5f146100da57806375490a8514610127578063cfb5192814610147575b600080fd5b34801561007d57600080fd5b506100866101a0565b60408051918252519081900360200190f35b3480156100a457600080fd5b506100c673ffffffffffffffffffffffffffffffffffffffff600435166101a6565b604080519115158252519081900360200190f35b3480156100e657600080fd5b506100f26004356101ac565b6040805173ffffffffffffffffffffffffffffffffffffffff9094168452602084019290925282820152519081900360600190f35b34801561013357600080fd5b506100c660048035602481019101356101f4565b34801561015357600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100869436949293602493928401919081908401838280828437509497506103189650505050505050565b60005481565b50600090565b60018054829081106101ba57fe5b600091825260209091206003909102018054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff909116925083565b600080600061023285858080601f01602080910402602001604051908101604052809392919081815260200183838082843750610318945050505050565b60408051606081018252338152426020820190815291810192835260018054808201808355600092835292517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf66003909202918201805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff90921691909117905592517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf784015592517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf8909201919091551195945050505050565b600060208251111515156103b357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4572726f723a204c6f636174696f6e206d75737420626520333220636861726160448201527f6374657273206f72206c65737300000000000000000000000000000000000000606482015290519081900360840190fd5b5060200151905600a165627a7a7230582068f69840c553861402a6a8758dac7b659457f20332f5cd3e0df5bfa2326feda40029`

// DeployArrayOfStructs deploys a new Ethereum contract, binding an instance of ArrayOfStructs to it.
func DeployArrayOfStructs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArrayOfStructs, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayOfStructsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArrayOfStructsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArrayOfStructs{ArrayOfStructsCaller: ArrayOfStructsCaller{contract: contract}, ArrayOfStructsTransactor: ArrayOfStructsTransactor{contract: contract}, ArrayOfStructsFilterer: ArrayOfStructsFilterer{contract: contract}}, nil
}

// ArrayOfStructs is an auto generated Go binding around an Ethereum contract.
type ArrayOfStructs struct {
	ArrayOfStructsCaller     // Read-only binding to the contract
	ArrayOfStructsTransactor // Write-only binding to the contract
	ArrayOfStructsFilterer   // Log filterer for contract events
}

// ArrayOfStructsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArrayOfStructsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayOfStructsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArrayOfStructsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayOfStructsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArrayOfStructsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayOfStructsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArrayOfStructsSession struct {
	Contract     *ArrayOfStructs   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayOfStructsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArrayOfStructsCallerSession struct {
	Contract *ArrayOfStructsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ArrayOfStructsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArrayOfStructsTransactorSession struct {
	Contract     *ArrayOfStructsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ArrayOfStructsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArrayOfStructsRaw struct {
	Contract *ArrayOfStructs // Generic contract binding to access the raw methods on
}

// ArrayOfStructsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArrayOfStructsCallerRaw struct {
	Contract *ArrayOfStructsCaller // Generic read-only contract binding to access the raw methods on
}

// ArrayOfStructsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArrayOfStructsTransactorRaw struct {
	Contract *ArrayOfStructsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrayOfStructs creates a new instance of ArrayOfStructs, bound to a specific deployed contract.
func NewArrayOfStructs(address common.Address, backend bind.ContractBackend) (*ArrayOfStructs, error) {
	contract, err := bindArrayOfStructs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArrayOfStructs{ArrayOfStructsCaller: ArrayOfStructsCaller{contract: contract}, ArrayOfStructsTransactor: ArrayOfStructsTransactor{contract: contract}, ArrayOfStructsFilterer: ArrayOfStructsFilterer{contract: contract}}, nil
}

// NewArrayOfStructsCaller creates a new read-only instance of ArrayOfStructs, bound to a specific deployed contract.
func NewArrayOfStructsCaller(address common.Address, caller bind.ContractCaller) (*ArrayOfStructsCaller, error) {
	contract, err := bindArrayOfStructs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayOfStructsCaller{contract: contract}, nil
}

// NewArrayOfStructsTransactor creates a new write-only instance of ArrayOfStructs, bound to a specific deployed contract.
func NewArrayOfStructsTransactor(address common.Address, transactor bind.ContractTransactor) (*ArrayOfStructsTransactor, error) {
	contract, err := bindArrayOfStructs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayOfStructsTransactor{contract: contract}, nil
}

// NewArrayOfStructsFilterer creates a new log filterer instance of ArrayOfStructs, bound to a specific deployed contract.
func NewArrayOfStructsFilterer(address common.Address, filterer bind.ContractFilterer) (*ArrayOfStructsFilterer, error) {
	contract, err := bindArrayOfStructs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArrayOfStructsFilterer{contract: contract}, nil
}

// bindArrayOfStructs binds a generic wrapper to an already deployed contract.
func bindArrayOfStructs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayOfStructsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayOfStructs *ArrayOfStructsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArrayOfStructs.Contract.ArrayOfStructsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayOfStructs *ArrayOfStructsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayOfStructs.Contract.ArrayOfStructsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayOfStructs *ArrayOfStructsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayOfStructs.Contract.ArrayOfStructsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayOfStructs *ArrayOfStructsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArrayOfStructs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayOfStructs *ArrayOfStructsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayOfStructs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayOfStructs *ArrayOfStructsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayOfStructs.Contract.contract.Transact(opts, method, params...)
}

// MAXKOOKOOROOS is a free data retrieval call binding the contract method 0x111bf23b.
//
// Solidity: function MAX_KOO_KOO_ROOS() constant returns(uint256)
func (_ArrayOfStructs *ArrayOfStructsCaller) MAXKOOKOOROOS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArrayOfStructs.contract.Call(opts, out, "MAX_KOO_KOO_ROOS")
	return *ret0, err
}

// MAXKOOKOOROOS is a free data retrieval call binding the contract method 0x111bf23b.
//
// Solidity: function MAX_KOO_KOO_ROOS() constant returns(uint256)
func (_ArrayOfStructs *ArrayOfStructsSession) MAXKOOKOOROOS() (*big.Int, error) {
	return _ArrayOfStructs.Contract.MAXKOOKOOROOS(&_ArrayOfStructs.CallOpts)
}

// MAXKOOKOOROOS is a free data retrieval call binding the contract method 0x111bf23b.
//
// Solidity: function MAX_KOO_KOO_ROOS() constant returns(uint256)
func (_ArrayOfStructs *ArrayOfStructsCallerSession) MAXKOOKOOROOS() (*big.Int, error) {
	return _ArrayOfStructs.Contract.MAXKOOKOOROOS(&_ArrayOfStructs.CallOpts)
}

// IsKooKooRooOwner is a free data retrieval call binding the contract method 0x30b81557.
//
// Solidity: function isKooKooRooOwner(candidate address) constant returns(bool)
func (_ArrayOfStructs *ArrayOfStructsCaller) IsKooKooRooOwner(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArrayOfStructs.contract.Call(opts, out, "isKooKooRooOwner", candidate)
	return *ret0, err
}

// IsKooKooRooOwner is a free data retrieval call binding the contract method 0x30b81557.
//
// Solidity: function isKooKooRooOwner(candidate address) constant returns(bool)
func (_ArrayOfStructs *ArrayOfStructsSession) IsKooKooRooOwner(candidate common.Address) (bool, error) {
	return _ArrayOfStructs.Contract.IsKooKooRooOwner(&_ArrayOfStructs.CallOpts, candidate)
}

// IsKooKooRooOwner is a free data retrieval call binding the contract method 0x30b81557.
//
// Solidity: function isKooKooRooOwner(candidate address) constant returns(bool)
func (_ArrayOfStructs *ArrayOfStructsCallerSession) IsKooKooRooOwner(candidate common.Address) (bool, error) {
	return _ArrayOfStructs.Contract.IsKooKooRooOwner(&_ArrayOfStructs.CallOpts, candidate)
}

// KooKooRoos is a free data retrieval call binding the contract method 0x550b9f5f.
//
// Solidity: function kooKooRoos( uint256) constant returns(owner address, listed uint256, location bytes32)
func (_ArrayOfStructs *ArrayOfStructsCaller) KooKooRoos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner    common.Address
	Listed   *big.Int
	Location [32]byte
}, error) {
	ret := new(struct {
		Owner    common.Address
		Listed   *big.Int
		Location [32]byte
	})
	out := ret
	err := _ArrayOfStructs.contract.Call(opts, out, "kooKooRoos", arg0)
	return *ret, err
}

// KooKooRoos is a free data retrieval call binding the contract method 0x550b9f5f.
//
// Solidity: function kooKooRoos( uint256) constant returns(owner address, listed uint256, location bytes32)
func (_ArrayOfStructs *ArrayOfStructsSession) KooKooRoos(arg0 *big.Int) (struct {
	Owner    common.Address
	Listed   *big.Int
	Location [32]byte
}, error) {
	return _ArrayOfStructs.Contract.KooKooRoos(&_ArrayOfStructs.CallOpts, arg0)
}

// KooKooRoos is a free data retrieval call binding the contract method 0x550b9f5f.
//
// Solidity: function kooKooRoos( uint256) constant returns(owner address, listed uint256, location bytes32)
func (_ArrayOfStructs *ArrayOfStructsCallerSession) KooKooRoos(arg0 *big.Int) (struct {
	Owner    common.Address
	Listed   *big.Int
	Location [32]byte
}, error) {
	return _ArrayOfStructs.Contract.KooKooRoos(&_ArrayOfStructs.CallOpts, arg0)
}

// StringToBytes32 is a free data retrieval call binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(src string) constant returns(result bytes32)
func (_ArrayOfStructs *ArrayOfStructsCaller) StringToBytes32(opts *bind.CallOpts, src string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArrayOfStructs.contract.Call(opts, out, "stringToBytes32", src)
	return *ret0, err
}

// StringToBytes32 is a free data retrieval call binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(src string) constant returns(result bytes32)
func (_ArrayOfStructs *ArrayOfStructsSession) StringToBytes32(src string) ([32]byte, error) {
	return _ArrayOfStructs.Contract.StringToBytes32(&_ArrayOfStructs.CallOpts, src)
}

// StringToBytes32 is a free data retrieval call binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(src string) constant returns(result bytes32)
func (_ArrayOfStructs *ArrayOfStructsCallerSession) StringToBytes32(src string) ([32]byte, error) {
	return _ArrayOfStructs.Contract.StringToBytes32(&_ArrayOfStructs.CallOpts, src)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0x75490a85.
//
// Solidity: function becomeOwner(location string) returns(bool)
func (_ArrayOfStructs *ArrayOfStructsTransactor) BecomeOwner(opts *bind.TransactOpts, location string) (*types.Transaction, error) {
	return _ArrayOfStructs.contract.Transact(opts, "becomeOwner", location)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0x75490a85.
//
// Solidity: function becomeOwner(location string) returns(bool)
func (_ArrayOfStructs *ArrayOfStructsSession) BecomeOwner(location string) (*types.Transaction, error) {
	return _ArrayOfStructs.Contract.BecomeOwner(&_ArrayOfStructs.TransactOpts, location)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0x75490a85.
//
// Solidity: function becomeOwner(location string) returns(bool)
func (_ArrayOfStructs *ArrayOfStructsTransactorSession) BecomeOwner(location string) (*types.Transaction, error) {
	return _ArrayOfStructs.Contract.BecomeOwner(&_ArrayOfStructs.TransactOpts, location)
}
