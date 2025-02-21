package keeper

import (
	"github.com/amovidhussaini/ybtcclone/x/finality/types"
)

var _ types.QueryServer = Keeper{}
