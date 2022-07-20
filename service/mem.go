package service

import (
	"level0/model"
)

type MemStorageImpl struct {
	Data map[uint]model.WbOrder
}

func (m *MemStorageImpl) Save(order *model.WbOrder) (uint, error) {
	m.Data[order.ID] = *order
	return order.ID, nil
}

func (m *MemStorageImpl) SaveAll(order map[uint]model.WbOrder) {
	m.Data = order
}

func (m *MemStorageImpl) GetAll() (map[uint]model.WbOrder, error) {
	return m.Data, nil
}

func (m *MemStorageImpl) Get(id uint) model.WbOrder {
	return m.Data[id]
}
