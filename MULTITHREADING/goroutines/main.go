package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running.\n", i, name)
		time.Sleep(time.Second * 1)
	}
}

// thread main, essa é a principal quando terminar
// encerrará todas as outras threads(executadas ou não)
func main() {
	// thread 2
	go task("A")
	// thread 3
	go task("B")

	// thread 4
	func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running.\n", i, "anonymous")
			time.Sleep(time.Second * 1)
		}
	}()

	// foi colocado o sleep abaixo para dar tempo de executar
	// as go routines
	time.Sleep(time.Second * 15)

}
