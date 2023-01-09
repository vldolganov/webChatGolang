package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"webSocketChatGo/database"
	"webSocketChatGo/database/models"
	"webSocketChatGo/utilities"
)

func EditUserInfo(c *gin.Context) {

	var payload RequestPayload
	var db = database.Connection
	var user models.Users
	var userAvatarName string

	getCookie, _ := c.Cookie("refresh_token")
	refreshToken := []byte(getCookie)
	userId := utilities.CheckToken(refreshToken, []byte(os.Getenv("REFRESH_SECRET")))
	db.Where("id", userId).Find(&user)
	avatarFromDb := user.Avatar
	saveAvatarPath := "static/default/user-icon"

	if payload.Avatar != nil {
		if avatarFromDb != payload.Avatar.Filename {
			os.Remove("./static/images/" + avatarFromDb)
			c.SaveUploadedFile(payload.Avatar, saveAvatarPath)
			userAvatarName = payload.Avatar.Filename
		}
	}
	//
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, "should bind")
		return
	}

	result := db.Model(&user).Where("id=?", userId).Updates(models.Users{Login: payload.Login, Email: payload.Email, Avatar: userAvatarName})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not update",
		})
	} else if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "database error",
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}
