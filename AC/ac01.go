package main

import "fmt"

/* 01. Crie uma função calculaMedia que receba um número variável de parâmetros de tipo ponto flutuante e retorne a média aritmética desses valores.*/

func calculaMedia(numeros ...float64) float64{
	total := 0.0
	for _, num := range numeros{
		total += num
	}
	return total / float64(len(numeros))
}

/* 02. Construa uma função verificaParidade que receba um inteiro e retorne se o número é par ou ímpar.*/

func verificaParidade(a int){
	if a % 2 == 0 {
		fmt.Println(a, "é um número par.")
	} else {
		fmt.Println(a, "é um número ímpar.")
	}
}

/* 03. Desenvolva uma função minhaPotencia que receba dois inteiros, base e expoente, e retorne o resultado de base elevado ao expoente, sem usar funções prontas da linguagem.*/

func minhaPotencia(base, expoente int) int{
	resultado := 1
	for i := 0; expoente > i; i++{
		resultado *= base 
	}
	return resultado
}

/* 04. Implemente uma função converteCelsiusParaFahrenheit que receba uma temperatura em Celsius e retorne a temperatura convertida em Fahrenheit.*/

func converteCelsiusParaFahrenheit(){
	
}

func main(){
	verificaParidade(3)
	fmt.Println(minhaPotencia(2, 3))
	fmt.Println(calculaMedia(10.0, 9.5, 8.7))
}