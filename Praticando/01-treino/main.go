package main

import "fmt"

type Aluno struct {
	Nome      string
	CPF       string
	Idade     int
	Matricula string
}

type Sistema interface {
	CadastrarAluno() []Aluno
	PegarAlunos() Aluno
}

type Usuario struct {
	Nome  string
	Email string
	Senha string
}

func imprimirNome(u *Usuario) string {

	nome := u.Nome

	return nome

}

func main() {

	a := "Erick"
	var ponteiro *string = &a
	print(&a)

	fmt.Println()

	usuario := Usuario{
		Nome:  "Jose",
		Email: "gmail.com",
		Senha: "123456",
	}

	fmt.Println(imprimirNome(&usuario))

	fmt.Println(*ponteiro)

}

func calculo(a, b, c int, resultado string) int {
	fmt.Println(resultado)
	return (a + b) - c

}
