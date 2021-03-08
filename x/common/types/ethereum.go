package types

import (
	"fmt"
	"reflect"

	"github.com/Evrynetlabs/evrynet-node/common"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const PeggedCoinPrefix = "peggy"

// EthereumAddress defines a standard ethereum address
type EthereumAddress common.Address

// NewEthereumAddress is a constructor function for EthereumAddress
func NewEthereumAddress(address string) EthereumAddress {
	return EthereumAddress(common.HexToAddress(address))
}

// Route should return the name of the module
func (ethAddr EthereumAddress) String() string {
	return ethcommon.Address(ethAddr).String()
}

// MarshalJSON marshals the ethereum address to JSON
func (ethAddr EthereumAddress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", ethAddr.String())), nil
}

// UnmarshalJSON unmarshal an ethereum address
func (ethAddr *EthereumAddress) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(ethcommon.Address{}), input, ethAddr[:])
}
