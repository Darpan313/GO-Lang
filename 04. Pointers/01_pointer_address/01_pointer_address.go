package main

import "fmt"

func main() {
	coffee := "Espresso"

	//pointer points to location in memory( type is *string)
	pointer := &coffee

	fmt.Println("Coffee name for coffee variable:", coffee)
	fmt.Println("Memory location:", pointer)

	fmt.Printf("Pointer address: %p\n", pointer)

	fmt.Println("----------------------")

	// value is copied from one location to another
	coffeeCopy := coffee

	fmt.Println("Coffee name for coffeeCopy variable:", coffeeCopy)
	fmt.Println("Memory location:", &coffeeCopy)

	fmt.Printf("Pointer address: %p\n", &coffeeCopy)
}