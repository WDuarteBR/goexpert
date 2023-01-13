package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/") // faço a requição e recebo um tipo http.Response
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requição. Err_msg: %v", err)
		}

		defer req.Body.Close()

		var viacep ViaCep
		res, err := io.ReadAll(req.Body) // passo um tipo Reader e recebo []byte
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler  a resposta. Err_msg: %v", err)
		}

		err = json.Unmarshal(res, &viacep) // converto os []byte(json) para o  tipo Struct passado por ponteiro
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta. Err_msg: %v", err)
		}

		arq_cep, err := os.Create("data_cep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao salvar no arquivo. Err_msg: %v", err)
		}

		defer arq_cep.Close()

		arq_cep.WriteString(fmt.Sprintf("CEP: %s - Cidade: %s - UF: %s", viacep.Cep, viacep.Localidade, viacep.Uf))

		fmt.Println("Arquivo criado com sucesso.")

	}

}
