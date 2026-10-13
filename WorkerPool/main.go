package main

import (
	"fmt"

	"github.com/Darpan313/worker-demo/pool"
)

func main() {
	jobs := []pool.Job{
		{ID: 1, Multiply: 5},
		{ID: 2, Multiply: 5},
		{ID: 3, Multiply: 5},
		{ID: 4, Multiply: 5},
	}

	fmt.Println("Starting worker pool...")
	
	// Delegate the complex concurrency to your tested package
	results := pool.Process(3, jobs)

	for _, r := range results {
		fmt.Printf("Job %d output: %d\n", r.JobID, r.Value)
	}
	
	fmt.Println("Shutdown complete.")
}
