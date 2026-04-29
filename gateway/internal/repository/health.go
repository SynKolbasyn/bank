package repository

import (
	"context"
	"net/http"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Health struct {
	pool *pgxpool.Pool
}

func NewHealth(pool *pgxpool.Pool) *Health {
	return &Health{
		pool: pool,
	}
}

func (h *Health) Health(ctx context.Context) error {
	err := h.pool.Ping(ctx)
	if err != nil {
		return domain.NewAppError(http.StatusServiceUnavailable, err)
	}
	return nil
} 
