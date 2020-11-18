package txs

import (
	"github.com/Evrynetlabs/evrsdk/client/context"
	"github.com/Evrynetlabs/evrsdk/client/keys"
	"github.com/Evrynetlabs/evrsdk/codec"
	sdk "github.com/Evrynetlabs/evrsdk/types"
	"github.com/Evrynetlabs/evrsdk/x/auth/client/utils"
	authtypes "github.com/Evrynetlabs/evrsdk/x/auth/types"

	"github.com/Evrynetlabs/evrhub/x/ethbridge"
	"github.com/Evrynetlabs/evrhub/x/ethbridge/types"
)

// RelayToCosmos applies validator's signature to an EthBridgeClaim message containing
// information about an event on the Ethereum blockchain before relaying to the Bridge
func RelayToCosmos(cdc *codec.Codec, moniker string, claim *types.EthBridgeClaim, cliCtx context.CLIContext,
	txBldr authtypes.TxBuilder) error {
	// Packages the claim as a Tendermint message
	msg := ethbridge.NewMsgCreateEthBridgeClaim(*claim)

	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	// Prepare tx
	txBldr, err = utils.PrepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return err
	}

	// Build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(moniker, keys.DefaultKeyPass, []sdk.Msg{msg})
	if err != nil {
		return err
	}

	// Broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTxSync(txBytes)
	if err != nil {
		return err
	}

	if err = cliCtx.PrintOutput(res); err != nil {
		return err
	}
	return nil
}
