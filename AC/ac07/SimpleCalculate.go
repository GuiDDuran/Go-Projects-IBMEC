package main

import "fmt"

func main(){
	var code1, quant1, code2, quant2 int
	var price1, price2, total float32
	fmt.Scanln(&code1, &quant1, &price1)
	fmt.Scanln(&code2, &quant2, &price2)
	total = float32(quant1) * price1 + float32(quant2) * price2
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}