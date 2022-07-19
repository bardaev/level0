package model

import "gorm.io/gorm"

type WbOrder struct {
	gorm.Model
	Order_uid    string
	Track_number string
	Entry        string

	Delivery Delivery `gorm:"embedded;embeddedPrefix:delivery_"`
	Payment  Payment  `gorm:"embedded;embeddedPrefix:payment_"`
	Items    []Items  `gorm:"many2many:wborder_items"`

	Locale             string
	Internal_signature string
	Customer_id        string
	Delivery_service   string
	Shardkey           string
	Sm_id              int64
	Date_created       string
	Oof_shard          string
}
