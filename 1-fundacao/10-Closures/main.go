package main

import "fmt"

func main() {

	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 999, 0, -8) * 2
	}()

	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
