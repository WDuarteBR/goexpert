package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome  string
	Carga int
}

func main() {

	// curso := Curso{Nome: "Go", Carga: 40}
	// forma detalhada
	// tmp := template.New("CursoTmp")
	// tmp, _ = tmp.Parse("Linguagem: {{ .Nome }} - Carga Horária: {{ .Carga }}")

	// forma reduzida (usando o método Must
	// tmp := template.Must(template.New("CursoTmp").Parse("Linguagem: {{ .Nome }} - Carga Horária: {{ .Carga }}"))
	// err := tmp.Execute(os.Stdout, curso)

	// usando arquivo externo
	type Cursos []Curso

	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	cursos := Cursos{
		{"Go", 48},
		{"Python", 36},
		{"Rust", 72},
	}
	err := tmp.Execute(os.Stdout, cursos)

	if err != nil {
		panic(err)
	}

}
