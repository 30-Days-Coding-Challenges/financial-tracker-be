package item

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllItem() ([]Item, error)
	GetItemByID(ID string) (Item, error)
	CreateItem(item Item) (Item, error)
	DeleteItem(ID string) error
	UpdateItem(item Item) (Item, error)
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

func (r repository) GetItemByID(ID string) (Item, error) {
	var item Item

	err := r.db.First(&item, "id = ?", ID).Error

	return item, err
}

func (r repository) CreateItem(item Item) (Item, error) {
	err := r.db.Create(&item).Error

	return item, err
}
func (r repository) UpdateItem(item Item) (Item, error) {
	err := r.db.Save(&item).Error

	return item, err
}
func (r repository) DeleteItem(ID string) error {
	err := r.db.Delete(&Item{}, "id = ?", ID).Error

	return err
}
