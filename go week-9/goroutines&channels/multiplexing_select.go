package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		fmt.Println("goroutine 1 started processing")
		time.Sleep(7 * time.Second)
		c1 <- 999
	}()

	go func() {
		fmt.Println("goroutine 2 started processing")
		time.Sleep(3 * time.Second)
		c2 <- 111
	}()

	var msg int
	select {
	case msg = <-c1:
		fmt.Println("value from c1 :", msg)
	case msg = <-c2:
		fmt.Println("value from c2 :", msg)

	}

	// fmt.Println("Receiving first value...")
	// msg = <-c1
	// fmt.Println("value from c1 :", msg)
	// fmt.Println("First value received..")

	// fmt.Println("Receiving second value ..")
	// msg = <-c2
	// fmt.Println("value from c2 :", msg)
	// fmt.Println("Second value received..")

}
