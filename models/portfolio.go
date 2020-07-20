package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Portfolio struct {
	UserID uuid.UUID
	Assets []Asset
	Status *string
}

type Asset struct {
	Symbol          string
	Amount          decimal.Decimal
	TransactionDate time.Time
	Status          string
}
