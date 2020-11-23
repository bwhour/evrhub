package relayer

import (

	"context"
	"crypto/ecdsa"
	ctypes "github.com/ethereum/go-ethereum/core/types"
	"os"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	evrtypes "github.com/Evrynetlabs/evrynet-node/core/types"
	tmKv "github.com/tendermint/tendermint/libs/kv"
	tmLog "github.com/tendermint/tendermint/libs/log"
       "github.com/Evrynetlabs/evrhub/cmd/evrrelayer/txs"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"
)

// TODO: Move relay functionality out of EvrnetSub into a new Relayer parent struct

// EvrnetSub defines a Cosmos listener that relays events to Ethereum and Evrnet
type EvrnetSub struct {
	EvrProvider           string
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

// Start a Cosmos chain subscription
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
	for {
		select {
		// vLog is raw event data
		case vLog := <-logs:
			sub.Logger.Info(fmt.Sprintf("Witnessed tx %s on block %d\n", vLog.TxHash.Hex(), vLog.BlockNumber))
			var err error
			switch vLog.Topics[0].Hex() {
			//case eventLogBurnSignature:
			//	err = sub.handleEthereumEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
			//		types.LogBurn.String(), vLog)
			//case eventLogLockSignature:
			//	err = sub.handleEthereumEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
			//		types.LogLock.String(), vLog)
			//case eventLogNewProphecyClaimSignature:
			//	err = sub.handleLogNewProphecyClaim(cosmosBridgeAddress, cosmosBridgeContractABI,
			//		types.LogNewProphecyClaim.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
			}
		}
	}
			// Iterate over each event in the transaction
			//for _, event := range tx.Result.Events {
			//	claimType := getOracleClaimType(event.GetType())
			//
			//	switch claimType {
			//	case types.MsgBurn, types.MsgLock:
			//		// Parse event data, then package it as a ProphecyClaim and relay to the Ethereum Network
			//		err := sub.handleBurnLockMsg(event.GetAttributes(), claimType)
			//		if err != nil {
			//			sub.Logger.Error(err.Error())
			//		}
			//	}
			//}

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
	cosmosMsg := txs.BurnLockEventToCosmosMsg(claimType, attributes)
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
