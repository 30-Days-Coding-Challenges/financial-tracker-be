package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                       uuid.UUID `gorm:"type:string;primaryKey"`
	Name                     string
	Email                    string
	DigestUserAuth           string
	ConfirmationToken        string
	ConfirmationTokenExpired time.Time
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
