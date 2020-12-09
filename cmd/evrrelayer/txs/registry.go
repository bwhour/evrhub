package txs

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	bridgeregistry "github.com/Evrynetlabs/evrhub/cmd/evrrelayer/ethContract/generated/bindings/bridgeregistry"
)

// TODO: Update BridgeRegistry ethContract so that all bridge ethContract addresses can be queried
//		in one transaction. Then refactor ContractRegistry to a map and store it under new
//		Relayer struct.

// ContractRegistry is an enum for the bridge ethContract types
type ContractRegistry byte

const (
	// Valset valset ethContract
	Valset ContractRegistry = iota + 1
	// Oracle ethContract
	Oracle
	// BridgeBank bridgeBank ethContract
	BridgeBank
	// EvrnetBridge cosmosBridge ethContract
	EvrnetBridge
)

// String returns the event type as a string
func (d ContractRegistry) String() string {
	return [...]string{"valset", "oracle", "bridgebank", "evrnetbridge"}[d-1]
}

// GetAddressFromBridgeRegistry queries the requested ethContract address from the BridgeRegistry ethContract
func GetAddressFromBridgeRegistry(client *ethclient.Client, registry common.Address, target ContractRegistry,
) (common.Address, error) {
	sender, err := LoadSender()
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set up CallOpts auth
	auth := bind.CallOpts{
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
		panic("invalid target ethContract address")
	}

	if err != nil {
		log.Fatal(err)
	}

	return address, nil
}
