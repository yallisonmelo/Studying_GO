package main

import "fmt"

func main() {

	//Generic Variable
	var name = "Name"
	fmt.Println(name)

	//int variables and sum
	var num1, num2 int = 10, 20
	fmt.Println(num1 + num2)

	// typed variable
	var is_multiply bool = false
	//Conditional
	if is_multiply {
		var result = num1 * num2
		fmt.Printf("result multiplication: %d", result)
	} else {
		fmt.Println(num2 - num1)
		fmt.Printf("result subtraction: %d", num2-num1)
		fmt.Println("")
	}

	//short declaration variable
	mail := "mail@gmail.com"
	fmt.Println(mail)

	//Multiple definition
	var a, b, c int32
	a = 2
	b = 3
	c = 4
	fmt.Println(a, b, c)

	//Multiple attribuiton
	var user, active, age = "yallison", true, 31
	fmt.Println("The user: ", user)
	fmt.Println("The status: ", active)
	fmt.Println("The age: ", age)

}
