package main

import (
	"io"

	"github.com/ddddddO/gtree"
)

func mkdir(in io.Reader, options []gtree.Option) error {
	return gtree.MkdirFromMarkdown(in, options...)
}
