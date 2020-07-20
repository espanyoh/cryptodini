package robo_server

import (
	"cryptodini/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *service) Withdraw(userID uuid.UUID, amount decimal.Decimal) models.SellOrders {

	// Get port (from Cryptodini SDK)

	// Get list of assets based on performance (DESC)

	// Looping and remove for the list of current port until expected amount in sufficient - and create Selling order in the same time

	// if asset not enough still selling with less amount

	// if asset in portfolio is empty remove userID from Robo Server Database

	return nil
}
