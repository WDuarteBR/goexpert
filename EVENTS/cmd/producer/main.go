package main

import (
	"github.com/wduartebr/goexpert/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "from my app", "amq.direct")
}
