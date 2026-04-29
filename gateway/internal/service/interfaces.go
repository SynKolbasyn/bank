package service

import (
	"context"

	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/google/uuid"
)

type IHealth interface {
	Health(ctx context.Context) model.HealthResponse
}

type INotificationManager interface {
	Notify(ctx context.Context, paymentID uuid.UUID) error
}

type IAuth interface {
	SignUp(ctx context.Context, user *model.SignRequest) (string, error)
	SignIn(ctx context.Context, user *model.SignRequest) (string, error)
}

type IPayments interface{
	Create(ctx context.Context, userID uuid.UUID, payment model.PaymentRequest) error
	Get(ctx context.Context, userID uuid.UUID) ([]model.Payment, error)
}
