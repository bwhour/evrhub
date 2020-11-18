package types // noalias

import (
	sdk "github.com/Evrynetlabs/evrsdk/types"
	"github.com/Evrynetlabs/evrsdk/x/staking"
)

// StakingKeeper defines the expected staking keeper
type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator staking.Validator, found bool)
	GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64)
	GetLastTotalPower(ctx sdk.Context) (power sdk.Int)
	GetBondedValidatorsByPower(ctx sdk.Context) []staking.Validator
}
