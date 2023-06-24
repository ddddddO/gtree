//go:build !wasm

package gtree

import "io"

type tree interface {
	output(io.Writer, io.Reader, *config) error
	outputProgrammably(io.Writer, *Node, *config) error
	makedir(io.Reader, *config) error
	makedirProgrammably(*Node, *config) error
}
