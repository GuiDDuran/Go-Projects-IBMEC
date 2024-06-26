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
	// topo    *No
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

// func (f *Fila) percorre() {
// 	if f.topo == nil {
// 		fmt.Println("Fila vazia")
// 	} else {
// 		no := f.topo
// 		for no.prox != nil {
// 			fmt.Print(no.valor, " -> ")
// 			no = no.prox
// 		}
// 		fmt.Println(no.valor)
// 	}
// }

// func (f *Fila) insere(novoValor int) {
// 	novoNo := &No{valor: novoValor}

// 	if f.topo == nil {
// 		f.topo = novoNo
// 	} else {
// 		aux := f.topo
// 		for aux.prox != nil {
// 			aux = aux.prox
// 		}
// 		aux.prox = novoNo
// 	}

// 	f.tamanho++
// }

// func (f *Fila) remove() error {
// 	if f.topo == nil {
// 		return fmt.Errorf("Fila vazia")
// 	}

// 	aux := f.topo
// 	f.topo = f.topo.prox
// 	aux.prox = nil

// 	f.tamanho--
// 	return nil
// }

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