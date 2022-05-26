package item

import (
	"github.com/google/uuid"
)

type Service interface {
	GetAllItem() ([]Item, error)
	GetItemByID(ID string) (Item, error)
	CreateItem(itemRequest ItemRequest) (Item, error)
	DeleteItem(ID string) error
	UpdateItem(ID string, item ItemRequest) (Item, error)
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

func (s service) GetItemByID(ID string) (Item, error) {
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

func (s service) DeleteItem(ID string) error {
	err := s.repository.DeleteItem(ID)

	return err
}

func (s service) UpdateItem(ID string, itemReq ItemRequest) (Item, error) {

	item, err := s.repository.GetItemByID(ID)

	item.Name = itemReq.Name
	item.Notes = itemReq.Notes
	item.Value = itemReq.Value
	item.Type = itemReq.Type

	updatedItem, err := s.repository.UpdateItem(item)

	return updatedItem, err
}
