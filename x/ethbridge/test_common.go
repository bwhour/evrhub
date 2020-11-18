package ethbridge

import (
	"testing"

	sdk "github.com/Evrynetlabs/evrsdk/types"

	"github.com/Evrynetlabs/evrsdk/x/auth"
	"github.com/Evrynetlabs/evrsdk/x/bank"
	"github.com/Evrynetlabs/evrsdk/x/supply"
	oracle "github.com/Evrynetlabs/evrhub/x/oracle"
	keeperLib "github.com/Evrynetlabs/evrhub/x/oracle/keeper"
)

func CreateTestHandler(
	t *testing.T, consensusNeeded float64, validatorAmounts []int64,
) (sdk.Context, oracle.Keeper, bank.Keeper, supply.Keeper, auth.AccountKeeper, []sdk.ValAddress, sdk.Handler) {
	ctx, oracleKeeper, bankKeeper, supplyKeeper,
		accountKeeper, validatorAddresses := oracle.CreateTestKeepers(t, consensusNeeded, validatorAmounts, ModuleName)
	bridgeAccount := supply.NewEmptyModuleAccount(ModuleName, supply.Burner, supply.Minter)
	supplyKeeper.SetModuleAccount(ctx, bridgeAccount)

	cdc := keeperLib.MakeTestCodec()
	bridgeKeeper := NewKeeper(cdc, supplyKeeper, oracleKeeper)
	handler := NewHandler(accountKeeper, bridgeKeeper, cdc)

	return ctx, oracleKeeper, bankKeeper, supplyKeeper, accountKeeper, validatorAddresses, handler
}
