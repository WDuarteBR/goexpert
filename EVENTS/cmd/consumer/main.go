package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wduartebr/goexpert/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // indica que após a mensagem ser lida ela não retornará para fila
	}

}
