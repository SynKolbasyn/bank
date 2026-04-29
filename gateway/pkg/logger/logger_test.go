package logger_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/SynKolbasyn/bank/gateway/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()
	data := []slog.Level{
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}
	for _, level := range data {
		t.Run(fmt.Sprintf("level-%s", level), func(t *testing.T) {
			logger := logger.NewLogger(level)
			require.NotNil(t, logger)
		})
	}
}
