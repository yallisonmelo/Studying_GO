package main

import (
	"fmt"
)

func generateDeveloper() Developer {
	d := Developer{
		Employee{"Jack", 25, 2300, "XXXXXXXXXXXXXX"},
		"Go"}

	return d
}

func generateSeller() Seller {
	s := Seller{
		Employee{"Jonh", 23, 1800, "XXXXXXXXXXXXXX"},
		2}

	return s
}

func main() {

	fmt.Println("Developer:", generateDeveloper())
	fmt.Println("Seller:", generateSeller())
}
