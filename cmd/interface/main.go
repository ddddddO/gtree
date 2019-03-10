package main

import (
	"fmt"
)

type Car interface {
	Introduce()
	Move() string
}

type MiniCar struct {
	Name      string
	Direction string
}

type BigCar struct {
	Name      string
	Direction string
}

func (m MiniCar) Introduce() {
	fmt.Printf("I am %s!\n", m.Name)
}

func (m MiniCar) Move() string {
	return m.Direction
}

func canMove(c Car) {
	if c.Move() != "north" {
		fmt.Println("cat not move...")
	} else {
		fmt.Println("let's go!")
	}
}

func main() {
	var vamos Car
	vamos = MiniCar{Name: "Vamos", Direction: "east"}

	vamos.Introduce()
	fmt.Println(vamos.Move())

	canMove(vamos)

	/*
		pajero := BigCar{Name: "Pajero", Direction: "north"}
		canMove(pajero)
	*/
}
