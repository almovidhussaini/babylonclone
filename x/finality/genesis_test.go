package finality_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/almovidhussaini/babylonclone/testutil/keeper"
	"github.com/almovidhussaini/babylonclone/testutil/nullify"
	"github.com/almovidhussaini/babylonclone/x/finality"
	"github.com/almovidhussaini/babylonclone/x/finality/types"
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
