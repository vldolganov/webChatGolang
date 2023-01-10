package messages

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"webSocketChatGo/database"
	"webSocketChatGo/database/models"
	"webSocketChatGo/utilities"
)

func GetMessages(c *gin.Context) {
	var db = database.Connection
	var payload MsgPayload
	var messages models.Messages
	getCookie, _ := c.Cookie("refresh_token")
	refreshToken := []byte(getCookie)
	userId := utilities.CheckToken(refreshToken, []byte(os.Getenv("REFRESH_SECRET")))

	result := db.Table("messages").Where("from_user=?", userId).Where("to_user", payload.ToUser).Find(&messages)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "database error")
	} else if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "user not found")
	} else {
		c.JSON(http.StatusOK, messages)
	}
}
