package model

import "gorm.io/gorm"

type ItemMaster struct {
	gorm.Model
	ItemName   string
	ItemNumber string
}
