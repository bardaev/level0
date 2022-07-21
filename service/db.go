package service

import (
	"level0/model"

	"gorm.io/gorm"
)

type DbStorageImpl struct {
	DB *gorm.DB
}

func (ds *DbStorageImpl) Save(order *model.WbOrder) (uint, error) {
	result := ds.DB.Create(&order)
	return order.ID, result.Error
}

func (ds *DbStorageImpl) SaveAll(orders map[uint]model.WbOrder) {
	panic("Not implement")
}

func (ds *DbStorageImpl) GetAll() (map[uint]model.WbOrder, error) {

	var orders []model.WbOrder
	result := ds.DB.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return ds.convertSliceToWbOrder(orders), nil
}

func (ds *DbStorageImpl) Get(id uint) model.WbOrder {
	var order model.WbOrder
	ds.DB.First(&order, id)
	return order
}

func (ds *DbStorageImpl) convertSliceToWbOrder(arrOrders []model.WbOrder) map[uint]model.WbOrder {
	var mapResult map[uint]model.WbOrder = make(map[uint]model.WbOrder, 0)

	for _, val := range arrOrders {
		mapResult[val.ID] = val
	}

	return mapResult
}
