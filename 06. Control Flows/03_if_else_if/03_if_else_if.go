package main

import "fmt"

func main() {
	// var points int // zero values for int is 0

	// points = 5

	// if points > 0 {
	// 	fmt.Println("Client earned some points")
	// }

	points := 75

	if points >= 100 { //[100, ...]
		fmt.Println("Platinium member: Free coffee every day!")
	} else if points >= 50  && points < 100 { // [50,100)
		fmt.Println("Gold member: 20% discount on latte");
	} else if points >= 20 { // [20, 50)]
		fmt.Println("Silver member: Free cookie on Monday")
	} else { // [0,20)]
		fmt.Println("Broonze member: keep sipping to earn rewards")
	}
}
