package main

import (
	"github.com/ddddddO/work/go/gof/singleton"
)

func main() {
	desc1 := "only one instance"
	instance1 := singleton.GetInstance(desc1)

	desc2 := "not only one instance!"
	instance2 := singleton.GetInstance(desc2)

	instance1.Speak()
	instance2.Speak()

	// Output:
	// only one instance.
	// only one instance.
}
