package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, res chan<- int) {

	for task := range tasks {
		fmt.Printf("Worker %d assigned task %d ..\n", id, task)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d has done task %d\n", id, task)
		res <- task * 2
	}
}

func main() {
	fmt.Println("Worker Pools")

	jobs := 10
	workers := 3
	tasks := make(chan int, jobs)
	results := make(chan int, jobs)

	for i := range workers {
		go worker(i, tasks, results)
	}

	go func() {

		for i := range jobs {
			tasks <- i

		}
		close(tasks)
	}()

	for range jobs {
		res := <-results
		fmt.Println("result :", res)
	}

}
