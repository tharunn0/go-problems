package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(a int, wg *sync.WaitGroup) {
	fmt.Printf("worker %d started the job\n", a)
	time.Sleep(time.Second)
	fmt.Printf("Joh finished by worker %d\n", a)
	defer wg.Done()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(5)
	for i := range 5 {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println()

}
