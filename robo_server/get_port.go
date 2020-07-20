package robo_server

import (
	"cryptodini/models"
	"github.com/google/uuid"
)

func (s *service) GetPort(userID uuid.UUID) models.Portfolio {

	// Always get port from Cryptodini server (Single source of truth)

	return models.Portfolio{}
}
