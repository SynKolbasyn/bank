package app

import (
	"log/slog"

	"github.com/SynKolbasyn/bank/gateway/config"
	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewServer(
	cfg *config.Config,
	logger *slog.Logger,
	pool *pgxpool.Pool,
	clientRedpanda *kgo.Client,
) *echo.Echo {
	server := echo.New()
	server.Logger = logger
	server.Validator = domain.NewValidator()

	repositories := NewRepositories(pool)
	services := NewServices(cfg, repositories, clientRedpanda)
	handlers := NewHandlers(services)

	setRoutes(server, cfg, handlers)

	return server
}
