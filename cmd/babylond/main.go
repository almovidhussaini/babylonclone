package main

import (
	"os"

	"cosmossdk.io/log"

	"github.com/almovidhussaini/babylonclone/app"
	"github.com/almovidhussaini/babylonclone/cmd/babylond/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/almovidhussaini/babylonclone/app/params"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.BabylonAppEnvPrefix, app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)
	}
}
