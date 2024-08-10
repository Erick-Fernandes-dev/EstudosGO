package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	caminhoArquivo := "/home/wolf/Documents/MinhasFormacoes/FormacaoGoLang/2-PacotesImportantes/1-ManipulacaoDeArquivos/arquivo.txt"

	// Criar um arquivo
	f, err := os.Create(caminhoArquivo)

	if err != nil {
		panic(err)
	}

	// Colocando valor dento do arquivo
	//tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	println()

	f.Close()

	// Lendo arquivo
	arquivo, err := os.ReadFile(caminhoArquivo)
	if err != nil {
		panic(err)
	}

	// Convertendo de byte para string
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open(caminhoArquivo)
	if err != nil {
		panic(err)
	}
	// Vai ler o conteúdo que está aberto no arquivo2
	reader := bufio.NewReader(arquivo2)
	// vai ler de 10 em 10 bytes
	buffer := make([]byte, 10)

	for {
		// vai pegar o reader e vai fazer a leitura baseada no tamanho do buffer
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// juntando o conteúdo no slice, vai grudar os pedaços e imprimir na tela
		fmt.Println(string(buffer[:n]))
	}

	// Removendo arquivo
	//err = os.Remove(caminhoArquivo)
	//if err != nil {
	//	panic(err)
	//}

}
