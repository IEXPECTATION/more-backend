package dao

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Password   string
	Cash       float64
	DreamFund  float64
	Goose      float64
	GoldrenEgg float64
	SilverEgg  float64
}
