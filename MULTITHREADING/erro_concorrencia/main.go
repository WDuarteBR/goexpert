package main

import (
	"fmt"
	"net/http"
	"time"
)

var number int64 = 0

/*
Para simular várias requisições para esta aplicação
foi instalado o apache bench no linux

	instalação:
	sudo apt install -y apache2-ultis

	comando de uso:
	ab -n 10000 -c 100 http://localhost:8000/

Com o comando acima estou solicitando 10000 requisições
com dividido entre 100 clientes

Quando o comando terminar eu executo:

	curl http://localhost:8000/

Na saída vejo que o numero retornado é inferior a 10000,
desta forma é mostrado um erro de concorrência
*/
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(time.Millisecond * 300)
		w.Write([]byte(fmt.Sprintf("Você o visitante de n° %d", number)))
	})
	http.ListenAndServe(":8000", nil)
}
