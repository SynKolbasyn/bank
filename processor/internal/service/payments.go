package service

import (
	"context"

	"github.com/SynKolbasyn/bank/processor/internal/model"
	"github.com/SynKolbasyn/bank/processor/internal/repository"
	"github.com/google/uuid"
)

type Payments struct {
	transactionManager repository.TransactionManager
	repositoryUser     repository.IUser
	repositoryPayments repository.IPayments
}

func NewPayments(
	transactionManager repository.TransactionManager,
	repositoryUser repository.IUser,
	repositoryPayments repository.IPayments,
) *Payments {
	return &Payments{
		transactionManager: transactionManager,
		repositoryUser:     repositoryUser,
		repositoryPayments: repositoryPayments,
	}
}

func (p *Payments) Process(ctx context.Context, paymentID uuid.UUID) error {
	err := p.transactionManager.WithTransaction(ctx, func(ctx context.Context) error {
		status, err := p.repositoryPayments.GetStatusForUpdate(ctx, paymentID)
		if err != nil {
			return err
		}

		if status == model.PaymentStatusPending {
			return p.repositoryPayments.SetStatus(ctx, paymentID, model.PaymentStatusProcessing)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return p.processPayment(ctx, paymentID)
}

func (p *Payments) processPayment(ctx context.Context, paymentID uuid.UUID) error {
	return p.transactionManager.WithTransaction(ctx, func(ctx context.Context) error {
		payment, err := p.repositoryPayments.GetForUpdate(ctx, paymentID)
		if err != nil {
			return err
		}

		if payment.Status != model.PaymentStatusProcessing {
			return nil
		}

		senderMoney, err := p.repositoryUser.GetBalanceForUpdate(ctx, payment.Sender)
		if err != nil {
			return err
		}

		recipientMoney, err := p.repositoryUser.GetBalanceForUpdate(ctx, payment.Recipient)
		if err != nil {
			return err
		}

		if senderMoney.LessThan(payment.Amount) {
			return p.repositoryPayments.SetStatus(ctx, paymentID, model.PaymentStatusCanceled)
		}

		err = p.repositoryUser.SetBalance(ctx, payment.Sender, senderMoney.Sub(payment.Amount))
		if err != nil {
			return err
		}

		err = p.repositoryUser.SetBalance(
			ctx,
			payment.Recipient,
			recipientMoney.Add(payment.Amount),
		)
		if err != nil {
			return err
		}

		return p.repositoryPayments.SetStatus(ctx, paymentID, model.PaymentStatusSuccess)
	})
}
