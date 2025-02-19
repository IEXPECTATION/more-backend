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

func (user *User) Validate(other *User) bool {
	if user.Password != other.Password {
		return false
	}
	return true
}
