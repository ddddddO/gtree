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

	tree := initializeTree(conf, roots)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(w)
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

	tree := initializeTree(conf, roots)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.mkdir()
}
