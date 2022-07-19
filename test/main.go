package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"level0/model"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "test1")

	if err != nil {
		panic(err)
	}

	file, err := os.Open("test/model.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// for i := 1; i < 10; i++ {
	// 	sc.Publish("foo", []byte(strconv.Itoa(i)))
	// 	time.Sleep(time.Second)
	// }

	defer sc.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	var order model.WbOrder

	jsonErr := json.Unmarshal(data, &order)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(order)

	sc.Publish("foo", []byte(data))
}
