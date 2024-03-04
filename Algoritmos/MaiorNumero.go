package main

import "fmt"

func MaiorNumero(){
	maior := 0
	var num int
	for {
		fmt.Scanln(&num)
		if num == 0{
			break
		}
		if num > maior{
			maior = num
		}
	}
	fmt.Println(maior)
}

func main(){
	MaiorNumero()
}