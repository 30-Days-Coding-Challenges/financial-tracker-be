package income

import "github.com/google/uuid"

type IncomeRequest struct {
	Name           string    `json:"name" binding:"required"`
	Value          int       `json:"value" binding:"required"`
	Notes          string    `json:"notes"`
	IncomeSourceID uuid.UUID `json:"source_id" binding:"required"`
}
