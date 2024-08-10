package main

import (
	"fmt"
)

func main() {

	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// :0 -> tudo que estiver a direita vai sumir, mostrar 0 itens
	// A capacidade ainda é 9
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	// Vai pegar os 4 primeiro itens, e dropar o restante
	// A capacidade ainda é 9
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	// Capacidade agora 7, pois os primeiro valores foram ignorados
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Função de append - adcionar
	s = append(s, 110) // Toda vez que adicionar algo no slice e caso ele não tenha essa capacidade,
	// ele vai pegar o tamanho inicial e vai dobrar a capacidade
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	// :4 -> o dois pontos antes, ele vai utilizar o que tem antes pela quantidade de vezes colocada e ignoraar o restante
	// 2: -> o dois pontos depois, vai ignorar os primeiros itens e utilizar o restante
	// O Slice sempre vai apontar para um Array, ou seja, quando um Slice é aumentado de tamanho ele vai ter um Array eumentado tmb

	//Dica: Se voce for trabalhar com slice com muitos dados no slice tente inicializar a capacidade num valor próximo daquilo que voce
	// acha que vai trabalhar, porque se voce trabalhar com valor muito diferente, e todas que tiver que redimensionar, vai ser um processo
	// muito doloroso por debaixo ali pro go.

}
