package messages

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"webSocketChatGo/database"
	"webSocketChatGo/database/models"
	"webSocketChatGo/utilities"
)

func SendMessage(c *gin.Context) {
	db := database.Connection
	var payload MsgPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
		})
	}
	getCookie, _ := c.Cookie("refresh_token")
	refreshToken := []byte(getCookie)

	checkTokenId := utilities.CheckToken(refreshToken, []byte(os.Getenv("REFRESH_SECRET")))
	fromUserId := checkTokenId.(float64)
	var message = models.Messages{
		ToUser:   payload.ToUser,
		FromUser: fromUserId,
		Text:     payload.Text,
	}

	result := db.Create(&message)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "database error",
		})
	} else {
		c.JSON(http.StatusCreated, message)
	}
}
