package main

import (
	"fmt"
	"sync"
)

// thread 1
func main() {
	canal := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(10)
	// thread 2
	go publish(canal)
	// thread 3
	go recieve(canal, &wg)

	// a thread 1 vai aguardar as
	// outras 2 threads por conta do
	// waitgroups
	wg.Wait()
}

func recieve(ch chan int, wg *sync.WaitGroup) {
	for c := range ch {
		fmt.Printf("Recieved: %d \n", c)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
