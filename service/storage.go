package service

import (
	"level0/model"
	"log"
)

type Storage interface {
	Save(*model.WbOrder) (uint, error)
	SaveAll(map[uint]model.WbOrder)
	GetAll() (map[uint]model.WbOrder, error)
	Get(uint) model.WbOrder
}

type StorageData struct {
	DB  Storage
	MEM Storage
}

func (s *StorageData) InitMemCache() {
	orders, err := s.DB.GetAll()

	if err != nil {
		log.Fatal(err)
		return
	}

	s.MEM.SaveAll(orders)
}

func (s *StorageData) Save(order *model.WbOrder) {
	id, err := s.DB.Save(order)
	if err != nil {
		log.Fatal(err)
		return
	}

	wbOrder := s.DB.Get(id)

	s.MEM.Save(&wbOrder)
}

func (s *StorageData) Get(id uint) model.WbOrder {
	order := s.MEM.Get(id)
	return order
}

func (s *StorageData) GetAll() map[uint]model.WbOrder {
	orders, _ := s.MEM.GetAll()
	return orders
}
