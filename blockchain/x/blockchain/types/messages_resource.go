package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateResource{}

func NewMsgCreateResource(creator string, name string, data string) *MsgCreateResource {
	return &MsgCreateResource{
		Creator: creator,
		Name:    name,
		Data:    data,
	}
}

func (msg *MsgCreateResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateResource{}

func NewMsgUpdateResource(creator string, id uint64, name string, data string) *MsgUpdateResource {
	return &MsgUpdateResource{
		Id:      id,
		Creator: creator,
		Name:    name,
		Data:    data,
	}
}

func (msg *MsgUpdateResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteResource{}

func NewMsgDeleteResource(creator string, id uint64) *MsgDeleteResource {
	return &MsgDeleteResource{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
