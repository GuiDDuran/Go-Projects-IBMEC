package main

import "fmt"

func main() {
	var n int
	var l, q float32

	fmt.Scanln(&n, &l, &q)

	names := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&names[i])
	}

	i := 0
	for l - q > 0{
		l -= q
		i++
	}
	i %= n
	fmt.Printf("%s %.1f\n", names[i], l)
}
