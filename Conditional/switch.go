package main

import "fmt"

func getViewByStatus(status int) (string, error) {
	switch status {

	case 0:
		return "This is product in preparation", nil
	case 1:
		return "This is product in stock", nil
	case 2:
		return "This is product out of stock", nil
	case 3:
		return "This is product in delivery route", nil
	case 4:
		return "This is product delivered", nil
	default:
		return "", nil

	}
}

func main() {
	result, err := getViewByStatus(0)
	if err != nil {
		fmt.Println("Problem in status")
	}
	fmt.Println(result)
}
