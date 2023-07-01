//go:build !wasm

package gtree

import (
	"context"
	"sync"
)

func newGrowerPipeline(
	lastNodeFormat, intermedialNodeFormat branchFormat,
	enabledValidation bool,
) growerPipeline {
	return &defaultGrowerPipeline{
		defaultGrowerSimple: newGrowerSimple(lastNodeFormat, intermedialNodeFormat, enabledValidation).(*defaultGrowerSimple),
	}
}

type defaultGrowerPipeline struct {
	*defaultGrowerSimple
}

const workerGrowNum = 10

func (dg *defaultGrowerPipeline) grow(ctx context.Context, roots <-chan *Node) (<-chan *Node, <-chan error) {
	nodes := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(nodes)
			close(errc)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < workerGrowNum; i++ {
			wg.Add(1)
			go dg.worker(ctx, wg, roots, nodes, errc)
		}
		wg.Wait()
	}()

	return nodes, errc
}

func (dg *defaultGrowerPipeline) worker(ctx context.Context, wg *sync.WaitGroup, roots <-chan *Node, nodes chan<- *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			if err := dg.assemble(root); err != nil {
				errc <- err
				return
			}
			select {
			case <-ctx.Done():
				return
			case nodes <- root:
			}
		}
	}
}

func newNopGrowerPipeline() growerPipeline {
	return &nopGrowerPipeline{
		nopGrowerSimple: newNopGrowerSimple().(*nopGrowerSimple),
	}
}

type nopGrowerPipeline struct {
	*nopGrowerSimple
}

func (*nopGrowerPipeline) grow(ctx context.Context, roots <-chan *Node) (<-chan *Node, <-chan error) {
	nodes := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(nodes)
			close(errc)
		}()

	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				nodes <- root
			}
		}
	}()

	return nodes, errc
}

var (
	_ growerPipeline = (*defaultGrowerPipeline)(nil)
	_ growerPipeline = (*nopGrowerPipeline)(nil)
)
