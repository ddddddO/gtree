//go:build !wasm

package gtree

import (
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, options ...Option) error {
	cfg := newConfig(options)

	tree := newTreeSimple(cfg)
	if cfg.massive {
		tree = newTreePipeline(cfg)
	}
	return tree.output(w, r, cfg)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, options ...Option) error {
	cfg := newConfig(options)

	tree := newTreeSimple(cfg)
	if cfg.massive {
		tree = newTreePipeline(cfg)
	}
	return tree.mkdir(r, cfg)
}

// Verify verifies directories.
func Verify(r io.Reader, options ...Option) error {
	cfg := newConfig(options)

	tree := newTreeSimple(cfg)
	if cfg.massive {
		tree = newTreePipeline(cfg)
	}
	return tree.verify(r, cfg)
}
