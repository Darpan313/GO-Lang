package main

import "fmt"


func processPayment(orderTotal float64, tip float64, amountPaid float64) (float64, float64) {
	totalAmoutDue := orderTotal + tip

	change := amountPaid - totalAmoutDue

	return totalAmoutDue, change
}
func main() {

	totalCost, change := processPayment(6.50, 2.00, 10.00)

	fmt.Printf("TotalCost(with tip):$%.2f Change:$%.2f\n", totalCost, change)

	// 6.50 (orderTotal) + 2.00 (tip) = 8.50 (totalAmountDue)
	// 10.00

	// 10.00 - 8.50 = 1.50 (change)

	// We need to calc totalAmountDue and change


}
