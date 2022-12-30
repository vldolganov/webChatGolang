package users

import (
	"fmt"
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
	getCookie, _ := c.Cookie("refresh_token")
	refreshToken := []byte(getCookie)
	userId := utilities.CheckToken(refreshToken, []byte(os.Getenv("REFRESH_SECRET")))
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "BAD REQ",
		})
	}

	fmt.Println(payload.FirstName)

	result := db.Model(&user).Where("id=?", userId).Updates(models.Users{Login: payload.Login, Email: payload.Email})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists",
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
