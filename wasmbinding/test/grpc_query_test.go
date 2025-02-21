package wasmbinding

import (
	"encoding/json"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/amovidhussaini/ybtcclone/app"
	epochingtypes "github.com/amovidhussaini/ybtcclone/x/epoching/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var pathToGrpcContract = "../testdata/artifacts/testgrpc.wasm"

func TestGrpcQueryEpoch(t *testing.T) {
	acc := RandomAccountAddress()
	ybtcApp, ctx := SetupAppWithContext(t)
	FundAccount(t, ctx, ybtcApp, acc)

	contractAddress := deployGrpcContract(t, ctx, ybtcApp, acc, pathToGrpcContract)
	require.NotEmpty(t, contractAddress)

	query := TestQuery{
		QueryCurrentEpoch: &struct{}{},
	}
	resp := epochingtypes.QueryCurrentEpochResponse{}
	testGrpcQuery(t, ctx, ybtcApp, contractAddress, query, &resp)
	require.Equal(t, resp.CurrentEpoch, uint64(1))

	newEpoch := ybtcApp.EpochingKeeper.IncEpoch(ctx)

	resp = epochingtypes.QueryCurrentEpochResponse{}
	testGrpcQuery(t, ctx, ybtcApp, contractAddress, query, &resp)
	require.Equal(t, resp.CurrentEpoch, newEpoch.EpochNumber)
}

func instantiateGrpcContract(
	t *testing.T,
	ctx sdk.Context,
	bbn *app.ybtcApp,
	funder sdk.AccAddress,
	codeId uint64,
) sdk.AccAddress {
	initMsgBz := []byte(`{"admin":"bbn1kghr9hekuxj0tqa9pfnpxym4x6z0k0x77qxa79", "consumer_id":"op-stack-l2-1"}`)
	contractKeeper := keeper.NewDefaultPermissionKeeper(bbn.WasmKeeper)
	addr, _, err := contractKeeper.Instantiate(ctx, codeId, funder, funder, initMsgBz, "test grpc contract", nil)
	require.NoError(t, err)
	return addr
}

func deployGrpcContract(
	t *testing.T,
	ctx sdk.Context,
	bbn *app.ybtcApp,
	deployer sdk.AccAddress,
	codePath string,
) sdk.AccAddress {
	codeId, _ := StoreTestCode(t, ctx, bbn, deployer, codePath)
	contractAddr := instantiateGrpcContract(t, ctx, bbn, deployer, codeId)
	return contractAddr
}

type TestQuery struct {
	QueryCurrentEpoch *struct{} `json:"query_current_epoch,omitempty"`
}

func testGrpcQuery(
	t *testing.T,
	ctx sdk.Context,
	bbn *app.ybtcApp,
	contract sdk.AccAddress,
	request TestQuery,
	response interface{},
) {
	msgBz, err := json.Marshal(request)
	require.NoError(t, err)

	resBz, err := bbn.WasmKeeper.QuerySmart(ctx, contract, msgBz)
	require.NoError(t, err)

	err = json.Unmarshal(resBz, &response)
	require.NoError(t, err)
}
