package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(time.Second * 5):
		// print no console do servidor
		log.Println("Request processada com sucesso")
		// renderisa no browser
		w.Write([]byte("Request bem sucedida"))

	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
