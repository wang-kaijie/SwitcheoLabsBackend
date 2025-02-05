package keeper

import (
	"fmt"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"crude/x/crude/types"
)

// AppendUser adds a new user to the store, assigns a unique ID, and updates the user count.
func (k Keeper) AppendUser(ctx sdk.Context, user types.User) uint64 {
	count := k.GetUserCount(ctx)
	user.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))
	appendedValue := k.cdc.MustMarshal(&user)
	store.Set(GetUserIDBytes(user.Id), appendedValue)
	k.SetUserCount(ctx, count+1)
	return count
}

// GetUserCount retrieves the total number of users stored on the blockchain.
func (k Keeper) GetUserCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.UserCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetUserIDBytes converts a uint64 user ID into a byte array for storage.
func GetUserIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// SetUserCount updates the total user count stored on the blockchain.
func (k Keeper) SetUserCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.UserCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetUser retrieves a user from the blockchain store using their ID.
func (k Keeper) GetUser(ctx sdk.Context, id uint64) (val types.User, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))
	b := store.Get(GetUserIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetUsersByAge retrieves all the users from the blockchain store that is equal to `age`.
func (k Keeper) GetUsersByAge(ctx sdk.Context, id uint64) (val types.User, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))
	b := store.Get(GetUserIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetUser update the user
func (k Keeper) SetUser(ctx sdk.Context, user types.User) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))
	fmt.Println("Updating User in Store:", user)
	b := k.cdc.MustMarshal(&user)
	store.Set(GetUserIDBytes(user.Id), b)
	return nil
}

// SetUser remove (delele) the user
func (k Keeper) RemoveUser(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))
	store.Delete(GetUserIDBytes(id))
}
