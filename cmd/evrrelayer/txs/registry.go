package txs

import (
	"context"
	bridgeregistry "github.com/Evrynetlabs/evrhub/cmd/evrrelayer/evrcontract/generated/bindings/bridgeregistry"
	evrBind "github.com/Evrynetlabs/evrynet-node/accounts/abi/bind"
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/evrclient"
	"log"
)

// TODO: Update BridgeRegistry evrcontract so that all bridge evrcontract addresses can be queried
//		in one transaction. Then refactor ContractRegistry to a map and store it under new
//		Relayer struct.

// ContractRegistry is an enum for the bridge evrcontract types
type ContractRegistry byte

const (
	// Valset valset evrcontract
	Valset ContractRegistry = iota + 1
	// Oracle evrcontract
	Oracle
	// BridgeBank bridgeBank evrcontract
	BridgeBank
	// EvrnetBridge evrnetBridge evrcontract
	EvrnetBridge
)

// String returns the event type as a string
func (d ContractRegistry) String() string {
	return [...]string{"valset", "oracle", "bridgebank", "evrnetbridge"}[d-1]
}

// GetAddressFromBridgeRegistry queries the requested evrcontract address from the BridgeRegistry evrcontract
func GetAddressFromBridgeRegistry(client *evrclient.Client, registry common.Address, target ContractRegistry,
) (common.Address, error) {
	sender, err := LoadEvrSender()
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set up CallOpts auth
	//var xxx *big.Int
	//var sender .Address
	auth := evrBind.CallOpts{
		Pending:     true,
		From:        sender,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	// Initialize BridgeRegistry instance
	registryInstance, err := bridgeregistry.NewBridgeRegistry(registry, client)
	if err != nil {
		log.Fatal(err)
	}

	var address common.Address
	switch target {
	case Valset:
		address, err = registryInstance.Valset(&auth)
	case Oracle:
		address, err = registryInstance.Oracle(&auth)
	case BridgeBank:
		address, err = registryInstance.BridgeBank(&auth)
	case EvrnetBridge:
		address, err = registryInstance.EvrnetBridge(&auth)
	default:
		panic("invalid target evrcontract address")
	}

	if err != nil {
		log.Fatal(err)
	}

	return address, nil
}


