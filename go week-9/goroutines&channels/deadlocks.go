package main

import (
	"fmt"
	"sync"
)

func main() {

	var m1, m2 sync.Mutex

	go func() {
		m1.Lock()
		fmt.Println("G 1 locked m1")
		m2.Lock()
		fmt.Println("G 1 locked m2")
		m1.Unlock()
		fmt.Println("G1 unlocked m1")
		m2.Unlock()
		fmt.Println("G1 unlocked m2")

	}()
	go func() {
		m2.Lock()
		fmt.Println("G 2 locked m2")
		m1.Lock()
		fmt.Println("G 2 locked m1")
		m2.Unlock()
		fmt.Println("G2 unlocked m2")
		m1.Unlock()
		fmt.Println("G2 unlocked m1")
	}()

	// time.Sleep(3 * time.Second)
	for {
	}
}
