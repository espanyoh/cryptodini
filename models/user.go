package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID    uuid.UUID
	Username  string
	Status    string
	CreatedAt time.Time
}
