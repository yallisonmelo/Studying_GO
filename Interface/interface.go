package main

import "fmt"

func main() {
	printers := []Printer{Epson{}, HP{}}
	for _, printer := range printers {
		switch printer.(type) {
		case Epson:
			fmt.Println("Epson printer")
		case HP:
			fmt.Println("HP printer")
		default:
			fmt.Println("Unknown printer")
		}
		fmt.Println(printer.printing())
	}
}
