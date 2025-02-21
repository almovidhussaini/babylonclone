package keeper_test

import (
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/amovidhussaini/ybtcclone/testutil/datagen"
	keepertest "github.com/amovidhussaini/ybtcclone/testutil/keeper"
	btclctypes "github.com/amovidhussaini/ybtcclone/x/btclightclient/types"
	"github.com/amovidhussaini/ybtcclone/x/btcstaking/types"
)

func FuzzBTCHeightIndex(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// mock BTC light client
		btclcKeeper := types.NewMockBTCLightClientKeeper(ctrl)
		keeper, ctx := keepertest.BTCStakingKeeper(t, btclcKeeper, nil, nil)

		// randomise ybtc height and BTC height
		ybtcHeight := datagen.RandomInt(r, 100)
		ctx = datagen.WithCtxHeight(ctx, ybtcHeight)
		btcHeight := uint32(datagen.RandomInt(r, 100))
		btclcKeeper.EXPECT().GetTipInfo(gomock.Any()).Return(&btclctypes.BTCHeaderInfo{Height: btcHeight}).Times(1)
		keeper.IndexBTCHeight(ctx)

		// assert BTC height
		actualBtcHeight := keeper.GetBTCHeightAtybtcHeight(ctx, ybtcHeight)
		require.Equal(t, btcHeight, actualBtcHeight)
		// assert current BTC height
		curBtcHeight := keeper.GetCurrentBTCHeight(ctx)
		require.Equal(t, btcHeight, curBtcHeight)
	})
}
