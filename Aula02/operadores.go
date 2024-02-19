package main

import "fmt"

func main(){
	var x = 2 
	y := 4

	// Operadores aritméticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// Operadores unários
	x++ // -> x = x + 1
	y-- // -> y = y - 1

	// Operadores de comparação
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	// Operadores lógicos
	var a, b = true, false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!b)

	// Bitwise
	// * e & -> ponteiros
}