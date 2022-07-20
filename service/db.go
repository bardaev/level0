package service

import (
	"level0/model"

	"gorm.io/gorm"
)

type DbStorageImpl struct {
	DB *gorm.DB
}

func (ds *DbStorageImpl) Save(order ...model.WbOrder) error {
	result := ds.DB.Create(&order)
	return result.Error
}

func (ds *DbStorageImpl) GetAll() ([]model.WbOrder, error) {
	var orders []model.WbOrder
	result := ds.DB.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
