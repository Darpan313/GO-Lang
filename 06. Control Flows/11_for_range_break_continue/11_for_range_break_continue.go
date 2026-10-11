package main

import "fmt"

func main() {
	drinks := []string{"Expired Milk", "Latte", "Espresso", "Matcha", "Cappucino"}

	for _, drink := range drinks {
		if drink == "Expired Milk" {
			fmt.Println("Warning! Skipping expired milk")
			continue
		}
		if drink == "Matcha" {
			fmt.Println("Matcha sold out! stopping service...")
			break
		}
		fmt.Println("Preparing drink:", drink)
	}
}
