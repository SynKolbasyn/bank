package service

import (
	"context"
	"net/http"

	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/SynKolbasyn/bank/gateway/internal/repository"
)

type Health struct {
	repositoryHealth repository.IHealth
}

func NewHealth(repositoryHealth repository.IHealth) *Health {
	return &Health{
		repositoryHealth: repositoryHealth,
	}
}

func (h *Health) Health(ctx context.Context) model.HealthResponse {
	database := http.StatusOK

	err := h.repositoryHealth.Health(ctx)
	if err != nil {
		database = http.StatusServiceUnavailable
	}

	return model.HealthResponse{
		Databse: http.StatusText(database),
	}
}
