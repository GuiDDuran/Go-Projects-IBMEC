package main

import (
	"fmt"
)

func main(){
	var cases int
	fmt.Scan(&cases)

	for i := 0; i < cases; i++ {
		var qnt_products int
		fmt.Scan(&qnt_products)

		products := make(map[string]float64)

		for j := 0; j < qnt_products; j++ {
			var product string
			var price float64
			fmt.Scan(&product, &price)
			products[product] = price
		}

		var qnt_shopping int
		fmt.Scan(&qnt_shopping)
		total := 0.0

		for j := 0; j < qnt_shopping; j++ {
			var product string
			var qnt int
			fmt.Scan(&product, &qnt)
			total += products[product] * float64(qnt)
		}
		
		fmt.Printf("R$ %.2f\n", total)
	}
}