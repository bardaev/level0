package main

import (
	"encoding/json"
	"io/ioutil"
	"level0/model"
	"level0/service"
	"os"
	"testing"
)

func TestMemStorageSave(t *testing.T) {
	file, err := os.Open("model.json")
	if err != nil {
		t.Error("Cannot open file")
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		t.Error("Cant read data")
	}

	var wbOrder model.WbOrder
	jsonErr := json.Unmarshal(data, &wbOrder)

	if jsonErr != nil {
		t.Error("Not valid json")
	}

	mem := service.MemStorageImpl{Data: make(map[uint]model.WbOrder)}
	id, _ := mem.Save(&wbOrder)

	got := mem.Data[id].ID
	want := id

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMemStorageGet(t *testing.T) {
	file, err := os.Open("model.json")
	if err != nil {
		t.Error("Cannot open file")
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		t.Error("Cant read data")
	}

	var wbOrder model.WbOrder
	jsonErr := json.Unmarshal(data, &wbOrder)

	if jsonErr != nil {
		t.Error("Not valid json")
	}

	mem := service.MemStorageImpl{Data: make(map[uint]model.WbOrder)}
	id, _ := mem.Save(&wbOrder)

	got := mem.Get(id)
	want := id

	if got.ID != want {
		t.Errorf("got %q, want %q", got.ID, want)
	}
}
