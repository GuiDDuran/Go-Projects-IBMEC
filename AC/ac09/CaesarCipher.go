package main

import "fmt"

func main() {
	var key, n_cases int
	var text string

	fmt.Scanln(&n_cases)

	for i := 0; i < n_cases; i++{
		fmt.Scanln(&text)
		fmt.Scanln(&key)
		result := ""
		for _, char := range text {
			if char == 32 {
				result += " "
			} else {
				decrypted := ((int(char) - 'A' - key + 26) % 26) + 'A'
				result += string(decrypted)
			}
		}
		fmt.Println(result)
	}
}