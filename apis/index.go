package apis

import (
	"github.com/gin-gonic/gin"
	"webSocketChatGo/apis/auth"
	"webSocketChatGo/apis/messages"

	"webSocketChatGo/apis/users"
)

func InitRoutes(app *gin.Engine) {
	authorization := app.Group("/auth")
	{
		authorization.POST("/sign-up", auth.SignUp)
		authorization.POST("/sign-in", auth.SignIn)
		authorization.GET("/log-out", auth.LogOut)
		authorization.GET("/google-login", auth.GoogleLogin)
		authorization.GET("/google-callback", auth.GoogleCallback)
	}

	user := app.Group("/user")
	{
		user.GET("/", users.GetUserInfo)
		user.GET("/list", users.GetUserList)
		user.PUT("/edit", users.EditUserInfo)
	}

	chat := app.Group("/chat")

	{
		chat.POST("/", messages.SendMessage)
		chat.GET("/", messages.GetMessages)
	}
}
