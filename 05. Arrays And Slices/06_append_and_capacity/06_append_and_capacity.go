package main

import "fmt"

func main() {
	menu := []string{"Cake", "Pie"}

	fmt.Println("Initial menu:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) //2,2
	fmt.Printf("Memory Location:%p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])
	fmt.Println("--------")

	menu = append(menu, "Donut")

	fmt.Println("Menu after adding donut:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) //3,4
	fmt.Printf("Memory Location:%p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])
	fmt.Println("--------")

	menu = append(menu, "Ice Cream")
	fmt.Println("Menu after adding icecream:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) //4,4
	fmt.Printf("Memory Location:%p\n", &menu)
	fmt.Printf("Memory location of \"Ice Cream\": %p\n", &menu[0])
	fmt.Println("--------")

	menu = append(menu, "Cream")
	fmt.Println("Menu after adding Cream:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) //4,4
	fmt.Printf("Memory Location:%p\n", &menu)
	fmt.Printf("Memory location of \"Cream\": %p\n", &menu[0])
	fmt.Println("--------")

	cupSizes := make([]string, 0, 5) // len will be 0 and capacity will be 5
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSize:", cap(cupSizes))
	// cupSizes[0] = "Small" // panic: runtime error: index out of range [0] with length 0
	cupSizes = append(cupSizes, "Small", "Meduim")
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSize:", cap(cupSizes))
	cupSizes[0] = "Extra Small"
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSize:", cap(cupSizes))
}
