package main

// this is sample main generated using CodeWhisperer
import (
	"fmt"
)

func countUntilInput(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func printListOfNames(list []string) {
	for number, name := range list {
		fmt.Println(number, name)
	}
}

func main() {
	defer fmt.Println("end run") // this is function is called after main function is finished
	countUntilInput(10)
	fmt.Println("-----")
	names := []string{"John", "Paul", "George", "Ringo"}
	printListOfNames(names)
}
