package singleton

import (
	"fmt"
)

var singleton *instance

type instance struct {
	Description string
}

func GetInstance(description string) *instance {
	if singleton == nil {
		singleton = &instance{
			Description: description,
		}
	}

	return singleton
}

func (i *instance) Speak() {
	fmt.Println(i.Description + ".")
}
