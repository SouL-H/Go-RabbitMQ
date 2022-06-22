package main

import (
	"fmt"
	. "gorabbit/checkErr"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("-----Consumer-----")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	CheckErr(err)
	defer conn.Close()
	ch, err := conn.Channel()
	CheckErr(err)
	defer ch.Close()

	msgs, err := ch.Consume(
		"test_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Success")
	<-forever

}
