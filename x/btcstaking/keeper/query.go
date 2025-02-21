package keeper

import (
	"github.com/amovidhussaini/ybtcclone/x/btcstaking/types"
)

var _ types.QueryServer = Keeper{}
