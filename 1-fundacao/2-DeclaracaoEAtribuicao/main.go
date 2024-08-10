package main

// Definindo tipos de variáveis

// valor imutável
const a = "Hello world"

// Declarando de variável global
// por padrão a var b ela recebe um valor false
var (
	b bool   // false
	c int    // 0
	d string // vazio
	e float64
)

func main() {
	// Variável de escopo local
	var k string

	a := "X"
	a = "Z"

	println(k + " oi " + a)

	b = true
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)

}
