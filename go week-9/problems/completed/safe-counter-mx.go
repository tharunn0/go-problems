package main

import (
	"fmt"
	"sync"
)

/* Goal: Create a counter that multiple goroutines can increment safely. */

type Counter struct {
	count int
	mx    sync.Mutex
}

func (c *Counter) increment() {

	c.mx.Lock()
	c.count++
	c.mx.Unlock()
}
func (c *Counter) getval() int {
	return c.count
}

func main() {

	counter1 := Counter{}
	var wg sync.WaitGroup

	noG := 100

	for range noG {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for range 1000 {
				// counter1.count++ // printed 31763 at the end
				counter1.increment() // printed exact 100000
			}
		}(&wg)
	}

	wg.Wait()

	total := counter1.getval()
	fmt.Println("Total Count :", total)

}
