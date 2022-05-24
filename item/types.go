package item

import "github.com/google/uuid"

type ItemRequest struct {
	Name     string    `json:"name" binding:"required"`
	Value    int       `json:"value" binding:"required"`
	Notes    string    `json:"notes"`
	Type     string    `json:"type" binding:"required"`
	SourceID uuid.UUID `json:"source_id" binding:"required"`
}
