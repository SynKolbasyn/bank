package config_test

import (
	"testing"

	"github.com/SynKolbasyn/bank/config"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	config, err := config.LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, config)
	require.NotNil(t, config.Server)
	require.NotNil(t, config.Auth)
	require.NotNil(t, config.Postgres)
	require.NotNil(t, config.Redpanda)
}

func TestLoadConfigErr(t *testing.T) {
	t.Setenv(string(config.KeyServerLogLevel), "UNKNOWN")
	config, err := config.LoadConfig()
	require.Error(t, err)
	require.Nil(t, config)
}
