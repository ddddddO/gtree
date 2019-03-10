package main

import (
	"fmt"
)

func incrementGen() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func foo(params ...int) {
	for _, p := range params {
		fmt.Println(p)
	}
}

func main() {
	counter := incrementGen()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	counter2 := incrementGen()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())

	fmt.Println(counter())

	fmt.Println()

	foo(10, 20)
	foo(11, 22, 33)
	ps := []int{40, 50, 60}
	foo(ps...)


}
