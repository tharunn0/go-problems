package main

import (
	"fmt"
)

func printtxt() int {
	//time.Sleep(1 * time.Second)
	fmt.Println("goroutines")
	return 999
}

func main() {
	// fmt.Println("first text")

	// go func() {
	// 	r := printtxt()
	// 	fmt.Println(r)
	// }()
	// fmt.Println("second text")
	// fmt.Println("third text")
	// time.Sleep(3 * time.Second)
	// fmt.Println("fourth text")

	ch := make(chan string)
	str := "goroutines"

	go func() {
		ch <- str
	}()

	rec := <-ch

	fmt.Println(rec)
}
