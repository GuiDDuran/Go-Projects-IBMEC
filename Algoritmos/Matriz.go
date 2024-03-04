package main

import "fmt"

func buscaMatriz(matriz [][]int, n int, x int) bool{
	for i := 0; i < n ; i++ {
		for j := 0; j < n ; j++ {
			if matriz[i][j] == x{
				return true
			}
		}
	}
	return false
}

func verificaSoma(vetor []int, n int, soma int) (int, int){
	for i := 0; i < n - 1 ; i++ {
		for j := 0; j < n ; j++ {
			if vetor[i] + vetor[j] == soma{
				return vetor[i], vetor[j]
			}
		}
	}
	return -1, -1
}

func main(){
	matriz := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    fmt.Println(buscaMatriz(matriz, 3, 2)) 
	fmt.Println(buscaMatriz(matriz, 3, 10)) 

	vetor := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(vetor)
	soma := 12
	num1, num2 := verificaSoma(vetor, n, soma)
	fmt.Println(num1, num2)
} 