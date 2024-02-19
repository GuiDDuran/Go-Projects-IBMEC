package main

import "fmt"

// estrutura de decisão.

// if/else
func main(){

	idade := 19

	if idade > 18 {
		fmt.Println("Entrada liberada")
	} else if idade < 0{
		fmt.Println("Dado inválido")
	} else{
		fmt.Println("Entrada proibid")
	}

	// switch/case

	diaSemana := "segunda"

	switch diaSemana{
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("dia útil")
	case "sábado", "domingo":
		fmt.Println("final de semana")
	default:
		fmt.Println("dia inválido")
	}

	switch {
	case idade > 18:
		fmt.Println("Entrada liberada")
	case idade < 0:
		fmt.Println("Dado inválido")
	default:
		fmt.Println("Entrada proibida")
	}

	// Estruturas de repetição -> for

	// Comportamento do while.
	var x = 0
	for x < 5{
		fmt.Println(x)
		x++
	}

	// Comportamento do for.
	for i := 0; i < 5; i++{
		if i == 3 { continue } 
		if i == 4 { break }
		fmt.Println(i)

	}
}
