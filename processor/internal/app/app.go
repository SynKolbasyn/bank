package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/SynKolbasyn/bank/processor/config"
	"github.com/SynKolbasyn/bank/processor/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/twmb/franz-go/pkg/kgo"
)

func StartServer(
	ctx context.Context,
	cfg *config.Config,
	logger *slog.Logger,
	pool *pgxpool.Pool,
	clientRedpanda *kgo.Client,
) error {
	server := echo.New()
	server.Logger = logger
	server.Validator = domain.NewValidator()

	repositories := NewRepositories(pool)
	services := NewServices(repositories)
	handlers := NewHandlers(services)

	setRoutes(server, handlers)

	go StartProcessing(ctx, clientRedpanda, handlers.payments)

	startConfig := echo.StartConfig{
		Address:         cfg.Server.Address(),
		GracefulTimeout: time.Second,
	}

	return startConfig.Start(ctx, server)
}
