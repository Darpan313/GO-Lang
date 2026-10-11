package main

import "fmt"

func main() {
	// When you use constant, you cannot declare first and assign later(no recommendation for constant camelCase or PascalCase)
	// Only variable defined using Pascal Case can be accessed in another package
	const shopName = "Brew & Beans"

	fmt.Println("Welocome to", shopName)

}
