package main

import "fmt"

func main() {
	fmt.Println("Welcome!")

	for {
		var order string
		fmt.Println("Enter coffee name (or type 'exit' to quit):")
		fmt.Scanln(&order) 

		if order == "exit" {
			fmt.Println("Thank you!")
			break;
		} 
		if order == "" {
			fmt.Println("Please enter valid order")
			continue
		}
		fmt.Println("Preparing your order")
	}
	fmt.Println("Finishing program")
}
