package main

import (
	"fmt"
)

func getSliceCustomer() {
	customers := []string{"John", "Jane", "Joe"}
	fmt.Println(customers)

	//Add new Customer in Slice
	customers = append(customers, "Mary")
	fmt.Println(customers)
	fmt.Println(len(customers)) //view length of slice
	fmt.Println(cap(customers)) //view capacity of slice
}

func main() {
	getSliceCustomer()
}
