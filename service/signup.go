package service

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/database"
	"github.com/iexpectation/more/back-end/database/dao"
)

func SignupService(ctx *gin.Context) {
	user := dao.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
	}

	// TODO: Handle the error
	name, _ := base64.StdEncoding.DecodeString(user.Name)
	password, _ := base64.StdEncoding.DecodeString(user.Password)

	fmt.Printf("username: %s\t", name)
	fmt.Printf("userpasswd: %s\n", password)

	db := database.Instance()
	targetUser := dao.User{
		PeopleId: user.PeopleId,
		Name:     string(name),
		Password: string(password),
	}

	result := db.Where(dao.User{Name: string(name)}).FirstOrCreate(&targetUser)
	if result.RowsAffected == 0 {
		// The user is existed.
		fmt.Println("The user is existed!")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Sign up successfully."})
}
