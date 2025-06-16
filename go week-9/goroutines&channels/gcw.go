package main

import (
	"fmt"
	"sync"
)

func main() {

	res := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		go func() {
			for c := range res {
				fmt.Println(c)
			}
		}()
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			res <- v
		}(i)
	}

	wg.Wait()
	close(res)
	fmt.Println("Completed")

}
