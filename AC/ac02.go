package main

import "fmt"

/* 01. Faça um programa que cria um vetor de inteiros com 10 elementos. Então inicialize este vetor com números quaisquer e imprima na tela todos os seus elementos, um abaixo do outro. */

func Vetor() {
	outrosNumeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(outrosNumeros)

	for i := 0; i < len(outrosNumeros); i++ {
		fmt.Println(outrosNumeros[i])
	}

}

/* 02. aça uma função/método que receba uma string como parâmetro e que retorne uma nova string, onde a sequência dos caracteres foi invertida. Dentro da parte principal (main), leia uma string digitada pelo usuário e passe para a função/método criada, imprimindo em seguida a string devolvida. Para esse exercício, pesquise sobre o comportamento e a interação entre dados do tipo rune e do tipo string. */

func main() {
	Vetor()
}