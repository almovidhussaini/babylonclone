package monitor_test

import (
	"testing"

	"github.com/amovidhussaini/ybtcclone/x/monitor"
	"github.com/stretchr/testify/require"

	simapp "github.com/amovidhussaini/ybtcclone/app"
	"github.com/amovidhussaini/ybtcclone/x/monitor/types"
)

func TestExportGenesis(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false)
	genesisState := monitor.ExportGenesis(ctx, app.MonitorKeeper)
	require.Equal(t, genesisState, types.DefaultGenesis())
}
