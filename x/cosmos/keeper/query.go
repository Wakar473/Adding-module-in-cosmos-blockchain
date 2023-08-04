package keeper

import (
	"cosmos/x/cosmos/types"
)

var _ types.QueryServer = Keeper{}
