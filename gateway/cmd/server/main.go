package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SynKolbasyn/bank/gateway/config"
	"github.com/SynKolbasyn/bank/gateway/internal/app"
	"github.com/SynKolbasyn/bank/gateway/migrations"
	"github.com/SynKolbasyn/bank/gateway/pkg/logger"
	"github.com/SynKolbasyn/bank/gateway/pkg/redpanda"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "config.LoadConfig", slog.String("error", err.Error()))

		return err
	}

	logger := logger.NewLogger(cfg.Server.LogLevel)
	slog.SetDefault(logger)

	pool, err := pgxpool.New(ctx, cfg.Postgres.DSN())
	if err != nil {
		slog.ErrorContext(ctx, "pgxpool.New", slog.String("error", err.Error()))

		return err
	}
	defer pool.Close()

	err = migrations.Up(ctx, pool)
	if err != nil {
		slog.ErrorContext(ctx, "migrations.Up", slog.String("error", err.Error()))

		return err
	}

	clientRedpanda, err := redpanda.NewClient(cfg.Redpanda.Hosts, nil)
	if err != nil {
		slog.ErrorContext(ctx, "redpanda.NewClient", slog.String("error", err.Error()))

		return err
	}
	defer clientRedpanda.Close()

	serverConfig := echo.StartConfig{
		Address:         cfg.Server.Address(),
		GracefulTimeout: time.Second,
	}

	err = serverConfig.Start(ctx, app.NewServer(cfg, logger, pool, clientRedpanda))
	if err != nil {
		slog.ErrorContext(ctx, "serverConfig.Start", slog.String("error", err.Error()))

		return err
	}

	return nil
}
