//go:build !tinywasm

package gtree

import (
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).output(w, r, cfg)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).mkdir(r, cfg)
}

// Verify verifies directories.
func Verify(r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).verify(r, cfg)
}

// Walk executes user-defined function while traversing tree structure recursively.
func Walk(r io.Reader, cb func(*WalkerNode) error, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).walk(r, cb, cfg)
}
