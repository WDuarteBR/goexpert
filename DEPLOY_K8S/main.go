package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server on!"))

	})
	http.ListenAndServe("localhost:8080", nil)
}
