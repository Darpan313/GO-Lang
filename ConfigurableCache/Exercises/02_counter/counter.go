package main

import "fmt"

/*
	Structs, methods and pointer receivers

	Build: a Counter struct with Inc() and Value().

	Try: write Inc with a value receiver (c Counter) first, see that the count stays 0, then switch to (c *Counter).

	You’ve got it when: you can explain why every cache method uses *Cache.
*/

type Counter struct {
	Count int
}

// without pointer receiver prints Value 0 after Inc()
// c Counter changes value of a copy

// func (c Counter) Inc() {
// 	c.Count++
// }

// func (c Counter) Value() int {
// 	return c.Count
// }

// Pointer reciever changes actual value
func (c *Counter) Inc() {
	c.Count++
}

func (c Counter) Value() int {
	return c.Count
}

func main() {
	counter := Counter{
		Count: 0,
	}

	counter.Inc()

	fmt.Println(counter.Value())
}
