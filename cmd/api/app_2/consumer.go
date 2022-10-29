package app_2

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Consumer() {

	conn, err := amqp.Dial("amqp://user:password@localhost:7006/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"MyQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
			// err := gredis.Set(otpCode, d.Body, 180)
			// if err != nil {

			// }
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
