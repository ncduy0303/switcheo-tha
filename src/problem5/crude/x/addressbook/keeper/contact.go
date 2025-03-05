package keeper

import (
	"context"
	"encoding/binary"

	"crude/x/addressbook/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetContactCount get the total number of contact
func (k Keeper) GetContactCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ContactCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetContactCount set the total number of contact
func (k Keeper) SetContactCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ContactCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendContact appends a contact in the store with a new id and update the count
func (k Keeper) AppendContact(
	ctx context.Context,
	contact types.Contact,
) uint64 {
	// Create the contact
	count := k.GetContactCount(ctx)

	// Set the ID of the appended value
	contact.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	appendedValue := k.cdc.MustMarshal(&contact)
	store.Set(GetContactIDBytes(contact.Id), appendedValue)

	// Update contact count
	k.SetContactCount(ctx, count+1)

	return count
}

// SetContact set a specific contact in the store
func (k Keeper) SetContact(ctx context.Context, contact types.Contact) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	b := k.cdc.MustMarshal(&contact)
	store.Set(GetContactIDBytes(contact.Id), b)
}

// GetContact returns a contact from its id
func (k Keeper) GetContact(ctx context.Context, id uint64) (val types.Contact, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	b := store.Get(GetContactIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContact removes a contact from the store
func (k Keeper) RemoveContact(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	store.Delete(GetContactIDBytes(id))
}

// GetAllContact returns all contact
func (k Keeper) GetAllContact(ctx context.Context) (list []types.Contact) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Contact
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetContactIDBytes returns the byte representation of the ID
func GetContactIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ContactKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
