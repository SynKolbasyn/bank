package service

import (
	"context"

	"github.com/google/uuid"
)

type IHealth interface {
	Health(ctx context.Context) error
}

type IPayments interface {
	Process(ctx context.Context, paymentID uuid.UUID) error
}
