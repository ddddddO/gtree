package gtree

import (
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, options ...Option) error {
	conf, err := newConfig(options)
	if err != nil {
		return err
	}

	rg := newRootGenerator(r, conf.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	tree := newTree(conf)
	growingStream, errcg := tree.grow(roots)
	errcs := tree.spread(w, growingStream)

	return handleErr(errcg, errcs)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, options ...Option) error {
	conf, err := newConfig(options)
	if err != nil {
		return err
	}

	rg := newRootGenerator(r, conf.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	tree := newTree(conf)
	growingStream, errcg := tree.grow(roots)
	errcm := tree.mkdir(growingStream)

	return handleErr(errcg, errcm)
}
