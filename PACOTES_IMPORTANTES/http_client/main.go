package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second}
	resp, err := client.Get("https://google.com.br")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
