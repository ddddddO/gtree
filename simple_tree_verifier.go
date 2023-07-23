package gtree

import (
	"fmt"
)

func newVerifierSimple() verifierSimple {
	return &defaultVerifierSimple{}
}

type defaultVerifierSimple struct{}

func (dv *defaultVerifierSimple) verify(roots []*Node) error {
	fmt.Println("in verify!")
	return nil
}
