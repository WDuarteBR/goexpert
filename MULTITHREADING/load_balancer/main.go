package main

import (
	"fmt"
	"time"
)

func worker(wkId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d - Received %d\n", wkId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	qtdWorker := 10000

	// inicializa os workers
	for i := 0; i < qtdWorker; i++ {
		go worker(i, ch)
	}

	// passa a carga para os workers
	// atravez do channel
	for i := 0; i < 1000000; i++ {
		ch <- i
	}

}
