package main

import (
	"io"

	"github.com/ddddddO/gtree"
)

func mkdir(in io.Reader, indentation gtree.OptFn, extensions []string) error {
	var options []gtree.OptFn
	if indentation != nil {
		options = append(options, indentation)
	}
	options = append(options, gtree.WithFileExtension(extensions))

	return gtree.Mkdir(in, options...)
}
