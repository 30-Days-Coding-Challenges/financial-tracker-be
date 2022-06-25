package itemsource

import "gorm.io/gorm"

type Repository interface {
	CreateItemSource(source ItemSource) (ItemSource, error)
	GetAllItemSource() ([]ItemSource, error)
	GetItemSourceByID(ID string) (ItemSource, error)
	DeleteItemSource(itemSource ItemSource) (ItemSource, error)
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

func (r repository) GetItemSourceByID(ID string) (ItemSource, error) {
	var itemSource ItemSource

	err := r.db.Where("id = ?", ID).First(&itemSource).Error

	return itemSource, err
}

func (r repository) DeleteItemSource(item ItemSource) (ItemSource, error) {
	err := r.db.Delete(&item).Error

	return item, err

}
