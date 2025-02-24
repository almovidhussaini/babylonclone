package keeper

import (
	"github.com/almovidhussaini/babylonclone/x/btclightclient/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

func headerInfoFromStoredBytes(cdc codec.BinaryCodec, bz []byte) *types.BTCHeaderInfo {
	headerInfo := new(types.BTCHeaderInfo)
	cdc.MustUnmarshal(bz, headerInfo)
	return headerInfo
}
