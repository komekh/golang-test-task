package app_1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"twitch_chat_analysis/cmd/api/models"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func SendMessage(c *gin.Context) {

	conn, err := amqp.Dial("amqp://user:password@localhost:7006/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	defer conn.Close()

	var message models.MessageModel

	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MyQueue",
		true,
		false,
		false,
		false,
		nil,
	)

	fmt.Println("QUEUE: ", q)

	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	msgBytes, _ := json.Marshal(message)

	// attempt to publish a message to the queue!
	err = ch.Publish(
		"",
		"MyQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msgBytes,
		},
	)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	c.JSON(200, "OK")
}
