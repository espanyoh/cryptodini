package cryptodini_server

import (
	"cryptodini/models"
	"github.com/google/uuid"
)

type CryptodiniService interface {
	Adjust(UserID uuid.UUID, desiredPort *models.Portfolio) error
	GetPort(UserID uuid.UUID) (*models.Portfolio, error)
}

func Init(r RepoService) CryptodiniService {
	return &service{repo: r}
}

type service struct {
	repo RepoService
}
