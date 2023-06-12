// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wager

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

// WagerMetaData contains all meta data concerning the Wager contract.
var WagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressDidntWin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressIsNotBetCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetAlreadyTaken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetDoesntExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetHasNotYetExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetIsActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetIsExpiredAlready\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotYourBet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"betId\",\"type\":\"uint256\"}],\"name\":\"BetCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"betId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"expiration\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"openingPrice\",\"type\":\"uint128\"}],\"name\":\"BetMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"joiner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"betId\",\"type\":\"uint256\"}],\"name\":\"JoinBet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"betId\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"long\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"short\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"expiration\",\"type\":\"uint96\"},{\"internalType\":\"uint120\",\"name\":\"createdAt\",\"type\":\"uint120\"},{\"internalType\":\"uint128\",\"name\":\"openingPrice\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betId\",\"type\":\"uint256\"}],\"name\":\"cancelBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betId\",\"type\":\"uint256\"}],\"name\":\"joinBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_amount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_expiration\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"_long\",\"type\":\"bool\"}],\"name\":\"openBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betId\",\"type\":\"uint256\"}],\"name\":\"resolveAndWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WagerMetaData.ABI instead.
var WagerABI = WagerMetaData.ABI

// Wager is an auto generated Go binding around an Ethereum contract.
type Wager struct {
	WagerCaller     // Read-only binding to the contract
	WagerTransactor // Write-only binding to the contract
	WagerFilterer   // Log filterer for contract events
}

// WagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WagerSession struct {
	Contract     *Wager            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WagerCallerSession struct {
	Contract *WagerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WagerTransactorSession struct {
	Contract     *WagerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WagerRaw struct {
	Contract *Wager // Generic contract binding to access the raw methods on
}

// WagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WagerCallerRaw struct {
	Contract *WagerCaller // Generic read-only contract binding to access the raw methods on
}

// WagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WagerTransactorRaw struct {
	Contract *WagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWager creates a new instance of Wager, bound to a specific deployed contract.
func NewWager(address common.Address, backend bind.ContractBackend) (*Wager, error) {
	contract, err := bindWager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wager{WagerCaller: WagerCaller{contract: contract}, WagerTransactor: WagerTransactor{contract: contract}, WagerFilterer: WagerFilterer{contract: contract}}, nil
}

// NewWagerCaller creates a new read-only instance of Wager, bound to a specific deployed contract.
func NewWagerCaller(address common.Address, caller bind.ContractCaller) (*WagerCaller, error) {
	contract, err := bindWager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WagerCaller{contract: contract}, nil
}

// NewWagerTransactor creates a new write-only instance of Wager, bound to a specific deployed contract.
func NewWagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WagerTransactor, error) {
	contract, err := bindWager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WagerTransactor{contract: contract}, nil
}

// NewWagerFilterer creates a new log filterer instance of Wager, bound to a specific deployed contract.
func NewWagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WagerFilterer, error) {
	contract, err := bindWager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WagerFilterer{contract: contract}, nil
}

// bindWager binds a generic wrapper to an already deployed contract.
func bindWager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wager *WagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wager.Contract.WagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wager *WagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wager.Contract.WagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wager *WagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wager.Contract.WagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wager *WagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wager *WagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wager *WagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wager.Contract.contract.Transact(opts, method, params...)
}

// Bets is a free data retrieval call binding the contract method 0x22af00fa.
//
// Solidity: function bets(uint256 ) view returns(address long, address short, uint96 amount, uint96 expiration, uint120 createdAt, uint128 openingPrice, bool isActive)
func (_Wager *WagerCaller) Bets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Long         common.Address
	Short        common.Address
	Amount       *big.Int
	Expiration   *big.Int
	CreatedAt    *big.Int
	OpeningPrice *big.Int
	IsActive     bool
}, error) {
	var out []interface{}
	err := _Wager.contract.Call(opts, &out, "bets", arg0)

	outstruct := new(struct {
		Long         common.Address
		Short        common.Address
		Amount       *big.Int
		Expiration   *big.Int
		CreatedAt    *big.Int
		OpeningPrice *big.Int
		IsActive     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Long = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Short = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OpeningPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Bets is a free data retrieval call binding the contract method 0x22af00fa.
//
// Solidity: function bets(uint256 ) view returns(address long, address short, uint96 amount, uint96 expiration, uint120 createdAt, uint128 openingPrice, bool isActive)
func (_Wager *WagerSession) Bets(arg0 *big.Int) (struct {
	Long         common.Address
	Short        common.Address
	Amount       *big.Int
	Expiration   *big.Int
	CreatedAt    *big.Int
	OpeningPrice *big.Int
	IsActive     bool
}, error) {
	return _Wager.Contract.Bets(&_Wager.CallOpts, arg0)
}

// Bets is a free data retrieval call binding the contract method 0x22af00fa.
//
// Solidity: function bets(uint256 ) view returns(address long, address short, uint96 amount, uint96 expiration, uint120 createdAt, uint128 openingPrice, bool isActive)
func (_Wager *WagerCallerSession) Bets(arg0 *big.Int) (struct {
	Long         common.Address
	Short        common.Address
	Amount       *big.Int
	Expiration   *big.Int
	CreatedAt    *big.Int
	OpeningPrice *big.Int
	IsActive     bool
}, error) {
	return _Wager.Contract.Bets(&_Wager.CallOpts, arg0)
}

// CancelBet is a paid mutator transaction binding the contract method 0x357401f5.
//
// Solidity: function cancelBet(uint256 _betId) returns()
func (_Wager *WagerTransactor) CancelBet(opts *bind.TransactOpts, _betId *big.Int) (*types.Transaction, error) {
	return _Wager.contract.Transact(opts, "cancelBet", _betId)
}

// CancelBet is a paid mutator transaction binding the contract method 0x357401f5.
//
// Solidity: function cancelBet(uint256 _betId) returns()
func (_Wager *WagerSession) CancelBet(_betId *big.Int) (*types.Transaction, error) {
	return _Wager.Contract.CancelBet(&_Wager.TransactOpts, _betId)
}

// CancelBet is a paid mutator transaction binding the contract method 0x357401f5.
//
// Solidity: function cancelBet(uint256 _betId) returns()
func (_Wager *WagerTransactorSession) CancelBet(_betId *big.Int) (*types.Transaction, error) {
	return _Wager.Contract.CancelBet(&_Wager.TransactOpts, _betId)
}

// JoinBet is a paid mutator transaction binding the contract method 0x39fb3f45.
//
// Solidity: function joinBet(uint256 _betId) returns()
func (_Wager *WagerTransactor) JoinBet(opts *bind.TransactOpts, _betId *big.Int) (*types.Transaction, error) {
	return _Wager.contract.Transact(opts, "joinBet", _betId)
}

// JoinBet is a paid mutator transaction binding the contract method 0x39fb3f45.
//
// Solidity: function joinBet(uint256 _betId) returns()
func (_Wager *WagerSession) JoinBet(_betId *big.Int) (*types.Transaction, error) {
	return _Wager.Contract.JoinBet(&_Wager.TransactOpts, _betId)
}

// JoinBet is a paid mutator transaction binding the contract method 0x39fb3f45.
//
// Solidity: function joinBet(uint256 _betId) returns()
func (_Wager *WagerTransactorSession) JoinBet(_betId *big.Int) (*types.Transaction, error) {
	return _Wager.Contract.JoinBet(&_Wager.TransactOpts, _betId)
}

// OpenBet is a paid mutator transaction binding the contract method 0x30fb61ae.
//
// Solidity: function openBet(uint96 _amount, uint96 _expiration, bool _long) returns()
func (_Wager *WagerTransactor) OpenBet(opts *bind.TransactOpts, _amount *big.Int, _expiration *big.Int, _long bool) (*types.Transaction, error) {
	return _Wager.contract.Transact(opts, "openBet", _amount, _expiration, _long)
}

// OpenBet is a paid mutator transaction binding the contract method 0x30fb61ae.
//
// Solidity: function openBet(uint96 _amount, uint96 _expiration, bool _long) returns()
func (_Wager *WagerSession) OpenBet(_amount *big.Int, _expiration *big.Int, _long bool) (*types.Transaction, error) {
	return _Wager.Contract.OpenBet(&_Wager.TransactOpts, _amount, _expiration, _long)
}

// OpenBet is a paid mutator transaction binding the contract method 0x30fb61ae.
//
// Solidity: function openBet(uint96 _amount, uint96 _expiration, bool _long) returns()
func (_Wager *WagerTransactorSession) OpenBet(_amount *big.Int, _expiration *big.Int, _long bool) (*types.Transaction, error) {
	return _Wager.Contract.OpenBet(&_Wager.TransactOpts, _amount, _expiration, _long)
}

// ResolveAndWithdraw is a paid mutator transaction binding the contract method 0xefa1a87b.
//
// Solidity: function resolveAndWithdraw(uint256 _betId) returns()
func (_Wager *WagerTransactor) ResolveAndWithdraw(opts *bind.TransactOpts, _betId *big.Int) (*types.Transaction, error) {
	return _Wager.contract.Transact(opts, "resolveAndWithdraw", _betId)
}

// ResolveAndWithdraw is a paid mutator transaction binding the contract method 0xefa1a87b.
//
// Solidity: function resolveAndWithdraw(uint256 _betId) returns()
func (_Wager *WagerSession) ResolveAndWithdraw(_betId *big.Int) (*types.Transaction, error) {
	return _Wager.Contract.ResolveAndWithdraw(&_Wager.TransactOpts, _betId)
}

// ResolveAndWithdraw is a paid mutator transaction binding the contract method 0xefa1a87b.
//
// Solidity: function resolveAndWithdraw(uint256 _betId) returns()
func (_Wager *WagerTransactorSession) ResolveAndWithdraw(_betId *big.Int) (*types.Transaction, error) {
	return _Wager.Contract.ResolveAndWithdraw(&_Wager.TransactOpts, _betId)
}

// WagerBetCanceledIterator is returned from FilterBetCanceled and is used to iterate over the raw logs and unpacked data for BetCanceled events raised by the Wager contract.
type WagerBetCanceledIterator struct {
	Event *WagerBetCanceled // Event containing the contract specifics and raw log

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
func (it *WagerBetCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerBetCanceled)
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
		it.Event = new(WagerBetCanceled)
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
func (it *WagerBetCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerBetCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerBetCanceled represents a BetCanceled event raised by the Wager contract.
type WagerBetCanceled struct {
	Creator common.Address
	BetId   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBetCanceled is a free log retrieval operation binding the contract event 0xa93e870117c34319bff8abace21e29dab4a1d5d0ffc0c00ab598fcb0f185e657.
//
// Solidity: event BetCanceled(address creator, uint256 indexed betId)
func (_Wager *WagerFilterer) FilterBetCanceled(opts *bind.FilterOpts, betId []*big.Int) (*WagerBetCanceledIterator, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.FilterLogs(opts, "BetCanceled", betIdRule)
	if err != nil {
		return nil, err
	}
	return &WagerBetCanceledIterator{contract: _Wager.contract, event: "BetCanceled", logs: logs, sub: sub}, nil
}

// WatchBetCanceled is a free log subscription operation binding the contract event 0xa93e870117c34319bff8abace21e29dab4a1d5d0ffc0c00ab598fcb0f185e657.
//
// Solidity: event BetCanceled(address creator, uint256 indexed betId)
func (_Wager *WagerFilterer) WatchBetCanceled(opts *bind.WatchOpts, sink chan<- *WagerBetCanceled, betId []*big.Int) (event.Subscription, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.WatchLogs(opts, "BetCanceled", betIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerBetCanceled)
				if err := _Wager.contract.UnpackLog(event, "BetCanceled", log); err != nil {
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

// ParseBetCanceled is a log parse operation binding the contract event 0xa93e870117c34319bff8abace21e29dab4a1d5d0ffc0c00ab598fcb0f185e657.
//
// Solidity: event BetCanceled(address creator, uint256 indexed betId)
func (_Wager *WagerFilterer) ParseBetCanceled(log types.Log) (*WagerBetCanceled, error) {
	event := new(WagerBetCanceled)
	if err := _Wager.contract.UnpackLog(event, "BetCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerBetMadeIterator is returned from FilterBetMade and is used to iterate over the raw logs and unpacked data for BetMade events raised by the Wager contract.
type WagerBetMadeIterator struct {
	Event *WagerBetMade // Event containing the contract specifics and raw log

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
func (it *WagerBetMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerBetMade)
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
		it.Event = new(WagerBetMade)
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
func (it *WagerBetMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerBetMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerBetMade represents a BetMade event raised by the Wager contract.
type WagerBetMade struct {
	Initiator    common.Address
	Long         bool
	BetId        *big.Int
	Amount       *big.Int
	Expiration   *big.Int
	OpeningPrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBetMade is a free log retrieval operation binding the contract event 0xf60b3be4dffb12fa8420965eab772374f918f64962695601f560605ed203ad15.
//
// Solidity: event BetMade(address initiator, bool long, uint256 indexed betId, uint96 amount, uint96 expiration, uint128 openingPrice)
func (_Wager *WagerFilterer) FilterBetMade(opts *bind.FilterOpts, betId []*big.Int) (*WagerBetMadeIterator, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.FilterLogs(opts, "BetMade", betIdRule)
	if err != nil {
		return nil, err
	}
	return &WagerBetMadeIterator{contract: _Wager.contract, event: "BetMade", logs: logs, sub: sub}, nil
}

// WatchBetMade is a free log subscription operation binding the contract event 0xf60b3be4dffb12fa8420965eab772374f918f64962695601f560605ed203ad15.
//
// Solidity: event BetMade(address initiator, bool long, uint256 indexed betId, uint96 amount, uint96 expiration, uint128 openingPrice)
func (_Wager *WagerFilterer) WatchBetMade(opts *bind.WatchOpts, sink chan<- *WagerBetMade, betId []*big.Int) (event.Subscription, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.WatchLogs(opts, "BetMade", betIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerBetMade)
				if err := _Wager.contract.UnpackLog(event, "BetMade", log); err != nil {
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

// ParseBetMade is a log parse operation binding the contract event 0xf60b3be4dffb12fa8420965eab772374f918f64962695601f560605ed203ad15.
//
// Solidity: event BetMade(address initiator, bool long, uint256 indexed betId, uint96 amount, uint96 expiration, uint128 openingPrice)
func (_Wager *WagerFilterer) ParseBetMade(log types.Log) (*WagerBetMade, error) {
	event := new(WagerBetMade)
	if err := _Wager.contract.UnpackLog(event, "BetMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerJoinBetIterator is returned from FilterJoinBet and is used to iterate over the raw logs and unpacked data for JoinBet events raised by the Wager contract.
type WagerJoinBetIterator struct {
	Event *WagerJoinBet // Event containing the contract specifics and raw log

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
func (it *WagerJoinBetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerJoinBet)
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
		it.Event = new(WagerJoinBet)
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
func (it *WagerJoinBetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerJoinBetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerJoinBet represents a JoinBet event raised by the Wager contract.
type WagerJoinBet struct {
	Joiner common.Address
	BetId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJoinBet is a free log retrieval operation binding the contract event 0xf2509873e9d8d5d4131d23265f901646624b38685111497537a0d27fe4a0e56d.
//
// Solidity: event JoinBet(address indexed joiner, uint256 indexed betId)
func (_Wager *WagerFilterer) FilterJoinBet(opts *bind.FilterOpts, joiner []common.Address, betId []*big.Int) (*WagerJoinBetIterator, error) {

	var joinerRule []interface{}
	for _, joinerItem := range joiner {
		joinerRule = append(joinerRule, joinerItem)
	}
	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.FilterLogs(opts, "JoinBet", joinerRule, betIdRule)
	if err != nil {
		return nil, err
	}
	return &WagerJoinBetIterator{contract: _Wager.contract, event: "JoinBet", logs: logs, sub: sub}, nil
}

// WatchJoinBet is a free log subscription operation binding the contract event 0xf2509873e9d8d5d4131d23265f901646624b38685111497537a0d27fe4a0e56d.
//
// Solidity: event JoinBet(address indexed joiner, uint256 indexed betId)
func (_Wager *WagerFilterer) WatchJoinBet(opts *bind.WatchOpts, sink chan<- *WagerJoinBet, joiner []common.Address, betId []*big.Int) (event.Subscription, error) {

	var joinerRule []interface{}
	for _, joinerItem := range joiner {
		joinerRule = append(joinerRule, joinerItem)
	}
	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.WatchLogs(opts, "JoinBet", joinerRule, betIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerJoinBet)
				if err := _Wager.contract.UnpackLog(event, "JoinBet", log); err != nil {
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

// ParseJoinBet is a log parse operation binding the contract event 0xf2509873e9d8d5d4131d23265f901646624b38685111497537a0d27fe4a0e56d.
//
// Solidity: event JoinBet(address indexed joiner, uint256 indexed betId)
func (_Wager *WagerFilterer) ParseJoinBet(log types.Log) (*WagerJoinBet, error) {
	event := new(WagerJoinBet)
	if err := _Wager.contract.UnpackLog(event, "JoinBet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Wager contract.
type WagerWithdrawnIterator struct {
	Event *WagerWithdrawn // Event containing the contract specifics and raw log

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
func (it *WagerWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerWithdrawn)
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
		it.Event = new(WagerWithdrawn)
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
func (it *WagerWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerWithdrawn represents a Withdrawn event raised by the Wager contract.
type WagerWithdrawn struct {
	Winner common.Address
	BetId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address winner, uint256 indexed betId)
func (_Wager *WagerFilterer) FilterWithdrawn(opts *bind.FilterOpts, betId []*big.Int) (*WagerWithdrawnIterator, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.FilterLogs(opts, "Withdrawn", betIdRule)
	if err != nil {
		return nil, err
	}
	return &WagerWithdrawnIterator{contract: _Wager.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address winner, uint256 indexed betId)
func (_Wager *WagerFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *WagerWithdrawn, betId []*big.Int) (event.Subscription, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Wager.contract.WatchLogs(opts, "Withdrawn", betIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerWithdrawn)
				if err := _Wager.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address winner, uint256 indexed betId)
func (_Wager *WagerFilterer) ParseWithdrawn(log types.Log) (*WagerWithdrawn, error) {
	event := new(WagerWithdrawn)
	if err := _Wager.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
