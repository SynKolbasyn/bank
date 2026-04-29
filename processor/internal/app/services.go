package app

import (
	"github.com/SynKolbasyn/bank/processor/config"
	"github.com/SynKolbasyn/bank/processor/internal/service"
)

type Services struct {
	health   service.IHealth
	payments service.IPayments
}

func NewServices(config *config.Config, repositories *Repositories) *Services {
	return &Services{
		health:   service.NewHealth(repositories.health),
		payments: service.NewPayments(repositories.transactionManager, repositories.user, repositories.payments),
	}
}
