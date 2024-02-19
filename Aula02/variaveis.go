package main

import "fmt"

/* 
- Inteiros

int8, int16, int32, int64
uint8, uint16, uint32, uint64 -> unsigned, não tem números negativos
int -> tipo genérico, ocupa 32 ou 64bits, dependendo da arquitetura (padrão)
byte -> uint8
rune -> int32 -> caractere Unicode

- Ponto flutuante

float32, float64
float64 -> (padrão)
complex64, complex128

- Texto

string (cada caractere ocupa 8bits em memória)

- Booleanos

bool -> true ou false

- Outros tipos

nil
ponteiros

*/

// Declaração de variáveis globais.

var a int // declaração completa
var b = 15 // declaração implícita
//c := 10 -> não posso fazer
const PI = 3.1415 // preciso fazer declaração implícita para constantes

func main(){
	var msg1, msg2 string
	var x int
	var y = 14.5 // infere que y é uma variável do tipo float
	z := 20 // infere que z é uma variável do tipo int

	x = 18
	msg1 = "oi"
	msg2 = "olá"

	fmt.Println(x, y, z)
	fmt.Println(msg1, msg2)
}