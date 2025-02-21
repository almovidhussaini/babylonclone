package config_test

import (
	"testing"

	"github.com/amovidhussaini/ybtcclone/client/config"
	"github.com/stretchr/testify/require"
)

// TestybtcQueryConfig ensures that the default ybtc query config is valid
func TestybtcQueryConfig(t *testing.T) {
	defaultConfig := config.DefaultybtcQueryConfig()
	err := defaultConfig.Validate()
	require.NoError(t, err)
}
