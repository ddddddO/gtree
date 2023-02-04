package gtree

import (
	"context"
	"io"
	"log"
	"os"
	"runtime/trace"
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

	// trace start
	f, err := os.Create("trace_output_func.out")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err := trace.Start(f); err != nil {
		return err
	}
	defer trace.Stop()

	ctx, task := trace.NewTask(ctx, "Output tree")
	defer task.End()

	rootStream, errcr := newRootGenerator(r, conf.space).generate(ctx)
	growStream, errcg := tree.grow(ctx, rootStream)
	errcs := tree.spread(ctx, w, growStream)
	return handlePipelineErr(ctx, errcr, errcg, errcs)
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
	growStream, errcg := tree.grow(ctx, rootStream)
	errcm := tree.mkdir(ctx, growStream)
	return handlePipelineErr(ctx, errcr, errcg, errcm)
}
