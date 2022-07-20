package service

import (
	"level0/model"
	"log"
)

type Storage interface {
	Save(...model.WbOrder) error
	GetAll() ([]model.WbOrder, error)
}

type StorageData struct {
	DB  Storage
	MEM Storage
}

func (s *StorageData) InitMemCache() {
	orders, err := s.DB.GetAll()

	if err != nil {
		log.Fatal(err)
	}

	s.MEM.Save(orders...)
}

func (s *StorageData) Save(order model.WbOrder) {
	err := s.DB.Save(order)
	if err != nil {
		log.Fatal(err)
	}
}
