package main

import (
	"fmt"
)

func main(){
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	count := 0
	for i := 0; i < size; i++ {
		var num int
		fmt.Scan(&num)
		numbers[i] = num
		if i == 0 {
			count += 1
		} else {
			if numbers[i] == numbers[i-1] {
				continue
			} else {
				count += 1
			}
		}
	}
	fmt.Println(count)
}