package main

import (
	"fmt"
)

func main() {

	salarios := map[string]int{
		"Erick": 1000,
		"Jose":  4000,
		"Maria": 2500,
	}

	// Imprimindo valor com a chave
	fmt.Println(salarios["Erick"])

	//Deletando valor passando a chave
	delete(salarios, "Erick")

	//Adicionando valor
	salarios["Wes"] = 7000
	fmt.Println(salarios["Wes"])

	// A função make, vai criar aquele mapa inicial

	// preparou um map
	_ = make(map[string]int)
	//ou
	sal1 := map[string]int{}
	sal1["Wesley"] = 1000
	fmt.Println(sal1["Wesley"])

	for nome, salario := range salarios {

		fmt.Printf("O salario de %s é %d\n", nome, salario)

	}

	//Ignora os nomes com ' _ ' que se chama Blank Identifier:
	for _, salario := range salarios {

		fmt.Printf("O salario é de  %d\n", salario)

	}

}
