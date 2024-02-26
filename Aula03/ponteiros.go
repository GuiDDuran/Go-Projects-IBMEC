package main

import (
	"fmt"
	"strings"
)

func main(){
	x := 2
	y := x

	fmt.Println(x, y)

	x = 3 // Altero o valor da variável x.
	fmt.Println(x, y)

	z := &x  // Referência, indica o endereço da memória.
	fmt.Println(x, z)
	fmt.Println(x, *z)  // Dereferência.

	x = 4
	fmt.Println(x, z)
	fmt.Println(x, *z)

	msg1 := "Olá, mundo"
	fmt.Println(msg1)
	alteraMensagem(&msg1)
	fmt.Println(msg1)

	n1 := Numero{Valor: 10}
	fmt.Println(n1)
	n1.Incremento()
	fmt.Println(n1)
}

/* 
Usamos ponteiros como parâmetros de funções quando:
1. É necessário reduzir o espaço consumido em memória;
2. Queremos alterar o valor dos parâmetros.
*/
func alteraMensagem(msg *string) {
	novaMensagem := strings.ReplaceAll(*msg, "mundo", "turma")
	*msg = novaMensagem
}

type Numero struct {  // Ao trabalhar com tipos, não é necessário usar o & para fazer referência.
	Valor int
}

func (n *Numero) Incremento() {
	n.Valor++
}