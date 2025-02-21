package mint_test

import (
	"testing"

	appparams "github.com/amovidhussaini/ybtcclone/app/params"
	"github.com/amovidhussaini/ybtcclone/testutil/helper"
	"github.com/amovidhussaini/ybtcclone/x/mint/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	h := helper.NewHelper(t)
	app, ctx := h.App, h.Ctx

	acc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)

	feeColl := app.AccountKeeper.GetAccount(ctx, appparams.AccFeeCollector)
	require.Equal(t, "bbn17xpfvakm2amg962yls6f84z3kell8c5l88j35y", feeColl.GetAddress().String())
}
