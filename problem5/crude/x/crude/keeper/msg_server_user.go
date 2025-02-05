package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)
// create a new user
func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var user = types.User{
		Id:      msg.Id,
		Name:    msg.Name,
		Email:   msg.Email,
		Age:     msg.Age,
		Gender:  msg.Gender,
		Address: msg.Address,
	}
	id := k.AppendUser(
		ctx,
		user,
	)
	return &types.MsgCreateUserResponse{
		Id: id,
	}, nil
}

// update current user
func (k msgServer) UpdateUser(goCtx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var user = types.User{
        Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Email:   msg.Email,
		Age:     msg.Age,
		Gender:  msg.Gender,
		Address: msg.Address,
	}
	val, found := k.GetUser(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
    fmt.Println("Stored Creator:", val.Creator)
    fmt.Println("Incoming Creator:", msg.Creator)
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	err := k.SetUser(ctx, user)
    if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrPanic, "Unable to update")
	}
	return &types.MsgUpdateUserResponse{}, nil
}


// delete current user
func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetUser(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.RemoveUser(ctx, msg.Id)
	return &types.MsgDeleteUserResponse{}, nil
}
