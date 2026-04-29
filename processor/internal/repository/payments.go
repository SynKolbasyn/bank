package repository

import (
	"context"

	"github.com/SynKolbasyn/bank/processor/internal/model"
	"github.com/google/uuid"
)

type Payments struct {
	executor *Executor
}

func NewPayments(pool *Executor) *Payments {
	return &Payments{
		executor: pool,
	}
}

func (p *Payments) GetStatusForUpdate(ctx context.Context, paymentID uuid.UUID) (string, error) {
	query := `
		SELECT status
		WHERE id = $1::UUID
		FOR UPDATE;
	`

	var status string
	err := p.executor.GetExecutor(ctx).QueryRow(ctx, query, paymentID).Scan(&status)
	return status, err
}

func (p *Payments) SetStatus(ctx context.Context, paymentID uuid.UUID, status string) error {
	query := `
		UPDATE payments
		SET status = $2::PAYMENT_STATUS
		WHERE id = $1::UUID;
	`

	_, err := p.executor.GetExecutor(ctx).Exec(ctx, query, paymentID, status)
	return err
}

func (p *Payments) GetForUpdate(ctx context.Context, paymentID uuid.UUID) (model.Payment, error) {
	query := `
		SELECT sender_id, recipient_id, amount, status
		FROM payments
		WHERE id = $1::UUID
		FOR UPDATE;
	`

	var payment model.Payment
	err := p.executor.GetExecutor(ctx).QueryRow(ctx, query, paymentID).Scan(&payment.Sender, &payment.Recipient, &payment.Amount, &payment.Status)
	return payment, err
}
