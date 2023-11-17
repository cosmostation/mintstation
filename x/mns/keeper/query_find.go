package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmostation/mintstation/x/mns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Find(goCtx context.Context, req *types.QueryFindRequest) (*types.QueryFindResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.FindOwnership(
		ctx,
		req.Owner,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFindResponse{Ownership: val}, nil
}
