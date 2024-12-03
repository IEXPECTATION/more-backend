package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/database"
	"github.com/iexpectation/more/back-end/database/dao"
	"github.com/iexpectation/more/back-end/router"
)

func main() {
	db, err := database.Init()
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&dao.User{})

	// Init the gin framework.
	g := gin.Default()

	router.Register(g)

	g.Run(":5000")
}
