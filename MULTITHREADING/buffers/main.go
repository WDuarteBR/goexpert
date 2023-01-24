package main

func main() {
	ch := make(chan string, 2)

	ch <- "Olá "
	ch <- "mundo"

	println(<-ch)
	println(<-ch)

}
