package cryptodini_server

import (
	"cryptodini/models"
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *service) Adjust(UserID uuid.UUID, desiredPort *models.Portfolio) error {

	// check userID format and desiredPort - this step normally should be handled by controller layer

	// check userID from our system (exist and valid)
	user, err := s.repo.GetUser(UserID)
	if err != nil {
		return err
	}
	if user.Status != "ACTIVE" {
		return errors.New("invalid user status")
	}

	// Get current port
	currentPortfolio, err := s.repo.GetPort(UserID)
	if err != nil {
		return err
	}

	// Calculate which one should sell/buy (order)
	orders := CalculateBuyAndSellOrders(*currentPortfolio, *desiredPort)

	// Execute order
	err = s.repo.ExecuteOrders(orders)
	if err != nil {
		return err
	}

	// Save portfolio
	err = s.repo.SavePort(*desiredPort)
	if err != nil {
		//Retry save report since order is completed
		return err
	}

	return nil
}

//FAKE Calculate
func CalculateBuyAndSellOrders(currentPortfolio models.Portfolio, desiredPort models.Portfolio) []models.Order {
	return []models.Order{
		{
			Symbol: "BTC",
			Type:   "BUY",
			Amount: decimal.NewFromFloat(0.2201),
		}, {
			Symbol: "USDT",
			Type:   "SELL",
			Amount: decimal.NewFromFloat(100.20),
		},
	}

}
