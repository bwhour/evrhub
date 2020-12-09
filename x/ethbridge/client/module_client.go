package client

import (
	"github.com/Evrynetlabs/evrhub/x/ethbridge/client/cli"
	"github.com/Evrynetlabs/evrsdk/codec"
	"github.com/spf13/cobra"

	"github.com/Evrynetlabs/evrsdk/client/flags"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	// Group ethbridge queries under a subcommand
	ethBridgeQueryCmd := &cobra.Command{
		Use:   "ethbridge",
		Short: "Querying commands for the ethbridge module",
	}

	ethBridgeQueryCmd.AddCommand(flags.GetCommands(
		cli.GetCmdGetEthBridgeProphecy(storeKey, cdc),
	)...)

	return ethBridgeQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	ethBridgeTxCmd := &cobra.Command{
		Use:   "ethbridge",
		Short: "EthBridge transactions subcommands",
	}

	ethBridgeTxCmd.AddCommand(flags.PostCommands(
		cli.GetCmdCreateEthBridgeClaim(cdc),
		cli.GetCmdBurn(cdc),
		cli.GetCmdLock(cdc),
	)...)

	return ethBridgeTxCmd
}

