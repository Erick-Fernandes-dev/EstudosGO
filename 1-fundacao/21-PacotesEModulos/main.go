package main

//go mod init curso-go

import (
	"curso-go/matematica"
	"fmt"
)

func main() {

	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Andar())

	fmt.Println("Resultado", s)
	fmt.Println(matematica.A)

}

// Tudo que comeca com letra maiscula e exportado, e minuscula n e exportado
