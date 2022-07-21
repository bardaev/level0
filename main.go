package main

import (
	"encoding/json"
	"fmt"
	"level0/handler"
	"level0/model"
	"level0/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Подключаемся к БД
	dsn := "host=localhost user=go password=go dbname=go port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Автомиграция схемы
	db.AutoMigrate(&model.WbOrder{})

	// Подключение к nats streaming
	sc, err := stan.Connect("test-cluster", "test")

	if err != nil {
		panic(err)
	}

	// Создание объекта хранилища
	var storage service.StorageData = service.StorageData{
		DB:  &service.DbStorageImpl{DB: db},
		MEM: &service.MemStorageImpl{Data: make(map[uint]model.WbOrder)},
	}

	// Заполняем кэш из БД
	storage.InitMemCache()

	handler := handler.NewHandler(storage)

	// Подписываемся на канал
	// Тип канала: Durable, это даст возможность при сбое приложения не терять данные
	sub, err := sc.Subscribe("foo", func(msg *stan.Msg) {
		var wborder model.WbOrder

		// Маппим входной json на объект
		err := json.Unmarshal(msg.Data, &wborder)

		if err != err {
			log.Fatal(err)
		}

		// Сохраняем в бд и в кэш
		storage.Save(&wborder)

		fmt.Printf("Receive message %s\n", string(msg.Data))
	}, stan.DurableName("my-durable"))

	if err != nil {
		panic(err)
	}

	defer sub.Unsubscribe()
	defer sc.Close()

	router := gin.Default()

	// Rest-контроллер для получения записи по id
	router.GET("/order/:id", handler.GetOrder)
	// Rest-контроллер для получения всех записей
	router.GET("/orders", handler.GetAllOrders)

	router.Run()

}
