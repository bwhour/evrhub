package types

// Ethbridge module event types
var (
	EventTypeCreateClaim    = "create_claim"
	EventTypeProphecyStatus = "prophecy_status"
	EventTypeBurn           = "burn"
	EventTypeLock           = "lock"

	AttributeKeyEthereumSender = "ethereum_sender"
	AttributeKeyEvrnetReceiver = "evrnet_Receiver"
	AttributeKeyAmount         = "amount"
	AttributeKeySymbol         = "symbol"
	AttributeKeyCoins          = "coins"
	AttributeKeyStatus         = "status"
	AttributeKeyClaimType      = "claim_type"

	AttributeKeyEthereumChainID  = "ethereum_chain_id"
	AttributeKeyTokenContract    = "token_contract_address"
	AttributeKeyCosmosSender     = "evrnet_sender"
	AttributeKeyEthereumReceiver = "ethereum_receiver"

	AttributeValueCategory = ModuleName
)
