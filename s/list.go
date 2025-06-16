package main

import (
	"container/list"
	f "fmt"
)

func closure() func() {
	name := "Tharun"
	return func() {
		f.Println("Hello", name)
	}
}

func main() {
	li := list.New()

	li.PushFront("Tharun")
	li.PushBack(20)
	li.PushBack("Golang")

	// Print objects in list
	for e := li.Front(); e != nil; e = e.Next() {
		f.Println(e.Value)
	}

	f.Println(li.Len())

	cl := closure()
	cl()
}
