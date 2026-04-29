package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type User struct {
	executor *Executor
}

func NewUser(executor *Executor) *User {
	return &User{
		executor: executor,
	}
}

func (u *User) GetBalanceForUpdate(ctx context.Context, userID uuid.UUID) (decimal.Decimal, error) {
	query := `
		SELECT money
		FROM users
		WHERE id = $1::UUID
		FOR UPDATE;
	`

	var money decimal.Decimal

	err := u.executor.GetExecutor(ctx).QueryRow(ctx, query, userID).Scan(&money)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return money, nil
}

func (u *User) SetBalance(ctx context.Context, userID uuid.UUID, money decimal.Decimal) error {
	query := `
		UPDATE users
		SET money = $2::NUMERIC(17, 2)
		WHERE id = $1::UUID;
	`
	_, err := u.executor.GetExecutor(ctx).Exec(ctx, query, userID, money)

	return err
}
