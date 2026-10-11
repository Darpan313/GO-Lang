package main

import "fmt"

func main() {
	// 15:00 - 17:00
	// isMember
	// orderAmount > 10

	hour := 16
	isMember := true
	orderAmount := 13.50

	if hour >= 15 && hour <= 17 && isMember && orderAmount > 10 {
		fmt.Println("You get 30% off!")
	} else {
		fmt.Println("No happy hours deal available")
	}
 }
