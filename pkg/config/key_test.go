package config_test

import (
	"testing"

	"github.com/SynKolbasyn/bank/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestGetValue(t *testing.T) {
	var key config.Key = "KEY"
	require.Empty(t, key.GetValue())
	t.Setenv(string(key), "VALUE")
	require.Equal(t, "VALUE", key.GetValue())
}

func TestGetValueDefault(t *testing.T) {
	var key config.Key = "KEY"
	require.Equal(t, "ANOTHER_VALUE", key.GetValueDefault("ANOTHER_VALUE"))
	t.Setenv(string(key), "VALUE")
	require.Equal(t, "VALUE", key.GetValueDefault("ANOTHER_VALUE"))
}
