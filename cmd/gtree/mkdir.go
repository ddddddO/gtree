package main

import (
	"io"

	"github.com/ddddddO/gtree"
)

func mkdir(in io.Reader, indentation gtree.Option, extensions []string) error {
	options := []gtree.Option{gtree.WithFileExtensions(extensions)}
	if indentation != nil {
		options = append(options, indentation)
	}

	return gtree.Mkdir(in, options...)
}
