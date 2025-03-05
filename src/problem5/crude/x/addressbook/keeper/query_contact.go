package keeper

import (
	"context"

	"crude/x/addressbook/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContactAll(ctx context.Context, req *types.QueryAllContactRequest) (*types.QueryAllContactResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contacts []types.Contact

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	contactStore := prefix.NewStore(store, types.KeyPrefix(types.ContactKey))

	pageRes, err := query.Paginate(contactStore, req.Pagination, func(key []byte, value []byte) error {
		var contact types.Contact
		if err := k.cdc.Unmarshal(value, &contact); err != nil {
			return err
		}

		contacts = append(contacts, contact)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContactResponse{Contact: contacts, Pagination: pageRes}, nil
}

func (k Keeper) Contact(ctx context.Context, req *types.QueryGetContactRequest) (*types.QueryGetContactResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	contact, found := k.GetContact(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetContactResponse{Contact: contact}, nil
}
