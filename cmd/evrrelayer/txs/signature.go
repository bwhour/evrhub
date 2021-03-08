package txs

import (
	"crypto/ecdsa"
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/crypto"
	"github.com/joho/godotenv"

	"log"
	"os"
	"strings"

	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"
)

// LoadEvrPrivateKey loads the validator's private key from environment variables
func LoadEvrPrivateKey() (key *ecdsa.PrivateKey, err error) {
	// Load config file containing environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Private key for validator's Ethereum address must be set as an environment variable
	rawPrivateKey := os.Getenv("EVRNET_PRIVATE_KEY")
	if strings.TrimSpace(rawPrivateKey) == "" {
		log.Fatal("Error loading EVRNET_PRIVATE_KEY from .env file")
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(rawPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	return privateKey, nil
}

// LoadEvrSender uses the validator's private key to load the validator's address
func LoadEvrSender() (address common.Address, err error) {
	key, err := LoadEvrPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Bytes()
	outAddress := common.BytesToAddress(fromAddress)
	return outAddress, nil
}

// GenerateClaimMessage Generates a hashed message containing a ProphecyClaim event's data
func GenerateClaimMessage(event types.ProphecyClaimEvent) []byte {
	prophecyID := Int256(event.ProphecyID)
	sender := String(event.EthereumSender)
	recipient := Int256(event.EvrnetReceiver.Hex())
	token := String(event.TokenAddress.Hex())
	amount := Int256(event.Amount)

	// Generate claim message using ProphecyClaim data
	return SoliditySHA3(prophecyID, sender, recipient, token, amount)
}

// PrefixMsg prefixes a message for verification, mimics behavior of web3.eth.sign
func PrefixMsg(msg []byte) []byte {
	return SoliditySHA3(String("\x19Evrnet Signed Message:\n32"), msg)
}

// SignClaim Signs the prepared message with validator's private key
func SignClaim(msg []byte, key *ecdsa.PrivateKey) ([]byte, error) {
	// Sign the message
	sig, err := crypto.Sign(msg, key)
	if err != nil {
		panic(err)
	}
	return sig, nil
}
