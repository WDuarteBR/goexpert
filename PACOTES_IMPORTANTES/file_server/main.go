package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8081", mux))

}
