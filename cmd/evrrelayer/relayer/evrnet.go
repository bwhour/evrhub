package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/evrcontract"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/rpc/client"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/txs"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"
	"github.com/Evrynetlabs/evrynet-node"
	"github.com/Evrynetlabs/evrynet-node/accounts/abi"
	"github.com/Evrynetlabs/evrynet-node/common"
	ctypes "github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/Evrynetlabs/evrynet-node/evrclient"

	xcommon "github.com/Evrynetlabs/evrhub/x/common/types"
	evrcommon "github.com/Evrynetlabs/evrynet-node/common"
	tmLog "github.com/tendermint/tendermint/libs/log"
)

// TODO: Move relay functionality out of EvrnetSub into a new Relayer parent struct

// EvrnetSub defines a Evrnet listener that relays events to Ethereum and Evrnet
type EvrnetSub struct {
	EvrProvider             string
	RegistryContractAddress evrcommon.Address
	PrivateKey              *ecdsa.PrivateKey
	Logger                  tmLog.Logger
}

// NewEvrnetSub initializes a new EvrnetSub
func NewEvrnetSub(evrProvider string, registryContractAddress evrcommon.Address,
	privateKey *ecdsa.PrivateKey, logger tmLog.Logger) EvrnetSub {
	return EvrnetSub{
		EvrProvider:             evrProvider,
		RegistryContractAddress: registryContractAddress,
		PrivateKey:              privateKey,
		Logger:                  logger,
	}
}

// Start a Evrnet chain subscription
func (sub EvrnetSub) Start() {
	client, err := SetupWebsocketEvrClient(sub.EvrProvider)
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	sub.Logger.Info("Started Ethereum websocket with provider:", sub.EvrProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		println(clientChainID)
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	//We will check logs for new events
	logs := make(chan ctypes.Log)
	//Start BridgeBank subscription, prepare evrcontract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := evrcontract.LoadABI(txs.BridgeBank)
	eventLogLockSignature := bridgeBankContractABI.Events[types.LogLock.String()].Id().Hex()
	eventLogBurnSignature := bridgeBankContractABI.Events[types.LogBurn.String()].Id().Hex()

	// Start EvrnetBridge subscription, prepare ethcontract ABI and LogNewProphecyClaim event signature
	evrnetBridgeAddress, subEvrnetBridge := sub.startContractEventSub(logs, client, txs.EvrnetBridge)
	evrnetBridgeContractABI := evrcontract.LoadABI(txs.EvrnetBridge)
	eventLogNewProphecyClaimSignature := evrnetBridgeContractABI.Events[types.LogNewProphecyClaim.String()].Id().Hex()

	for {
		select {
		// Handle any errors
		case err := <-subBridgeBank.Err():
			sub.Logger.Error(err.Error())
		case err := <-subEvrnetBridge.Err():
			sub.Logger.Error(err.Error())
		//vLog is raw event data
		case vLog := <-logs:
			sub.Logger.Info(fmt.Sprintf("Witnessed tx %s on block %d\n", vLog.TxHash.Hex(), vLog.BlockNumber))
			var err error
			switch vLog.Topics[0].Hex() {
			case eventLogBurnSignature:
				err = sub.handleEvrnetEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.LogBurn.String(), vLog)
			case eventLogLockSignature:
				err = sub.handleEvrnetEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.LogLock.String(), vLog)
			case eventLogNewProphecyClaimSignature:
				err = sub.handleLogNewProphecyClaim(evrnetBridgeAddress, evrnetBridgeContractABI,
					types.LogNewProphecyClaim.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
			}
		}
	}
}

// startContractEventSub : starts an event subscription on the specified Peggy evrcontract
func (sub EvrnetSub) startContractEventSub(logs chan ctypes.Log, client *evrclient.Client,
	contractName txs.ContractRegistry) (common.Address, evrynet.Subscription) {
	// Get the evrcontract address for this subscription
	subContractAddress, err := txs.GetAddressFromBridgeRegistry(client, sub.RegistryContractAddress, contractName)
	if err != nil {
		sub.Logger.Error(err.Error())
	}
	//var subContractAddress common.Address
	// We need the address in []bytes for the query
	subQuery := evrynet.FilterQuery{
		Addresses: []common.Address{subContractAddress},
	}
	println(&subQuery)

	// Start the evrcontract subscription
	contractSub, err := client.SubscribeFilterLogs(context.Background(), subQuery, logs)
	if err != nil {
		sub.Logger.Error(err.Error())
	}
	//var contractSub evrynet.Subscription
	sub.Logger.Info(fmt.Sprintf("Subscribed to %v evrcontract at address: %s", contractName, subContractAddress.Hex()))
	return subContractAddress, contractSub
}

// handleEthereumEvent unpacks an Ethereum event, converts it to a ProphecyClaim, and relays a tx to Evrnet
func (sub EvrnetSub) handleEvrnetEvent(clientChainID *big.Int, contractAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via evrcontract ABI
	event := types.EvrnetEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeContractAddress = contractAddress
	event.EvrnetChainID = clientChainID

	if eventName == types.LogBurn.String() {
		event.ClaimType = xcommon.BurnText
	} else {
		event.ClaimType = xcommon.LockText
	}
	sub.Logger.Info(event.String())

	// Add the event to the record
	types.NewEventWrite(cLog.TxHash.Hex(), event)

	prophecyClaim := xcommon.EvrProphecyClaim{
		ClaimType:        event.ClaimType,
		EvrnetSender:     xcommon.NewEvrnetAddress(event.From.Hex()),
		EthereumReceiver: xcommon.NewEthereumAddress(event.To.Hex()),
		Symbol:           event.Symbol,
		Amount:           event.Value.String(),
	}

	//send claim to EthRelayer
	_, err = client.SendProphecyClaimToEthereum(prophecyClaim)
	if err != nil {
		return err
	}

	return nil
}

// Unpacks a handleLogNewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub EvrnetSub) handleLogNewProphecyClaim(contractAddress common.Address, contractABI abi.ABI,
	eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via ethcontract ABI
	event := types.ProphecyClaimEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	sub.Logger.Info(event.String())

	oracleClaim, err := txs.ProphecyClaimToSignedOracleClaim(event, sub.PrivateKey)
	if err != nil {
		return err
	}
	return txs.RelayOracleClaimToEvrnet(sub.EvrProvider, contractAddress, types.LogNewProphecyClaim,
		oracleClaim, sub.PrivateKey)
}
