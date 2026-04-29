package handler

import (
	"context"

	"github.com/SynKolbasyn/bank/processor/internal/service"
	"github.com/google/uuid"
)

type Payments struct {
	servicePayments service.IPayments
}

func NewPayments(servicePayments service.IPayments) *Payments {
	return &Payments{
		servicePayments: servicePayments,
	}
}

func (a *Payments) Process(ctx context.Context, data []byte) error {
	paymentID, err := uuid.Parse(string(data))
	if err != nil {
		return err
	}

	return a.servicePayments.Process(ctx, paymentID)
}
