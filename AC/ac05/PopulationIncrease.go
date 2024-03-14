package main

import "fmt"

func main(){
	var t int
	fmt.Scanln(&t)
	for i := 0; i < t; i++{
		var pa, pb int
		var  g1, g2 float32
		var anos = 0
		fmt.Scanln(&pa, &pb, &g1, &g2)
		for pa <= pb{
			pa += int(float32(pa) * (g1 / 100))
			pb += int(float32(pb) * (g2 / 100))
			anos++
		}
		if anos <= 100{
			fmt.Printf("%v anos.\n", anos)
		}else{
			fmt.Println("Mais de 1 seculo.")
		}
	}
}