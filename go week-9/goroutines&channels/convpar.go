package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(a int, wg *sync.WaitGroup) {
	defer wg.Done()

	t := time.Now()
	fmt.Printf("worker %d started working...\n", a)
	for range 10_000_000_000 {
	}

	fmt.Printf("worker %d finished the work in %v\n", a, time.Since(t))
}

func main() {

	var wg sync.WaitGroup

	for i := range 50 {
		wg.Add(1)
		go func(x int) {
			worker(x, &wg)
		}(i)
	}

	wg.Wait()

	fmt.Println("All have finished the job...")
}
