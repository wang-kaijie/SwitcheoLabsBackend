package keeper

import (
	"userBlock/x/userblock/types"
)

var _ types.QueryServer = Keeper{}
