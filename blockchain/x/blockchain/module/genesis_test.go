package blockchain_test

import (
	"testing"

	keepertest "github.com/LiHaixin000/blockchain/testutil/keeper"
	"github.com/LiHaixin000/blockchain/testutil/nullify"
	blockchain "github.com/LiHaixin000/blockchain/x/blockchain/module"
	"github.com/LiHaixin000/blockchain/x/blockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlockchainKeeper(t)
	blockchain.InitGenesis(ctx, k, genesisState)
	got := blockchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
