/*
Problem: Given a large slice of integers, divide it into N chunks
and sum each chunk concurrently. Finally, print the total sum.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func createarr(len int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, len)

	for i, _ := range arr {
		arr[i] = rand.Intn(90) + 10
	}
	return arr
}

func sum(ar []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var s int
	for _, v := range ar {
		s += v
	}
	c <- s
}

func main() {

	arr := createarr(200)
	csize := 40
	var chnks [][]int

	ch := make(chan int, 50)
	var wg sync.WaitGroup

	for i := 0; i < len(arr); i += csize {
		end := i + csize
		if end > len(arr) {
			end = len(arr)
		}
		slice := arr[i:end]
		chnks = append(chnks, slice)
	}

	for i := range chnks {
		wg.Add(1)

		go sum(chnks[i], ch, &wg)
	}

	wg.Wait()
	close(ch)

	var tot int
	for v := range ch {
		tot += v
	}
	fmt.Println("Total sum :", tot)
}
