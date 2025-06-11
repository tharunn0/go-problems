package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) performTask(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Printf("Worker ID:%d started to do %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d has finished %s in %v\n", w.ID, w.Task, time.Since(start))
}

func main() {
	var wg sync.WaitGroup
	tasks := []string{"painting", "laying brics", "cleaning", "planning", "mixing concrete", "wiring"}

	rand.Seed(time.Now().UnixNano())
	for _, v := range tasks {
		i := rand.Intn(9999) + 1111
		worker := Worker{ID: i, Task: v}
		wg.Add(1)
		go worker.performTask(&wg)
	}

	wg.Wait()

	fmt.Println("Construction Completed !!!")

}
