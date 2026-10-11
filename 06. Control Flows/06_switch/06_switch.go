package main

import "fmt"

func main() {
	day := "Saturday"
	
	switch day {
	case "Saturday", "Sunday" :
		fmt.Println("Weekend special")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Firday": 
		fmt.Println("Weekday special")
	default:
		fmt.Println("Unknown day")
	}
}
