package main

import "fmt"

func main() {
	//price of 1 cup of coffee
	price := 4.50 //float64

	//cups sold in a day
	quantity := 15 //int

	// total income
	// income := price * quantity //Go doesn't convert types

	income := price * float64(quantity)

	fmt.Printf("Total income during day %.2f\n", income)

	total := int(price) * quantity

	fmt.Printf("Total income during day %d\n", total)
}
