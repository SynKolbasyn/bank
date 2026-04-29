package repository

import (
	"context"

	"github.com/SynKolbasyn/bank/processor/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shopspring/decimal"
)

type repositoryCtxtKey string

const KeyTx repositoryCtxtKey = "pgx_tx"

type IExecutor interface {
	Exec(
		ctx context.Context,
		sql string,
		arguments ...any,
	) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type TransactionManager interface {
	WithTransaction(ctx context.Context, function func(ctx context.Context) error) error
}

type IHealth interface {
	Health(ctx context.Context) error
}

type IUser interface {
	GetBalanceForUpdate(ctx context.Context, userID uuid.UUID) (decimal.Decimal, error)
	SetBalance(ctx context.Context, userID uuid.UUID, money decimal.Decimal) error
}

type IPayments interface {
	GetStatusForUpdate(ctx context.Context, paymentID uuid.UUID) (string, error)
	SetStatus(ctx context.Context, paymentID uuid.UUID, status string) error
	GetForUpdate(ctx context.Context, paymentID uuid.UUID) (model.Payment, error)
}
