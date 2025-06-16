package main

import (
	"fmt"
	"sync"
	"time"
)

func sendonly(ch chan<- int) {
	for i := range 10 {
		time.Sleep(200 * time.Millisecond)
		ch <- i
	}
	close(ch)
}

func receiveonly(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		fmt.Println("Received", <-ch)
	}
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go sendonly(ch)
	go receiveonly(ch, &wg)
	wg.Wait()
}
