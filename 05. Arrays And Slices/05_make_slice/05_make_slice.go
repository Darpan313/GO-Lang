package main

import "fmt"

func main() {
	// we are using here slice literal
	ratings := []int{5, 4, 5, 5, 3}
	fmt.Println("Original ratings:", ratings)

	ratings[2] = 3
	fmt.Println("Changed ratings:", ratings)

	fmt.Println("Length of the slice is:", len(ratings))

	coffeeTypes := make([]string, 3)
	fmt.Println("CoffeeTypes:", coffeeTypes)

	coffeeTypes[0] = "Cappuccino"
	coffeeTypes[1] = "Latte"
	coffeeTypes[2] = "Espresso"

	fmt.Println("Coffee type after assignment:", coffeeTypes)

}
