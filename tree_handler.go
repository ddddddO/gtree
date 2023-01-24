package gtree

import (
	"context"
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, options ...Option) error {
	conf, err := newConfig(options)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tree := newTree(conf)
	rootStream, errcr := newRootGenerator(r, conf.space).generate(ctx)
	growingStream, errcg := tree.grow(ctx, rootStream)
	errcs := tree.spread(ctx, w, growingStream)

	return handlePipelineErr(errcr, errcg, errcs)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, options ...Option) error {
	conf, err := newConfig(options)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tree := newTree(conf)
	rootStream, errcr := newRootGenerator(r, conf.space).generate(ctx)
	growingStream, errcg := tree.grow(ctx, rootStream)
	errcm := tree.mkdir(ctx, growingStream)

	return handlePipelineErr(errcr, errcg, errcm)
}
