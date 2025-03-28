//go:build !tinywasm

package gtree

import (
	"io"
	"iter"
)

type tree interface {
	output(io.Writer, io.Reader, *config) error
	outputProgrammably(io.Writer, *Node, *config) error
	mkdir(io.Reader, *config) error
	mkdirProgrammably(*Node, *config) error
	verify(io.Reader, *config) error
	verifyProgrammably(*Node, *config) error
	walk(io.Reader, func(*WalkerNode) error, *config) error
	walkProgrammably(*Node, func(*WalkerNode) error, *config) error
	walkIterProgrammably(*Node, *config) iter.Seq2[*WalkerNode, error]
}

func initializeTree(cfg *config) tree {
	if cfg.massive {
		return newTreePipeline(cfg)
	}
	return newTreeSimple(cfg)
}
