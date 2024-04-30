package main

import (
  "fmt"
)

func main() {
  for {
    var digit byte
    var num string
	fmt.Scanf("%c %s\n", &digit, &num)
    count := 0

    if digit == '0' && num == "0" {
      break
    }
	
    for i := 0; i < len(num); i++ {
      if num[i] == '0' {
        if count != 0 {
          fmt.Printf("%c", num[i])
        }
      } else if num[i] != digit {
        fmt.Printf("%c", num[i])
        count = 1
      }
    }

    if count == 0 {
      fmt.Println(0)
    } else {
      fmt.Println()
    }
  }
}