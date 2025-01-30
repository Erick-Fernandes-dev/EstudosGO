package main

import "fmt"

type Pessoa struct {
	Name string
}

func (p *Pessoa) ChangeName(name string) {
	p.Name = name
	fmt.Println("Nome alterado para: ", p.Name)

}

func main() {

	pessoa := Pessoa{Name: "João"}
	pessoa.ChangeName("Maria")
	println(pessoa.Name)

	a := 10
	var ponteiro *int = &a

	*ponteiro = 20 // alterando o valor do ponteiro
	b := &a

	fmt.Println(b)

	*b = 30

	fmt.Println(*ponteiro)

	println(ponteiro) // 0xc000058730
	println(*ponteiro)

}

// Saber como as variáveis são armazenadas na memória
