package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/router"
)

func main() {
	g := gin.Default()

	router.RegisterAll(g)

	g.Run(":5000")
}
