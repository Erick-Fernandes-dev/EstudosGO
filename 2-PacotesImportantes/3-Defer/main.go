package main

import "fmt"

func main() {

	fmt.Println("Primeira linha")
	// vai rodar no final
	defer fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")

}
