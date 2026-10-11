package main

import "fmt"

func main() {
	customerName := "Darpan"
	orderedSize := "Large"

	if orderedSize == "Small" {
		fmt.Println(customerName, "ordered a small coffee.")
	} else {
		fmt.Println(customerName, "ordered something bigger.")
	}
}
