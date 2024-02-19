package main

import (
	"fmt"
	"math"
)

type Aluno struct {
	Nome string
	Matricula string
	Curso string
	Disciplinas []string
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main(){
	circ := Circulo{ Raio: 4.5}
	fmt.Println(circ.Area())


	al1 := Aluno{
		Nome: "abc",
		Matricula: "1234",
		Curso: "ADS",
		Disciplinas: []string{"Estrutura de Dados", "Desenvolvimento Web"},
	}

	fmt.Println(al1)
	fmt.Println(al1.Nome)

}