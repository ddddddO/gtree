//go:build !tinywasm

package gtree

import (
	"context"
	"sync"
)

type defaultVerifierPipeline struct {
	*defaultVerifierSimple
}

func newVerifierPipeline(dir string, strict bool) verifierPipeline {
	return &defaultVerifierPipeline{
		defaultVerifierSimple: newVerifierSimple(dir, strict).(*defaultVerifierSimple),
	}
}

const workerVerifyNum = 10

func (dv *defaultVerifierPipeline) verify(ctx context.Context, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(errc)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < workerVerifyNum; i++ {
			wg.Add(1)
			go dv.worker(ctx, wg, roots, errc)
		}
		wg.Wait()
	}()

	return errc
}

func (dv *defaultVerifierPipeline) worker(ctx context.Context, wg *sync.WaitGroup, roots <-chan *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			extra, noExists, err := dv.verifyRoot(root)
			if err != nil {
				errc <- err
			}
			// TODO: 1Root分のエラーしか出力しないようになってるから、全Root分の検査結果を出力する方がいいかも
			if err := dv.handleErr(extra, noExists); err != nil {
				errc <- err
			}
		}
	}
}
