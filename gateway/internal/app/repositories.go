package app

import (
	"github.com/SynKolbasyn/bank/gateway/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	health   repository.IHealth
	user     repository.IUser
	payments repository.IPayments
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		health:   repository.NewHealth(pool),
		user:     repository.NewUser(pool),
		payments: repository.NewPayments(pool),
	}
}
