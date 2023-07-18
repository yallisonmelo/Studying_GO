package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Status bool
}

func (p Person) getInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Status: %t", p.Name, p.Age, p.Status)
}

func (p *Person) changeStatus(person *Person) {
	person.Status = !person.Status
}
