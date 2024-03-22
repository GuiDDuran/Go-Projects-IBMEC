package main

import "fmt"

func mdc(num, den int) int {
	for den != 0 {
		num, den = den, num % den
	}
	return num
}

func absolute_value(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func simplify(num, den int) (int, int) {
	if den < 0 {
        num = -num
        den = -den
    }

	common_divisor := mdc(absolute_value(num), den)

	num /= common_divisor
    den /= common_divisor

	return num, den
}

func sum (n1, d1, n2, d2 int) (int, int){
	num := n1 * d2 + n2 * d1
	den := d1 * d2
	return num, den
	
}

func sub(n1, d1, n2, d2 int) (int, int) {
    num := n1*d2 - n2*d1
    den := d1 * d2
    return num, den
}

func mult(n1, d1, n2, d2 int) (int, int) {
    num := n1 * n2
    den := d1 * d2
    return num, den
}

func div(n1, d1, n2, d2 int) (int, int) {
    num := n1 * d2
    den := n2 * d1
    return num, den
}

func main(){
	var cases int 
	fmt.Scanln(&cases)
	for i := 0; i < cases; i++{
		var n1, d1, n2, d2 int
		var op rune
		fmt.Scanf("%d / %d %c %d / %d", &n1, &d1, &op, &n2, &d2)
		var result_n, result_d int
		switch op{
		case '+':
			result_n, result_d = sum(n1, d1, n2, d2)
		case '-':
			result_n, result_d = sub(n1, d1, n2, d2)
		case '*':
			result_n, result_d = mult(n1, d1, n2, d2)
		case '/':
			result_n, result_d = div(n1, d1, n2, d2)
		}
		result_n_simplified, result_d_simplified := simplify(result_n, result_d)
		fmt.Printf("%d/%d = %d/%d\n", result_n, result_d, result_n_simplified, result_d_simplified)
	}
}