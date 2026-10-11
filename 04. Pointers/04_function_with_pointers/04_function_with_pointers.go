package main

import "fmt"

func applyDiscount(price *float64, discountRate float64) {
	fmt.Println("Address of price:", &price)
	*price = *price - (*price * discountRate)
}
func main() {
	var coffeePrice float64 = 5.00
	fmt.Println("Address of coffeePrice in main:", &coffeePrice)
	var discount float64 = 0.10
	fmt.Printf("Basic coffee price:$%.2f\n", coffeePrice)

	applyDiscount(&coffeePrice, discount)
	fmt.Printf("Price with Discount applied:$%.2f\n", coffeePrice)

	// 5.00
	// 10%
	// 5-(5.00*0.10) = 4.50
}
