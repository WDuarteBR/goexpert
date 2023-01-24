package main

import "fmt"

func main() {

	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		// se a operação abaixo fosse
		// feita na mesma go routine
		// onde o canal foi declarado
		// teríamos um deadlock
		forever <- true
	}()

	<-forever
}
