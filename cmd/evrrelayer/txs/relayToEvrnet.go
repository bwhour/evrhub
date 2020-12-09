package txs

import (
	"github.com/Evrynetlabs/evrhub/x/ethbridge"
	"github.com/Evrynetlabs/evrhub/x/ethbridge/types"
)

// RelayToEvrnet applies validator's signature to an EthBridgeClaim message containing
// information about an event on the Ethereum blockchain before relaying to the Bridge
func RelayToEvrnet(moniker string, claim *types.EthBridgeClaim ) error {
	// Packages the claim as a Tendermint message
	msg := ethbridge.NewMsgCreateEthBridgeClaim(*claim)

	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

    //TODO send to evrnet

	return nil
}
