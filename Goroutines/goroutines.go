package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func printLetters() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Millisecond * 160)
	}
}

func main() {
	go printNumbers()
	go printLetters()
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
