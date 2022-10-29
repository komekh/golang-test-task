package main

import (
	"twitch_chat_analysis/cmd/api/app_1"
	"twitch_chat_analysis/cmd/api/app_2"
	"twitch_chat_analysis/cmd/api/app_3"
	"twitch_chat_analysis/cmd/api/pkg/gredis"

	"github.com/gin-gonic/gin"
)

func init() {
	app_2.Consumer()
	gredis.Setup()
}

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	r.POST("/message", app_1.SendMessage)
	r.GET("/message/list", app_3.GetMessages)

	r.Run()
}
