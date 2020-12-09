package relayer

import (

	"context"
	"crypto/ecdsa"
	ethContract "github.com/Evrynetlabs/evrhub/cmd/evrrelayer/ethcontract"
	ethbridge "github.com/Evrynetlabs/evrhub/x/ethbridge/types"
	"github.com/Evrynetlabs/evrynet-node/evrclient"
	"github.com/Evrynetlabs/evrynet-node"
	ctypes "github.com/Evrynetlabs/evrynet-node/core/types"
	tmKv "github.com/tendermint/tendermint/libs/kv"
	tmLog "github.com/tendermint/tendermint/libs/log"
	"math/big"
	"os"
	"fmt"
	"github.com/Evrynetlabs/evrynet-node/common"
    "github.com/Evrynetlabs/evrhub/cmd/evrrelayer/txs"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"
)

// TODO: Move relay functionality out of EvrnetSub into a new Relayer parent struct

// EvrnetSub defines a Evrnet listener that relays events to Ethereum and Evrnet
type EvrnetSub struct {
	EvrProvider             string
	EthProvider             string
	RegistryContractAddress common.Address
	PrivateKey              *ecdsa.PrivateKey
	Logger                  tmLog.Logger
}

// NewEvrnetSub initializes a new EvrnetSub
func NewEvrnetSub(evrProvider, ethProvider string, registryContractAddress common.Address,
	privateKey *ecdsa.PrivateKey, logger tmLog.Logger) EvrnetSub {
	return EvrnetSub{
		EvrProvider:             evrProvider,
		EthProvider:             ethProvider,
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
	sub.Logger.Info("Started Ethereum websocket with provider:", sub.EthProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	// We will check logs for new events
	logs := make(chan ctypes.Log)
	// Start BridgeBank subscription, prepare ethContract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := ethContract.LoadABI(txs.BridgeBank)
	eventMsgLockSignature := bridgeBankContractABI.Events[types.LogLock.String()].ID().Hex()
	eventMsgBurnSignature := bridgeBankContractABI.Events[types.LogBurn.String()].ID().Hex()

	for {
		select {
		// Handle any errors
		case err := <-subBridgeBank.Err():
			sub.Logger.Error(err.Error())
		// vLog is raw event data
		case vLog := <-logs:
			sub.Logger.Info(fmt.Sprintf("Witnessed tx %s on block %d\n", vLog.TxHash.Hex(), vLog.BlockNumber))
			var err error
			switch vLog.Topics[0].Hex() {
			case eventMsgBurnSignature:
				err = sub.handleEvrnetEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.MsgBurn.String(), vLog)
			case eventMsgLockSignature:
				err = sub.handleEvrnetEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.LogLock.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
			}
		}
	}
}

// getOracleClaimType sets the OracleClaim's claim type based upon the witnessed event type
func getOracleClaimType(eventType string) types.Event {
	var claimType types.Event
	switch eventType {
	case types.MsgBurn.String():
		claimType = types.MsgBurn
	case types.MsgLock.String():
		claimType = types.MsgLock
	default:
		claimType = types.Unsupported
	}
	return claimType
}

// Parses event data from the msg, event, builds a new ProphecyClaim, and relays it to Ethereum
func (sub EvrnetSub) handleBurnLockMsg(attributes []tmKv.Pair, claimType types.Event) error {
	cosmosMsg := txs.BurnLockEventToEvrnetMsg(claimType, attributes)
	sub.Logger.Info(cosmosMsg.String())

	// TODO: Ideally one validator should relay the prophecy and other validators make oracle claims upon that prophecy
	prophecyClaim := txs.EvrnetMsgToProphecyClaim(cosmosMsg)
	err := txs.RelayProphecyClaimToEthereum(sub.EthProvider, sub.RegistryContractAddress,
		claimType, prophecyClaim, sub.PrivateKey)
	if err != nil {
		return err
	}
	return nil
}
// startContractEventSub : starts an event subscription on the specified Peggy ethContract
func (sub EvrnetSub) startContractEventSub(logs chan ctypes.Log, client *evrclient.Client,
	contractName txs.ContractRegistry) (common.Address, evrynet.Subscription) {
	// Get the ethContract address for this subscription
	subContractAddress, err := txs.GetAddressFromBridgeRegistry(client, sub.RegistryContractAddress, contractName)
	if err != nil {
		sub.Logger.Error(err.Error())
	}

	// We need the address in []bytes for the query
	subQuery := evrynet.FilterQuery{
		Addresses: []common.Address{subContractAddress},
	}

	// Start the ethContract subscription
	contractSub, err := client.SubscribeFilterLogs(context.Background(), subQuery, logs)
	if err != nil {
		sub.Logger.Error(err.Error())
	}
	sub.Logger.Info(fmt.Sprintf("Subscribed to %v ethContract at address: %s", contractName, subContractAddress.Hex()))
	return subContractAddress, contractSub
}
// handleEthereumEvent unpacks an Ethereum event, converts it to a ProphecyClaim, and relays a tx to Evrnet
func (sub EvrnetSub) handleEvrnetEvent(clientChainID *big.Int, contractAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via ethContract ABI
	event := types.EvrnetEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeContractAddress = contractAddress
	event.EvrnetChainID = clientChainID
	if eventName == types.LogBurn.String() {
		event.ClaimType = ethbridge.BurnText
	} else {
		event.ClaimType = ethbridge.LockText
	}
	sub.Logger.Info(event.String())

	// Add the event to the record
	types.NewEventWrite(cLog.TxHash.Hex(), event)

	prophecyClaim, err := txs.EventToEthBridgeClaim(sub.ValidatorAddress, &event)
	if err != nil {
		return err
	}
	return txs.RelayToEthereum(sub.ValidatorName, &prophecyClaim)
}
