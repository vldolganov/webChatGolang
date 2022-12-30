package auth

import (
	"webSocketChatGo/config"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GoogleLogin(c *gin.Context) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("random")
	c.Redirect(http.StatusPermanentRedirect, url)
}