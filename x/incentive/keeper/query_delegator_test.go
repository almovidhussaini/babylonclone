package keeper_test

import (
	"github.com/amovidhussaini/ybtcclone/testutil/datagen"
	testkeeper "github.com/amovidhussaini/ybtcclone/testutil/keeper"
	"github.com/amovidhussaini/ybtcclone/x/incentive/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelegatorAddressQuery(t *testing.T) {
	keeper, ctx := testkeeper.IncentiveKeeper(t, nil, nil, nil)
	withdrawalAddr := datagen.GenRandomAccount().GetAddress()
	delegatorAddr := datagen.GenRandomAccount().GetAddress()
	err := keeper.SetWithdrawAddr(ctx, delegatorAddr, withdrawalAddr)
	require.NoError(t, err)

	response, err := keeper.DelegatorWithdrawAddress(ctx, &types.QueryDelegatorWithdrawAddressRequest{DelegatorAddress: delegatorAddr.String()})
	require.NoError(t, err)
	require.Equal(t, &types.QueryDelegatorWithdrawAddressResponse{WithdrawAddress: withdrawalAddr.String()}, response)
}
