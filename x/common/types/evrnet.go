package types

import (
	"fmt"
	"reflect"

	evrcommon "github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/common/hexutil"
)

// EvrnetAddress defines a standard ethereum address
type EvrnetAddress evrcommon.Address

// NewEvrnetAddress is a constructor function for EvrnetAddress
func NewEvrnetAddress(address string) EvrnetAddress {
	return EvrnetAddress(evrcommon.HexToAddress(address))
}

// Route should return the name of the module
func (evrAddr EvrnetAddress) String() string {
	return evrcommon.Address(evrAddr).String()
}

// MarshalJSON marshals the ethereum address to JSON
func (evrAddr EvrnetAddress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", evrAddr.String())), nil
}

// UnmarshalJSON unmarshal an ethereum address
func (evrAddr *EvrnetAddress) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(evrcommon.Address{}), input, evrAddr[:])
}
