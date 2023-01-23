package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number int64 = 0

func main() {
	// com o sync.Mutex posso controlar o acesso à variáves,
	// por multiplas threads
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()   // bloqueio o acesso

		// outra alternativa é utilizar a atomic.Add()
		// pois ele tbm implementa o Mutex internamente
		atomic.AddInt64(&number, 1)

		// m.Unlock() // libero o acesso
		time.Sleep(time.Millisecond * 300)
		w.Write([]byte(fmt.Sprintf("Você é o visitante de n° %d", number)))
	})
	http.ListenAndServe(":8000", nil)
}
