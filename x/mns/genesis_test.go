package mns_test

import (
	"testing"

	keepertest "github.com/cosmostation/mintstation/testutil/keeper"
	"github.com/cosmostation/mintstation/testutil/nullify"
	"github.com/cosmostation/mintstation/x/mns"
	"github.com/cosmostation/mintstation/x/mns/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		OwnershipList: []types.Ownership{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MnsKeeper(t)
	mns.InitGenesis(ctx, *k, genesisState)
	got := mns.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OwnershipList, got.OwnershipList)
	// this line is used by starport scaffolding # genesis/test/assert
}
