package main

import (
	"fmt"
)

const M = 10

func bubbleSort(x *[M]int) {
	limite := M
	trocou := true
	for trocou == true {
		trocou = false
		limite--
		for i := 0; i < limite; i++ {
			if x[i] > x[i+1] {
				x[i], x[i+1] = x[i+1], x[i]
				trocou = true
			}
		}
	}
}

func selectionSort(x *[M]int) {
	for i := 0; i < M-1; i++ {
		menor := i
		for j := i; j < M; j++ {
			if x[j] < x[menor] {
				menor = j
			}
		}

		if menor != i {
			x[i], x[menor] = x[menor], x[i]
		}
	}
}

func insertionSort(x *[M]int) {
	for j := 1; j < M-1; j++ {
		i := j - 1
		chave := x[j]
		for x[i] > chave && i >= 0 {
			x[i+1] = x[i]
			i--
		}
		x[i+1] = chave
	}
}

func quickSort(x *[M]int, inicio, fim int) {
	i, j := inicio, fim
	pivo := x[i]

	for i <= j {
		for x[i] < pivo {
			i++
		}
		for x[j] > pivo {
			j--
		}

		if i <= j {
			x[i], x[j] = x[j], x[i]
			i++
			j--
		}
	}

	if j > inicio {
		quickSort(x, inicio, j)
	}
	if i < fim {
		quickSort(x, i, fim)
	}
}

func combina(x *[M]int, inicio, meio, fim int, aux *[M]int) {
	i, j := inicio, meio+1
	for k := i; k <= fim; k++ {
		aux[k] = x[k]
	}

	for k := inicio; k <= fim; k++ {
		if i > meio {
			x[k] = aux[j]
			j++
		} else if j > fim {
			x[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			x[k] = aux[i]
			i++
		} else {
			x[k] = aux[j]
			j++
		}
	}
}

func mergeSort(x *[M]int, inicio, fim int, aux *[M]int) {
	meio := (inicio + fim) / 2
	mergeSort(x, inicio, meio, aux)
	mergeSort(x, meio+1, fim, aux)
	combina(x, inicio, meio, fim, aux)
}

func iniciaMergeSort(x *[M]int) {
	var aux [M]int
	mergeSort(x, 0, M-1, aux)
}

func main() {
	lista := [M]int{5, 3, 2, 4, 7, 1, 0, 6, 9, 8}

	bubbleSort(&lista)
	fmt.Println(lista)

	selectionSort(&lista)
	fmt.Println(lista)

	insertionSort(&lista)
	fmt.Println(lista)

	quickSort(&lista, 0, M-1)
	fmt.Println(lista)
}
