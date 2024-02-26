package main

import (
	"fmt"
	m "projeto/modelos"  // Python -> from projeto import modelos as m
	"projeto/modelos/academico"
)

func main() {
	fmt.Println("Olá, mundo")

	aluno := m.Aluno{}
	aluno.Nome = "Guilherme"
	aluno.Matricula = "1234"

	curso := academico.Curso{Nome: "Ciência de Dados"}

	fmt.Println(aluno, curso)
}