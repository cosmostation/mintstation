package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/cosmostation/mintstation/testutil/keeper"
	"github.com/cosmostation/mintstation/testutil/nullify"
	"github.com/cosmostation/mintstation/x/mns/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOwnershipQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MnsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOwnership(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetOwnershipRequest
		response *types.QueryGetOwnershipResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetOwnershipRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetOwnershipResponse{Ownership: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetOwnershipRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetOwnershipResponse{Ownership: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetOwnershipRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Ownership(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestOwnershipQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MnsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOwnership(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOwnershipRequest {
		return &types.QueryAllOwnershipRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OwnershipAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Ownership), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Ownership),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OwnershipAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Ownership), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Ownership),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.OwnershipAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Ownership),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.OwnershipAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
