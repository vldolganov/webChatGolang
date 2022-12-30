package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"webSocketChatGo/database"
	"webSocketChatGo/database/models"
	"webSocketChatGo/utilities"
)

func GetUserInfo(c *gin.Context) {
	var db = database.Connection
	var user models.Users
	getCookie, _ := c.Cookie("refresh_token")
	refreshToken := []byte(getCookie)
	userId := utilities.CheckToken(refreshToken, []byte(os.Getenv("REFRESH_SECRET")))

	result := db.Where("id=?", userId).Find(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "database error")
	} else if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "user not found")
	} else {
		c.JSON(http.StatusOK, user)
	}
}
