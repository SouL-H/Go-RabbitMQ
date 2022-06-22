package main

import (
	"fmt"

	"github.com/streadway/amqp"
	. "gorabbit/checkErr"
)

func main() {
	fmt.Println("Hello,Rabbit!")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	CheckErr(err)
	defer conn.Close()

	fmt.Println("Successfully")
	ch, err := conn.Channel()
	CheckErr(err)
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"test_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	CheckErr(err)
	fmt.Println(q)

	err = ch.Publish(
		"",
		"test_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello Rabbit!"),
		},
	)

	CheckErr(err)
	fmt.Println("Successfully published message")
}
