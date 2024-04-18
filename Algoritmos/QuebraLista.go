package main

import "fmt"

func quebraLista(v []int, n int){
	i := int(n/2)
	for {
		fmt.Println(v[i])

		if i == 1{
			break
		}
		i = int(i/2)
	}
}

func main(){
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	n := 10
	quebraLista(v, n)
}