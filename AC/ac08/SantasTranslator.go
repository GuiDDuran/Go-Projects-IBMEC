package main

import (
    "fmt"
    "strings"
)

func main() {
    var country string

    for {
        _, err := fmt.Scan(&country)
        if err != nil {
            break
        }

        switch strings.ToLower(country) {
        case "brasil", "portugal":
            fmt.Println("Feliz Natal!")
        case "alemanha":
            fmt.Println("Frohliche Weihnachten!")
        case "austria":
            fmt.Println("Frohe Weihnacht!")
        case "coreia":
            fmt.Println("Chuk Sung Tan!")
        case "espanha", "argentina", "chile", "mexico":
            fmt.Println("Feliz Navidad!")
        case "grecia":
            fmt.Println("Kala Christougena!")
        case "estados-unidos", "inglaterra", "australia", "antardida", "canada":
            fmt.Println("Merry Christmas!")
        case "suecia":
            fmt.Println("God Jul!")
        case "turquia":
            fmt.Println("Mutlu Noeller")
        case "irlanda":
            fmt.Println("Nollaig Shona Dhuit!")
        case "belgica":
            fmt.Println("Zalig Kerstfeest!")
        case "italia", "libia":
            fmt.Println("Buon Natale!")
        case "siria", "marrocos":
            fmt.Println("Milad Mubarak!")
        case "japao":
            fmt.Println("Merii Kurisumasu!")
        default:
            fmt.Println("--- NOT FOUND ---")
        }
    }
}