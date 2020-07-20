package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID     uuid.UUID
	Symbol string
	Amount decimal.Decimal
	Type   string
}

type BuyOrders []Order
type SellOrders []Order
