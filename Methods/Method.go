package main

import "fmt"

func generatePersonActive() Person {
	p := Person{
		Name:   "Jack",
		Age:    33,
		Status: false,
	}
	p.changeStatus(&p)
	return p
}

func generatePersonInactive() Person {
	p := Person{
		Name:   "John",
		Age:    30,
		Status: true,
	}
	p.changeStatus(&p)
	return p
}

func main() {
	fmt.Println(generatePersonActive())
	fmt.Println("====================")
	fmt.Println(generatePersonInactive())
}
