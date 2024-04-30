package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var m, n int

	for {
		_, err := fmt.Scanln(&m, &n)
		if err != nil {
			break
		}
		fmt.Println(factorial(m) + factorial(n))
	}
}