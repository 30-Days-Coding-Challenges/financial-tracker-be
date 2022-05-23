package income

import (
	"time"

	"github.com/google/uuid"
)

type Income struct {
	IncomeID       uuid.UUID `gorm:"type:string;primaryKey"`
	Name           string
	Value          int
	Notes          string
	IncomeSourceID int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
