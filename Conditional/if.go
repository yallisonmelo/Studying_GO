package main

import (
	"fmt"
	"log"
	"os"
)

func verifyTimeExpirationDriverLicense(age int) {
	if age > 18 {
		fmt.Println("You driver license expire in 5 years")
	} else if age > 60 {
		fmt.Println("You driver license expire in 3 years")
	} else {
		fmt.Println("you don't have a driver's license")
	}
}

func readFile(file string) {
	content, err := os.Open(file)
	if err != nil {
		log.Panic("Error:", err)
	}
	data := make([]byte, 100)
	if _, err := content.Read(data); err != nil {
		log.Panic("Error:", err)
	}
	fmt.Println("Content in file: ", string(data))
}

func main() {
	verifyTimeExpirationDriverLicense(20)
	readFile("hello.txt")
}
