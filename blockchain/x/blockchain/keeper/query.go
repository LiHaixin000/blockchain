package keeper

import (
	"github.com/LiHaixin000/blockchain/x/blockchain/types"
)

var _ types.QueryServer = Keeper{}
