package main

import "fmt"

func main() {
	var coffeeSizes [3]string // default element ""
	fmt.Println(coffeeSizes)

	coffeeSizes[0] = "Small"
	coffeeSizes[1] = "Meduim"
	coffeeSizes[2] = "Large"

	fmt.Println(coffeeSizes)

	fmt.Println(len(coffeeSizes))

	// coffeeSizes[0] = true //compilation error. type is not string
	// coffeeSizes[5] = "Super Extra Large" // index 5 out of bounds
}
