package incentive_test

import (
	"testing"

	keepertest "github.com/amovidhussaini/ybtcclone/testutil/keeper"
	"github.com/amovidhussaini/ybtcclone/testutil/nullify"
	"github.com/amovidhussaini/ybtcclone/x/incentive"
	"github.com/amovidhussaini/ybtcclone/x/incentive/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.IncentiveKeeper(t, nil, nil, nil)
	incentive.InitGenesis(ctx, *k, genesisState)
	got := incentive.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
