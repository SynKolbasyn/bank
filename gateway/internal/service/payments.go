package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/SynKolbasyn/bank/gateway/internal/repository"
	"github.com/google/uuid"
)

var ErrSelfPayment = errors.New("userID == payment.RecipientID")

type Payments struct {
	repositoryPayments  repository.IPayments
	notificationManager INotificationManager
}

func NewPayments(
	repositoryPayments repository.IPayments,
	notificationManager INotificationManager,
) *Payments {
	return &Payments{
		repositoryPayments:  repositoryPayments,
		notificationManager: notificationManager,
	}
}

func (p *Payments) Create(
	ctx context.Context,
	userID uuid.UUID,
	payment model.PaymentRequest,
) error {
	if userID == payment.RecipientID {
		return domain.NewAppError(
			http.StatusBadRequest,
			ErrSelfPayment,
		)
	}

	id, err := p.repositoryPayments.Create(ctx, userID, payment)
	if err != nil {
		return err
	}

	return p.notificationManager.Notify(ctx, id)
}

func (p *Payments) Get(ctx context.Context, userID uuid.UUID) ([]model.Payment, error) {
	return p.repositoryPayments.GetByUserID(ctx, userID)
}
