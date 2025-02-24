package blockchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/LiHaixin000/blockchain/x/blockchain/keeper"
	"github.com/LiHaixin000/blockchain/x/blockchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the resource
	for _, elem := range genState.ResourceList {
		k.SetResource(ctx, elem)
	}

	// Set resource count
	k.SetResourceCount(ctx, genState.ResourceCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ResourceList = k.GetAllResource(ctx)
	genesis.ResourceCount = k.GetResourceCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
