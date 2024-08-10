package main

func main() {

	// Memória -> Endereço -> Valor

	// variavel -> ponteiro que tem um endereco na memoria -> valor
	a := 10
	var ponteiro *int = &a
	println(ponteiro)
	println(*ponteiro)
	*ponteiro = 20
	println(ponteiro)
	b := &a
	*b = 30
	println(a)

	//*b - '*' desreferenciar um ponteiro.

	println()

	c := 25
	var ponteiro2 *int = &c
	println(ponteiro2)
	d := &c
	*d = 50
	println(c)
	//println(*ponteiro2)

	//& -> vai trazer o valor de endereçamento na memória / referenciar um endereçamento na memória
	// Uma variável aponta para um ponteiro que aponta para um endereçamento na memória

}
