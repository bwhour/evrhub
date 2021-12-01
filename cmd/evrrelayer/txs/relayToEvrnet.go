package txs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Evrynetlabs/evrynet-node/accounts/abi/bind"
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/evrclient"

	evrnetbridge "github.com/Evrynetlabs/evrhub/cmd/evrrelayer/evrcontract/generated/bindings/evrnetbridge"
	oracle "github.com/Evrynetlabs/evrhub/cmd/evrrelayer/evrcontract/generated/bindings/oracle"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"

	xcommon "github.com/Evrynetlabs/evrhub/x/common/types"
)

const (
	// GasLimit the gas limit in Gwei used for transactions sent with TransactOpts
	GasLimit      = uint64(3000000)
	DefaultPrefix = "PEGGY"
)

// RelayProphecyClaimToEvrnet relays the provided ProphecyClaim to EvrnetBridge evrcontract on the Evrnet network
func RelayProphecyClaimToEvrnet(provider string, contractAddress common.Address, event types.Event,
	claim xcommon.EthProphecyClaim, key *ecdsa.PrivateKey) error {
	// Initialize client service, validator's tx auth, and target evrcontract address
	client, auth, target := initRelayConfig(provider, contractAddress, event, key)

	// Initialize EvrnetBridge instance
	fmt.Println("\nFetching EvrnetBridge evrcontract...")
	evrnetBridgeInstance, err := evrnetbridge.NewEvrnetBridge(target, client)
	if err != nil {
		return err
	}

	// Send transaction
	fmt.Println("Sending new ProphecyClaim to EvrnetBridge...")
	if event == types.MsgBurn {
		if !strings.Contains(claim.Symbol, DefaultPrefix) {
			log.Fatal("Can only relay burns of 'PEGGY' prefixed coins")
		}
		strs := strings.SplitAfter(claim.Symbol, DefaultPrefix)
		claim.Symbol = strings.ToUpper(strings.Join(strs[1:], ""))
	} else {
		claim.Symbol = strings.ToUpper(claim.Symbol)
	}

	sender := common.HexToAddress(claim.EthereumSender.String())
	receiver := common.HexToAddress(claim.EvrnetReceiver.String())
	amount := big.NewInt(0)
	amount.SetString(claim.Amount, 10)

	tx, err := evrnetBridgeInstance.NewProphecyClaim(auth, uint8(event),
		sender, receiver, claim.Symbol, amount)
	if err != nil {
		return err
	}
	fmt.Println("NewProphecyClaim tx hash:", tx.Hash().Hex())

	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return err
	}

	switch receipt.Status {
	case 0:
		fmt.Println("Tx Status: 0 - Failed")
	case 1:
		fmt.Println("Tx Status: 1 - Successful")
	}
	return nil
}

// RelayOracleClaimToEvrnet relays the provided OracleClaim to Oracle evrcontract on the Evrnet network
func RelayOracleClaimToEvrnet(provider string, contractAddress common.Address, event types.Event,
	claim OracleClaim, key *ecdsa.PrivateKey) error {
	// Initialize client service, validator's tx auth, and target evrcontract address
	client, auth, target := initRelayConfig(provider, contractAddress, event, key)

	// Initialize Oracle instance
	fmt.Println("\nFetching Oracle evrcontract...")
	oracleInstance, err := oracle.NewOracle(target, client)
	if err != nil {
		return err
	}

	// Send transaction
	fmt.Println("Sending new OracleClaim to Oracle...")
	tx, err := oracleInstance.NewOracleClaim(auth, claim.ProphecyID, claim.Message, claim.Signature)
	if err != nil {
		return err
	}
	fmt.Println("NewOracleClaim tx hash:", tx.Hash().Hex())

	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return err
	}

	switch receipt.Status {
	case 0:
		fmt.Println("Tx Status: 0 - Failed")
	case 1:
		fmt.Println("Tx Status: 1 - Successful")
	}

	return nil
}

// initRelayConfig set up Evrnet client, validator's transaction auth, and the target evrcontract's address
func initRelayConfig(provider string, registry common.Address, event types.Event, key *ecdsa.PrivateKey,
) (*evrclient.Client, *bind.TransactOpts, common.Address) {
	// Start Evrnet client
	client, err := evrclient.Dial(provider)
	if err != nil {
		log.Fatal(err)
	}

	// Load the validator's address
	sender, err := LoadEvrSender()
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Set up TransactOpts auth's tx signature authorization
	transactOptsAuth := bind.NewKeyedTransactor(key)
	transactOptsAuth.Nonce = big.NewInt(int64(nonce))
	transactOptsAuth.Value = big.NewInt(0) // in wei
	transactOptsAuth.GasLimit = GasLimit
	transactOptsAuth.GasPrice = gasPrice

	var targetContract ContractRegistry
	switch event {
	// ProphecyClaims are sent to the EvrnetBridge evrcontract
	case types.MsgBurn, types.MsgLock:
		targetContract = EvrnetBridge
	// OracleClaims are sent to the Oracle evrcontract
	case types.LogNewProphecyClaim:
		targetContract = Oracle
	default:
		panic("invalid target evrcontract address")
	}

	// Get the specific evrcontract's address
	target, err := GetAddressFromBridgeRegistry(client, registry, targetContract)
	if err != nil {
		log.Fatal(err)
	}
	return client, transactOptsAuth, target
}
