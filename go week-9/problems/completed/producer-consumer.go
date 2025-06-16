package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, n int) {
	for i := range n {
		ch <- i
		fmt.Println("Produced :", i)
		time.Sleep(300 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("Consumed :", i)
	}
}

func main() {

	var wg sync.WaitGroup
	buff := make(chan int, 5)

	wg.Add(1)
	go producer(buff, 10)
	go consumer(buff, &wg)

	wg.Wait()

}
