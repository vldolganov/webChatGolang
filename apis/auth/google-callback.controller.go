package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"webSocketChatGo/database/models"
	"webSocketChatGo/utilities"

	"webSocketChatGo/config"
	"webSocketChatGo/database"
)

func GoogleCallback(c *gin.Context) {

	var db = database.Connection
	if c.Query("state") != "random" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "error state",
		})
	}

	token, err := config.SetupConfig().Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "token state",
		})
	}

	url := os.Getenv("GOOGLE_APIS") + token.AccessToken

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "USER NOT FOUND",
		})
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	var response GooglePayload
	json.Unmarshal(content, &response)

	var googleRes = Response{
		Login:       response.Email,
		Avatar:      response.Picture,
		FirstName:   response.GivenName,
		LastName:    response.FamilyName,
		AccessToken: token.AccessToken,
	}

	var user = models.Users{
		Email:     response.Email,
		FirstName: response.GivenName,
		LastName:  response.FamilyName,
		Avatar:    response.Picture,
	}

	refreshSecret := os.Getenv("REFRESH_SECRET")
	refreshToken := utilities.CreateToken(user.ID, 240*time.Hour, refreshSecret)

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "db error",
		})
	} else if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "user already exists",
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
		
		c.JSON(http.StatusCreated, gin.H{
			"user": googleRes,
		})
	}
}
