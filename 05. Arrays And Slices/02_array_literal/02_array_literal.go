package main

import "fmt"

func main() {
	// type detected implicitly
	var coffeeTypes = [3]string{"Espresso", "Latte", "Cappuccino"}
	// coffeeTypes := [3]string{"Espresso", "Latte", "Cappuccino"}

	// var coffeeTypes [3]string= [3]string{"Espresso", "Latte", "Cappuccino"} // type detected explicitly

	fmt.Println("Types of coffee:", coffeeTypes)
	fmt.Println(len(coffeeTypes))

	fmt.Println("String length is:", len("This is a coffee String!"))
}
