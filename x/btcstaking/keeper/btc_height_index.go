package keeper

import (
	"context"
	"math"

	"cosmossdk.io/store/prefix"
	"github.com/amovidhussaini/ybtcclone/x/btcstaking/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IndexBTCHeight indexes the current BTC height, and saves it to KVStore
func (k Keeper) IndexBTCHeight(ctx context.Context) {
	ybtcHeight := uint64(sdk.UnwrapSDKContext(ctx).HeaderInfo().Height)
	btcTip := k.btclcKeeper.GetTipInfo(ctx)
	if btcTip == nil {
		return
	}
	btcHeight := btcTip.Height
	store := k.btcHeightStore(ctx)
	store.Set(sdk.Uint64ToBigEndian(ybtcHeight), sdk.Uint64ToBigEndian(uint64(btcHeight)))
}

func (k Keeper) GetBTCHeightAtybtcHeight(ctx context.Context, ybtcHeight uint64) uint32 {
	store := k.btcHeightStore(ctx)
	btcHeightBytes := store.Get(sdk.Uint64ToBigEndian(ybtcHeight))
	if len(btcHeightBytes) == 0 {
		// if the previous height is not indexed (which might happen at genesis),
		// use the base header
		return k.btclcKeeper.GetBaseBTCHeader(ctx).Height
	}
	btcHeightUint64 := sdk.BigEndianToUint64(btcHeightBytes)
	if btcHeightUint64 > math.MaxUint32 {
		panic("Storage involves a btc height that is larger than math.MaxUint32")
	}
	return uint32(btcHeightUint64)
}

func (k Keeper) GetCurrentBTCHeight(ctx context.Context) uint32 {
	ybtcHeight := uint64(sdk.UnwrapSDKContext(ctx).HeaderInfo().Height)
	return k.GetBTCHeightAtybtcHeight(ctx, ybtcHeight)
}

// btcHeightStore returns the KVStore of the BTC heights
// prefix: BTCHeightKey
// key: ybtc block height
// value: BTC block height
func (k Keeper) btcHeightStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BTCHeightKey)
}
