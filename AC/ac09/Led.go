package main

import "fmt"

func main() {
	var n_cases int
	fmt.Scanln(&n_cases)
	for i := 0; i < n_cases; i++{
		var num string
		fmt.Scanln(&num)
		leds := 0
		for j := 0; j < len(num); j++ {
			digit := int(num[j] - '0')

			switch digit {
			case 1:
				leds += 2
			case 2, 3, 5:
				leds += 5
			case 4:
				leds += 4
			case 6, 9, 0:
				leds += 6
			case 7:
				leds += 3
			case 8:
				leds += 7
			}
		}
		fmt.Println(leds, "leds")
	}
}