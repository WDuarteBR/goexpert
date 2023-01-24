package main

import "fmt"

// thread 1
func main() {
	canal := make(chan string) // canal vazio

	// thread 2
	go func() {
		canal <- "Finalmente channels!" // canal cheio
	}()

	// thread 1
	msg := <-canal // canal vazio, agora pode ser reutilizado
	fmt.Printf("Message: %s", msg)

}
