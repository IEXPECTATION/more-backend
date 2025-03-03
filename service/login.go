package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/database"
	"github.com/iexpectation/more/back-end/database/dao"
	"gorm.io/gorm"
)

func LoginService(ctx *gin.Context) {
	targetUser := dao.User{}
	err := ctx.ShouldBindJSON(&targetUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	db := database.Instance()
	user := dao.User{}
	result := db.Where("people_id = ? and name = ?", targetUser.PeopleId, targetUser.Name).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{"status": "failure", "message": "The information of target user is incorrect!"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": result.Error.Error()})
		}

		ctx.Abort()
		return
	}

	if !user.ValidatePassword(targetUser.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": "The information of target user is incorrect!"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": nil})
}
