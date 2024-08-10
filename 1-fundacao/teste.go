package main

import (
	"fmt"
	"strconv"
)

type nome string
type idade int

var (
	name nome  = "Erick Fernandes"
	age  idade = 20
)

func main() {

	//fmt.Println("Seja bem vindo a brincadeira!!!")
	//fmt.Println(name+" ", age)

	fmt.Println(descricao(string(name), int(age), "01/01/2000", "123456"))

	pessoas()

}

func descricao(nome string, idade int, dataNascimento string, cpf string) string {

	// Converter o inteiro em uma String
	ida := strconv.Itoa(idade)

	desc := fmt.Sprintf(
		"Nome: %s\nIdade: %s\nData de nascimento: %s\nCPF: %s",
		nome,
		ida,
		dataNascimento,
		cpf,
	)

	return desc
}

func pessoas() {

	var nomes [4]string
	nomes[0] = "Erick"
	nomes[1] = "Juliana"
	nomes[2] = "Diodai"
	nomes[3] = "Gabriel"

	for _, v := range nomes {
		fmt.Println(v)
		//fmt.Printf("Indice: %d, nome: %s\n", i, v)

	}

}
