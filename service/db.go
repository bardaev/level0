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
	orders := map[uint]interface{}{}
	result := ds.DB.Table("wb_orders").Take(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return ds.convertMapToWbOrder(&orders), nil
}

func (ds *DbStorageImpl) Get(id uint) model.WbOrder {
	var order model.WbOrder
	ds.DB.First(&order, id)
	return order
}

func (ds *DbStorageImpl) convertMapToWbOrder(mapOrders *map[uint]interface{}) map[uint]model.WbOrder {
	result := make(map[uint]model.WbOrder)

	for k, v := range *mapOrders {
		if obj, ok := v.(model.WbOrder); ok {
			result[k] = obj
		}
	}

	return result
}
