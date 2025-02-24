package resource_test

import (
	"testing"

	keepertest "github.com/LiHaixin000/blockchain/testutil/keeper"
	"github.com/LiHaixin000/blockchain/testutil/nullify"
	resource "github.com/LiHaixin000/blockchain/x/resource/module"
	"github.com/LiHaixin000/blockchain/x/resource/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ResourceKeeper(t)
	resource.InitGenesis(ctx, k, genesisState)
	got := resource.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
