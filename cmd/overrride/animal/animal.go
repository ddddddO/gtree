package animal

import (
	"fmt"
)

type Creature interface {
	Run()
	Cry()
}

type Animal struct {
	Name  string
	Voice string
}

func (a *Animal) Run() {
	fmt.Println(a.Name + "is Run")
}

func (a *Animal) Cry() {
	fmt.Printf("%s, %s!\n", a.Voice, a.Voice)
}
