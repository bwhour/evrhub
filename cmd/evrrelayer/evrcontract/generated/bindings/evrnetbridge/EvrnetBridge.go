// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EvrnetBridge

import (
	"math/big"
	"strings"

	ethereum "github.com/Evrynetlabs/evrynet-node"
	"github.com/Evrynetlabs/evrynet-node/accounts/abi"
	"github.com/Evrynetlabs/evrynet-node/accounts/abi/bind"
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/Evrynetlabs/evrynet-node/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EvrnetBridgeABI is the input ABI used to generate the binding from.
const EvrnetBridgeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"isProphecyClaimValidatorActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasBridgeBank\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"completeProphecyClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimType\",\"type\":\"uint8\"},{\"name\":\"_ethereumSender\",\"type\":\"address\"},{\"name\":\"_evrnetReceiver\",\"type\":\"address\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"newProphecyClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"setBridgeBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prophecyClaimCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"isProphecyClaimActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prophecyClaims\",\"outputs\":[{\"name\":\"claimType\",\"type\":\"uint8\"},{\"name\":\"ethereumSender\",\"type\":\"address\"},{\"name\":\"evrnetReceiver\",\"type\":\"address\"},{\"name\":\"originalValidator\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_valset\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"LogOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"LogBridgeBankSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_claimType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_ethereumSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_evrnetReceiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogNewProphecyClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"LogProphecyCompleted\",\"type\":\"event\"}]"

// EvrnetBridgeBin is the compiled bytecode used for deploying new contracts.
const EvrnetBridgeBin = `60806040526040805190810160405280600581526020017f5045474759000000000000000000000000000000000000000000000000000000815250600090805190602001906200005192919062000164565b503480156200005f57600080fd5b5060405160408062002c93833981018060405260408110156200008157600080fd5b810190808051906020019092919080519060200190929190505050600060058190555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600360146101000a81548160ff0219169083151502179055506000600460146101000a81548160ff021916908315150217905550505062000213565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001a757805160ff1916838001178555620001d8565b82800160010185558215620001d8579182015b82811115620001d7578251825591602001919060010190620001ba565b5b509050620001e79190620001eb565b5090565b6200021091905b808211156200020c576000816000905550600101620001f2565b5090565b90565b612a7080620002236000396000f3fe6080604052600436106100ca576000357c0100000000000000000000000000000000000000000000000000000000900480630e41f373146100cf578063529f3dd214610126578063570ca7351461017957806369294a4e146101d05780636b3ce98c146101ff5780637691ac0a1461023a5780637adbf973146103595780637dc0d1d0146103aa5780637f54af0c14610401578063814c92c3146104585780638ea5352d146104a9578063d8da69ea146104d4578063db4237af14610527578063fb7831f2146106d8575b600080fd5b3480156100db57600080fd5b506100e4610707565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561013257600080fd5b5061015f6004803603602081101561014957600080fd5b810190808035906020019092919050505061072d565b604051808215151515815260200191505060405180910390f35b34801561018557600080fd5b5061018e610862565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101dc57600080fd5b506101e5610888565b604051808215151515815260200191505060405180910390f35b34801561020b57600080fd5b506102386004803603602081101561022257600080fd5b810190808035906020019092919050505061089b565b005b34801561024657600080fd5b50610357600480360360a081101561025d57600080fd5b81019080803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156102c757600080fd5b8201836020820111156102d957600080fd5b803590602001918460018302840111640100000000831117156102fb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610aed565b005b34801561036557600080fd5b506103a86004803603602081101561037c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611894565b005b3480156103b657600080fd5b506103bf611ae8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561040d57600080fd5b50610416611b0e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046457600080fd5b506104a76004803603602081101561047b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b34565b005b3480156104b557600080fd5b506104be611d88565b6040518082815260200191505060405180910390f35b3480156104e057600080fd5b5061050d600480360360208110156104f757600080fd5b8101908080359060200190929190505050611d8e565b604051808215151515815260200191505060405180910390f35b34801561053357600080fd5b506105606004803603602081101561054a57600080fd5b8101908080359060200190929190505050611dd4565b6040518089600281111561057057fe5b60ff1681526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200184815260200183600381111561065657fe5b60ff168152602001828103825285818151815260200191508051906020019080838360005b8381101561069657808201518184015260208101905061067b565b50505050905090810190601f1680156106c35780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b3480156106e457600080fd5b506106ed611f4e565b604051808215151515815260200191505060405180910390f35b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c6006600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561082057600080fd5b505afa158015610834573d6000803e3d6000fd5b505050506040513d602081101561084a57600080fd5b81019080805190602001909291905050509050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460149054906101000a900460ff1681565b806108a581611d8e565b1515610919576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f50726f706865637920636c61696d206973206e6f74206163746976650000000081525060200191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a04576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001807f4f6e6c7920746865204f7261636c65206d617920636f6d706c6574652070726f81526020017f706865636965730000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60026006600084815260200190815260200160002060060160006101000a81548160ff02191690836003811115610a3757fe5b021790555060006006600084815260200190815260200160002060000160009054906101000a900460ff16905060016002811115610a7157fe5b816002811115610a7d57fe5b1415610a9157610a8c83611f61565b610a9b565b610a9a83612343565b5b7f79e7c1c0bd54f11809c3bf6023c242783602d61ceff272c6bba6f8559c24ad0d838260405180838152602001826002811115610ad457fe5b60ff1681526020019250505060405180910390a1505050565b60011515600360149054906101000a900460ff161515148015610b23575060011515600460149054906101000a900460ff161515145b1515610be3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260468152602001807f546865204f70657261746f72206d7573742073657420746865206f7261636c6581526020017f20616e64206272696467652062616e6b20666f7220627269646765206163746981526020017f766174696f6e000000000000000000000000000000000000000000000000000081525060600191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610c9e57600080fd5b505afa158015610cb2573d6000803e3d6000fd5b505050506040513d6020811015610cc857600080fd5b81019080805190602001909291905050501515610d4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b6000606060016002811115610d5e57fe5b876002811115610d6a57fe5b14156110755782600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635acba655866040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e1b578082015181840152602081019050610e00565b50505050905090810190601f168015610e485780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610e6557600080fd5b505afa158015610e79573d6000803e3d6000fd5b505050506040513d6020811015610e8f57600080fd5b810190808051906020019092919050505010151515610f3c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a8152602001807f4e6f7420656e6f756768206c6f636b65642061737365747320746f20636f6d7081526020017f6c657465207468652070726f706f7365642070726f706865637900000000000081525060400191505060405180910390fd5b839050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630a1f9b66856040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610fe9578082015181840152602081019050610fce565b50505050905090810190601f1680156110165780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561103357600080fd5b505afa158015611047573d6000803e3d6000fd5b505050506040513d602081101561105d57600080fd5b81019080805190602001909291905050509150611478565b60028081111561108157fe5b87600281111561108d57fe5b14156113e35761113760008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561112c5780601f106111015761010080835404028352916020019161112c565b820191906000526020600020905b81548152906001019060200180831161110f57829003601f168201915b505050505085612795565b90506000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebb73ca9836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156111e55780820151818401526020810190506111ca565b50505050905090810190601f1680156112125780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561122f57600080fd5b505afa158015611243573d6000803e3d6000fd5b505050506040513d602081101561125957600080fd5b81019080805190602001909291905050509050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156113d957600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166350b06e4d836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561134b578082015181840152602081019050611330565b50505050905090810190601f1680156113785780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561139757600080fd5b505af11580156113ab573d6000803e3d6000fd5b505050506040513d60208110156113c157600080fd5b810190808051906020019092919050505092506113dd565b8092505b50611477565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001807f496e76616c696420636c61696d20747970652c206f6e6c79206275726e20616e81526020017f64206c6f636b2061726520737570706f727465642e000000000000000000000081525060400191505060405180910390fd5b5b6114806128eb565b6101006040519081016040528089600281111561149957fe5b81526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018581526020016001600381111561152757fe5b8152509050611542600160055461286190919063ffffffff16565b6005819055508060066000600554815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600281111561158257fe5b021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040190805190602001906116bf92919061299f565b5060c0820151816005015560e08201518160060160006101000a81548160ff021916908360038111156116ee57fe5b02179055509050507f48d624cafeb4b91b2e3554842481d71d4f3f4a564ade66bbe6af69c196b27bd86005548989893388888b6040518089815260200188600281111561173757fe5b60ff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561184957808201518184015260208101905061182e565b50505050905090810190601f1680156118765780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a15050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611959576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600360149054906101000a900460ff16151515611a04576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260318152602001807f546865204f7261636c652063616e6e6f742062652075706461746564206f6e6381526020017f6520697420686173206265656e2073657400000000000000000000000000000081525060400191505060405180910390fd5b6001600360146101000a81548160ff02191690831515021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611bf9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600460149054906101000a900460ff16151515611ca4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f546865204272696467652042616e6b2063616e6e6f742062652075706461746581526020017f64206f6e636520697420686173206265656e207365740000000000000000000081525060400191505060405180910390fd5b6001600460146101000a81548160ff02191690831515021790555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60055481565b600060016003811115611d9d57fe5b6006600084815260200190815260200160002060060160009054906101000a900460ff166003811115611dcc57fe5b149050919050565b60066020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611f2b5780601f10611f0057610100808354040283529160200191611f2b565b820191906000526020600020905b815481529060010190602001808311611f0e57829003601f168201915b5050505050908060050154908060060160009054906101000a900460ff16905088565b600360149054906101000a900460ff1681565b611f696128eb565b6006600083815260200190815260200160002061010060405190810160405290816000820160009054906101000a900460ff166002811115611fa757fe5b6002811115611fb257fe5b81526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156121a75780601f1061217c576101008083540402835291602001916121a7565b820191906000526020600020905b81548152906001019060200180831161218a57829003601f168201915b50505050508152602001600582015481526020016006820160009054906101000a900460ff1660038111156121d857fe5b60038111156121e357fe5b815250509050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e05988a482604001518360a001518460c001516040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156122d95780820151818401526020810190506122be565b50505050905090810190601f1680156123065780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561232757600080fd5b505af115801561233b573d6000803e3d6000fd5b505050505050565b61234b6128eb565b6006600083815260200190815260200160002061010060405190810160405290816000820160009054906101000a900460ff16600281111561238957fe5b600281111561239457fe5b81526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156125895780601f1061255e57610100808354040283529160200191612589565b820191906000526020600020905b81548152906001019060200180831161256c57829003601f168201915b50505050508152602001600582015481526020016006820160009054906101000a900460ff1660038111156125ba57fe5b60038111156125c557fe5b815250509050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f9e6f5548260200151836040015184608001518560a001518660c001516040518663ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561272957808201518184015260208101905061270e565b50505050905090810190601f1680156127565780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b15801561277957600080fd5b505af115801561278d573d6000803e3d6000fd5b505050505050565b606082826040516020018083805190602001908083835b6020831015156127d157805182526020820191506020810190506020830392506127ac565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b60208310151561282457805182526020820191506020810190506020830392506127ff565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052905092915050565b60008082840190508381101515156128e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b610100604051908101604052806000600281111561290557fe5b8152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081526020016000600381111561299957fe5b81525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106129e057805160ff1916838001178555612a0e565b82800160010185558215612a0e579182015b82811115612a0d5782518255916020019190600101906129f2565b5b509050612a1b9190612a1f565b5090565b612a4191905b80821115612a3d576000816000905550600101612a25565b5090565b9056fea165627a7a723058205ab02ffb9261898323b52724041b5d4ae858efe228566e3784687e18bc3b673a0029`

// DeployEvrnetBridge deploys a new Evrynet contract, binding an instance of EvrnetBridge to it.
func DeployEvrnetBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _valset common.Address) (common.Address, *types.Transaction, *EvrnetBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(EvrnetBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EvrnetBridgeBin), backend, _operator, _valset)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EvrnetBridge{EvrnetBridgeCaller: EvrnetBridgeCaller{contract: contract}, EvrnetBridgeTransactor: EvrnetBridgeTransactor{contract: contract}, EvrnetBridgeFilterer: EvrnetBridgeFilterer{contract: contract}}, nil
}

// EvrnetBridge is an auto generated Go binding around an Evrynet contract.
type EvrnetBridge struct {
	EvrnetBridgeCaller     // Read-only binding to the contract
	EvrnetBridgeTransactor // Write-only binding to the contract
	EvrnetBridgeFilterer   // Log filterer for contract events
}

// EvrnetBridgeCaller is an auto generated read-only Go binding around an Evrynet contract.
type EvrnetBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvrnetBridgeTransactor is an auto generated write-only Go binding around an Evrynet contract.
type EvrnetBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvrnetBridgeFilterer is an auto generated log filtering Go binding around an Evrynet contract events.
type EvrnetBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvrnetBridgeSession is an auto generated Go binding around an Evrynet contract,
// with pre-set call and transact options.
type EvrnetBridgeSession struct {
	Contract     *EvrnetBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvrnetBridgeCallerSession is an auto generated read-only Go binding around an Evrynet contract,
// with pre-set call options.
type EvrnetBridgeCallerSession struct {
	Contract *EvrnetBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EvrnetBridgeTransactorSession is an auto generated write-only Go binding around an Evrynet contract,
// with pre-set transact options.
type EvrnetBridgeTransactorSession struct {
	Contract     *EvrnetBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EvrnetBridgeRaw is an auto generated low-level Go binding around an Evrynet contract.
type EvrnetBridgeRaw struct {
	Contract *EvrnetBridge // Generic contract binding to access the raw methods on
}

// EvrnetBridgeCallerRaw is an auto generated low-level read-only Go binding around an Evrynet contract.
type EvrnetBridgeCallerRaw struct {
	Contract *EvrnetBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// EvrnetBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Evrynet contract.
type EvrnetBridgeTransactorRaw struct {
	Contract *EvrnetBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvrnetBridge creates a new instance of EvrnetBridge, bound to a specific deployed contract.
func NewEvrnetBridge(address common.Address, backend bind.ContractBackend) (*EvrnetBridge, error) {
	contract, err := bindEvrnetBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EvrnetBridge{EvrnetBridgeCaller: EvrnetBridgeCaller{contract: contract}, EvrnetBridgeTransactor: EvrnetBridgeTransactor{contract: contract}, EvrnetBridgeFilterer: EvrnetBridgeFilterer{contract: contract}}, nil
}

// NewEvrnetBridgeCaller creates a new read-only instance of EvrnetBridge, bound to a specific deployed contract.
func NewEvrnetBridgeCaller(address common.Address, caller bind.ContractCaller) (*EvrnetBridgeCaller, error) {
	contract, err := bindEvrnetBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvrnetBridgeCaller{contract: contract}, nil
}

// NewEvrnetBridgeTransactor creates a new write-only instance of EvrnetBridge, bound to a specific deployed contract.
func NewEvrnetBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*EvrnetBridgeTransactor, error) {
	contract, err := bindEvrnetBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvrnetBridgeTransactor{contract: contract}, nil
}

// NewEvrnetBridgeFilterer creates a new log filterer instance of EvrnetBridge, bound to a specific deployed contract.
func NewEvrnetBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*EvrnetBridgeFilterer, error) {
	contract, err := bindEvrnetBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvrnetBridgeFilterer{contract: contract}, nil
}

// bindEvrnetBridge binds a generic wrapper to an already deployed contract.
func bindEvrnetBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EvrnetBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvrnetBridge *EvrnetBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EvrnetBridge.Contract.EvrnetBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvrnetBridge *EvrnetBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.EvrnetBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvrnetBridge *EvrnetBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.EvrnetBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvrnetBridge *EvrnetBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EvrnetBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvrnetBridge *EvrnetBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvrnetBridge *EvrnetBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "bridgeBank")
	return *ret0, err
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeSession) BridgeBank() (common.Address, error) {
	return _EvrnetBridge.Contract.BridgeBank(&_EvrnetBridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCallerSession) BridgeBank() (common.Address, error) {
	return _EvrnetBridge.Contract.BridgeBank(&_EvrnetBridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCaller) HasBridgeBank(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "hasBridgeBank")
	return *ret0, err
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeSession) HasBridgeBank() (bool, error) {
	return _EvrnetBridge.Contract.HasBridgeBank(&_EvrnetBridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCallerSession) HasBridgeBank() (bool, error) {
	return _EvrnetBridge.Contract.HasBridgeBank(&_EvrnetBridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCaller) HasOracle(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "hasOracle")
	return *ret0, err
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeSession) HasOracle() (bool, error) {
	return _EvrnetBridge.Contract.HasOracle(&_EvrnetBridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCallerSession) HasOracle() (bool, error) {
	return _EvrnetBridge.Contract.HasOracle(&_EvrnetBridge.CallOpts)
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xd8da69ea.
//
// Solidity: function isProphecyClaimActive(uint256 _prophecyID) constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCaller) IsProphecyClaimActive(opts *bind.CallOpts, _prophecyID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "isProphecyClaimActive", _prophecyID)
	return *ret0, err
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xd8da69ea.
//
// Solidity: function isProphecyClaimActive(uint256 _prophecyID) constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeSession) IsProphecyClaimActive(_prophecyID *big.Int) (bool, error) {
	return _EvrnetBridge.Contract.IsProphecyClaimActive(&_EvrnetBridge.CallOpts, _prophecyID)
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xd8da69ea.
//
// Solidity: function isProphecyClaimActive(uint256 _prophecyID) constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCallerSession) IsProphecyClaimActive(_prophecyID *big.Int) (bool, error) {
	return _EvrnetBridge.Contract.IsProphecyClaimActive(&_EvrnetBridge.CallOpts, _prophecyID)
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x529f3dd2.
//
// Solidity: function isProphecyClaimValidatorActive(uint256 _prophecyID) constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCaller) IsProphecyClaimValidatorActive(opts *bind.CallOpts, _prophecyID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "isProphecyClaimValidatorActive", _prophecyID)
	return *ret0, err
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x529f3dd2.
//
// Solidity: function isProphecyClaimValidatorActive(uint256 _prophecyID) constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeSession) IsProphecyClaimValidatorActive(_prophecyID *big.Int) (bool, error) {
	return _EvrnetBridge.Contract.IsProphecyClaimValidatorActive(&_EvrnetBridge.CallOpts, _prophecyID)
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x529f3dd2.
//
// Solidity: function isProphecyClaimValidatorActive(uint256 _prophecyID) constant returns(bool)
func (_EvrnetBridge *EvrnetBridgeCallerSession) IsProphecyClaimValidatorActive(_prophecyID *big.Int) (bool, error) {
	return _EvrnetBridge.Contract.IsProphecyClaimValidatorActive(&_EvrnetBridge.CallOpts, _prophecyID)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeSession) Operator() (common.Address, error) {
	return _EvrnetBridge.Contract.Operator(&_EvrnetBridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCallerSession) Operator() (common.Address, error) {
	return _EvrnetBridge.Contract.Operator(&_EvrnetBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeSession) Oracle() (common.Address, error) {
	return _EvrnetBridge.Contract.Oracle(&_EvrnetBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCallerSession) Oracle() (common.Address, error) {
	return _EvrnetBridge.Contract.Oracle(&_EvrnetBridge.CallOpts)
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() constant returns(uint256)
func (_EvrnetBridge *EvrnetBridgeCaller) ProphecyClaimCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "prophecyClaimCount")
	return *ret0, err
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() constant returns(uint256)
func (_EvrnetBridge *EvrnetBridgeSession) ProphecyClaimCount() (*big.Int, error) {
	return _EvrnetBridge.Contract.ProphecyClaimCount(&_EvrnetBridge.CallOpts)
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() constant returns(uint256)
func (_EvrnetBridge *EvrnetBridgeCallerSession) ProphecyClaimCount() (*big.Int, error) {
	return _EvrnetBridge.Contract.ProphecyClaimCount(&_EvrnetBridge.CallOpts)
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xdb4237af.
//
// Solidity: function prophecyClaims(uint256 ) constant returns(uint8 claimType, address ethereumSender, address evrnetReceiver, address originalValidator, address tokenAddress, string symbol, uint256 amount, uint8 status)
func (_EvrnetBridge *EvrnetBridgeCaller) ProphecyClaims(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ClaimType         uint8
	EthereumSender    common.Address
	EvrnetReceiver    common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Symbol            string
	Amount            *big.Int
	Status            uint8
}, error) {
	ret := new(struct {
		ClaimType         uint8
		EthereumSender    common.Address
		EvrnetReceiver    common.Address
		OriginalValidator common.Address
		TokenAddress      common.Address
		Symbol            string
		Amount            *big.Int
		Status            uint8
	})
	out := ret
	err := _EvrnetBridge.contract.Call(opts, out, "prophecyClaims", arg0)
	return *ret, err
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xdb4237af.
//
// Solidity: function prophecyClaims(uint256 ) constant returns(uint8 claimType, address ethereumSender, address evrnetReceiver, address originalValidator, address tokenAddress, string symbol, uint256 amount, uint8 status)
func (_EvrnetBridge *EvrnetBridgeSession) ProphecyClaims(arg0 *big.Int) (struct {
	ClaimType         uint8
	EthereumSender    common.Address
	EvrnetReceiver    common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Symbol            string
	Amount            *big.Int
	Status            uint8
}, error) {
	return _EvrnetBridge.Contract.ProphecyClaims(&_EvrnetBridge.CallOpts, arg0)
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xdb4237af.
//
// Solidity: function prophecyClaims(uint256 ) constant returns(uint8 claimType, address ethereumSender, address evrnetReceiver, address originalValidator, address tokenAddress, string symbol, uint256 amount, uint8 status)
func (_EvrnetBridge *EvrnetBridgeCallerSession) ProphecyClaims(arg0 *big.Int) (struct {
	ClaimType         uint8
	EthereumSender    common.Address
	EvrnetReceiver    common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Symbol            string
	Amount            *big.Int
	Status            uint8
}, error) {
	return _EvrnetBridge.Contract.ProphecyClaims(&_EvrnetBridge.CallOpts, arg0)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EvrnetBridge.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeSession) Valset() (common.Address, error) {
	return _EvrnetBridge.Contract.Valset(&_EvrnetBridge.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() constant returns(address)
func (_EvrnetBridge *EvrnetBridgeCallerSession) Valset() (common.Address, error) {
	return _EvrnetBridge.Contract.Valset(&_EvrnetBridge.CallOpts)
}

// CompleteProphecyClaim is a paid mutator transaction binding the contract method 0x6b3ce98c.
//
// Solidity: function completeProphecyClaim(uint256 _prophecyID) returns()
func (_EvrnetBridge *EvrnetBridgeTransactor) CompleteProphecyClaim(opts *bind.TransactOpts, _prophecyID *big.Int) (*types.Transaction, error) {
	return _EvrnetBridge.contract.Transact(opts, "completeProphecyClaim", _prophecyID)
}

// CompleteProphecyClaim is a paid mutator transaction binding the contract method 0x6b3ce98c.
//
// Solidity: function completeProphecyClaim(uint256 _prophecyID) returns()
func (_EvrnetBridge *EvrnetBridgeSession) CompleteProphecyClaim(_prophecyID *big.Int) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.CompleteProphecyClaim(&_EvrnetBridge.TransactOpts, _prophecyID)
}

// CompleteProphecyClaim is a paid mutator transaction binding the contract method 0x6b3ce98c.
//
// Solidity: function completeProphecyClaim(uint256 _prophecyID) returns()
func (_EvrnetBridge *EvrnetBridgeTransactorSession) CompleteProphecyClaim(_prophecyID *big.Int) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.CompleteProphecyClaim(&_EvrnetBridge.TransactOpts, _prophecyID)
}

// NewProphecyClaim is a paid mutator transaction binding the contract method 0x7691ac0a.
//
// Solidity: function newProphecyClaim(uint8 _claimType, address _ethereumSender, address _evrnetReceiver, string _symbol, uint256 _amount) returns()
func (_EvrnetBridge *EvrnetBridgeTransactor) NewProphecyClaim(opts *bind.TransactOpts, _claimType uint8, _ethereumSender common.Address, _evrnetReceiver common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _EvrnetBridge.contract.Transact(opts, "newProphecyClaim", _claimType, _ethereumSender, _evrnetReceiver, _symbol, _amount)
}

// NewProphecyClaim is a paid mutator transaction binding the contract method 0x7691ac0a.
//
// Solidity: function newProphecyClaim(uint8 _claimType, address _ethereumSender, address _evrnetReceiver, string _symbol, uint256 _amount) returns()
func (_EvrnetBridge *EvrnetBridgeSession) NewProphecyClaim(_claimType uint8, _ethereumSender common.Address, _evrnetReceiver common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.NewProphecyClaim(&_EvrnetBridge.TransactOpts, _claimType, _ethereumSender, _evrnetReceiver, _symbol, _amount)
}

// NewProphecyClaim is a paid mutator transaction binding the contract method 0x7691ac0a.
//
// Solidity: function newProphecyClaim(uint8 _claimType, address _ethereumSender, address _evrnetReceiver, string _symbol, uint256 _amount) returns()
func (_EvrnetBridge *EvrnetBridgeTransactorSession) NewProphecyClaim(_claimType uint8, _ethereumSender common.Address, _evrnetReceiver common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.NewProphecyClaim(&_EvrnetBridge.TransactOpts, _claimType, _ethereumSender, _evrnetReceiver, _symbol, _amount)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_EvrnetBridge *EvrnetBridgeTransactor) SetBridgeBank(opts *bind.TransactOpts, _bridgeBank common.Address) (*types.Transaction, error) {
	return _EvrnetBridge.contract.Transact(opts, "setBridgeBank", _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_EvrnetBridge *EvrnetBridgeSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.SetBridgeBank(&_EvrnetBridge.TransactOpts, _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_EvrnetBridge *EvrnetBridgeTransactorSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.SetBridgeBank(&_EvrnetBridge.TransactOpts, _bridgeBank)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_EvrnetBridge *EvrnetBridgeTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _EvrnetBridge.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_EvrnetBridge *EvrnetBridgeSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.SetOracle(&_EvrnetBridge.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_EvrnetBridge *EvrnetBridgeTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _EvrnetBridge.Contract.SetOracle(&_EvrnetBridge.TransactOpts, _oracle)
}

// EvrnetBridgeLogBridgeBankSetIterator is returned from FilterLogBridgeBankSet and is used to iterate over the raw logs and unpacked data for LogBridgeBankSet events raised by the EvrnetBridge contract.
type EvrnetBridgeLogBridgeBankSetIterator struct {
	Event *EvrnetBridgeLogBridgeBankSet // Event containing the contract specifics and raw log

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
func (it *EvrnetBridgeLogBridgeBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvrnetBridgeLogBridgeBankSet)
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
		it.Event = new(EvrnetBridgeLogBridgeBankSet)
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
func (it *EvrnetBridgeLogBridgeBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvrnetBridgeLogBridgeBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvrnetBridgeLogBridgeBankSet represents a LogBridgeBankSet event raised by the EvrnetBridge contract.
type EvrnetBridgeLogBridgeBankSet struct {
	BridgeBank common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeBankSet is a free log retrieval operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_EvrnetBridge *EvrnetBridgeFilterer) FilterLogBridgeBankSet(opts *bind.FilterOpts) (*EvrnetBridgeLogBridgeBankSetIterator, error) {

	logs, sub, err := _EvrnetBridge.contract.FilterLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return &EvrnetBridgeLogBridgeBankSetIterator{contract: _EvrnetBridge.contract, event: "LogBridgeBankSet", logs: logs, sub: sub}, nil
}

// WatchLogBridgeBankSet is a free log subscription operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_EvrnetBridge *EvrnetBridgeFilterer) WatchLogBridgeBankSet(opts *bind.WatchOpts, sink chan<- *EvrnetBridgeLogBridgeBankSet) (event.Subscription, error) {

	logs, sub, err := _EvrnetBridge.contract.WatchLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvrnetBridgeLogBridgeBankSet)
				if err := _EvrnetBridge.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
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

// EvrnetBridgeLogNewProphecyClaimIterator is returned from FilterLogNewProphecyClaim and is used to iterate over the raw logs and unpacked data for LogNewProphecyClaim events raised by the EvrnetBridge contract.
type EvrnetBridgeLogNewProphecyClaimIterator struct {
	Event *EvrnetBridgeLogNewProphecyClaim // Event containing the contract specifics and raw log

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
func (it *EvrnetBridgeLogNewProphecyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvrnetBridgeLogNewProphecyClaim)
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
		it.Event = new(EvrnetBridgeLogNewProphecyClaim)
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
func (it *EvrnetBridgeLogNewProphecyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvrnetBridgeLogNewProphecyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvrnetBridgeLogNewProphecyClaim represents a LogNewProphecyClaim event raised by the EvrnetBridge contract.
type EvrnetBridgeLogNewProphecyClaim struct {
	ProphecyID       *big.Int
	ClaimType        uint8
	EthereumSender   common.Address
	EvrnetReceiver   common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Symbol           string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewProphecyClaim is a free log retrieval operation binding the contract event 0x48d624cafeb4b91b2e3554842481d71d4f3f4a564ade66bbe6af69c196b27bd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, address _ethereumSender, address _evrnetReceiver, address _validatorAddress, address _tokenAddress, string _symbol, uint256 _amount)
func (_EvrnetBridge *EvrnetBridgeFilterer) FilterLogNewProphecyClaim(opts *bind.FilterOpts) (*EvrnetBridgeLogNewProphecyClaimIterator, error) {

	logs, sub, err := _EvrnetBridge.contract.FilterLogs(opts, "LogNewProphecyClaim")
	if err != nil {
		return nil, err
	}
	return &EvrnetBridgeLogNewProphecyClaimIterator{contract: _EvrnetBridge.contract, event: "LogNewProphecyClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewProphecyClaim is a free log subscription operation binding the contract event 0x48d624cafeb4b91b2e3554842481d71d4f3f4a564ade66bbe6af69c196b27bd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, address _ethereumSender, address _evrnetReceiver, address _validatorAddress, address _tokenAddress, string _symbol, uint256 _amount)
func (_EvrnetBridge *EvrnetBridgeFilterer) WatchLogNewProphecyClaim(opts *bind.WatchOpts, sink chan<- *EvrnetBridgeLogNewProphecyClaim) (event.Subscription, error) {

	logs, sub, err := _EvrnetBridge.contract.WatchLogs(opts, "LogNewProphecyClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvrnetBridgeLogNewProphecyClaim)
				if err := _EvrnetBridge.contract.UnpackLog(event, "LogNewProphecyClaim", log); err != nil {
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

// EvrnetBridgeLogOracleSetIterator is returned from FilterLogOracleSet and is used to iterate over the raw logs and unpacked data for LogOracleSet events raised by the EvrnetBridge contract.
type EvrnetBridgeLogOracleSetIterator struct {
	Event *EvrnetBridgeLogOracleSet // Event containing the contract specifics and raw log

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
func (it *EvrnetBridgeLogOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvrnetBridgeLogOracleSet)
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
		it.Event = new(EvrnetBridgeLogOracleSet)
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
func (it *EvrnetBridgeLogOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvrnetBridgeLogOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvrnetBridgeLogOracleSet represents a LogOracleSet event raised by the EvrnetBridge contract.
type EvrnetBridgeLogOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogOracleSet is a free log retrieval operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_EvrnetBridge *EvrnetBridgeFilterer) FilterLogOracleSet(opts *bind.FilterOpts) (*EvrnetBridgeLogOracleSetIterator, error) {

	logs, sub, err := _EvrnetBridge.contract.FilterLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return &EvrnetBridgeLogOracleSetIterator{contract: _EvrnetBridge.contract, event: "LogOracleSet", logs: logs, sub: sub}, nil
}

// WatchLogOracleSet is a free log subscription operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_EvrnetBridge *EvrnetBridgeFilterer) WatchLogOracleSet(opts *bind.WatchOpts, sink chan<- *EvrnetBridgeLogOracleSet) (event.Subscription, error) {

	logs, sub, err := _EvrnetBridge.contract.WatchLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvrnetBridgeLogOracleSet)
				if err := _EvrnetBridge.contract.UnpackLog(event, "LogOracleSet", log); err != nil {
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

// EvrnetBridgeLogProphecyCompletedIterator is returned from FilterLogProphecyCompleted and is used to iterate over the raw logs and unpacked data for LogProphecyCompleted events raised by the EvrnetBridge contract.
type EvrnetBridgeLogProphecyCompletedIterator struct {
	Event *EvrnetBridgeLogProphecyCompleted // Event containing the contract specifics and raw log

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
func (it *EvrnetBridgeLogProphecyCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvrnetBridgeLogProphecyCompleted)
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
		it.Event = new(EvrnetBridgeLogProphecyCompleted)
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
func (it *EvrnetBridgeLogProphecyCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvrnetBridgeLogProphecyCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvrnetBridgeLogProphecyCompleted represents a LogProphecyCompleted event raised by the EvrnetBridge contract.
type EvrnetBridgeLogProphecyCompleted struct {
	ProphecyID *big.Int
	ClaimType  uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogProphecyCompleted is a free log retrieval operation binding the contract event 0x79e7c1c0bd54f11809c3bf6023c242783602d61ceff272c6bba6f8559c24ad0d.
//
// Solidity: event LogProphecyCompleted(uint256 _prophecyID, uint8 _claimType)
func (_EvrnetBridge *EvrnetBridgeFilterer) FilterLogProphecyCompleted(opts *bind.FilterOpts) (*EvrnetBridgeLogProphecyCompletedIterator, error) {

	logs, sub, err := _EvrnetBridge.contract.FilterLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return &EvrnetBridgeLogProphecyCompletedIterator{contract: _EvrnetBridge.contract, event: "LogProphecyCompleted", logs: logs, sub: sub}, nil
}

// WatchLogProphecyCompleted is a free log subscription operation binding the contract event 0x79e7c1c0bd54f11809c3bf6023c242783602d61ceff272c6bba6f8559c24ad0d.
//
// Solidity: event LogProphecyCompleted(uint256 _prophecyID, uint8 _claimType)
func (_EvrnetBridge *EvrnetBridgeFilterer) WatchLogProphecyCompleted(opts *bind.WatchOpts, sink chan<- *EvrnetBridgeLogProphecyCompleted) (event.Subscription, error) {

	logs, sub, err := _EvrnetBridge.contract.WatchLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvrnetBridgeLogProphecyCompleted)
				if err := _EvrnetBridge.contract.UnpackLog(event, "LogProphecyCompleted", log); err != nil {
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
