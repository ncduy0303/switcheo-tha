package keeper

import (
	"context"
	"strings"

	"crude/x/addressbook/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListContactFilter(goCtx context.Context, req *types.QueryListContactFilterRequest) (*types.QueryListContactFilterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	contacts := k.GetAllContact(ctx)
	var filteredContacts []types.Contact
	for _, contact := range contacts {
		if filterContact(&contact, req) {
			filteredContacts = append(filteredContacts, contact)
		}
	}

	return &types.QueryListContactFilterResponse{Contact: filteredContacts}, nil
}

func filterContact(contact *types.Contact, req *types.QueryListContactFilterRequest) bool {
	// Filter by name, phone, email, address (substring match, case-insensitive)
	if req.Name != "" && !strings.Contains(strings.ToLower(contact.Name), strings.ToLower(req.Name)) {
		return false
	}
	if req.Phone != "" && !strings.Contains(strings.ToLower(contact.Phone), strings.ToLower(req.Phone)) {
		return false
	}
	if req.Email != "" && !strings.Contains(strings.ToLower(contact.Email), strings.ToLower(req.Email)) {
		return false
	}
	if req.Address != "" && !strings.Contains(strings.ToLower(contact.Address), strings.ToLower(req.Address)) {
		return false
	}
	return true
}
