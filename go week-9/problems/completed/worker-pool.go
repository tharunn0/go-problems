package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, job chan string, res chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range job {
		fmt.Printf("Worker %d started to do %s\n", id, task)
		time.Sleep(3 * time.Second)
		fmt.Printf("Worker %d has finished %s\n", id, task)
		res <- task
	}

}

func main() {
	jobs := []string{"painting", "laying bricks", "mixing", "resting", "eating", "washing", "checking", "unloading", "loading"}

	jobch := make(chan string, len(jobs))
	resultch := make(chan string, len(jobs))

	var wg sync.WaitGroup

	for _, j := range jobs {
		jobch <- j
	}

	for i := range 3 {
		wg.Add(1)
		go worker(i, jobch, resultch, &wg)
	}

	close(jobch)

	wg.Wait()

	close(resultch)

	fmt.Printf("All Jobs done are :")
	for r := range resultch {
		fmt.Printf("%s, ", r)
	}
}
