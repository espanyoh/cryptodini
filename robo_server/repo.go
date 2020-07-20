package robo_server

import (
	"cryptodini/models"
)

// Simulate to list the asset performance for our service
type RepoService interface {
	GetBestAssets() ([]models.Asset, error)
	ListWorstAssets() ([]models.Asset, error)
}
