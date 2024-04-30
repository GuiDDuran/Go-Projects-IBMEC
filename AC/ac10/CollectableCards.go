package main

import (
	"fmt"
)

func divider(f1, f2 int) int {
	for f2 != 0 {
		f1, f2 = f2, f1%f2
	}
	return f1
}

func main() {
	var cases int
	fmt.Scanln(&cases)

	for i := 0; i < cases; i++ {
		var f1, f2 int
		fmt.Scanln(&f1, &f2)
		fmt.Println(divider(f1, f2))
	}
}