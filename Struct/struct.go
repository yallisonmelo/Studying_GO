package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	status bool // this is prorperty is private, if first letter is capital, it is public
}

type Dead struct {
	Name   string
	Status bool
	Father *Dead
}

func testStructDead() {
	w := Dead{
		Name:   "John",
		Status: true,
	}
	x := Dead{
		Name:   "Mark",
		Status: true,
		Father: &w,
	}
	fmt.Println("Dead: ", x.Name)
	fmt.Println("GrandFather: ", x.Father.Name)

}

func (p Person) getInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Status: %t", p.Name, p.Age, p.status)
}

func main() {
	y := Person{
		Name: "Mark",
		Age:  40,
	}
	y.status = true
	fmt.Println(y.getInfo())

	var p Person
	p.Name = "John"
	p.Age = 30
	p.status = true
	fmt.Println(p.getInfo())

	x := Person{"Mary", 20, false}
	fmt.Println(x.getInfo())

	fmt.Println("=========================")
	testStructDead()
}
