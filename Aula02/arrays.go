package main

import "fmt"

const MAX_ARRAY = 5

func main(){
	// não se pode declarar um array como abaixo
	//x := 2
	//var numeros [x] int

	// Declaração explicita.
	var numeros [5]int // o tamanho do array precisa ser definido e ele sempre é inicializado com 0
	var nomes [MAX_ARRAY]string

	fmt.Println(numeros)
	fmt.Println(nomes)

	numeros[2] = 5
	fmt.Println(numeros)

	// Declaração curta.
	outrosNumeros := [4]int{4, 8, 1, -5}
	fmt.Println(outrosNumeros)

	// Percorrendo arrays.
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	for i, numero := range numeros { // equivalente ao enumerate
		fmt.Println(i, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

	for i := range numeros {
		fmt.Println(i)
	}

	// Slices 
	var nums []int // como não definimos um tamanho, estamos criando um slice
	nums = numeros[:3]
	fmt.Println(nums)

	maisNumeros := []float64{1.4, 5.8, 4.56}
	fmt.Println(maisNumeros)

	exemplos := make([]int, 0)
	fmt.Println(len(exemplos), cap(exemplos))
	exemplos = append(exemplos, 1)
	fmt.Println(len(exemplos), cap(exemplos))
	exemplos = append(exemplos, 1)
	fmt.Println(len(exemplos), cap(exemplos))
	exemplos = append(exemplos, 1)
	fmt.Println(len(exemplos), cap(exemplos))
	exemplos = append(exemplos, 1)
	fmt.Println(len(exemplos), cap(exemplos))
	exemplos = append(exemplos, 1)
	fmt.Println(len(exemplos), cap(exemplos)) // a capacidade do slice dobra
	exemplos = append(exemplos, 1)

	fmt.Println(exemplos)
	alteraElemento(exemplos, 10)
	fmt.Println(exemplos)
}

func alteraElemento(slice []int, valor int){
	slice[0] = valor
}