package main

import (
	"io"

	"github.com/ddddddO/gtree"
)

func mkdir(in io.Reader, indentation gtree.OptFn, extensions []string) error {
	options := []gtree.OptFn{gtree.WithFileExtension(extensions)}
	if indentation != nil {
		options = append(options, indentation)
	}

	return gtree.Mkdir(in, options...)
}
