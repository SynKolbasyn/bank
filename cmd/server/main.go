package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SynKolbasyn/bank/config"
	"github.com/SynKolbasyn/bank/internal/app"
	"github.com/SynKolbasyn/bank/migrations"
	"github.com/SynKolbasyn/bank/pkg/redpanda"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	pool, err := pgxpool.New(ctx, config.Postgres.DSN())
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	err = migrations.Up(ctx, pool)
	if err != nil {
		panic(err)
	}

	clientRedpanda, err := redpanda.NewClient(config.Redpanda.Hosts, nil)
	if err != nil {
		panic(err)
	}
	defer clientRedpanda.Close()

	serverConfig := echo.StartConfig{
		Address: config.Server.Address(),
		GracefulTimeout: 1 * time.Second,
	}
	err = serverConfig.Start(ctx, app.NewServer(config, pool, clientRedpanda))
	if err != nil {
		panic(err)
	}
}
