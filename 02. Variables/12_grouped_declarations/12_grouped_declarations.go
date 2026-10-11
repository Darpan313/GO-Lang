package main

import "fmt"

func main() {
	var coffeeType string = "Latte"
	var quantity int = 3
	var unitPrice float64 = 4.25

	fmt.Printf("Ordered %d %s priced $%.2f each \n", coffeeType, quantity, unitPrice)

	// In Go there's possibility to group variables
	var (
		customerName string = "Darpan"
		tableNumber  int    = 8
		isReadyToPay bool   = false
	)

	fmt.Printf("Customer %s at table %d is ready to pay: %t\n", customerName, tableNumber, isReadyToPay)

	// No compilation error if const defined but not used, Go don't allocate memory until used for var it does allocate memory
	const (
		SizeSmall  = "S"
		SizeMeduim = "M"
		SizeLarge  = "L"
	)

}
