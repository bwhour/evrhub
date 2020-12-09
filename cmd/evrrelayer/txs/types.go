package txs

import (
	"math/big"

	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"
	"github.com/ethereum/go-ethereum/common"
	evrcommon "github.com/Evrynetlabs/evrynet-node/common"
)

// OracleClaim contains data required to make an OracleClaim
type OracleClaim struct {
	ProphecyID *big.Int
	Message    [32]byte
	Signature  []byte
}

// ProphecyClaim contains data required to make an ProphecyClaim
type ProphecyClaim struct {
	ClaimType            types.Event
	EvrnetSender         evrcommon.Address
	EthereumReceiver     common.Address
	Symbol               string
	Amount               *big.Int
}
