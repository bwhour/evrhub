package types

import (
	"fmt"
	"reflect"

	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	evrCommon "github.com/Evrynetlabs/evrynet-node/common"
	evrhexutil "github.com/Evrynetlabs/evrynet-node/common/hexutil"
)

const PeggedCoinPrefix = "peggy"

// EthereumAddress defines a standard ethereum address
type EthereumAddress gethCommon.Address

// NewEthereumAddress is a constructor function for EthereumAddress
func NewEthereumAddress(address string) EthereumAddress {
	return EthereumAddress(gethCommon.HexToAddress(address))
}

// Route should return the name of the module
func (ethAddr EthereumAddress) String() string {
	return gethCommon.Address(ethAddr).String()
}

// MarshalJSON marshals the etherum address to JSON
func (ethAddr EthereumAddress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", ethAddr.String())), nil
}

// UnmarshalJSON unmarshals an ethereum address
func (ethAddr *EthereumAddress) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(gethCommon.Address{}), input, ethAddr[:])
}
// EvrnetAddress defines a standard ethereum address
type EvrnetAddress evrCommon.Address

// NewEvrnetAddress is a constructor function for EthereumAddress
func NewEvrnetAddress(address string) EvrnetAddress {
	return EvrnetAddress(evrCommon.HexToAddress(address))
}

// Route should return the name of the module
func (evrAddr EvrnetAddress) String() string {
	return evrCommon.Address(ethAddr).String()
}

// MarshalJSON marshals the evrnet address to JSON
func (evrAddr EvrnetAddress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", evrAddr.String())), nil
}

// UnmarshalJSON unmarshals an evrnet address
func (evrAddr *EvrnetAddress) UnmarshalJSON(input []byte) error {
	return evrhexutil.UnmarshalFixedJSON(reflect.TypeOf(evrCommon.Address{}), input, evrAddr[:])
}
