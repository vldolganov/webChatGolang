package auth

import "github.com/gin-gonic/gin"

func LogOut(c *gin.Context) {
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"localhost",
		false,
		true)
}
