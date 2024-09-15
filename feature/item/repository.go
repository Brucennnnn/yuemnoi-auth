package item

import (
	"github.com/sds-2/model"
	"gorm.io/gorm"
)

type itemRepository interface {
	GetAll() ([]model.Item, error)
}

type itemRepositoryImpl struct {
	db *gorm.DB
}

func NewitemRepository(db *gorm.DB) *itemRepositoryImpl {
	return &itemRepositoryImpl{
		db: db,
	}
}
func (i itemRepositoryImpl) GetAll() ([]model.Item, error) {
	var items []model.Item
	if err := i.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
