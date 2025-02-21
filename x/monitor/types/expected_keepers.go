package types

import (
	"context"
	lc "github.com/amovidhussaini/ybtcclone/x/btclightclient/types"
)

type BTCLightClientKeeper interface {
	GetTipInfo(ctx context.Context) *lc.BTCHeaderInfo
	GetBaseBTCHeader(ctx context.Context) *lc.BTCHeaderInfo
}
