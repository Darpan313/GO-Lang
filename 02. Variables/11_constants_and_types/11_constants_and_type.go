package main

import "fmt"

func main() {
	// when declared variable using const Go treats as untyped constant
	// In Go constant can adapt to  other types
	const rewardPoints = 10

	fmt.Printf("Default type of revwardPoints is %T\n", rewardPoints) //int

	var totalRewardPoints float64 = 150.3

	// Adding untyped constant to a float64 - valid, constants adapts (it's not conversion of type, it's adaptation)
	// 10 -> 10.0
	totalRewardPoints = totalRewardPoints + rewardPoints

	fmt.Printf("Updated loyalty points %.2f\n", totalRewardPoints)

}
