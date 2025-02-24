package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bsctypes "github.com/almovidhussaini/babylonclone/x/btcstkconsumer/types"
	epochtypes "github.com/almovidhussaini/babylonclone/x/epoching/types"
	ftypes "github.com/almovidhussaini/babylonclone/x/finality/types"
)

// WhitelistedGrpcQuery returns the whitelisted Grpc queries
func WhitelistedGrpcQuery() wasmkeeper.AcceptedQueries {
	return wasmkeeper.AcceptedQueries{
		// btcstkconsumer
		"/babylon.btcstkconsumer.v1.Query/FinalityProvider": &bsctypes.QueryFinalityProviderResponse{},
		// btcstaking
		"/babylon.btcstaking.v1.Query/FinalityProviderCurrentPower": &ftypes.QueryFinalityProviderCurrentPowerResponse{},
		// for testing
		"/babylon.epoching.v1.Query/CurrentEpoch": &epochtypes.QueryCurrentEpochResponse{},
	}
}
