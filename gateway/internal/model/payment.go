package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentRequest struct {
	RecipientID uuid.UUID       `json:"recipient_id" validate:"required"`
	Amount      decimal.Decimal `json:"amount"       validate:"required"`
}

type Payment struct {
	SenderID    uuid.UUID       `json:"sender_id"`
	RecipientID uuid.UUID       `json:"recipient_id"`
	Amount      decimal.Decimal `json:"amount"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
