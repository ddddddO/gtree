//go:build wasm

package gtree

import (
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, options ...Option) error {
	cfg := newConfig(options)

	rg := newRootGenerator(r, cfg.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	tree := newTree(cfg, roots)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(w)
}
