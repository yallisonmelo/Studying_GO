package main

import (
	"fmt"
)

func getMapStudentsWithAge() map[string]int {
	students := make(map[string]int)
	students["John"] = 25
	students["Mary"] = 30
	students["Bob"] = 31
	students["Alice"] = 20
	students["Mark"] = 19
	return students
}

func main() {
	students := getMapStudentsWithAge()
	fmt.Println(students)

	fmt.Println(students["John"])
	delete(students, "John")
	fmt.Println(students)
	//if student dont exist it will return 0 of primitive type
	fmt.Println(students["Joshua"])
}
