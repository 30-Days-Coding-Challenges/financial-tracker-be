package income

import "github.com/google/uuid"

type Service interface {
	GetAllIncome() ([]Income, error)
	GetIncomeByID(ID uuid.UUID) (Income, error)
	CreateIncome(incomeRequest IncomeRequest) (Income, error)
}

type service struct {
	repository Repository
}

func IncomeService(repository Repository) *service {
	return &service{repository}
}

func (s service) GetAllIncome() ([]Income, error) {
	incomes, err := s.repository.GetAllIncome()

	return incomes, err
}

func (s service) GetIncomeByID(ID uuid.UUID) (Income, error) {
	income, err := s.repository.GetIncomeByID(ID)

	return income, err
}

func (s service) CreateIncome(incomeRequest IncomeRequest) (Income, error) {
	income := Income{
		Name:     incomeRequest.Name,
		Value:    incomeRequest.Value,
		Notes:    incomeRequest.Notes,
		IncomeID: incomeRequest.IncomeSourceID,
	}

	newIncome, err := s.repository.CreateIncome(income)

	return newIncome, err
}
