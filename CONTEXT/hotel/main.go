package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // se ultrapassar o timeout do context o Done será executado
		fmt.Println("Reserva não sucedida. Excedeu o tempo limite.")
		return

	case <-time.After(time.Second * 5): // se o context estiver ainda dentro do tempo a tarefa será executada
		fmt.Println("Reserva feita com sucesso")
		return
	}

}
