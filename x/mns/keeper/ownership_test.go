package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cosmostation/mintstation/testutil/keeper"
	"github.com/cosmostation/mintstation/testutil/nullify"
	"github.com/cosmostation/mintstation/x/mns/keeper"
	"github.com/cosmostation/mintstation/x/mns/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNOwnership(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ownership {
	items := make([]types.Ownership, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetOwnership(ctx, items[i])
	}
	return items
}

func TestOwnershipGet(t *testing.T) {
	keeper, ctx := keepertest.MnsKeeper(t)
	items := createNOwnership(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOwnership(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOwnershipRemove(t *testing.T) {
	keeper, ctx := keepertest.MnsKeeper(t)
	items := createNOwnership(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOwnership(ctx,
			item.Index,
		)
		_, found := keeper.GetOwnership(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestOwnershipGetAll(t *testing.T) {
	keeper, ctx := keepertest.MnsKeeper(t)
	items := createNOwnership(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOwnership(ctx)),
	)
}
