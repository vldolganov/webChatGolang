package main

import (
	"log"
	"webSocketChatGo/apis"
	"webSocketChatGo/database"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	apis.InitRoutes(r)

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	defer server.Close()

	r.GET("/socket.io/", gin.WrapH(server))
	r.POST("/socket.io/", gin.WrapH(server))

	database.InitConnection()

	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
