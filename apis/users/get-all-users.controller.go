package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webSocketChatGo/database"
	"webSocketChatGo/database/models"
)

func GetUserList(c *gin.Context) {
	var users = []models.Users{}
	var db = database.Connection
	db.Order("id").Find(&users)

	c.JSON(http.StatusOK, users)
}
