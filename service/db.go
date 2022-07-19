package service

import (
	"level0/model"

	"gorm.io/gorm"
)

type DbStorage interface {
	Save(model.WbOrder) error
}

type DbStorageImpl struct {
	DB *gorm.DB
}

func (ds DbStorageImpl) Save(order model.WbOrder) error {
	result := ds.DB.Create(&order)
	return result.Error
}
