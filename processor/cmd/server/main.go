package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SynKolbasyn/bank/processor/config"
	"github.com/SynKolbasyn/bank/processor/internal/app"
	"github.com/SynKolbasyn/bank/processor/pkg/logger"
	"github.com/SynKolbasyn/bank/processor/pkg/redpanda"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "config.LoadConfig", slog.String("error", err.Error()))
		os.Exit(1)
	}

	logger := logger.NewLogger(cfg.Server.LogLevel)
	slog.SetDefault(logger)

	pool, err := pgxpool.New(ctx, cfg.Postgres.DSN())
	if err != nil {
		slog.ErrorContext(ctx, "pgxpool.New", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer pool.Close()

	clientRedpanda, err := redpanda.NewClient(cfg.Redpanda.Hosts, cfg.Redpanda.Topics)
	if err != nil {
		slog.ErrorContext(ctx, "redpanda.NewClient", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer clientRedpanda.Close()

	err = app.StartServer(ctx, cfg, logger, pool, clientRedpanda)
	if err != nil {
		slog.ErrorContext(ctx, "app.StartServer", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
