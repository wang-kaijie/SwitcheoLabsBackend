package keeper

import (
	"context"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SpecificUser(goCtx context.Context, req *types.QuerySpecificUserRequest) (*types.QuerySpecificUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QuerySpecificUserResponse{}, nil
}
