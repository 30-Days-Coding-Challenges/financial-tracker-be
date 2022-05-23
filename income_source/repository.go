package incomesource

import "gorm.io/gorm"

type Repository interface {
	CreateSourceIncome(source IncomeSource) (IncomeSource, error)
	GetAllIncomeSource() ([]IncomeSource, error)
}

type repository struct {
	db *gorm.DB
}

func SourceIncomeRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) CreateSourceIncome(source IncomeSource) (IncomeSource, error) {
	err := r.db.Create(&source).Error

	return source, err
}

func (r repository) GetAllIncomeSource() ([]IncomeSource, error) {
	var incomeSources []IncomeSource

	err := r.db.Find(&incomeSources).Error

	return incomeSources, err
}
