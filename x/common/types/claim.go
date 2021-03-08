package types

// EthProphecyClaim contains data required to make an ProphecyClaim
type EthProphecyClaim struct {
	ClaimType      ClaimType       `json:"claim_type"`
	EthereumSender EthereumAddress `json:"ethereum_sender"`
	EvrnetReceiver EvrnetAddress   `json:"evrnet_receiver"`
	Symbol         string          `json:"symbol"`
	Amount         string          `json:"amount"`
}

// EvrProphecyClaim contains data required to make an ProphecyClaim
type EvrProphecyClaim struct {
	ClaimType        ClaimType       `json:"claim_type"`
	EvrnetSender     EvrnetAddress   `json:"evrnet_sender"`
	EthereumReceiver EthereumAddress `json:"ethereum_receiver"`
	Symbol           string          `json:"symbol"`
	Amount           string          `json:"amount"`
}
