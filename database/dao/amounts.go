package dao

import (
	"gorm.io/gorm"
)

type Amount struct {
	gorm.Model
	UserId     uint
	Cash       float64
	DreamFund  float64
	Goose      float64
	GoldrenEgg float64
	SilverEgg  float64
}
