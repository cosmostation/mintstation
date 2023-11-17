package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmostation/mintstation/x/mns/types"
)

// SetOwnership set a specific ownership in the store from its index
func (k Keeper) SetOwnership(ctx sdk.Context, ownership types.Ownership) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnershipKeyPrefix))
	b := k.cdc.MustMarshal(&ownership)
	store.Set(types.OwnershipKey(
		ownership.Index,
	), b)
}

// GetOwnership returns a ownership from its index
func (k Keeper) GetOwnership(
	ctx sdk.Context,
	index string,

) (val types.Ownership, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnershipKeyPrefix))

	b := store.Get(types.OwnershipKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// FindOwnership returns a ownership by owner
func (k Keeper) FindOwnership(
	ctx sdk.Context,
	address string,
) (val types.Ownership, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnershipKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ownership
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Owner == address {
			return val, true
		}
	}

	return val, false
}

// RemoveOwnership removes a ownership from the store
func (k Keeper) RemoveOwnership(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnershipKeyPrefix))
	store.Delete(types.OwnershipKey(
		index,
	))
}

// GetAllOwnership returns all ownership
func (k Keeper) GetAllOwnership(ctx sdk.Context) (list []types.Ownership) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnershipKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ownership
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
