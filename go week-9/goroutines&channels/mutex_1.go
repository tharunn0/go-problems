package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getval() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	var wg sync.WaitGroup

	counter := &counter{}
	ngr := 10

	for range ngr {

		for range 10000 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counter.increment()
			}()
		}
	}

	wg.Wait()

	fmt.Println("total count :", counter.getval())

}
