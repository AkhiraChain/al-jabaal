package ethbridge

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	oracle "github.com/akhirachain/al-jabaal/x/oracle"
	keeperLib "github.com/akhirachain/al-jabaal/x/oracle/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
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
