package animal

import (
	"fmt"
)

type Dog struct {
	*Animal
}

func NewDog(name string) *Dog {
	return &Dog{
		&Animal{
			Name:  name,
			Voice: "Wan",
		},
	}
}

func (d *Dog) Run() {
	fmt.Println(d.Name + " is DOG." + d.Name + " is Run")
}
