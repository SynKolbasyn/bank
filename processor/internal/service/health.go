package service

import (
	"context"

	"github.com/SynKolbasyn/bank/processor/internal/repository"
)

type Health struct {
	repositoryHealth repository.IHealth
}

func NewHealth(repositoryHealth repository.IHealth) *Health {
	return &Health{
		repositoryHealth: repositoryHealth,
	}
}

func (h *Health) Health(ctx context.Context) error {
	return h.repositoryHealth.Health(ctx)
}
