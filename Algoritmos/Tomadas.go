package main

import "fmt"

func main(){
	var t1, t2, t3, t4 int
	fmt.Scanln(&t1, &t2, &t3, &t4)
	total := (t1 + t2 + t3 + t4) - 3
	fmt.Print(total)
}