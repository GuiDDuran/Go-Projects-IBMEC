package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"projeto/geometria"
)

/* 01. Faça um programa que cria um vetor de inteiros com 10 elementos. Então inicialize este vetor com números quaisquer e imprima na tela todos os seus elementos, um abaixo do outro. */

func vetor() {
	outrosNumeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(outrosNumeros)

	for i := 0; i < len(outrosNumeros); i++ {
		fmt.Println(outrosNumeros[i])
	}

}

/* 02. Faça uma função/método que receba uma string como parâmetro e que retorne uma nova string, onde a sequência dos caracteres foi invertida. Dentro da parte principal (main), leia uma string digitada pelo usuário e passe para a função/método criada, imprimindo em seguida a string devolvida. Para esse exercício, pesquise sobre o comportamento e a interação entre dados do tipo rune e do tipo string. */

func inverteString(msg string) string{
	caracteres := []rune(msg)
	for i, j := 0, len(caracteres) - 1; i < j; i, j = i + 1, j - 1{
		caracteres[i], caracteres[j] = caracteres[j], caracteres[i]
	}
	return string(caracteres)
}

/* 03. Crie um tipo chamado Ponto que represente um ponto no plano cartesiano, com coordenadas X e Y. Em seguida, implemente um método chamado DistanciaOrigem() que calcule a distância desse ponto até a origem (0,0). */

type Ponto struct{
	X float64
	Y float64
}

func (p *Ponto) distanciaOrigem() float64{
	//d=√((x_2-x_1)²+(y_2-y_1)²)
	return math.Sqrt((p.X - 0)*(p.X - 0) + (p.Y - 0)*(p.Y - 0))
}

/* 04. Crie um pacote chamado geometria que contenha funções para calcular a área e o perímetro de um retângulo. Em seguida, use esse pacote em um programa principal para calcular a área e o perímetro de um retângulo com dimensões fornecidas pelo usuário. */

func main() {

	vetor()

	leitor := bufio.NewReader(os.Stdin)
	var msg string
	fmt.Println("Digite uma mensagem: ")
	msg, _ = leitor.ReadString('\n')
	msgInvertida := inverteString(msg)
	fmt.Println(msgInvertida)

	pontos := Ponto{X: 1, Y: 1}
	distancia := pontos.distanciaOrigem()
	fmt.Printf("%.3f\n", distancia)

	var base, altura float64
	fmt.Println("Digite a base do retângulo: ")
	fmt.Scanln(&base)
	fmt.Println("Digite a altura do retângulo: ")
	fmt.Scanln(&altura)

	area := geometria.AreaRetangulo(base, altura)
	perimetro := geometria.PerimetroRetangulo(base, altura)
	fmt.Printf("A área do retângulo é %.1f e o perímetro é %.1f\n", area, perimetro)
}