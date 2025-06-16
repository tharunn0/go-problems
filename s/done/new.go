package main

import (
	"fmt"
)

func escape() int {
	x := 99
	fmt.Printf("address of %d in the function : %p\n", x, &x)
	return x
}

func Newescape() *int {
	x := 99
	fmt.Printf("address of %d in the function : %p\n", x, &x)
	return &x
}

func main() {

	x := escape()

	fmt.Printf("address of %d in main function : %p\n", x, &x)

	fmt.Println("=====================")

	y := Newescape()
	fmt.Printf("address of %d in main function : %p\n", *y, y)

	p := new(int)

	*p = 99

	fmt.Printf("val is %d and add is %p\n", *p, p)

}
