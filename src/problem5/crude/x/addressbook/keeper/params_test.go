package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "crude/testutil/keeper"
	"crude/x/addressbook/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AddressbookKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
