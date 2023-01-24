package main

import "fmt"

func main() {
	ch := make(chan string)
	go ler("Fala aí", ch)
	escrever(ch)

}

// com notação (chan<-) indica que
// o channel em questão somente recebe(receive only) a informação
func ler(palavra string, ch chan<- string) {
	ch <- palavra
}

// com a notação(<-chan ) indica que este
// channel somente envia(send only) a informação
func escrever(ch <-chan string) {
	fmt.Println(<-ch)
}
