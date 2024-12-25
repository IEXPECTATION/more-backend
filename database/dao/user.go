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

func GetUser(db *gorm.DB, name string) (u User) {
	db.First(&u, "name = ?", name)
	return
}

func (u *User) Update(db *gorm.DB) {

}
