package main

import (
	"fmt"
)

const hello = "Hello wold"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Erick"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e)
	fmt.Println()
	fmt.Printf("O tipo de C é %T", c)
	fmt.Println()
	fmt.Printf("O valor de f é %T", f)

}
