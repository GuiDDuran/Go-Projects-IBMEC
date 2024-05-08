package main

import (
	"fmt"
)

func main (){
	var competitors, papers, qnt_papers int
	fmt.Scan(&competitors, &papers, &qnt_papers)
	if (papers >= competitors * qnt_papers) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}