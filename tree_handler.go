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
	splitStream, errcsl := split(ctx, r)
	rootStream, errcr := newRootGenerator(conf.space).generate(ctx, splitStream)
	growStream, errcg := tree.grow(ctx, rootStream)
	errcs := tree.spread(ctx, w, growStream)
	return handlePipelineErr(errcsl, errcr, errcg, errcs)
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
	splitStream, errcsl := split(ctx, r)
	rootStream, errcr := newRootGenerator(conf.space).generate(ctx, splitStream)
	growStream, errcg := tree.grow(ctx, rootStream)
	errcm := tree.mkdir(ctx, growStream)
	return handlePipelineErr(errcsl, errcr, errcg, errcm)
}
