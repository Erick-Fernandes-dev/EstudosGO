package main

import "fmt"

func main() {

	fmt.Println(sum(1, 22, 33, 44, 55, 66, 77, 88, 99, 22))
}

// Passar quantos parametros que for preciso
func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {

		total += numero
	}
	return total
}
