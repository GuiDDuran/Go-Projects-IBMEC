package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	if x == 1 {
		return false
	}

	if x == 2 {
		return true
	}

	if x%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var cases int
	fmt.Scanln(&cases)

	for i := 0; i < cases; i++ {
		var x int
		fmt.Scanln(&x)
		if isPrime(x) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
	}
}