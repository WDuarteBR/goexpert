package main

import "net/http"

func main() {

	mux := http.NewServeMux()
	http.ListenAndServe(":8081", mux)
	mux.HandleFunc("/", handlerCheck)
	mux.Handle("/blog", blog{title: "Blog is Open"})

}

func handlerCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("O pai tah on"))

}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))

}
