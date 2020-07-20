package robo_server

import (
	"cryptodini/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AssetManagerService interface {
	Deposit(userID uuid.UUID, amount decimal.Decimal) models.BuyOrders
	Withdraw(userID uuid.UUID, amount decimal.Decimal) models.SellOrders
	GetPort(userID uuid.UUID) models.Portfolio
}

func Init(r RepoService) AssetManagerService {
	return &service{repo: r}
}

type service struct {
	repo RepoService
}
