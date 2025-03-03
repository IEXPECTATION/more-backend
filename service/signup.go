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

func SignupService(ctx *gin.Context) {
	targetUser := dao.User{}
	err := ctx.ShouldBindJSON(&targetUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	db := database.Instance()
	user := dao.User{}

	result := db.Where("people_id = ? and name = ?", targetUser.PeopleId, targetUser.Name).First(&user)
	if result.Error != nil {
		fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Add the target user to db
			targetUser.HashPassword()
			if db.Create(&targetUser).Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": "failed to create the new user!"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "signup successful!"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": result.Error.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "failure", "message": "The target user has already signed up!"})
}
