package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	body := bytes.NewBuffer([]byte(`{ "nome": "wanderson" }`))
	r, err := c.Post("https://google.com.br", "applicatio/json", body)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	io.CopyBuffer(os.Stdout, r.Body, nil)

}
