package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmostation/mintstation/x/mns/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting a name from the store
	_, isFound := k.GetOwnership(ctx, msg.Name)
	caller, _ := sdk.AccAddressFromBech32(msg.Creator)

	_, isFoundByName := k.FindOwnership(ctx, caller.String())

	// If a name is found in store
	if isFound || isFoundByName {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Already has ownership")
	}

	// Create an updated whois record
	newOwnership := types.Ownership{
		Index: msg.Name,
		Name:  msg.Name,
		Owner: caller.String(),
	}

	// Write whois information to the store
	k.SetOwnership(ctx, newOwnership)
	return &types.MsgSetNameResponse{}, nil
}
