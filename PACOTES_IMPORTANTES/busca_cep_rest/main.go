package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/", BuscaCepHandler) // anexo uma função à uma rota
	http.ListenAndServe(":8081", nil)     // levanta um servidor web na porta 8081
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	route := r.URL.Path
	if route != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	param := r.URL.Query().Get("cep")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	viacep, err := GetbyCep(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Forma detalhada de retornar
	// retorno, err := json.Marshal(viacep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(retorno)

	//Forma compacta (somente redireciona a resposta)
	json.NewEncoder(w).Encode(viacep)

}

func GetbyCep(cep string) (*ViaCep, error) {

	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var via ViaCep
	err = json.Unmarshal(res, &via)
	if err != nil {
		return nil, err
	}

	return &via, nil

}
