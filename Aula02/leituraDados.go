package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var x, y, z int
	fmt.Print("Insira um valor inteiro a seguir: ")
	fmt.Scanln(&x)
	fmt.Println(x)

	fmt.Print("Insira dois valores inteiros a seguir: ")
	fmt.Scanln(&y, &z) // pega os valores separados por espaço
	fmt.Println(y, z)

	var nome string
	fmt.Print("Informe seu nome: ")
	fmt.Scanln(&nome)
	fmt.Println(nome)

	leitor := bufio.NewReader(os.Stdin)
	fmt.Print("Informe seu nome: ")
	nome, _ = leitor.ReadString('\n') // utiliza o enter como critério de parada
	nome = strings.ReplaceAll(nome, "\n", "")
	fmt.Println(nome)

	// Formatando dados.
	a := 4.5
	fmt.Printf("O número de a é %.3f\n", a)

	texto := fmt.Sprintf("O número de a é %.2f\n", a)
	fmt.Println(texto)
}