//go:build !wasm

package gtree

import (
	"bufio"
	"context"
	"strings"
	"sync"
)

type rootGenerator struct {
	nodeGenerator *nodeGenerator
}

func newRootGenerator(st spaceType) *rootGenerator {
	return &rootGenerator{
		nodeGenerator: newNodeGenerator(st),
	}
}

const workerGenerateNum = 10

func (rg *rootGenerator) generate(ctx context.Context, blocks <-chan string) (<-chan *Node, <-chan error) {
	rootc := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(rootc)
			close(errc)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < workerGenerateNum; i++ {
			wg.Add(1)
			go rg.worker(ctx, wg, blocks, rootc, errc)
		}
		wg.Wait()
	}()

	return rootc, errc
}

func (rg *rootGenerator) worker(ctx context.Context, wg *sync.WaitGroup, blocks <-chan string, rootc chan<- *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case block, ok := <-blocks:
			if !ok {
				return
			}

			var (
				sc      = bufio.NewScanner(strings.NewReader(block))
				root    *Node
				nodes   = newStack()
				counter = newCounter()
			)
			for sc.Scan() {
				currentNode, err := rg.nodeGenerator.generate(sc.Text(), counter.next())
				if err != nil {
					errc <- err
					return
				}

				if currentNode == nil {
					continue
				}
				if currentNode.isRoot() {
					root = currentNode
					nodes.push(currentNode)
					continue
				}

				if nodes == nil {
					errc <- errNilStack
					return
				}

				nodes.dfs(currentNode)
			}
			if err := sc.Err(); err != nil {
				errc <- err
				return
			}

			rootc <- root
		}
	}
}
