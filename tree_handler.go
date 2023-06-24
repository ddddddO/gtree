//go:build !wasm

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

	if conf.massive {
		tree := newTreePipeline(conf)
		return tree.output(w, r, conf)
	}

	tree := newTreeSimple(conf)
	return tree.output(w, r, conf)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, options ...Option) error {
	conf, err := newConfig(options)
	if err != nil {
		return err
	}

	if conf.massive {
		tree := newTreePipeline(conf)
		return tree.makedir(r, conf)
	}

	tree := newTreeSimple(conf)
	return tree.makedir(r, conf)
}
