package app

import (
	"github.com/SynKolbasyn/bank/processor/internal/handler"
)

type Handlers struct {
	health   *handler.Health
	payments *handler.Payments
}

func NewHandlers(services *Services) *Handlers {
	return &Handlers{
		health:   handler.NewHealth(services.health),
		payments: handler.NewPayments(services.payments),
	}
}
