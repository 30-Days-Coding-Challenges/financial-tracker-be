package item

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllItem() ([]Item, error)
	GetItemByID(ID uuid.UUID) (Item, error)
	CreateItem(item Item) (Item, error)
	// DeleteItem(ID int) (Item, error)
	// UpdateItem(ID int, item Item) (Item, error)
}

type repository struct {
	db *gorm.DB
}

func ItemRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) GetAllItem() ([]Item, error) {
	var items []Item

	err := r.db.Find(&items).Error

	return items, err
}

func (r repository) GetItemByID(ID uuid.UUID) (Item, error) {
	var item Item

	err := r.db.Find(&item, ID).Error

	return item, err
}

func (r repository) CreateItem(item Item) (Item, error) {
	err := r.db.Create(&item).Error

	return item, err
}
