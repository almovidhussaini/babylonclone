package zoneconcierge

import (
	"context"
	"time"

	"github.com/almovidhussaini/babylonclone/x/zoneconcierge/keeper"
	"github.com/almovidhussaini/babylonclone/x/zoneconcierge/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
)

// BeginBlocker sends a pending packet for every channel upon each new block,
// so that the relayer is kept awake to relay headers
func BeginBlocker(ctx context.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	return nil
}

func EndBlocker(ctx context.Context, k keeper.Keeper) ([]abci.ValidatorUpdate, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.BroadcastBTCStakingConsumerEvents(ctx)
	return []abci.ValidatorUpdate{}, nil
}
