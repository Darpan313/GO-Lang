package main

import "fmt"

func main() {
	//Explicit type declaration
	var cupsQty int = 3

	// cupsQty = "four" // compile time error
	fmt.Println("Number of cups:", cupsQty)

	var wasProcessed = true

	fmt.Println("Order was processed:", wasProcessed)
}
