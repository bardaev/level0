package service

import "level0/model"

type MemStorageImpl struct {
	data []model.WbOrder
}

func (m *MemStorageImpl) Save(order ...model.WbOrder) error {
	m.data = append(m.data, order...)
	return nil
}

func (m *MemStorageImpl) GetAll() ([]model.WbOrder, error) {
	return m.data, nil
}
