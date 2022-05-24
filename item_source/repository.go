package itemsource

import "gorm.io/gorm"

type Repository interface {
	CreateItemSource(source ItemSource) (ItemSource, error)
	GetAllItemSource() ([]ItemSource, error)
}

type repository struct {
	db *gorm.DB
}

func ItemSourceRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) CreateItemSource(source ItemSource) (ItemSource, error) {
	err := r.db.Create(&source).Error

	return source, err
}

func (r repository) GetAllItemSource() ([]ItemSource, error) {
	var itemSources []ItemSource

	err := r.db.Find(&itemSources).Error

	return itemSources, err
}
