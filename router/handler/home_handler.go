package handler

import "github.com/gin-gonic/gin"

func HomePageHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "Ok",
	})
}
