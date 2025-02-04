package userblock_test

import (
	"testing"

	keepertest "userBlock/testutil/keeper"
	"userBlock/testutil/nullify"
	userblock "userBlock/x/userblock/module"
	"userBlock/x/userblock/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UserblockKeeper(t)
	userblock.InitGenesis(ctx, k, genesisState)
	got := userblock.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
