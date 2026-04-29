package repository

import (
	"context"

	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/google/uuid"
)

type IHealth interface {
	Health(ctx context.Context) error
}

type IUser interface {
	Create(ctx context.Context, email, password string) (uuid.UUID, error)
	Get(ctx context.Context, email string) (uuid.UUID, string, error)
}

type IPayments interface {
	Create(ctx context.Context, userID uuid.UUID, payment model.PaymentRequest) (uuid.UUID, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]model.Payment, error)
}
