// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Oracle

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oracleClaimValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"name\":\"_message\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"newOracleClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"processBridgeProphecy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasMadeClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"checkBridgeProphecy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"evrnetBridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consensusThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_valset\",\"type\":\"address\"},{\"name\":\"_evrnetBridge\",\"type\":\"address\"},{\"name\":\"_consensusThreshold\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_message\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"LogNewOracleClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_prophecyPowerCurrent\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_prophecyPowerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"LogProphecyProcessed\",\"type\":\"event\"}]"

// OracleBin is the compiled bytecode used for deploying new contracts.
var OracleBin = "0x608060405234801561001057600080fd5b50604051608080611b538339810180604052608081101561003057600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291905050506000811115156100fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f436f6e73656e737573207468726573686f6c64206d75737420626520706f736981526020017f746976652e00000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806003819055505050505061197a806101d96000396000f3fe608060405260043610610093576000357c01000000000000000000000000000000000000000000000000000000009004806336e4134114610098578063568b3c4f1461011d578063570ca735146101f95780637f54af0c1461025057806389ed70b7146102a7578063a219763e146102e2578063e33a8b2a14610355578063e48e227d146103b6578063f9b0b5b91461040d575b600080fd5b3480156100a457600080fd5b506100db600480360360408110156100bb57600080fd5b810190808035906020019092919080359060200190929190505050610438565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561012957600080fd5b506101f76004803603606081101561014057600080fd5b8101908080359060200190929190803590602001909291908035906020019064010000000081111561017157600080fd5b82018360208201111561018357600080fd5b803590602001918460018302840111640100000000831117156101a557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610485565b005b34801561020557600080fd5b5061020e610c9a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561025c57600080fd5b50610265610cc0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102b357600080fd5b506102e0600480360360208110156102ca57600080fd5b8101908080359060200190929190505050610ce6565b005b3480156102ee57600080fd5b5061033b6004803603604081101561030557600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fb1565b604051808215151515815260200191505060405180910390f35b34801561036157600080fd5b5061038e6004803603602081101561037857600080fd5b8101908080359060200190929190505050610fe0565b6040518084151515158152602001838152602001828152602001935050505060405180910390f35b3480156103c257600080fd5b506103cb611370565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561041957600080fd5b50610422611395565b6040518082815260200191505060405180910390f35b60046020528160005260406000208181548110151561045357fe5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561054057600080fd5b505afa158015610554573d6000803e3d6000fd5b505050506040513d602081101561056a57600080fd5b810190808051906020019092919050505015156105ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b82600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d8da69ea836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b15801561068257600080fd5b505afa158015610696573d6000803e3d6000fd5b505050506040513d60208110156106ac57600080fd5b8101908080519060200190929190505050151514151561075a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f5468652070726f7068656379206d7573742062652070656e64696e6720666f7281526020017f2074686973206f7065726174696f6e000000000000000000000000000000000081525060400191505060405180910390fd5b6000339050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319045a2585856040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156108105780820151818401526020810190506107f5565b50505050905090810190601f16801561083d5780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b15801561085b57600080fd5b505afa15801561086f573d6000803e3d6000fd5b505050506040513d602081101561088557600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610938576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f496e76616c6964206d657373616765207369676e61747572652e00000000000081525060200191505060405180910390fd5b6005600086815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515610a31576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a8152602001807f43616e6e6f74206d616b65206475706c6963617465206f7261636c6520636c6181526020017f696d732066726f6d207468652073616d6520616464726573732e00000000000081525060400191505060405180910390fd5b60016005600087815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600460008681526020019081526020016000208190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f50e466de4726c2437aa7498d554322f5599f31f0f69f9ce036ad96db7759049185858386604051808581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610bb4578082015181840152602081019050610b99565b50505050905090810190601f168015610be15780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a16000806000610bff8861139b565b9250925092508215610c9057610c148861174c565b7f1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc88838333604051808581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390a15b5050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d8da69ea836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b158015610d7957600080fd5b505afa158015610d8d573d6000803e3d6000fd5b505050506040513d6020811015610da357600080fd5b81019080805190602001909291905050501515141515610e51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f5468652070726f7068656379206d7573742062652070656e64696e6720666f7281526020017f2074686973206f7065726174696f6e000000000000000000000000000000000081525060400191505060405180910390fd5b6000806000610e5f8561139b565b925092509250821515610f26576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260488152602001807f5468652063756d756c617469766520706f776572206f66207369676e61746f7281526020017f792076616c696461746f727320646f6573206e6f74206d65657420746865207481526020017f68726573686f6c6400000000000000000000000000000000000000000000000081525060600191505060405180910390fd5b610f2f8561174c565b7f1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc85838333604051808581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390a15050505050565b60056020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b6000806000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156110aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b83600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d8da69ea836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b15801561113d57600080fd5b505afa158015611151573d6000803e3d6000fd5b505050506040513d602081101561116757600080fd5b81019080805190602001909291905050501515141515611215576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f5468652070726f7068656379206d7573742062652070656e64696e6720666f7281526020017f2074686973206f7065726174696f6e000000000000000000000000000000000081525060400191505060405180910390fd5b600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d8da69ea876040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b1580156112a757600080fd5b505afa1580156112bb573d6000803e3d6000fd5b505050506040513d60208110156112d157600080fd5b81019080805190602001909291905050501515141515611359576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f43616e206f6e6c7920636865636b206163746976652070726f7068656369657381525060200191505060405180910390fd5b6113628561139b565b935093509350509193909250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600080600080600090506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663db3ad22c6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561142b57600080fd5b505afa15801561143f573d6000803e3d6000fd5b505050506040513d602081101561145557600080fd5b8101908080519060200190929190505050905060008090505b60046000888152602001908152602001600020805490508110156116fd57600060046000898152602001908152602001600020828154811015156114ae57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561159657600080fd5b505afa1580156115aa573d6000803e3d6000fd5b505050506040513d60208110156115c057600080fd5b8101908080519060200190929190505050156116e1576116de600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663473691a4836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561169457600080fd5b505afa1580156116a8573d6000803e3d6000fd5b505050506040513d60208110156116be57600080fd5b8101908080519060200190929190505050856117f790919063ffffffff16565b93505b506116f66001826117f790919063ffffffff16565b905061146e565b5060006117156003548361188190919063ffffffff16565b9050600061172d60648561188190919063ffffffff16565b9050600082821015905080828497509750975050505050509193909250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636b3ce98c826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b1580156117dc57600080fd5b505af11580156117f0573d6000803e3d6000fd5b5050505050565b6000808284019050838110151515611877576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000808314156118945760009050611948565b600082840290508284828115156118a757fe5b04141515611943576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f81526020017f770000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b809150505b9291505056fea165627a7a723058205a0de51fc997a98f66d8ac968b0aa55aa51d3b4afe840c4c7b4fb4c150b047060029"

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _valset common.Address, _evrnetBridge common.Address, _consensusThreshold *big.Int) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleBin), backend, _operator, _valset, _evrnetBridge, _consensusThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// CheckBridgeProphecy is a free data retrieval call binding the contract method 0xe33a8b2a.
//
// Solidity: function checkBridgeProphecy(uint256 _prophecyID) view returns(bool, uint256, uint256)
func (_Oracle *OracleCaller) CheckBridgeProphecy(opts *bind.CallOpts, _prophecyID *big.Int) (bool, *big.Int, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Oracle.contract.Call(opts, out, "checkBridgeProphecy", _prophecyID)
	return *ret0, *ret1, *ret2, err
}

// CheckBridgeProphecy is a free data retrieval call binding the contract method 0xe33a8b2a.
//
// Solidity: function checkBridgeProphecy(uint256 _prophecyID) view returns(bool, uint256, uint256)
func (_Oracle *OracleSession) CheckBridgeProphecy(_prophecyID *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.CheckBridgeProphecy(&_Oracle.CallOpts, _prophecyID)
}

// CheckBridgeProphecy is a free data retrieval call binding the contract method 0xe33a8b2a.
//
// Solidity: function checkBridgeProphecy(uint256 _prophecyID) view returns(bool, uint256, uint256)
func (_Oracle *OracleCallerSession) CheckBridgeProphecy(_prophecyID *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.CheckBridgeProphecy(&_Oracle.CallOpts, _prophecyID)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_Oracle *OracleCaller) ConsensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "consensusThreshold")
	return *ret0, err
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_Oracle *OracleSession) ConsensusThreshold() (*big.Int, error) {
	return _Oracle.Contract.ConsensusThreshold(&_Oracle.CallOpts)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_Oracle *OracleCallerSession) ConsensusThreshold() (*big.Int, error) {
	return _Oracle.Contract.ConsensusThreshold(&_Oracle.CallOpts)
}

// EvrnetBridge is a free data retrieval call binding the contract method 0xe48e227d.
//
// Solidity: function evrnetBridge() view returns(address)
func (_Oracle *OracleCaller) EvrnetBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "evrnetBridge")
	return *ret0, err
}

// EvrnetBridge is a free data retrieval call binding the contract method 0xe48e227d.
//
// Solidity: function evrnetBridge() view returns(address)
func (_Oracle *OracleSession) EvrnetBridge() (common.Address, error) {
	return _Oracle.Contract.EvrnetBridge(&_Oracle.CallOpts)
}

// EvrnetBridge is a free data retrieval call binding the contract method 0xe48e227d.
//
// Solidity: function evrnetBridge() view returns(address)
func (_Oracle *OracleCallerSession) EvrnetBridge() (common.Address, error) {
	return _Oracle.Contract.EvrnetBridge(&_Oracle.CallOpts)
}

// HasMadeClaim is a free data retrieval call binding the contract method 0xa219763e.
//
// Solidity: function hasMadeClaim(uint256 , address ) view returns(bool)
func (_Oracle *OracleCaller) HasMadeClaim(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "hasMadeClaim", arg0, arg1)
	return *ret0, err
}

// HasMadeClaim is a free data retrieval call binding the contract method 0xa219763e.
//
// Solidity: function hasMadeClaim(uint256 , address ) view returns(bool)
func (_Oracle *OracleSession) HasMadeClaim(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Oracle.Contract.HasMadeClaim(&_Oracle.CallOpts, arg0, arg1)
}

// HasMadeClaim is a free data retrieval call binding the contract method 0xa219763e.
//
// Solidity: function hasMadeClaim(uint256 , address ) view returns(bool)
func (_Oracle *OracleCallerSession) HasMadeClaim(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Oracle.Contract.HasMadeClaim(&_Oracle.CallOpts, arg0, arg1)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleSession) Operator() (common.Address, error) {
	return _Oracle.Contract.Operator(&_Oracle.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleCallerSession) Operator() (common.Address, error) {
	return _Oracle.Contract.Operator(&_Oracle.CallOpts)
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x36e41341.
//
// Solidity: function oracleClaimValidators(uint256 , uint256 ) view returns(address)
func (_Oracle *OracleCaller) OracleClaimValidators(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "oracleClaimValidators", arg0, arg1)
	return *ret0, err
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x36e41341.
//
// Solidity: function oracleClaimValidators(uint256 , uint256 ) view returns(address)
func (_Oracle *OracleSession) OracleClaimValidators(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Oracle.Contract.OracleClaimValidators(&_Oracle.CallOpts, arg0, arg1)
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x36e41341.
//
// Solidity: function oracleClaimValidators(uint256 , uint256 ) view returns(address)
func (_Oracle *OracleCallerSession) OracleClaimValidators(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Oracle.Contract.OracleClaimValidators(&_Oracle.CallOpts, arg0, arg1)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleSession) Valset() (common.Address, error) {
	return _Oracle.Contract.Valset(&_Oracle.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleCallerSession) Valset() (common.Address, error) {
	return _Oracle.Contract.Valset(&_Oracle.CallOpts)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x568b3c4f.
//
// Solidity: function newOracleClaim(uint256 _prophecyID, bytes32 _message, bytes _signature) returns()
func (_Oracle *OracleTransactor) NewOracleClaim(opts *bind.TransactOpts, _prophecyID *big.Int, _message [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "newOracleClaim", _prophecyID, _message, _signature)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x568b3c4f.
//
// Solidity: function newOracleClaim(uint256 _prophecyID, bytes32 _message, bytes _signature) returns()
func (_Oracle *OracleSession) NewOracleClaim(_prophecyID *big.Int, _message [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleClaim(&_Oracle.TransactOpts, _prophecyID, _message, _signature)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x568b3c4f.
//
// Solidity: function newOracleClaim(uint256 _prophecyID, bytes32 _message, bytes _signature) returns()
func (_Oracle *OracleTransactorSession) NewOracleClaim(_prophecyID *big.Int, _message [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleClaim(&_Oracle.TransactOpts, _prophecyID, _message, _signature)
}

// ProcessBridgeProphecy is a paid mutator transaction binding the contract method 0x89ed70b7.
//
// Solidity: function processBridgeProphecy(uint256 _prophecyID) returns()
func (_Oracle *OracleTransactor) ProcessBridgeProphecy(opts *bind.TransactOpts, _prophecyID *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "processBridgeProphecy", _prophecyID)
}

// ProcessBridgeProphecy is a paid mutator transaction binding the contract method 0x89ed70b7.
//
// Solidity: function processBridgeProphecy(uint256 _prophecyID) returns()
func (_Oracle *OracleSession) ProcessBridgeProphecy(_prophecyID *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ProcessBridgeProphecy(&_Oracle.TransactOpts, _prophecyID)
}

// ProcessBridgeProphecy is a paid mutator transaction binding the contract method 0x89ed70b7.
//
// Solidity: function processBridgeProphecy(uint256 _prophecyID) returns()
func (_Oracle *OracleTransactorSession) ProcessBridgeProphecy(_prophecyID *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ProcessBridgeProphecy(&_Oracle.TransactOpts, _prophecyID)
}

// OracleLogNewOracleClaimIterator is returned from FilterLogNewOracleClaim and is used to iterate over the raw logs and unpacked data for LogNewOracleClaim events raised by the Oracle contract.
type OracleLogNewOracleClaimIterator struct {
	Event *OracleLogNewOracleClaim // Event containing the contract specifics and raw log

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
func (it *OracleLogNewOracleClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleLogNewOracleClaim)
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
		it.Event = new(OracleLogNewOracleClaim)
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
func (it *OracleLogNewOracleClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleLogNewOracleClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleLogNewOracleClaim represents a LogNewOracleClaim event raised by the Oracle contract.
type OracleLogNewOracleClaim struct {
	ProphecyID       *big.Int
	Message          [32]byte
	ValidatorAddress common.Address
	Signature        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewOracleClaim is a free log retrieval operation binding the contract event 0x50e466de4726c2437aa7498d554322f5599f31f0f69f9ce036ad96db77590491.
//
// Solidity: event LogNewOracleClaim(uint256 _prophecyID, bytes32 _message, address _validatorAddress, bytes _signature)
func (_Oracle *OracleFilterer) FilterLogNewOracleClaim(opts *bind.FilterOpts) (*OracleLogNewOracleClaimIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "LogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return &OracleLogNewOracleClaimIterator{contract: _Oracle.contract, event: "LogNewOracleClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewOracleClaim is a free log subscription operation binding the contract event 0x50e466de4726c2437aa7498d554322f5599f31f0f69f9ce036ad96db77590491.
//
// Solidity: event LogNewOracleClaim(uint256 _prophecyID, bytes32 _message, address _validatorAddress, bytes _signature)
func (_Oracle *OracleFilterer) WatchLogNewOracleClaim(opts *bind.WatchOpts, sink chan<- *OracleLogNewOracleClaim) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "LogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleLogNewOracleClaim)
				if err := _Oracle.contract.UnpackLog(event, "LogNewOracleClaim", log); err != nil {
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

// ParseLogNewOracleClaim is a log parse operation binding the contract event 0x50e466de4726c2437aa7498d554322f5599f31f0f69f9ce036ad96db77590491.
//
// Solidity: event LogNewOracleClaim(uint256 _prophecyID, bytes32 _message, address _validatorAddress, bytes _signature)
func (_Oracle *OracleFilterer) ParseLogNewOracleClaim(log types.Log) (*OracleLogNewOracleClaim, error) {
	event := new(OracleLogNewOracleClaim)
	if err := _Oracle.contract.UnpackLog(event, "LogNewOracleClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OracleLogProphecyProcessedIterator is returned from FilterLogProphecyProcessed and is used to iterate over the raw logs and unpacked data for LogProphecyProcessed events raised by the Oracle contract.
type OracleLogProphecyProcessedIterator struct {
	Event *OracleLogProphecyProcessed // Event containing the contract specifics and raw log

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
func (it *OracleLogProphecyProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleLogProphecyProcessed)
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
		it.Event = new(OracleLogProphecyProcessed)
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
func (it *OracleLogProphecyProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleLogProphecyProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleLogProphecyProcessed represents a LogProphecyProcessed event raised by the Oracle contract.
type OracleLogProphecyProcessed struct {
	ProphecyID             *big.Int
	ProphecyPowerCurrent   *big.Int
	ProphecyPowerThreshold *big.Int
	Submitter              common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterLogProphecyProcessed is a free log retrieval operation binding the contract event 0x1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc.
//
// Solidity: event LogProphecyProcessed(uint256 _prophecyID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_Oracle *OracleFilterer) FilterLogProphecyProcessed(opts *bind.FilterOpts) (*OracleLogProphecyProcessedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "LogProphecyProcessed")
	if err != nil {
		return nil, err
	}
	return &OracleLogProphecyProcessedIterator{contract: _Oracle.contract, event: "LogProphecyProcessed", logs: logs, sub: sub}, nil
}

// WatchLogProphecyProcessed is a free log subscription operation binding the contract event 0x1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc.
//
// Solidity: event LogProphecyProcessed(uint256 _prophecyID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_Oracle *OracleFilterer) WatchLogProphecyProcessed(opts *bind.WatchOpts, sink chan<- *OracleLogProphecyProcessed) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "LogProphecyProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleLogProphecyProcessed)
				if err := _Oracle.contract.UnpackLog(event, "LogProphecyProcessed", log); err != nil {
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

// ParseLogProphecyProcessed is a log parse operation binding the contract event 0x1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc.
//
// Solidity: event LogProphecyProcessed(uint256 _prophecyID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_Oracle *OracleFilterer) ParseLogProphecyProcessed(log types.Log) (*OracleLogProphecyProcessed, error) {
	event := new(OracleLogProphecyProcessed)
	if err := _Oracle.contract.UnpackLog(event, "LogProphecyProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}
