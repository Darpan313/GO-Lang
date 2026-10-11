package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
	fmt.Println("Dessert Menu:", dessertMenu)

	slice := dessertMenu[1:2]
	fmt.Println("Slice of Dessert Menu:", slice) // [1,2)

	slice = dessertMenu[:]
	fmt.Println("Slice of Dessert Menu:", slice) // all elements

	slice = dessertMenu[2:]
	fmt.Println("Slice of Dessert Menu:", slice) // all elements from index 2

	slice = dessertMenu[:3]
	fmt.Println("Slice of Dessert Menu:", slice) // all elements from beginning to index 3 
}
