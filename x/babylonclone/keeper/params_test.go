package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/amovidhussaini/babylonclone/testutil/keeper"
	"github.com/amovidhussaini/babylonclone/x/babylonclone/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BabyloncloneKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
