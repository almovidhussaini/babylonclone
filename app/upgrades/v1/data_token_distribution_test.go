package v1_test

import (
	"testing"

	v1 "github.com/almovidhussaini/babylonclone/app/upgrades/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCheckTokensDistributionFromData(t *testing.T) {
	for _, upgradeData := range UpgradeV1Data {
		d, err := v1.LoadTokenDistributionFromData(upgradeData.TokensDistributionStr)
		require.NoError(t, err)
		require.Greater(t, len(d.TokenDistribution), 1)

		for _, td := range d.TokenDistribution {
			sender, err := sdk.AccAddressFromBech32(td.AddressSender)
			require.NoError(t, err)
			require.Equal(t, sender.String(), td.AddressSender)

			receiver, err := sdk.AccAddressFromBech32(td.AddressReceiver)
			require.NoError(t, err)
			require.Equal(t, receiver.String(), td.AddressReceiver)

			require.True(t, td.Amount > 0)
		}
	}
}
