//go:build !tinywasm

package gtree

import (
	"context"
	"sync"
)

func newMkdirerPipeline(dir string, fileExtensions []string) mkdirerPipeline {
	return &defaultMkdirerPipeline{
		defaultMkdirerSimple: newMkdirerSimple(dir, fileExtensions).(*defaultMkdirerSimple),
	}
}

type defaultMkdirerPipeline struct {
	*defaultMkdirerSimple
}

const workerMkdirNum = 10

func (dm *defaultMkdirerPipeline) mkdir(ctx context.Context, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		wg := &sync.WaitGroup{}
		for i := 0; i < workerMkdirNum; i++ {
			wg.Add(1)
			go dm.worker(ctx, wg, roots, errc)
		}
		wg.Wait()
	}()

	return errc
}

func (dm *defaultMkdirerPipeline) worker(ctx context.Context, wg *sync.WaitGroup, roots <-chan *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			if dm.isExistRoot([]*Node{root}) {
				errc <- ErrExistPath
				return
			}
			if err := dm.makeDirectoriesAndFiles(root); err != nil {
				errc <- err
				return
			}
		}
	}
}

var _ mkdirerPipeline = (*defaultMkdirerPipeline)(nil)
