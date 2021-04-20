package main

import (
	"fmt"
)

type human struct {
	name        string
	age         int
	description [80]byte
}

func (c *human) echo() {
	fmt.Println("who am i ...")
}

func NewHumanEmpty() *human {
	return &human{}
}

func NewHumanNil() *human {
	return (*human)(nil)
}

func main() {
	c1 := &human{}
	c1.echo()

	c2 := (*human)(nil)
	c2.echo()

	// Output:
	// who am i ...
	// who am i ...
}
