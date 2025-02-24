package btcstaking_test

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	keepertest "github.com/almovidhussaini/babylonclone/testutil/keeper"
	"github.com/almovidhussaini/babylonclone/testutil/nullify"
	"github.com/almovidhussaini/babylonclone/x/btcstaking"
	"github.com/almovidhussaini/babylonclone/x/btcstaking/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	p := types.DefaultParams()
	genesisState := types.GenesisState{
		Params: []*types.Params{&p},
	}
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewTestLogger(t), storemetrics.NewNoOpMetrics())
	k, ctx := keepertest.BTCStakingKeeperWithStore(t, db, stateStore, nil, nil, nil)

	btcstaking.InitGenesis(ctx, *k, genesisState)
	got := btcstaking.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
