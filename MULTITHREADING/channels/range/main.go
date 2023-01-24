package main

import "fmt"

// thread 1
func main() {
	canal := make(chan int)

	// thread 2
	go publish(canal)

	// sem o close e publish,
	// após ler tudo no canal
	// será executado todos os prints
	// e no fim teremos um deadlock
	recieve(canal)
}

func recieve(ch chan int) {
	for c := range ch {
		fmt.Printf("Recieved: %d \n", c)
	}
}
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// com o método close o canal será fechado
	// "sinalizado" que não haverá mais dados
	// à serem lidos
	close(ch)
}
