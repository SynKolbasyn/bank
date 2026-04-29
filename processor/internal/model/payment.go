package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Payment struct {
	Sender uuid.UUID
	Recipient uuid.UUID
	Amount decimal.Decimal
	Status string
}
