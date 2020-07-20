package cryptodini_server

import (
	"cryptodini/models"
	"github.com/google/uuid"
)

// Simulate whether call to another microservices or database
type RepoService interface {
	GetUser(userID uuid.UUID) (*models.User, error)
	GetPort(userID uuid.UUID) (*models.Portfolio, error)
	ExecuteOrders([]models.Order) error
	SavePort(models.Portfolio) error
}
