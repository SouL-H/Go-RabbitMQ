package main

import (
	"fmt"

	. "gorabbit/checkErr"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Hello,Rabbit!")
	conn, err := amqp.Dial("amqp://test:123321123@localhost:5672/")
	CheckErr(err)
	defer conn.Close()

	fmt.Println("Successfully")
	ch, err := conn.Channel()
	CheckErr(err)
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"test_queuee",
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
		"test_queuee",
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
