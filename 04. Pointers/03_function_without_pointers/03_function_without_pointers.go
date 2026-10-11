package main

import "fmt"


func calculatePriceAfterDiscount(price float64, discountRate float64) float64{
	return price - (price * discountRate)
}
func main() {

	var coffeePrice float64 = 5.00
	var discount float64 = 0.10
	fmt.Printf("Basic coffee price:$%.2f\n", coffeePrice)

	coffeePrice = calculatePriceAfterDiscount(coffeePrice, discount)
	fmt.Printf("Price with Discount applied:$%.2f\n", coffeePrice)

	// 5.00
	// 10%
	// 5-(5.00*0.10) = 4.50
}
