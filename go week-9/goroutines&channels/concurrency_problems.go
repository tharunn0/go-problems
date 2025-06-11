package main

import (
	"fmt"
	"sync"
)

func worker(a int, tasks chan int, res chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("worker %d finding square of %d\n", a, task)
		res <- task * task
	}
}

func senddata(a []int, task chan int) {
	for _, v := range a {
		task <- v
	}
	close(task)
}

func main() {
	sl := []int{3, 4, 5, 10, 9}
	task := make(chan int)
	res := make(chan int)
	var wg sync.WaitGroup

	// creating worker
	for i := range 2 {
		fmt.Println(i)
		go worker(i, task, res, &wg)
		wg.Add(1)
	}

	//function to receive values
	go func() {
		for range len(sl) {
			fmt.Println("squares are :", <-res)
		}
	}()
	senddata(sl, task)

	wg.Wait()

}
