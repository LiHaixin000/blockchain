package blockchain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/LiHaixin000/blockchain/testutil/sample"
	blockchainsimulation "github.com/LiHaixin000/blockchain/x/blockchain/simulation"
	"github.com/LiHaixin000/blockchain/x/blockchain/types"
)

// avoid unused import issue
var (
	_ = blockchainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateResource = "op_weight_msg_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateResource int = 100

	opWeightMsgUpdateResource = "op_weight_msg_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateResource int = 100

	opWeightMsgDeleteResource = "op_weight_msg_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteResource int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blockchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ResourceList: []types.Resource{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ResourceCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blockchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateResource int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateResource, &weightMsgCreateResource, nil,
		func(_ *rand.Rand) {
			weightMsgCreateResource = defaultWeightMsgCreateResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateResource,
		blockchainsimulation.SimulateMsgCreateResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateResource int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateResource, &weightMsgUpdateResource, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateResource = defaultWeightMsgUpdateResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateResource,
		blockchainsimulation.SimulateMsgUpdateResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteResource int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteResource, &weightMsgDeleteResource, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteResource = defaultWeightMsgDeleteResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteResource,
		blockchainsimulation.SimulateMsgDeleteResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateResource,
			defaultWeightMsgCreateResource,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blockchainsimulation.SimulateMsgCreateResource(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateResource,
			defaultWeightMsgUpdateResource,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blockchainsimulation.SimulateMsgUpdateResource(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteResource,
			defaultWeightMsgDeleteResource,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blockchainsimulation.SimulateMsgDeleteResource(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
