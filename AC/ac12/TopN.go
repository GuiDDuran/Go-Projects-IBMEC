package main

import (
	"fmt"
)

func main(){
	var k int
	fmt.Scan(&k)
	if (k == 1){
		fmt.Println("Top 1")
	}
	if (k > 1 && k <= 3){
		fmt.Println("Top 3")
	}
	if (k > 3 && k <= 5){
		fmt.Println("Top 5")
	}
	if (k > 5 && k <= 10){
		fmt.Println("Top 10")
	}
	if (k > 10 && k <= 25){
		fmt.Println("Top 25")
	}
	if (k > 25 && k <= 50){
		fmt.Println("Top 50")
	}
	if (k > 50 && k <= 100){
		fmt.Println("Top 100")
	}
}