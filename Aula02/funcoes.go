package main

import "fmt"

// func 'nome'(parametros tipo) tipo

// função com letra maiúscula significa que ela é pública.
func soma(a, b int) int {
	return a + b
}

func informalidade(nome string, idade int){ // essa função não possui nenhum tipo de retorno
	fmt.Println(nome, "tem", idade, "anos")
}

func trocaValores(x, y float64) (float64, float64){
	return y, x
}

func anonima(){
	dobra := func(x int) int{
		return x * 2
	}
	fmt.Println(dobra(5))

	calcular := func(f func(int) int, x int) int {
		return f(x)
	}
	fmt.Println(calcular(dobra, 8))
}

func main(){
	fmt.Println(soma(4, 8))
	informalidade("Guilherme", 25)

	a, b := 4.4, 8.5
	fmt.Println(trocaValores(a, b))
}