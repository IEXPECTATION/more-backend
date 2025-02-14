package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/service"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func Register(g *gin.Engine) {
	g.Use(CORSMiddleware())
	g.GET("/", service.Test)
	g.POST("/login", service.LoginService)
	g.POST("/signup", service.SignupService)
}
