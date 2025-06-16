package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Account struct {
	balance float64
	mx      sync.Mutex
	log     []string
}

func (a *Account) deposit() {
	d := float64(rand.Intn(50) + 10)

	a.mx.Lock()
	a.balance += d
	fmt.Println("Deposited to account :", d)
	ds := strconv.FormatFloat(d, 'f', 2, 64)
	bs := strconv.FormatFloat(a.balance, 'f', 2, 64)
	log := ds + " deposited , balance : " + bs
	a.log = append(a.log, log)
	a.mx.Unlock()
}
func (a *Account) withdraw() {
	w := float64(rand.Intn(50) + 10)

	a.mx.Lock()
	if a.balance-float64(w) < 0 {
		fmt.Println("Not enough balance")
	} else {
		a.balance -= w
		fmt.Println("Withdrawed from account : ", w)

		ws := strconv.FormatFloat(w, 'f', 2, 64)
		bs := strconv.FormatFloat(a.balance, 'f', 2, 64)
		log := ws + " withdrawed , balance : " + bs
		a.log = append(a.log, log)
	}
	a.mx.Unlock()

}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	workers := 5
	acc := Account{}

	for range workers {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for range 50 {
				acc.deposit()
			}
		}()
		go func() {
			defer wg.Done()
			for range 50 {
				acc.withdraw()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Balance :", acc.balance)

	for _, v := range acc.log {
		fmt.Println(v)
	}
}
