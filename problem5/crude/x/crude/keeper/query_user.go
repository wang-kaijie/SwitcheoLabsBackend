package keeper

import (
	"context"

	"crude/x/crude/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (k Keeper) UsersByAge(ctx context.Context, req *types.QueryUsersByAgeRequest) (*types.QueryUsersByAgeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []types.User

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))
	reqAge := req.Age
	pageRes, err := query.Paginate(userStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return err
		}
		if user.Age == reqAge {
			users = append(users, user)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryUsersByAgeResponse{Users: users, Pagination: pageRes}, nil
}

func (k Keeper) UsersByGender(ctx context.Context, req *types.QueryUsersByGenderRequest) (*types.QueryUsersByGenderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []types.User

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))
	reqGender := req.Gender
	pageRes, err := query.Paginate(userStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return err
		}
		if user.Gender == reqGender {
			users = append(users, user)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryUsersByGenderResponse{Users: users, Pagination: pageRes}, nil
}
