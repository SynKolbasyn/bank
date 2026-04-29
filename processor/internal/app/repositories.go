package app

import (
	"github.com/SynKolbasyn/bank/processor/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	transactionManager repository.TransactionManager
	health             repository.IHealth
	user               repository.IUser
	payments           repository.IPayments
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	executor := repository.NewExecutor(pool)

	return &Repositories{
		transactionManager: executor,
		health:             repository.NewHealth(pool),
		user:               repository.NewUser(executor),
		payments:           repository.NewPayments(executor),
	}
}
