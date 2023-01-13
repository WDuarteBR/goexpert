package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome  string
	Carga int
}

func main() {
	http.HandleFunc("/", handlerRender)
	http.ListenAndServe(":8081", nil)

}
func ToBig(s string) string {
	return strings.ToUpper(s)
}

func handlerRender(w http.ResponseWriter, r *http.Request) {
	type Cursos []Curso

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// tmp := template.Must(template.New("content.html").ParseFiles(templates...))
	tmp := template.New("content.html")
	tmp.Funcs(template.FuncMap{"Tobig": ToBig})
	tmp = template.Must(tmp.ParseFiles(templates...))
	cursos := Cursos{
		{"Go", 48},
		{"Python", 36},
		{"Rust", 72},
	}
	err := tmp.Execute(w, cursos)

	if err != nil {
		panic(err)
	}

}
