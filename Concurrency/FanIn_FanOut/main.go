package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Generate data (Source)
	inputNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	jobs := generateJobs(inputNums)

	// 2. Fan-out: start 3 workers all reading from jobs
	worker1 := worker(1, jobs)
	worker2 := worker(2, jobs)
	worker3 := worker(3, jobs)

	// 3. Fan-In: Merge the 3 worker output channels into a single channel
	results := fanIn(worker1, worker2, worker3)

	// 4. Consume the merged result
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("All processing complete")
}

// generateJobs converts a slice of integers into a channel (our stream of work)
func generateJobs(nums []int) <-chan int{
	out := make(chan int)
	go func ()  {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// worker processes numbers. It takes an input channel and returns its own output channel.
func worker(id int, in <-chan int) <-chan string{
	out := make(chan string)
	go func ()  {
		for n := range in {
			// Simulate varying processing times
			sleepDuration := time.Duration(rand.Intn(500)+100) * time.Millisecond
			time.Sleep(sleepDuration)
			
			out <- fmt.Sprintf("Worker %d processed number %d", id, n*n)
		}
		close(out)
	}()
	return out
}

func fanIn(channels ...<-chan string) <-chan string{
	var wg sync.WaitGroup
	mergedChannel := make(chan string)

	// helper function to copy values from a single channel to the merged channel
	multiplex := func(c <-chan string) {
		defer wg.Done()
		for n := range c {
			mergedChannel<-n
		}
	}

	// Add 1 to WaitGroup for every channel we're listening to
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Start a single background go routine to wait for all workers to finish
	// Once they finished, close the merged channel so the main function's range loop terminates
	go func ()  {
		wg.Wait()
		close(mergedChannel)
	}()
	return mergedChannel
}