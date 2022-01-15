package main

import (
	"io"

	"github.com/ddddddO/gtree"
)

func mkdir(in io.Reader, indentation indentation, extensions []string) error {
	var options []gtree.OptFn

	switch indentation {
	case indentationTS:
		options = append(options, gtree.WithIndentTwoSpaces())
	case indentationFS:
		options = append(options, gtree.WithIndentFourSpaces())
	}

	options = append(options, gtree.WithFileExtension(extensions))

	return gtree.Mkdir(in, options...)
}
