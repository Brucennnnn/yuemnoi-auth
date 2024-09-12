package items

import "github.com/sds-2/models"

type ItemsDomain interface {
	GetAll() ([]models.Item, error)
}

type ItemsDomainImpl struct {
}

func (i ItemsDomainImpl) GetAll() ([]models.Item, error) {
	return nil, nil
}

func NewItemsDomain() *ItemsDomainImpl {
	return &ItemsDomainImpl{}
}
