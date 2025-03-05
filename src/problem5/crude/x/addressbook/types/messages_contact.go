package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateContact{}

func NewMsgCreateContact(creator string, name string, phone string, email string, address string) *MsgCreateContact {
	return &MsgCreateContact{
		Creator: creator,
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address,
	}
}

func (msg *MsgCreateContact) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateContact{}

func NewMsgUpdateContact(creator string, id uint64, name string, phone string, email string, address string) *MsgUpdateContact {
	return &MsgUpdateContact{
		Id:      id,
		Creator: creator,
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address,
	}
}

func (msg *MsgUpdateContact) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteContact{}

func NewMsgDeleteContact(creator string, id uint64) *MsgDeleteContact {
	return &MsgDeleteContact{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteContact) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
