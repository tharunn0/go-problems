package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	tid  int
	nos  int
	cost float64
}

func ticketProcess(w int, req <-chan Request, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range req {
		fmt.Printf("W:%d - Processing id %d , No of tickets : %d , Total cost : %.2f\n", w, i.tid, i.nos, i.cost)
		res <- i.tid
	}
}

func main() {
	wrk := 4
	req := make(chan Request)
	res := make(chan int, 25)
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// creating 4 workers
	for i := range wrk {
		go ticketProcess(i, req, res, &wg)
		wg.Add(1)
	}

	// function sending request data to the channel
	go func() {
		for range 25 {

			n := rand.Intn(6) + 1
			u := rand.Intn(9999) + 999
			req <- Request{
				tid:  u,
				nos:  n,
				cost: float64(n * 60),
			}
		}
		close(req)

	}()

	wg.Wait()

}
