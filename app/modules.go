package app

import (
	"github.com/CosmWasm/wasmd/x/wasm/ibctesting"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/client"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
)

var _ ibctesting.ChainApp = &ybtcApp{}

// The following functions are required by ibctesting
// (copied from https://github.com/osmosis-labs/osmosis/blob/main/app/modules.go)

func (app *ybtcApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

func (app *ybtcApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

func (app *ybtcApp) GetBankKeeper() bankkeeper.Keeper {
	return app.BankKeeper
}

func (app *ybtcApp) GetStakingKeeper() *stakingkeeper.Keeper {
	return app.StakingKeeper
}

func (app *ybtcApp) GetAccountKeeper() authkeeper.AccountKeeper {
	return app.AccountKeeper
}

func (app *ybtcApp) GetWasmKeeper() wasmkeeper.Keeper {
	return app.WasmKeeper
}

func (app *ybtcApp) GetTxConfig() client.TxConfig {
	return app.TxConfig()
}
