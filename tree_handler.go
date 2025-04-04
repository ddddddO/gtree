//go:build !tinywasm

package gtree

import (
	"io"
)

// OutputFromMarkdown outputs a tree to w with r as Markdown format input.
func OutputFromMarkdown(w io.Writer, r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).output(w, r, cfg)
}

// MkdirFromMarkdown makes directories.
func MkdirFromMarkdown(r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).mkdir(r, cfg)
}

// VerifyFromMarkdown verifies directories.
func VerifyFromMarkdown(r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).verify(r, cfg)
}

// WalkFromMarkdown executes user-defined function while traversing tree structure recursively.
func WalkFromMarkdown(r io.Reader, callback func(*WalkerNode) error, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).walk(r, callback, cfg)
}

// Output outputs a tree to w with r as Markdown format input.
//
// Deprecated: Call OutputFromMarkdown.
func Output(w io.Writer, r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).output(w, r, cfg)
}

// Mkdir makes directories.
//
// Deprecated: Call MkdirFromMarkdown.
func Mkdir(r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).mkdir(r, cfg)
}

// Verify verifies directories.
//
// Deprecated: Call VerifyFromMarkdown.
func Verify(r io.Reader, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).verify(r, cfg)
}

// Walk executes user-defined function while traversing tree structure recursively.
//
// Deprecated: Call WalkFromMarkdown.
func Walk(r io.Reader, callback func(*WalkerNode) error, options ...Option) error {
	cfg := newConfig(options)
	return initializeTree(cfg).walk(r, callback, cfg)
}
