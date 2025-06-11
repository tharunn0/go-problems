package main

import "fmt"

func worker(a int, tasks chan int, res chan int) {
	for range tasks {
		fmt.Printf("worker %d working\n", a)
		res <- a * 2
	}
}

func main() {
	tasks := make(chan int)
	res := make(chan int)

	for i := range 2 {
		go worker(i, tasks, res)
	}
	go func() {
		for range 5 {
			fmt.Println(<-res)

		}
	}()
	for i := range 5 {
		tasks <- i
	}
	close(tasks)

}
