package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso": 2.50,
		"Matcha": 3.75,
		"Capuccino": 3.50,
	}

	fmt.Printf("Espresso cost:$%.2f\n", menu["Espresso"])

	menu["Espresso"] = 4.25 // change value of the key

	fmt.Printf("New Espresso cost:$%.2f\n", menu["Espresso"])
	fmt.Println("Menu items quantity:", len(menu))

	menu["Mocha"] = 4.00 // add new key value pair
	fmt.Println("Updated Menu:", menu)
	fmt.Println("Updated Menu items quantity:", len(menu))
}
