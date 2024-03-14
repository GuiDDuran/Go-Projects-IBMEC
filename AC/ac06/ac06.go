package main

import "fmt"

func JuntaListas (v1, v2 []int, m, n int) []int{
	var v3 []int
	i := 0
	j := 0

	for i < m && j < n{
		if v1[i] < v2[j]{
			v3 = append(v3, v1[i])
			i++
		}else{
			v3 = append(v3, v2[j])
			j++
		}
	}

	for i < m{
		v3 = append(v3, v1[i])
		i++
	}

	for j < n{
		v3 = append(v3, v2[j])
		j++
	}

	return v3
}

func main(){
	var v1 = [7]int{2, 3, 5, 7, 9, 15, 17}
	var v2 = [7]int{1, 4, 6, 11, 12, 16, 22}
	m := len(v1)
	n := len(v2)
	v3 := JuntaListas(v1[:], v2[:], m, n) 
	fmt.Println(v3)
}