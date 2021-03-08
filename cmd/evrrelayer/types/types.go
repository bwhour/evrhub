package types

import (
	"fmt"
	"math/big"

	xcommon "github.com/Evrynetlabs/evrhub/x/common/types"
	"github.com/Evrynetlabs/evrynet-node/common"
)

// Event enum containing supported chain events
type Event byte

const (
	// Unsupported is an invalid Evrnet or Ethereum event
	Unsupported Event = iota
	// MsgBurn is a Ethereum msg of type MsgBurn
	MsgBurn
	// MsgLock is a Ethereum msg of type MsgLock
	MsgLock
	// LogLock is for Evrnet event LogLock
	LogLock
	// LogBurn is for Evrnet event LogBurn
	LogBurn
	// LogNewProphecyClaim is an Ethereum event named 'LogNewProphecyClaim'
	LogNewProphecyClaim
)

// String returns the event type as a string
func (d Event) String() string {
	return [...]string{"unsupported", "burn", "lock", "LogLock", "LogBurn", "LogNewProphecyClaim"}[d]
}

// EvrnetEvent struct is used by LogLock and LogBurn
type EvrnetEvent struct {
	EvrnetChainID         *big.Int
	BridgeContractAddress common.Address
	ID                    [32]byte
	From                  common.Address
	To                    common.Address
	Token                 common.Address
	Symbol                string
	Value                 *big.Int
	Nonce                 *big.Int
	ClaimType             xcommon.ClaimType
}

// String implements fmt.Stringer
func (e EvrnetEvent) String() string {
	return fmt.Sprintf("\nChain ID: %v\nBridge evrcontract address: %v\nToken symbol: %v\nToken "+
		"evrcontract address: %v\nSender: %v\nRecipient: %v\nValue: %v\nNonce: %v\nClaim type: %v",
		e.EvrnetChainID, e.BridgeContractAddress.Hex(), e.Symbol, e.Token.Hex(), e.From.Hex(),
		e.To.Hex(), e.Value, e.Nonce, e.ClaimType.String())
}

// ProphecyClaimEvent struct which represents a LogNewProphecyClaim event
type ProphecyClaimEvent struct {
	EthereumSender   common.Address
	Symbol           string
	ProphecyID       *big.Int
	Amount           *big.Int
	EvrnetReceiver   common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	ClaimType        uint8
}

// String implements fmt.Stringer
func (p ProphecyClaimEvent) String() string {
	return fmt.Sprintf("\nProphecy ID: %v\nClaim Type: %v\nSender: %v\n"+
		"Recipient: %v\nSymbol: %v\nToken: %v\nAmount: %v\nValidator: %v\n\n",
		p.ProphecyID, p.ClaimType, p.EthereumSender.Hex(), p.EvrnetReceiver.Hex(),
		p.Symbol, p.TokenAddress.Hex(), p.Amount, p.ValidatorAddress.Hex())
}
