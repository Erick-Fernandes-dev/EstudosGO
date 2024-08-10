package main

import "fmt"

func main() {

	// For com indice
	for i := 0; i < 10; i++ {
		println(i)
	}

	println()

	numeros := []string{"um", "dois", "três"}

	// For com chave valor
	for k, v := range numeros {
		fmt.Printf("chave: %d\nvalor: %s", k, v)
	}

	fmt.Println()

	//For estilo while

	indice := 0

	for indice < 10 {
		fmt.Println(indice)
		indice++
	}

	// loop infinito
	for {
		fmt.Println("Hello World!")
	}

}
