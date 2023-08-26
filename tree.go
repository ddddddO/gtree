//go:build !tinywasm

package gtree

import "io"

type tree interface {
	output(io.Writer, io.Reader, *config) error
	outputProgrammably(io.Writer, *Node, *config) error
	mkdir(io.Reader, *config) error
	mkdirProgrammably(*Node, *config) error
	verify(io.Reader, *config) error
	verifyProgrammably(*Node, *config) error
}

func initializeTree(cfg *config) tree {
	tree := newTreeSimple(cfg)
	if cfg.massive {
		tree = newTreePipeline(cfg)
	}
	return tree
}
