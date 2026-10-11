package main

import "fmt"

func main() {
	var order string = "Latte"
	fmt.Println("Initial order:", order)

	// since go is statically typed language we cannot assign variable to another type.
	order = "Cappucino"
	fmt.Println("Updated order:", order)
}
