package main

import (
	"fmt"
)

const hello = "Hello World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Erick"
	e float64 = 1.2
	f ID      = 1
)

func main() {

	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	// Ele pega o tamanho do array
	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])
	fmt.Printf("O valor de f é %T\n", f)

	for i, v := range meuArray {
		fmt.Printf("O valor de indice é %d e o valor é %d\n", i, v)
	}

	listaNome()

	//d -> digito / numero
	//OBS! Os arrays possuem tamanho fixo

}

func listaNome() {

	var nomes [4]string
	nomes[0] = "Erick"
	nomes[1] = "Jose"
	nomes[2] = "Gabriel"
	nomes[3] = "Maria"

	for i, v := range nomes {
		fmt.Printf("O valor do indice da lista de nomes é %d e o nome é %s\n", i, v)
	}

}
