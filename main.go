package main

import (
	"encoding/json"
	"fmt"
	"level0/model"
	"level0/service"
	"log"
	"time"

	"github.com/nats-io/stan.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=go password=go dbname=go port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.WbOrder{})

	sc, err := stan.Connect("test-cluster", "test")

	if err != nil {
		panic(err)
	}

	var storage service.StorageData = service.StorageData{
		DB:  &service.DbStorageImpl{DB: db},
		MEM: &service.MemStorageImpl{},
	}

	sub, err := sc.Subscribe("foo", func(msg *stan.Msg) {
		var wborder model.WbOrder

		err := json.Unmarshal(msg.Data, &wborder)

		if err != err {
			log.Fatal(err)
		}

		storage.Save(wborder)

		fmt.Printf("Receive message %s\n", string(msg.Data))
	}, stan.DurableName("my-durable"))

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 100)

	sub.Unsubscribe()

	sc.Close()

}
