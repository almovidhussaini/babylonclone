package app

import (
	"os"

	"cosmossdk.io/log"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client/flags"
	simsutils "github.com/cosmos/cosmos-sdk/testutil/sims"

	appparams "github.com/amovidhussaini/ybtcclone/app/params"
	"github.com/amovidhussaini/ybtcclone/testutil/signer"
	bbn "github.com/amovidhussaini/ybtcclone/types"
)

// TmpAppOptions returns an app option with tmp dir and btc network
func TmpAppOptions() simsutils.AppOptionsMap {
	dir, err := os.MkdirTemp("", "ybtc-tmp-app")
	if err != nil {
		panic(err)
	}
	appOpts := simsutils.AppOptionsMap{
		flags.FlagHome:       dir,
		"btc-config.network": string(bbn.BtcSimnet),
	}
	return appOpts
}

func NewTmpybtcApp() *ybtcApp {
	signer, _ := signer.SetupTestPrivSigner()
	return NewybtcApp(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		0,
		signer,
		TmpAppOptions(),
		[]wasmkeeper.Option{})
}

// GetEncodingConfig returns a *registered* encoding config
// Note that the only way to register configuration is through the app creation
func GetEncodingConfig() *appparams.EncodingConfig {
	tmpApp := NewTmpybtcApp()
	return &appparams.EncodingConfig{
		InterfaceRegistry: tmpApp.InterfaceRegistry(),
		Codec:             tmpApp.AppCodec(),
		TxConfig:          tmpApp.TxConfig(),
		Amino:             tmpApp.LegacyAmino(),
	}
}
