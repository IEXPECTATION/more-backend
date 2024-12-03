package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/service"
)

func Register(g *gin.Engine) {
	g.GET("/", service.HomePage)
	g.GET("/login", service.LoginService)
}
