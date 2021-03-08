package evrcontract

// BridgeContract is an enum containing supported evrcontract names
type BridgeContract int

const (
	// BridgeRegistry registers deployed addresses of the other contracts
	BridgeRegistry BridgeContract = iota + 1
	// Valset manages the validator set and signature verification
	Valset
	// Oracle enables validators to make OracleClaims and processes ProphecyClaims
	Oracle
	// EvrnetBridge enables validators to make ProphecyClaims
	EvrnetBridge
	// BridgeBank manages protocol assets on both Ethereum and Evrnet
	BridgeBank
)

// BridgeContractToString returns the string associated with a BridgeContract
var BridgeContractToString = [...]string{"BridgeRegistry", "Valset", "Oracle", "EvrnetBridge", "BridgeBank"}

// String returns the BridgeContract as a string
func (d BridgeContract) String() string {
	return BridgeContractToString[d-1]
}

// BridgeContracts is a slice of BridgeContract
type BridgeContracts []BridgeContract

// LoadBridgeContracts loads populated slice of BridgeContract
func LoadBridgeContracts() BridgeContracts {
	return BridgeContracts{
		BridgeRegistry,
		Valset,
		Oracle,
		EvrnetBridge,
		BridgeBank,
	}
}
