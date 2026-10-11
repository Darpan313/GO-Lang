package main

import "fmt"

func calculateLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

func updateTotalPoints(currentPoints int, newPoints int) int {
	return currentPoints + newPoints
}


func main() {
	totalPoints := 120
	newlyEarnedPoints := calculateLoyaltyPoints(9.30) // 18
	fmt.Println("Earned points today:", newlyEarnedPoints)

	totalPoints = updateTotalPoints(totalPoints, newlyEarnedPoints)

	fmt.Printf("Updated loyalty points:%d\n", totalPoints) //120 + 18 = 138
}
