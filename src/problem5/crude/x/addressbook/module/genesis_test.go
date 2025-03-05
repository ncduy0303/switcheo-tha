package addressbook_test

import (
	"testing"

	keepertest "crude/testutil/keeper"
	"crude/testutil/nullify"
	addressbook "crude/x/addressbook/module"
	"crude/x/addressbook/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ContactList: []types.Contact{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ContactCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AddressbookKeeper(t)
	addressbook.InitGenesis(ctx, k, genesisState)
	got := addressbook.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ContactList, got.ContactList)
	require.Equal(t, genesisState.ContactCount, got.ContactCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
