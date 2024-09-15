package item

import (
	"fmt"

	"github.com/sds-2/model"
)

type ItemDomain interface {
	GetAll() ([]model.Item, error)
}

type itemDomainImpl struct {
	itemRepository
}

func (i itemDomainImpl) GetAll() ([]model.Item, error) {
	items, err := i.itemRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all items from item domain: %w", err)
	}

	return items, nil
}

func NewitemDomain(itemRepository itemRepository) *itemDomainImpl {
	return &itemDomainImpl{itemRepository: itemRepository}
}
