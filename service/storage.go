package service

import (
	"level0/model"
	"log"
)

type Storage interface {
	Save(model.WbOrder)
}

type StorageImpl struct {
	DB DbStorage
}

func (s StorageImpl) Save(order model.WbOrder) {
	err := s.DB.Save(order)
	if err != nil {
		log.Fatal(err)
	}
}
