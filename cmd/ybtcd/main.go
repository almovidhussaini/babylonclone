package main

import (
	"cosmossdk.io/log"
	"os"

	"github.com/amovidhussaini/ybtcclone/app"
	"github.com/amovidhussaini/ybtcclone/cmd/ybtcd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/amovidhussaini/ybtcclone/app/params"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.ybtcAppEnvPrefix, app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)
	}
}
