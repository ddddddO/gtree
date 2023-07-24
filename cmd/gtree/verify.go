package main

import (
	"io"

	"github.com/ddddddO/gtree"
)

func verify(in io.Reader, options []gtree.Option) error {
	return gtree.Verify(in, options...)
}
