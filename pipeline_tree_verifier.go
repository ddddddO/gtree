package gtree

import (
	"context"
	"fmt"
)

type defaultVerifierPipeline struct {
	*defaultVerifierSimple
}

func newVerifierPipeline(dir string, strict bool) verifierPipeline {
	return &defaultVerifierPipeline{
		defaultVerifierSimple: newVerifierSimple(dir, strict).(*defaultVerifierSimple),
	}
}

func (dv *defaultVerifierPipeline) verify(ctx context.Context, roots <-chan *Node) <-chan error {
	fmt.Println("in verify pipeline")
	errc := make(chan error, 1)

	go func() {
		defer close(errc)
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					return
				}
				exists, noExists, err := dv.verifyRoot(root)
				if err != nil {
					errc <- err
				}
				if err := dv.handleErr(exists, noExists); err != nil {
					errc <- err
				}
			}
		}
	}()

	fmt.Println("in verify pipeline end")
	return errc
}
