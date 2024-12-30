package database

import (
	"github.com/iexpectation/more/back-end/database/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (*gorm.DB, error) {
	// Init the mariadb.
	url := "more:more@tcp(127.0.0.1:3306)/moredb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		return db, err
	}

	// Create a new people table for all basic users.
	db.AutoMigrate(&dao.People{})
	db.FirstOrCreate(&dao.People{Name: "BasicPeople"})

	db.AutoMigrate(&dao.User{})

	return db, nil
}

func UseDB() *gorm.DB {
	return db
}
