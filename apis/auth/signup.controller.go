package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
	"webSocketChatGo/database/models"
	"webSocketChatGo/utilities"

	"webSocketChatGo/database"
)

func SignUp(c *gin.Context) {
	var payload SignUpPayload
	var db = database.Connection

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, "should bind")
		return
	}

	var saveAvatarPath string
	var userAvatarName string

	if payload.Avatar == nil {
		saveAvatarPath = "static/default/user-icon"
		userAvatarName = "user-icon"
	} else {
		saveAvatarPath = "static/images/" + uuid.New().String() + ".jpg"
		userAvatarName = payload.Avatar.Filename
	}

	//TODO validation FIELDS!!!

	hashPassword := utilities.HashPassword(payload.Password)

	var user = models.Users{
		Login:     payload.Login,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Avatar:    saveAvatarPath,
		Password:  hashPassword,
	}

	result := db.Create(&user)

	refreshSecret := os.Getenv("REFRESH_SECRET")
	accessSecret := os.Getenv("ACCESS_SECRET")

	refreshToken := utilities.CreateToken(user.ID, 240*time.Hour, refreshSecret)
	accessToken := utilities.CreateToken(user.ID, 15*time.Minute, accessSecret)

	var res = Response{
		user.ID,
		user.Login,
		userAvatarName,
		payload.FirstName,
		payload.LastName,
		accessToken,
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists",
		})
	} else if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "database error",
		})
	} else {
		c.SetCookie(
			"refresh_token",
			refreshToken,
			3600,
			"/",
			"localhost",
			false,
			true)
		if payload.Avatar != nil {
			c.SaveUploadedFile(payload.Avatar, saveAvatarPath)
		}
		c.JSON(http.StatusCreated, gin.H{
			"user": res,
		})
	}
}
