package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/router/handler"
)

func RegisterAll(g *gin.Engine) {
	g.GET("/", handler.HomePageHandler)
}
