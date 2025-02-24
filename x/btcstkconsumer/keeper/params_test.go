package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/almovidhussaini/babylonclone/testutil/keeper"
	"github.com/almovidhussaini/babylonclone/x/btcstkconsumer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BTCStkConsumerKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
