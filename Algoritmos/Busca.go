package main

import "fmt"

func Busca1(v []int, n int, x int) int{
	i := 0
	posicao := -1
	for i < n{
		if v[i] == x{
			posicao = i
			break;
		}else{
			i = i + 1
		}
	}
	return posicao
}

func Busca2(v []int, n int, x int) int{
	for i:=0; i < n; i++{
		if v[i] == x{
			return i
		}
	}
	return -1
}

func main(){
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	n := 10
	x := 5
	fmt.Println(Busca1(v, n, x))
	fmt.Println(Busca2(v, n, x))
}