package cryptodini_server

import (
	"cryptodini/models"
	"errors"
	"github.com/google/uuid"
)

func (s *service) GetPort(UserID uuid.UUID) (*models.Portfolio, error) {
	user, err := s.repo.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	if user.Status != "ACTIVE" {
		return nil, errors.New("invalid user status")
	}

	// Get current port
	currentPortfolio, err := s.repo.GetPort(UserID)
	if err != nil {
		return nil, err
	}
	return currentPortfolio, nil
}
