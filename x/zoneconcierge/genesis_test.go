package zoneconcierge_test

import (
	"testing"

	keepertest "github.com/almovidhussaini/babylonclone/testutil/keeper"
	"github.com/almovidhussaini/babylonclone/testutil/nullify"
	"github.com/almovidhussaini/babylonclone/x/zoneconcierge"
	"github.com/almovidhussaini/babylonclone/x/zoneconcierge/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		PortId: types.PortID,
		Params: types.Params{IbcPacketTimeoutSeconds: 100},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	portKeeper := types.NewMockPortKeeper(ctrl)
	portKeeper.EXPECT().BindPort(gomock.Any(), gomock.Any()).Return(&capabilitytypes.Capability{}).AnyTimes()

	k, ctx := keepertest.ZoneConciergeKeeper(t, nil, portKeeper, nil, nil, nil, nil, nil, nil)
	zoneconcierge.InitGenesis(ctx, *k, genesisState)
	got := zoneconcierge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)
	require.Equal(t, genesisState.Params, got.Params)
}
