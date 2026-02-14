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
//
//go:fix inline
func Output(w io.Writer, r io.Reader, options ...Option) error {
	return OutputFromMarkdown(w, r, options...)
}

// Mkdir makes directories.
//
// Deprecated: Call MkdirFromMarkdown.
//
//go:fix inline
func Mkdir(r io.Reader, options ...Option) error {
	return MkdirFromMarkdown(r, options...)
}

// Verify verifies directories.
//
// Deprecated: Call VerifyFromMarkdown.
//
//go:fix inline
func Verify(r io.Reader, options ...Option) error {
	return VerifyFromMarkdown(r, options...)
}

// Walk executes user-defined function while traversing tree structure recursively.
//
// Deprecated: Call WalkFromMarkdown.
//
//go:fix inline
func Walk(r io.Reader, callback func(*WalkerNode) error, options ...Option) error {
	return WalkFromMarkdown(r, callback, options...)
}
