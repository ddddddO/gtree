package main

import (
	"fmt"
)

func switchType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%T: %s\n", v, v)
	case bool:
		fmt.Printf("%T: %t\n", v, v)
	default:
		fmt.Println("not match type...")
	}
}

func main() {
	var i1 interface{} = "interface"
	s := i1.(string) // type assertion
	fmt.Println(s)

	/*
		var i2 interface{} = 999
		ss := i2.(string) // panic!
		fmt.Println(ss)
	*/

	fmt.Println()

	var (
		i3 interface{} = 999
		i4 interface{} = true
	)
	switchType(i1)
	switchType(i3)
	switchType(i4)
}
