package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type Conta struct {
		Numero int `json:"num"` // use `json:"-"` para omitir campo
		Saldo  int `json:"sal"` // pode-se usar validate:gt=0 para que o saldo seja maior que zero
	}

	conta := Conta{Numero: 12, Saldo: 1000}

	res, err := json.Marshal(conta) // serializando conta para json
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) // serializando e redirecionando para saída padrão
	if err != nil {
		panic(err)
	}

	var contaW Conta
	jsonPuro := []byte(`{"num":15,"sal":1500}`)

	err = json.Unmarshal(jsonPuro, &contaW) // deserializando json para conta
	if err != nil {
		panic(err)
	}

	fmt.Println(contaW)

}
