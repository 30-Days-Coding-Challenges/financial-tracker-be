package income

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllIncome() ([]Income, error)
	GetIncomeByID(ID uuid.UUID) (Income, error)
	CreateIncome(income Income) (Income, error)
	// DeleteIncome(ID int) (Income, error)
	// UpdateIncome(ID int, income Income) (Income, error)
}

type repository struct {
	db *gorm.DB
}

func IncomeRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) GetAllIncome() ([]Income, error) {
	var incomes []Income

	err := r.db.Find(&incomes).Error

	return incomes, err
}

func (r repository) GetIncomeByID(ID uuid.UUID) (Income, error) {
	var income Income

	err := r.db.Find(&income, ID).Error

	return income, err
}

func (r repository) CreateIncome(income Income) (Income, error) {
	err := r.db.Create(&income).Error

	return income, err
}
