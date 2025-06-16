package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(id int, tasks chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for t := range tasks {
		fmt.Printf("worker %d started %s \n", id, t)
		time.Sleep(2 * time.Second)
		fmt.Printf("worker %d has finished working\n", id)
	}
}

func main() {

	// var wg sync.WaitGroup

	// tasks := []string{"building", "working", "painting", "resting"}
	// ch := make(chan string, 5)

	// for _, v := range tasks {
	// 	ch <- v
	// }
	// close(ch)

	// for i := range 3 {
	// 	wg.Add(1)
	// 	go workers(i, ch, &wg)
	// }

	// wg.Wait()

	st := "Hello"

	r := []rune(st)

	r[4] = 'O'

	fmt.Println(string(r))

}
