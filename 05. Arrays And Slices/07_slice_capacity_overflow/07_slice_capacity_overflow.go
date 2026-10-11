package main

import "fmt"

func main() {
	desserts := [3]string{"Cupcake", "Eclair", "Ice-Cream"}
	fmt.Println("Array:",desserts)
	fmt.Println()
	slice := desserts[:1] // cap will be same as original array len which we use to create slice
	fmt.Println("Length of slice:", len(slice), "Capacity of slice:", cap(slice), "Slice:", slice)

	fmt.Println()
	slice = append(slice, "Macron")
	fmt.Println("Array:",desserts)

	fmt.Println()
	slice = append(slice, "Cake")
	fmt.Println("Array:",desserts)
	fmt.Println("Length of slice:", len(slice), "Capacity of slice is:", cap(slice), "Slice:", slice)
	fmt.Println()

	// because len is already equa cap -> new array is allocated for the slice
	slice = append(slice, "Juice")
	fmt.Println("Array:",desserts)
	fmt.Println("Length of slice:", len(slice), "Capacity of slice is:", cap(slice), "Slice:", slice)

	fmt.Println()
	slice[0] = "Chocolate" // orginal array is not modified
	fmt.Println("Array:",desserts)
	fmt.Println("Length of slice:", len(slice), "Capacity of slice is:", cap(slice), "Slice:", slice)

}
