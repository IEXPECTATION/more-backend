package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

func LoginService(ctx *gin.Context) {
	user := user{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("username: %s\t", user.Name)
	fmt.Printf("userpasswd: %s\n", user.Passwd)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": "Resource created"})
}
