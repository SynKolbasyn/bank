package app

import (
	"github.com/SynKolbasyn/bank/config"
	"github.com/SynKolbasyn/bank/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewServer(config *config.Config, pool *pgxpool.Pool, clientRedpanda *kgo.Client) *echo.Echo {
	server := echo.New()
	server.Validator = domain.NewValidator()

	repositories := NewRepositories(pool)
	services := NewServices(config, repositories, clientRedpanda)
	handlers := NewHandlers(config, services)

	setRoutes(server, config, handlers)

	return server
}
