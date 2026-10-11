package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)

	// Compile time (code you write): 	var coffeePrice = 4.50
	// Runtime (what machine sees):		[some memory address] = 4.50

	// Compile time (code you write): 	fmt.Println("Coffee price", coffeePrice)
	// Runtime (what machine sees):		fmt.Println([some memory address], [mem address])
	fmt.Println("Memory address of price 4.50:", &coffeePrice)

	coffeePrice = 5.00
	fmt.Println("Memory address of price 5.00:", &coffeePrice)

	// pointerToCoffeePrice := &coffeePrice
	var pointerToCoffeePrice *float64 = &coffeePrice
	/*go to the memory location where pointerToCoffeePrice 
	points to and change value in this memory location*/
	*pointerToCoffeePrice = 5.50 

	fmt.Println("Coffee price", coffeePrice)
}
