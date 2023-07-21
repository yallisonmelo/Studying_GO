package main

import (
	"fmt"
	"strconv"
)

func hello(name string) {
	fmt.Println("Hello: ", name, "!")
}

// this is function with first letter upper is a function to use in other packages
func Hello(name string) {
	fmt.Println("Hello: ", name, "!")
}

func sum(value1, value2 int) int {
	return value1 + value2
}

func convertAndSumSupressedError(a int, b string) int {
	converted, _ := strconv.Atoi(b)
	return a + converted
}

func convertAndSumAnReturnError(a int, b string) (total int, err error) {
	converted, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	total = a + converted
	return
}

func printNumbers(numbers ...int) {
	for _, number := range numbers {
		println(number)
	}
}

func main() {
	hello("Yallison")
	fmt.Println(sum(23, 43))
	fmt.Println("total: ", convertAndSumSupressedError(25, "90"))
	total, err := convertAndSumAnReturnError(10, "34")
	fmt.Println("total: ", total, "Error: ", err)

	printNumbers(1, 2, 3, 4, 5)
}
