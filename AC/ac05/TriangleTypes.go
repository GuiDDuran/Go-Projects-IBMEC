package main

import "fmt"

func main() {
	var a, b, c float32
	fmt.Scanln(&a, &b, &c)
	
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}

	if a >= b + c {
		fmt.Println("NAO FORMA TRIANGULO")
	} else if a * a == (b * b) + (c * c) {
		fmt.Println("TRIANGULO RETANGULO")
	} else if a * a > (b * b) + (c * c) {
		fmt.Println("TRIANGULO OBTUSANGULO")
	} else if a * a < (b * b) + (c * c) {
		fmt.Println("TRIANGULO ACUTANGULO")
	}
	
	if a == b && a == c {
		fmt.Println("TRIANGULO EQUILATERO")
	} else if a == b || a == c || b == c {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}
