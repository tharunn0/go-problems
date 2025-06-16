package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, jobs chan string, res chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d has started doing %s\n", i, job)
		res <- job
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {

	var wg sync.WaitGroup
	jobs := make(chan string)
	comp := make(chan string)
	availablejobs := []string{
		"digging",
		"welding",
		"painting",
		"plastering",
		"bricklaying",
		"carpentry",
		"tiling",
		"roofing",
		"drilling",
		"measuring",
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, comp, &wg)
	}

	go func() {
		for c := range comp {
			fmt.Printf("%s has completed\n", c)
		}
	}()

	for _, v := range availablejobs {
		jobs <- v
	}

	close(jobs)
	wg.Wait()
	close(comp)
}
