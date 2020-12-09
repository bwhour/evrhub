package ethContract

// -------------------------------------------------------
//    Contract Contains functionality for loading the
//				 smart ethContract
// -------------------------------------------------------

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/txs"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// File paths to Peggy smart ethContract ABIs
const (
	BridgeBankABI   = "/generated/abi/BridgeBank/BridgeBank.abi"
	EvrnetBridgeABI = "/generated/abi/EvrnetBridge/EvrnetBridge.abi"
)

// LoadABI loads a smart ethContract as an abi.ABI
func LoadABI(contractType txs.ContractRegistry) abi.ABI {
	var (
		_, b, _, _ = runtime.Caller(0)
		dir        = filepath.Dir(b)
	)

	var filePath string
	switch contractType {
	case txs.EvrnetBridge:
		filePath = EvrnetBridgeABI
	case txs.BridgeBank:
		filePath = BridgeBankABI
	}

	// Read the file containing the ethContract's ABI
	contractRaw, err := ioutil.ReadFile(dir + filePath)
	if err != nil {
		panic(err)
	}

	// Convert the raw abi into a usable format
	contractABI, err := abi.JSON(strings.NewReader(string(contractRaw)))
	if err != nil {
		panic(err)
	}
	return contractABI
}
