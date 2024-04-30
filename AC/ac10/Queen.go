package main

import (
	"fmt"
	"math"
)

func minMoves(x1, y1, x2, y2 int) int {

	if x1 == x2 && y1 == y2 {
		return 0
	}

	if x1 == x2 || y1 == y2 || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
		return 1
	}
	return 2
}

func main() {
	for {
		var x1, y1, x2, y2 int
		fmt.Scanln(&x1, &y1, &x2, &y2)

		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		}

		fmt.Println(minMoves(x1, y1, x2, y2))
	}
}