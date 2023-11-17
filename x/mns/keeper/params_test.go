package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmostation/mintstation/testutil/keeper"
	"github.com/cosmostation/mintstation/x/mns/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MnsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
