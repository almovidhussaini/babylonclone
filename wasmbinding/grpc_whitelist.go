package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bsctypes "github.com/amovidhussaini/ybtcclone/x/btcstkconsumer/types"
	epochtypes "github.com/amovidhussaini/ybtcclone/x/epoching/types"
	ftypes "github.com/amovidhussaini/ybtcclone/x/finality/types"
)

// WhitelistedGrpcQuery returns the whitelisted Grpc queries
func WhitelistedGrpcQuery() wasmkeeper.AcceptedQueries {
	return wasmkeeper.AcceptedQueries{
		// btcstkconsumer
		"/ybtc.btcstkconsumer.v1.Query/FinalityProvider": &bsctypes.QueryFinalityProviderResponse{},
		// btcstaking
		"/ybtc.btcstaking.v1.Query/FinalityProviderCurrentPower": &ftypes.QueryFinalityProviderCurrentPowerResponse{},
		// for testing
		"/ybtc.epoching.v1.Query/CurrentEpoch": &epochtypes.QueryCurrentEpochResponse{},
	}
}
