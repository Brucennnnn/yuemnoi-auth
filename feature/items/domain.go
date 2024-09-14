package items

import "github.com/sds-2/model"

type ItemsDomain interface {
	GetAll() ([]model.Item, error)
}

type ItemsDomainImpl struct{}

func (i ItemsDomainImpl) GetAll() ([]model.Item, error) {
	return nil, nil
}

func NewItemsDomain() *ItemsDomainImpl {
	return &ItemsDomainImpl{}
}
