package adapter

import (
	"fmt"
)

// Adaptee役(適合される側。元々提供されている機能を持つ役。既存機能)
type Adaptee struct {
	Description string
}

func NewAdaptee(description string) *Adaptee {
	return &Adaptee{
		Description: description,
	}
}

func (adaptee *Adaptee) Output() {
	fmt.Println(adaptee.Description)
}
