package main

import "fmt"

const M = 3

func exibir(fila *[M]int, in, fim int){
	fmt.Println("Inicio: ", in, " Fim: ", fim)
	fmt.Println(fila)
}

func inserir(fila *[M]int, in, fim *int, valor int) error{
	aux := (*fim + 1) % M
	if aux == *in { return fmt.Errorf("Overflow")}

	fila[aux] = valor
	*fim = aux
	if *in == -1 {	*in = 0}
	return nil
}

func remover(fila *[M]int, in, fim *int) (int, error){
	if *in == -1 { return -1, fmt.Errorf("Underflow")}
	valorRemovido := fila[*in]
	fila[*in] = 0
	if *in == *fim{
		*in, *fim = -1, -1
	}else{
		*in = (*in + 1) % M
	}
	return valorRemovido, nil
}

func main(){
	var fila [M]int
	in, fim := -1, -1

	exibir(&fila, in, fim)
	inserir(&fila, &in, &fim, 2)
	inserir(&fila, &in, &fim, 3)
	inserir(&fila, &in, &fim, 4)
	exibir(&fila, in, fim)
	remover(&fila, &in, &fim)
	exibir(&fila, in, fim)
	inserir(&fila, &in, &fim, 5)
	err := inserir(&fila, &in, &fim, 6)
	fmt.Println(err)
}