package adapter

import (
	"fmt"
)

// Adapter役(適合する側。Target役(必要とされているもの)を満たそうとする役)
type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(description string) *Adapter {
	adaptee := NewAdaptee(description)
	return &Adapter{
		adaptee: adaptee,
	}
}

func (adapter *Adapter) OutputSharpFrame() {
	fmt.Println("#############")
	adapter.adaptee.Output()
	fmt.Println("#############")
}

func (adapter *Adapter) OutputHyphenFrame() {
	fmt.Println("-------------")
	adapter.adaptee.Output()
	fmt.Println("-------------")
}
