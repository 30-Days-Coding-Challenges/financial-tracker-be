package item

import "github.com/google/uuid"

type Service interface {
	GetAllItem() ([]Item, error)
	GetItemByID(ID uuid.UUID) (Item, error)
	CreateItem(itemRequest ItemRequest) (Item, error)
}

type service struct {
	repository Repository
}

func ItemService(repository Repository) *service {
	return &service{repository}
}

func (s service) GetAllItem() ([]Item, error) {
	items, err := s.repository.GetAllItem()

	return items, err
}

func (s service) GetItemByID(ID uuid.UUID) (Item, error) {
	item, err := s.repository.GetItemByID(ID)

	return item, err
}

func (s service) CreateItem(itemRequest ItemRequest) (Item, error) {
	item := Item{
		ID:       uuid.Must(uuid.NewRandom()),
		Name:     itemRequest.Name,
		Value:    itemRequest.Value,
		Notes:    itemRequest.Notes,
		SourceID: itemRequest.SourceID,
		Type:     itemRequest.Type,
	}

	newItem, err := s.repository.CreateItem(item)

	return newItem, err
}
