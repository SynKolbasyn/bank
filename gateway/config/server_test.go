package config_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/SynKolbasyn/bank/gateway/config"
	"github.com/stretchr/testify/require"
)

func TestLoadServerAddress(t *testing.T) {
	testData := []struct {
		Host     string
		Port     string
		Expected string
	}{
		{"localhost", "80", "localhost:80"},
		{"localhost", "8080", "localhost:8080"},
		{"specialhost", "443", "specialhost:443"},
		{"host.with.layer", "443", "host.with.layer:443"},
	}

	for i, data := range testData {
		t.Run(fmt.Sprintf("host-%d", i+1), func(t *testing.T) {
			t.Setenv(string(config.KeyServerHost), data.Host)
			t.Setenv(string(config.KeyServerPort), data.Port)

			server, err := config.LoadServer()
			require.NoError(t, err)
			require.NotNil(t, server)
			require.Equal(t, data.Expected, server.Address())
		})
	}
}

func TestLoadServerLogLevel(t *testing.T) {
	testData := []struct {
		LogLevel string
		Expected slog.Level
		IsError  bool
	}{
		{"DEBUG", slog.LevelDebug, false},
		{"INFO", slog.LevelInfo, false},
		{"WARN", slog.LevelWarn, false},
		{"ERROR", slog.LevelError, false},
		{"DeBuG", slog.LevelDebug, false},
		{"iNfO", slog.LevelInfo, false},
		{"warn", slog.LevelWarn, false},
		{"ErrOr", slog.LevelError, false},
		{"unknown", 0, true},
	}

	for i, data := range testData {
		t.Run(fmt.Sprintf("host-%d", i+1), func(t *testing.T) {
			t.Setenv(string(config.KeyServerLogLevel), data.LogLevel)

			server, err := config.LoadServer()
			if data.IsError {
				require.Error(t, err)
				require.Nil(t, server)
			} else {
				require.NoError(t, err)
				require.NotNil(t, server)
				require.Equal(t, data.Expected, server.LogLevel)
			}
		})
	}
}
