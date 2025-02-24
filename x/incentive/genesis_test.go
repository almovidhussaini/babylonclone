package incentive_test

import (
	"testing"

	keepertest "github.com/almovidhussaini/babylonclone/testutil/keeper"
	"github.com/almovidhussaini/babylonclone/testutil/nullify"
	"github.com/almovidhussaini/babylonclone/x/incentive"
	"github.com/almovidhussaini/babylonclone/x/incentive/types"
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
