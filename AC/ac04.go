package main

import "fmt"

/* 01. Implemente em Go a solução recursiva para o problema da Torre de Hanói. */

func torreDeHanoi(n int, origem, trabalho, destino string){
	if n > 0{
		torreDeHanoi(n - 1, origem, destino, trabalho)
		fmt.Printf("Move o disco %d do topo ddo pino %s para o pino %s\n", n, origem, destino)
		torreDeHanoi(n - 1, trabalho, origem, destino)
	}
}

/* 02. Implemente em Go um algoritmo que resolva o seguinte problema: dado um array de números inteiros positivos, considerado ordenado crescentemente, e um valor alvo, encontre um par de números no array cuja soma seja igual ao valor alvo. Se nenhum par for encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado. Resolva esse problema com um algoritmo cuja complexidade é O(n). */

func encontraPar(vetor []int, alvo int) (int, int){
	soma := 0
	num1 := 0
	num2 := len(vetor) - 1
	for num1 < num2{
		soma = vetor[num1] + vetor[num2]
		if soma == alvo{
			return vetor[num1], vetor[num2]
		}else if soma > alvo{
			num2--
		}else{
			num1++
		}
	}
	return -1, -1
}

func main(){
	var discos int
	fmt.Scanln(&discos)

	torreDeHanoi(discos, "Origem", "Trabalho", "Destino")

	vetor := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(encontraPar(vetor, 6)) 
	fmt.Println(encontraPar(vetor, 20))
}