package repository

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Payments struct {
	pool *pgxpool.Pool
}

func NewPayments(pool *pgxpool.Pool) *Payments {
	return &Payments{
		pool: pool,
	}
}

func (p *Payments) Create(
	ctx context.Context,
	userID uuid.UUID,
	payment model.PaymentRequest,
) (uuid.UUID, error) {
	query := `
		INSERT INTO payments(sender_id, recipient_id, amount)
		VALUES ($1, $2, $3)
		RETURNING id, created_at;
	`

	var (
		paymentID uuid.UUID
		createdAt time.Time
	)

	err := p.pool.QueryRow(ctx, query, userID, payment.RecipientID, payment.Amount).
		Scan(&paymentID, &createdAt)
	if err != nil {
		var pgErr *pgconn.PgError

		statusCode := http.StatusInternalServerError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			statusCode = http.StatusNotFound
		}

		return uuid.UUID{}, domain.NewAppError(statusCode, err)
	}

	return paymentID, nil
}

func (p *Payments) GetByUserID(ctx context.Context, userID uuid.UUID) ([]model.Payment, error) {
	query := `
		SELECT sender_id, recipient_id, amount, status, created_at, updated_at
		FROM payments
		WHERE recipient_id = $1 or sender_id = $1
		ORDER BY created_at DESC;
	`

	var payments []model.Payment

	rows, err := p.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, domain.NewAppError(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var payment model.Payment

		err := rows.Scan(
			&payment.SenderID,
			&payment.RecipientID,
			&payment.Amount,
			&payment.Status,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)
		if err != nil {
			return nil, domain.NewAppError(http.StatusInternalServerError, err)
		}

		payments = append(payments, payment)
	}

	err = rows.Err()
	if err != nil {
		return nil, domain.NewAppError(http.StatusInternalServerError, err)
	}

	return payments, nil
}
