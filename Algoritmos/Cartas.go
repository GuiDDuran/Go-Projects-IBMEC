package main

import "fmt"

func main(){
	var cartas [5]int
	for i := 0; i < len(cartas); i++{
		fmt.Scanf("%d", &cartas[i])
	}

	crescente := true
	decrescente := true
	for i := 0; i < len(cartas) - 1; i++{
		if cartas[i] > cartas[i + 1]{
			crescente = false
		}else if cartas[i] < cartas[i + 1]{
			decrescente = false
		}
	}

	if crescente{
		fmt.Print("C\n")
	}else if decrescente{
		fmt.Print("D\n")
	}else{
		fmt.Print("N\n")
	}
}