package main

type MyNumber int

// Constraints - Sao tipos especificos que voce cria para serem substituidos aqui nessa tipagem
type Number interface {
	~int | ~float64
}

// ~ - significa que considerar que vai usar qualquer tipo que possui o int por debaixo dos panos

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

//func Soma[T int | float64](m map[string]T) T {
//	var soma T
//	for _, v := range m {
//		soma += v
//	}
//
//	return soma
//}

//func SomaFloat(m map[string]float64) float64 {
//	var soma float64
//	for _, v := range m {
//		soma += v
//	}
//
//	return soma
//}

// Usando constraint ja pronta
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Erick": 1000, "Jose": 2000, "Maria": 3000}
	m2 := map[string]float64{"Erick": 100.20, "Jose": 2000.3, "Maria": 300.0}

	m3 := map[string]MyNumber{"Erick": 1000, "Jose": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))

}
