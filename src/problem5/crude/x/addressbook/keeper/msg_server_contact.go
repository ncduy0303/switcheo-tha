package keeper

import (
	"context"
	"fmt"

	"crude/x/addressbook/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateContact(goCtx context.Context, msg *types.MsgCreateContact) (*types.MsgCreateContactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var contact = types.Contact{
		Creator: msg.Creator,
		Name:    msg.Name,
		Phone:   msg.Phone,
		Email:   msg.Email,
		Address: msg.Address,
	}

	id := k.AppendContact(
		ctx,
		contact,
	)

	return &types.MsgCreateContactResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateContact(goCtx context.Context, msg *types.MsgUpdateContact) (*types.MsgUpdateContactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var contact = types.Contact{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Phone:   msg.Phone,
		Email:   msg.Email,
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetContact(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetContact(ctx, contact)

	return &types.MsgUpdateContactResponse{}, nil
}

func (k msgServer) DeleteContact(goCtx context.Context, msg *types.MsgDeleteContact) (*types.MsgDeleteContactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetContact(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveContact(ctx, msg.Id)

	return &types.MsgDeleteContactResponse{}, nil
}
