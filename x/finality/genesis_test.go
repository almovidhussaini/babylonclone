package finality_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/amovidhussaini/ybtcclone/testutil/keeper"
	"github.com/amovidhussaini/ybtcclone/testutil/nullify"
	"github.com/amovidhussaini/ybtcclone/x/finality"
	"github.com/amovidhussaini/ybtcclone/x/finality/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.FinalityKeeper(t, nil, nil, nil)
	finality.InitGenesis(ctx, *k, genesisState)
	got := finality.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
