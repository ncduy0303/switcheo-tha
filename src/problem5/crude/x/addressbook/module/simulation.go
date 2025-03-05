package addressbook

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"crude/testutil/sample"
	addressbooksimulation "crude/x/addressbook/simulation"
	"crude/x/addressbook/types"
)

// avoid unused import issue
var (
	_ = addressbooksimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateContact = "op_weight_msg_contact"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContact int = 100

	opWeightMsgUpdateContact = "op_weight_msg_contact"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateContact int = 100

	opWeightMsgDeleteContact = "op_weight_msg_contact"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteContact int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	addressbookGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ContactList: []types.Contact{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ContactCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&addressbookGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateContact int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateContact, &weightMsgCreateContact, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContact = defaultWeightMsgCreateContact
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContact,
		addressbooksimulation.SimulateMsgCreateContact(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateContact int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateContact, &weightMsgUpdateContact, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateContact = defaultWeightMsgUpdateContact
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateContact,
		addressbooksimulation.SimulateMsgUpdateContact(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteContact int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteContact, &weightMsgDeleteContact, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteContact = defaultWeightMsgDeleteContact
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteContact,
		addressbooksimulation.SimulateMsgDeleteContact(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateContact,
			defaultWeightMsgCreateContact,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				addressbooksimulation.SimulateMsgCreateContact(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateContact,
			defaultWeightMsgUpdateContact,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				addressbooksimulation.SimulateMsgUpdateContact(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteContact,
			defaultWeightMsgDeleteContact,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				addressbooksimulation.SimulateMsgDeleteContact(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
