package main

import "fmt"

func main() {

	a := "Rafael"

	switch a {
		case "Gabriela":
			fmt.Println("Oi Gabi")

		case "Rafael":
			fmt.Println("Oi Rafa")

		default:
			fmt.Println("Não te conheço")
	}


}

