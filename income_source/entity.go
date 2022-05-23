package incomesource

import (
	"time"

	"github.com/google/uuid"
)

type IncomeSource struct {
	Name      string
	ID        uuid.UUID `gorm:"type:string;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type Base struct {
// 	// DeletedAt *time.Time `sql:"index"`
// }

// func (base *Base) BeforeCreate(scope *gorm.DB) error {
// 	newUuid := uuid.NewV4()
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return scope.Statement.SetColumn("ID", newUuid)
// }
