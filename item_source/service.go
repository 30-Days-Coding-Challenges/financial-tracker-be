package itemsource

import (
	"github.com/google/uuid"
)

type Service interface {
	CreateItemSource(source ItemSourceRequest) (ItemSource, error)
	GetAllItemSource() ([]ItemSource, error)
}

type service struct {
	repository Repository
}

func ItemSourceService(repository Repository) *service {
	return &service{repository}
}

func (s service) CreateItemSource(source ItemSourceRequest) (ItemSource, error) {

	newUuid := uuid.Must(uuid.NewRandom())
	itemSource := ItemSource{
		ID:   newUuid,
		Name: source.Name,
		Type: source.Type,
	}

	newSource, err := s.repository.CreateItemSource(itemSource)

	return newSource, err
}

func (s service) GetAllItemSource() ([]ItemSource, error) {
	allItemSources, err := s.repository.GetAllItemSource()

	return allItemSources, err
}
