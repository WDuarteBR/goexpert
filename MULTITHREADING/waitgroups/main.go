package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running.\n", i, name)
		time.Sleep(time.Second * 1)
		// informa o término de uma tarefa
		// "subtrai" uma tarefa do total de 25
		wg.Done()
	}
}

func main() {
	wg := sync.WaitGroup{}
	// foram adicionadas 25 tarefas
	// a serem esperadas
	wg.Add(25)
	// 10 tarefas aqui
	go task("A", &wg)
	// 10 tarefas aqui
	go task("B", &wg)
	// 5 tarefas aqui
	func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running.\n", i, "anonymous")
			time.Sleep(time.Second * 1)
			// informa o término de uma tarefa
			// "subtrai" uma tarefa do total de 25
			wg.Done()
		}
	}()
	// a thread principal irá
	// esperar até todas as tarefas terminarem
	wg.Wait()

}
