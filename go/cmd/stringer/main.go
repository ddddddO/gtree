package main

import (
	"fmt"
)

type Car struct {
	Name string
}

func (c Car) String() string {
	return c.Name + "!!"
}

func main() {
	c := Car{Name: "Vamos"}
	fmt.Println(c)
}
