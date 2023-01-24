package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Msg from RabbitMQ"}
			c1 <- msg
		}
		// c1 <- 1
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Msg from Kafka"}
			c2 <- msg
		}
	}()

	// será executado o case em que
	// o channel primeiro receber valor

	// for i := 0; i < 2; i++ {
	for {
		select {
		case msg := <-c1:
			fmt.Printf("Msg id: %d - %s\n", msg.Id, msg.Msg)

		case msg := <-c2:
			fmt.Printf("Msg id: %d - %s\n", msg.Id, msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")

			// default:
			// 	println("defualt")
		}
	}

}
