package keeper

import (
	"context"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    "github.com/cosmos/cosmos-sdk/types/query"

)

func (k Keeper) SpecificUser(goCtx context.Context, req *types.QuerySpecificUserRequest) (*types.QuerySpecificUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	user, found := k.GetUser(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QuerySpecificUserResponse{User: &user}, nil
}


func (k Keeper) AllUsers(ctx context.Context, req *types.QueryAllUsersRequest) (*types.QueryAllUsersResponse, error) {
    if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []types.User

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))

	pageRes, err := query.Paginate(userStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return err
		}

		users = append(users, user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryAllUsersResponse{Users: users, Pagination: pageRes}, nil
}
