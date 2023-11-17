package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmostation/mintstation/x/mns/types"
)

func (k msgServer) DeleteName(goCtx context.Context, msg *types.MsgDeleteName) (*types.MsgDeleteNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting name information from the store
	ownership, isFound := k.GetOwnership(ctx, msg.Name)

	// If a name is not found, throw an error
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Name doesn't exist")
	}

	// If the message sender address doesn't match the name owner, throw an error
	if !(ownership.Owner == msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	// Otherwise, remove the name information from the store
	k.RemoveOwnership(ctx, msg.Name)
	return &types.MsgDeleteNameResponse{}, nil
}
