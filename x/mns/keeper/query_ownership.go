package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmostation/mintstation/x/mns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OwnershipAll(goCtx context.Context, req *types.QueryAllOwnershipRequest) (*types.QueryAllOwnershipResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ownerships []types.Ownership
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ownershipStore := prefix.NewStore(store, types.KeyPrefix(types.OwnershipKeyPrefix))

	pageRes, err := query.Paginate(ownershipStore, req.Pagination, func(key []byte, value []byte) error {
		var ownership types.Ownership
		if err := k.cdc.Unmarshal(value, &ownership); err != nil {
			return err
		}

		ownerships = append(ownerships, ownership)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOwnershipResponse{Ownership: ownerships, Pagination: pageRes}, nil
}

func (k Keeper) Ownership(goCtx context.Context, req *types.QueryGetOwnershipRequest) (*types.QueryGetOwnershipResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetOwnership(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOwnershipResponse{Ownership: val}, nil
}
