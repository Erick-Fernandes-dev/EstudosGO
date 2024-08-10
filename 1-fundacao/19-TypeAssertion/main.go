package main

import "fmt"

func main() {

	var minhaVar interface{} = "Erick Fernandes"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o valor de ok é %v", res, ok)

	// Vai dar um erro falsepanic
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)

}
