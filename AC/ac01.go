package main

import "fmt"

/* 01. Crie uma função calculaMedia que receba um número variável de parâmetros de tipo ponto flutuante e retorne a média aritmética desses valores.*/

func calculaMedia(numeros ...float64) float64 {
	soma := 0.0
	for _, num := range numeros {
		soma += num
	}
	return soma / float64(len(numeros))
}

/* 02. Construa uma função verificaParidade que receba um inteiro e retorne se o número é par ou ímpar.*/

func verificaParidade(num int) string{
	if num % 2 == 0 {
		return "par"
	} else {
		return "ímpar"
	}
}

/* 03. Desenvolva uma função minhaPotencia que receba dois inteiros, base e expoente, e retorne o resultado de base elevado ao expoente, sem usar funções prontas da linguagem.*/

func minhaPotencia(base, expoente int) int {
	resultado := 1
	for i := 0; expoente > i; i++ {
		resultado *= base
	}
	return resultado
}

/* 04. Implemente uma função converteCelsiusParaFahrenheit que receba uma temperatura em Celsius e retorne a temperatura convertida em Fahrenheit.*/

func converteCelsiusParaFahrenheit(celsius float64) float64{
	return (celsius * 9/5) + 32
}

func main() {
	media := calculaMedia(10.0, 9.5, 8.7)
	fmt.Println("A média é:", media)

	var numero int
	fmt.Print("Insira um valor inteiro a seguir: ")
	fmt.Scanln(&numero)
	fmt.Printf("O número %d é %s \n", numero, verificaParidade(numero))

	var base, expoente int
	fmt.Print("Insira dois valores inteiros a seguir: ")
	fmt.Scanln(&base, &expoente)
	fmt.Printf("A base é %d, o expoente é %d, e o resultado é %d \n", base, expoente, minhaPotencia(base, expoente))

	var temperatura float64
	fmt.Print("Insira uma temperatura em graus celsius: ")
	fmt.Scanln(&temperatura)
	fmt.Printf("A temperatura em graus celsius é %.1f e convertida para fahrenheint é %.1f \n", temperatura, converteCelsiusParaFahrenheit(temperatura))
}