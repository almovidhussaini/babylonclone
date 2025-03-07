package incentive

import (
	"context"

	"github.com/almovidhussaini/babylonclone/x/incentive/keeper"
	"github.com/almovidhussaini/babylonclone/x/incentive/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx context.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
	// TODO(rafilx): add gauge, reward tracker
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx context.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// TODO(rafilx): add gauge, reward tracker
	return genesis
}
