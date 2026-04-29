package app

import (
	"github.com/SynKolbasyn/bank/internal/handler"
	"github.com/SynKolbasyn/bank/config"
)

type Handlers struct {
	health *handler.Health
	auth *handler.Auth
	payments *handler.Payments
}

func NewHandlers(cfg *config.Config, services *Services) *Handlers {
	return &Handlers{
		health: handler.NewHealth(services.health),
		auth: handler.NewAuth(services.auth),
		payments: handler.NewPayments(services.payments),
	}
}
