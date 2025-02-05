package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateUser{}

func NewMsgUpdateUser(creator string, name string, email string, age uint32, gender string, address string) *MsgUpdateUser {
	return &MsgUpdateUser{
		Creator: creator,
		Name:    name,
		Email:   email,
		Age:     age,
		Gender:  gender,
		Address: address,
	}
}

func (msg *MsgUpdateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
