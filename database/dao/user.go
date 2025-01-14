package dao

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PeopleId uint
	Name     string
	Password string
}
