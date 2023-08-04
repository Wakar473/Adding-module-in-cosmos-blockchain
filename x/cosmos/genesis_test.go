package cosmos_test

import (
	"testing"

	keepertest "cosmos/testutil/keeper"
	"cosmos/testutil/nullify"
	"cosmos/x/cosmos"
	"cosmos/x/cosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosKeeper(t)
	cosmos.InitGenesis(ctx, *k, genesisState)
	got := cosmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
