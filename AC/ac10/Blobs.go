package main

import (
	"fmt"
)

func main() {
	var cases int
	fmt.Scanln(&cases)

	for i := 0; i < cases; i++ {
		var food float64
		fmt.Scanln(&food)
		count := 0
		for food > 1 {
			food = food / 2
			count++
		}
		fmt.Printf("%d dias\n", count)
	}
}