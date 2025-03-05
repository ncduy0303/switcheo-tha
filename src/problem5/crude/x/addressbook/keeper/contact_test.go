package keeper_test

import (
	"context"
	"testing"

	keepertest "crude/testutil/keeper"
	"crude/testutil/nullify"
	"crude/x/addressbook/keeper"
	"crude/x/addressbook/types"

	"github.com/stretchr/testify/require"
)

func createNContact(keeper keeper.Keeper, ctx context.Context, n int) []types.Contact {
	items := make([]types.Contact, n)
	for i := range items {
		items[i].Id = keeper.AppendContact(ctx, items[i])
	}
	return items
}

func TestContactGet(t *testing.T) {
	keeper, ctx := keepertest.AddressbookKeeper(t)
	items := createNContact(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetContact(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestContactRemove(t *testing.T) {
	keeper, ctx := keepertest.AddressbookKeeper(t)
	items := createNContact(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveContact(ctx, item.Id)
		_, found := keeper.GetContact(ctx, item.Id)
		require.False(t, found)
	}
}

func TestContactGetAll(t *testing.T) {
	keeper, ctx := keepertest.AddressbookKeeper(t)
	items := createNContact(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllContact(ctx)),
	)
}

func TestContactCount(t *testing.T) {
	keeper, ctx := keepertest.AddressbookKeeper(t)
	items := createNContact(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetContactCount(ctx))
}
