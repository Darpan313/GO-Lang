package pool

import "sync"

/*
	Add(delta int): Increases or decreases the internal counter by a specified number (usually 
	called with Add(1) before launching each goroutine).

	Done(): Decreases the internal counter by 1. It is typically called inside a defer statement 
	inside the goroutine.

	Wait(): Blocks the calling goroutine (such as main) until the internal counter drops to zero.
*/
type Job struct {
	ID int
	Multiply int
}

type Result struct {
	JobID int
	Value int
}

/*
	The range Loop: Inside the worker, for j := range jobs doesn't just read once; it blocks until a job arrives. 
	The magic of range on a channel is that it automatically exits the loop the exact moment the channel is close()'d and emptied.

	worker acts as a private processing enginer, notice it starts with lowercase 'w' - this hides it from other packages. 
*/
func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- Result{JobID: j.ID, Value: j.ID * j.Multiply}
	}
}

/*
	Process is the public API (starts with an uppercase 'P')
	It orchastrates the pool, blocks until completion, and returns all results.
*/
func Process(numWorkers int, jobsList []Job) []Result {
	jobsChan := make(chan Job, len(jobsList))
	resultsChan := make(chan Result, len(jobsList))
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(jobsChan, resultsChan, &wg)
	}

	// Enqueue jobs
	for _, j := range jobsList {
		jobsChan <- j
	}
	close(jobsChan)

	// Monitor and close results - why?
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var finalResults []Result
	for r := range resultsChan {
		finalResults = append(finalResults, r)
	}
	return finalResults
}
