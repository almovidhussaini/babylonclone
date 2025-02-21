package keeper_test

import (
	"testing"

	testkeeper "github.com/amovidhussaini/ybtcclone/testutil/keeper"
	"github.com/amovidhussaini/ybtcclone/x/incentive/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IncentiveKeeper(t, nil, nil, nil)
	params := types.DefaultParams()

	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	require.EqualValues(t, params, k.GetParams(ctx))
}
