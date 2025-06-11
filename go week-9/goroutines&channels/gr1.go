package main

import (
	"fmt"
	"time"
)

func calc(c chan int, a int) {
	res := a * a
	c <- res
}

func findsum(c chan int, arr []int) {
	var s int
	for _, v := range arr {
		s += v
	}
	c <- s
}

func passvalue(ch chan int, a int) {
	go func() {
		for i := range a {
			ch <- i
		}
		close(ch)
	}()
}

func receiveval(ch <-chan int) {
	for v := range ch {
		fmt.Println("Goroutine", v, "working...")
	}
}

func main() {
	ch := make(chan string)
	t := "tharun"
	go func(c chan string) {
		ch <- t
		fmt.Printf("val -> %d\n", ch)
	}(ch)
	<-ch
	time.Sleep(2 * time.Second)
	fmt.Println("main program finished")
}
