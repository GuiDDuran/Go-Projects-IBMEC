package main

import (
	"fmt"
)

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio  *No
	fim    *No
}

func (f *Fila) percorre() {
	no := f.inicio
	if no == nil {
		fmt.Println("Fila vazia!")
	} else {
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}

func (f *Fila) insere(novoValor int) {
	no := &No{valor: novoValor}
	if f.inicio == nil {
		f.inicio = no
		f.fim = no
	} else {
		f.fim.prox = no
		f.fim = no
	}

	f.tamanho++
}

func (f *Fila) remove() error {
	if f.inicio == nil {
		return fmt.Errorf("Fila vazia")
	}

	aux := f.inicio
	f.inicio = f.inicio.prox
	aux.prox = nil

	f.tamanho--
	return nil
}

func main() {
	fila := Fila{}

	fila.insere(1)
	fila.insere(2)
	fila.insere(3)

	fila.percorre()

	fila.remove()

	fila.percorre()

	fila.remove()

	fila.percorre()

	fila.remove()

	fila.percorre()
}