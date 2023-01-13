package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()                            // gera um context vazio
	ctx, cancel := context.WithTimeout(ctx, time.Second*2) // atribui 2 segundo de duração para o context
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com.br", nil) // passando o context para a request
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(data))

}
