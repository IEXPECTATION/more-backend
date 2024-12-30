package service

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/iexpectation/more/back-end/database"
	"github.com/iexpectation/more/back-end/database/dao"
)

func LoginService(ctx *gin.Context) {
	user := dao.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
	}

	name, _ := base64.StdEncoding.DecodeString(user.Name)
  password, _ := base64.StdEncoding.DecodeString(user.Password)

	fmt.Printf("username: %s\t", name)
	fmt.Printf("userpasswd: %s\n", password)

	// db := database.UseDB()
	// targetUser := dao.User{}
	// db.First(&targetUser, "name = ?", user.Name)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": "Resource created"})
}
