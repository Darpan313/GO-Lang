package main

import "fmt"

func main() {
	orderTotal := 15.0
	// minimumRequiredOrderAmount := 10 (error if this variable used)

	// here condition is true
	if orderTotal > 10 {
		fmt.Println("Offer discount!")
	}

	// now condition is false
	orderTotal = 7.50
	if orderTotal > 10 {
		fmt.Println("Offer discount!")
	}
}
