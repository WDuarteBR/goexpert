package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//criando arquivo
	arq, err := os.Create("arq.txt")
	if err != nil {
		panic(err)
	}

	//escrevendo no arquivo
	// tamanho, err := arq.WriteString("Fazendo limonada ...") //caso for inserida somente string
	tamanho, err := arq.Write([]byte("Ainda fazendo limonada ...")) //caso eu de não sabermos qual tipo de dado
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso com %d bytes \n", tamanho)

	arq.Close()

	// leitura de arquivo

	arq2, err := os.ReadFile("arq.txt") //Lê o arquivo de retorna o um []byte
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arq2))

	//lendo arquivos por n bytes
	arq3, err := os.Open("arq.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arq3)
	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	//deletando arquivo

	err = os.Remove("arq.txt")
	if err != nil {
		panic((err))
	}

	fmt.Println("Arquivo removido com sucesso")

}
