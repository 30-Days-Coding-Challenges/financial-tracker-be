package incomesource

import (
	"github.com/google/uuid"
)

// import "github.com/google/uuid"

type Service interface {
	CreateIncomeSource(source IncomeSourceRequest) (IncomeSource, error)
	GetAllIncomeSource() ([]IncomeSource, error)
}

type service struct {
	repository Repository
}

func IncomeSourceService(repository Repository) *service {
	return &service{repository}
}

func (s service) CreateIncomeSource(source IncomeSourceRequest) (IncomeSource, error) {

	newUuid := uuid.Must(uuid.NewRandom())
	incomeSource := IncomeSource{
		ID:   newUuid,
		Name: source.Name,
	}

	newSource, err := s.repository.CreateSourceIncome(incomeSource)

	return newSource, err
}

func (s service) GetAllIncomeSource() ([]IncomeSource, error) {
	allIncomeSources, err := s.repository.GetAllIncomeSource()

	return allIncomeSources, err
}
