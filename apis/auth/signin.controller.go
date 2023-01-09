package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	"webSocketChatGo/database/models"
	"webSocketChatGo/utilities"

	"webSocketChatGo/database"
)

func SignIn(c *gin.Context) {
	var payload RequestPayload
	var db = database.Connection

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
		})
	}
	var user = models.Users{
		Login:    payload.Login,
		Password: payload.Password,
	}

	result := db.Where("login", user.Login).Find(&user)

	checkPassHash := utilities.CheckPasswordHash(payload.Password, user.Password)

	refreshSecret := os.Getenv("REFRESH_SECRET")
	accessSecret := os.Getenv("ACCESS_SECRET")
	refreshToken := utilities.CreateToken(user.ID, 240*time.Hour, refreshSecret)
	accessToken := utilities.CreateToken(user.ID, 15*time.Minute, accessSecret)

	var res = ResponseSignIn{
		user.ID,
		user.Login,
		accessToken,
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"info": "user not found",
		})
	} else if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "database error",
		})
	} else if !checkPassHash {
		c.JSON(http.StatusUnauthorized, gin.H{
			"info": "wrong password",
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

		c.JSON(http.StatusOK, gin.H{
			"user": res,
		})
	}
}
