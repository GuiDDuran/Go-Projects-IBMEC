package main 

import (
    "fmt"
	"math"
)

func main(){
    var n, h, c, l int

    for{
        _, err := fmt.Scan(&n)
        if err != nil { return }

        fmt.Scan(&h, &c, &l)

		aux := math.Sqrt(float64(h*h + c*c)) * float64(n)

		totalAreaMeters := aux * float64(l) / 10000
        
        fmt.Printf("%.4f\n", totalAreaMeters)
    }
}