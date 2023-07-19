package main

import "fmt"

func main() {
	printers := []Printer{Epson{}, HP{}}
	for _, printer := range printers {
		fmt.Println(printer.printing())
	}
}
