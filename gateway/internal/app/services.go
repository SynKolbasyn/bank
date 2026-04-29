package app

import (
	"github.com/SynKolbasyn/bank/gateway/config"
	"github.com/SynKolbasyn/bank/gateway/internal/service"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Services struct {
	health service.IHealth
	auth service.IAuth
	payments service.IPayments
}

func NewServices(cfg *config.Config, repositories *Repositories, clientRedpanda *kgo.Client) *Services {
	return &Services{
		health: service.NewHealth(repositories.health),
		auth: service.NewAuth(repositories.user, cfg.Auth.Secret),
		payments: service.NewPayments(repositories.payments, service.NewNotificationManager(clientRedpanda, "payments")),
	}
}
