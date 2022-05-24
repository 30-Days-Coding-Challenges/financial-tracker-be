package item

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID `gorm:"type:string;primaryKey"`
	Name      string
	Value     int
	Notes     string
	SourceID  uuid.UUID `gorm:"type:string"`
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
