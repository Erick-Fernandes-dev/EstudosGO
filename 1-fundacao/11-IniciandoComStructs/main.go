package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	erick := Cliente{
		Nome:  "Erick",
		Idade: 23,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, idade %d, Ativo %t", erick.Nome, erick.Idade, erick.Ativo)

	//Alterando valor
	erick.Idade = 25
	erick.Ativo = false

	fmt.Printf("\nNova Idade %d, novo ativo %t", erick.Idade, erick.Ativo)

}
