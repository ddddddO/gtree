//go:build !tinywasm

package gtree

import (
	"context"
	"sync"
)

type defaultWalkerPipeline struct {
	*defaultWalkerSimple
}

func newWalkerPipeline() walkerPipeline {
	return &defaultWalkerPipeline{
		defaultWalkerSimple: &defaultWalkerSimple{},
	}
}

const workerWalkerNum = 10

func (dw *defaultWalkerPipeline) walk(ctx context.Context, roots <-chan *Node, cb func(*WalkerNode) error) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(errc)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < workerWalkerNum; i++ {
			wg.Add(1)
			go dw.worker(ctx, wg, roots, cb, errc)
		}
		wg.Wait()
	}()

	return errc
}

func (dw *defaultWalkerPipeline) worker(ctx context.Context, wg *sync.WaitGroup, roots <-chan *Node, cb func(*WalkerNode) error, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			if err := dw.walkNode(root, cb); err != nil {
				errc <- err
			}
		}
	}
}
