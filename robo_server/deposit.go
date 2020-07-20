package robo_server

import (
	"cryptodini/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *service) Deposit(userID uuid.UUID, amount decimal.Decimal) models.BuyOrders {

	// If userID not not our Robo Server - Save them

	// Get Best asset list - it could be multiple based on strategy , assume 3

	// Create Buy order and return

	return nil
}
