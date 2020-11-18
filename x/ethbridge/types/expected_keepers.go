package types

import (
	sdk "github.com/Evrynetlabs/evrsdk/types"
	authexported "github.com/Evrynetlabs/evrsdk/x/auth/exported"
	supplyexported "github.com/Evrynetlabs/evrsdk/x/supply/exported"

	"github.com/Evrynetlabs/evrhub/x/oracle"
)

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(sdk.Context, sdk.AccAddress) authexported.Account
}

// SupplyKeeper defines the expected supply keeper
type SupplyKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	SetModuleAccount(sdk.Context, supplyexported.ModuleAccountI)
}

// OracleKeeper defines the expected oracle keeper
type OracleKeeper interface {
	ProcessClaim(ctx sdk.Context, claim oracle.Claim) (oracle.Status, error)
	GetProphecy(ctx sdk.Context, id string) (oracle.Prophecy, bool)
}
