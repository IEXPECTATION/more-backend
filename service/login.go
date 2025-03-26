package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iexpectation/more/back-end/database"
	"github.com/iexpectation/more/back-end/database/dao"
	"github.com/iexpectation/more/back-end/utils"
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

		return
	}

	if !user.ValidatePassword(targetUser.Password) {
		ctx.JSON(http.StatusOK, gin.H{"status": "failure", "message": "The information of target user is incorrect!"})
		return
	}

	token, err := utils.GenerateJWT(targetUser.ID, targetUser.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": err.Error()})
		return
	}

	// TODO: Save the token into the redis
	rdb := utils.RedisInstance()
  err = rdb.Set(ctx, targetUser.Name, token, 1*time.Hour).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": nil})
}
