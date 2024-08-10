package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) desativar() {

	c.Ativo = false
	fmt.Printf("O Ativo do Cliente agora é %t", c.Ativo)

}

func main() {

	cliente := Cliente{
		Nome:  "Jose",
		Idade: 21,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", cliente.Nome, cliente.Idade, cliente.Ativo)

	cliente.Ativo = false
	cliente.Endereco.Estado = "Paraíba"
	cliente.Endereco.Numero = 51
	cliente.Endereco.Cidade = "Guarabira"
	cliente.Logradouro = "Rua Francisco Pimentel da Cunha"

	fmt.Println(cliente.Ativo)

	fmt.Printf("Endereço do cliente:\nEstado: %s\nCidade: %s\nLogradouro: %s\nNumero: %d", cliente.Estado, cliente.Cidade,
		cliente.Logradouro, cliente.Numero)

	cliente.desativar()

}
