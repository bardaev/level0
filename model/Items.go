package model

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	Chrt_id      int64
	Track_number string
	Price        int64
	Rid          string
	Name         string
	Sale         int64
	Size         string
	Total_price  int64
	Nm_id        int64
	Brand        string
	Status       int64
}
