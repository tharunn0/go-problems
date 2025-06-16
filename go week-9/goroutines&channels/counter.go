package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mx    sync.Mutex
}

func (c *Counter) add() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.count++
}

func main() {

	var wg sync.WaitGroup

	c1 := Counter{}
	fn := 20
	for range fn {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				c1.add()
			}
		}()
	}

	wg.Wait()

	fmt.Println(c1.count)

}
