package txs

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	sdk "github.com/Evrynetlabs/evrsdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tmKv "github.com/tendermint/tendermint/libs/kv"

	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"
	ethbridge "github.com/Evrynetlabs/evrhub/x/ethbridge/types"
)

const (
	nullAddress   = "0x0000000000000000000000000000000000000000"
	defaultPrefix = "peggy"
)

// EthereumEventToEthBridgeClaim parses and packages an Ethereum event struct with a validator address in an EthBridgeClaim msg
func EthereumEventToEthBridgeClaim(valAddr sdk.ValAddress, event *types.EthereumEvent) (ethbridge.EthBridgeClaim, error) {
	witnessClaim := ethbridge.EthBridgeClaim{}

	// chainID type casting (*big.Int -> int)
	chainID := int(event.EthereumChainID.Int64())

	bridgeContractAddress := ethbridge.NewEthereumAddress(event.BridgeContractAddress.Hex())

	// Sender type casting (address.common -> string)
	sender := ethbridge.NewEthereumAddress(event.From.Hex())

	// Recipient type casting ([]bytes -> sdk.AccAddress)
	recipient, err := sdk.AccAddressFromBech32(string(event.To))
	if err != nil {
		return witnessClaim, err
	}
	if recipient.Empty() {
		return witnessClaim, errors.New("empty recipient address")
	}

	// Sender type casting (address.common -> string)
	tokenContractAddress := ethbridge.NewEthereumAddress(event.Token.Hex())

	// Symbol formatted to lowercase
	symbol := strings.ToLower(event.Symbol)
	switch event.ClaimType {
	case ethbridge.LockText:
		if symbol == "eth" && !isZeroAddress(event.Token) {
			return witnessClaim, errors.New("symbol \"eth\" must have null address set as token address")
		}
	case ethbridge.BurnText:
		if !strings.Contains(symbol, defaultPrefix) {
			log.Fatal("Can only relay burns of 'PEGGY' prefixed tokens")
		}
		res := strings.SplitAfter(symbol, defaultPrefix)
		symbol = strings.Join(res[1:], "")
	}

	amount := event.Value.Int64()

	// Nonce type casting (*big.Int -> int)
	nonce := int(event.Nonce.Int64())

	// Package the information in a unique EthBridgeClaim
	witnessClaim.EthereumChainID = chainID
	witnessClaim.BridgeContractAddress = bridgeContractAddress
	witnessClaim.Nonce = nonce
	witnessClaim.TokenContractAddress = tokenContractAddress
	witnessClaim.Symbol = symbol
	witnessClaim.EthereumSender = sender
	witnessClaim.ValidatorAddress = valAddr
	witnessClaim.EvrnetReceiver = recipient
	witnessClaim.Amount = amount
	witnessClaim.ClaimType = event.ClaimType

	return witnessClaim, nil
}

// ProphecyClaimToSignedOracleClaim packages and signs a prophecy claim's data, returning a new oracle claim
func ProphecyClaimToSignedOracleClaim(event types.ProphecyClaimEvent, key *ecdsa.PrivateKey) (OracleClaim, error) {
	oracleClaim := OracleClaim{}

	// Generate a hashed claim message which contains ProphecyClaim's data
	fmt.Println("Generating unique message for ProphecyClaim", event.ProphecyID)
	message := GenerateClaimMessage(event)

	// Sign the message using the validator's private key
	fmt.Println("Signing message...")
	signature, err := SignClaim(PrefixMsg(message), key)
	if err != nil {
		return oracleClaim, err
	}
	fmt.Println("Signature generated:", hexutil.Encode(signature))

	oracleClaim.ProphecyID = event.ProphecyID
	var message32 [32]byte
	copy(message32[:], message)
	oracleClaim.Message = message32
	oracleClaim.Signature = signature
	return oracleClaim, nil
}

// EvrnetMsgToProphecyClaim parses event data from a EvrnetMsg, packaging it as a ProphecyClaim
func EvrnetMsgToProphecyClaim(event types.EvrnetMsg) ProphecyClaim {
	claimType := event.ClaimType
	evrnetSender := event.EvrnetSender
	ethereumReceiver := event.EthereumReceiver
	symbol := strings.ToUpper(event.Symbol)
	amount := event.Amount

	prophecyClaim := ProphecyClaim{
		ClaimType:        claimType,
		EvrnetSender:     evrnetSender,
		EthereumReceiver: ethereumReceiver,
		Symbol:           symbol,
		Amount:           amount,
	}
	return prophecyClaim
}

// BurnLockEventToEvrnetMsg parses data from a Burn/Lock event witnessed on Evrnet into a EvrnetMsg struct
func BurnLockEventToEvrnetMsg(claimType types.Event, attributes []tmKv.Pair) types.EvrnetMsg {
	var evrnetSender []byte
	var ethereumReceiver common.Address
	var symbol string
	var amount *big.Int

	for _, attribute := range attributes {
		key := string(attribute.GetKey())
		val := string(attribute.GetValue())

		// Set variable based on the attribute's key
		switch key {
		case types.EvrnetSender.String():
			evrnetSender = []byte(val)
		case types.EthereumReceiver.String():
			if !common.IsHexAddress(val) {
				log.Fatal("Invalid recipient address:", val)
			}
			ethereumReceiver = common.HexToAddress(val)
		case types.Symbol.String():
			if claimType == types.MsgBurn {
				if !strings.Contains(val, defaultPrefix) {
					log.Fatal("Can only relay burns of 'peggy' prefixed coins")
				}
				res := strings.SplitAfter(val, defaultPrefix)
				symbol = strings.ToUpper(strings.Join(res[1:], ""))
			} else {
				symbol = strings.ToUpper(val)
			}
		case types.Amount.String():
			tempAmount := new(big.Int)
			tempAmount, ok := tempAmount.SetString(val, 10)
			if !ok {
				log.Fatal("Invalid amount:", val)
			}
			amount = tempAmount
		}
	}
	return types.NewEvrnetMsg(claimType, evrnetSender, ethereumReceiver, symbol, amount)
}

// isZeroAddress checks an Ethereum address and returns a bool which indicates if it is the null address
func isZeroAddress(address common.Address) bool {
	return address == common.HexToAddress(nullAddress)
}
