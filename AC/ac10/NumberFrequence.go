package main

import (
	"fmt"
	"sort"
)

func main() {
	var cases int
	fmt.Scanln(&cases)

	count := make(map[int]int)

	var numbers []int

	for i := 0; i < cases; i++ {
		var x int
		fmt.Scanln(&x)
		count[x]++
		if count[x] == 1 {
			numbers = append(numbers, x)
		}
	}

	sort.Ints(numbers)

	for _, num := range numbers {
		fmt.Printf("%d aparece %d vez(es)\n", num, count[num])
	}
}