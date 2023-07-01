//go:build !wasm

package gtree

import (
	"bufio"
	"context"
	"io"
	"strings"
	"sync"
)

type rootGeneratorSimple struct {
	counter       *counter
	scanner       *bufio.Scanner
	nodeGenerator *nodeGenerator
}

func newRootGeneratorSimple(r io.Reader, st spaceType) *rootGeneratorSimple {
	return &rootGeneratorSimple{
		counter:       newCounter(),
		scanner:       bufio.NewScanner(r),
		nodeGenerator: newNodeGenerator(st),
	}
}

func (rg *rootGeneratorSimple) generate() ([]*Node, error) {
	var (
		stack *stack
		roots []*Node
	)

	for rg.scanner.Scan() {
		currentNode, err := rg.nodeGenerator.generate(rg.scanner.Text(), rg.counter.next())
		if err != nil {
			return nil, err
		}
		if currentNode == nil {
			continue
		}

		if currentNode.isRoot() {
			rg.counter.reset()
			roots = append(roots, currentNode)
			stack = newStack()
			stack.push(currentNode)
			continue
		}

		if stack == nil {
			return nil, errNilStack
		}

		stack.dfs(currentNode)
	}

	return roots, rg.scanner.Err()
}

type rootGeneratorPipeline struct {
	nodeGenerator *nodeGenerator
}

func newRootGeneratorPipeline(st spaceType) *rootGeneratorPipeline {
	return &rootGeneratorPipeline{
		nodeGenerator: newNodeGenerator(st),
	}
}

const workerGenerateNum = 10

func (rg *rootGeneratorPipeline) generate(ctx context.Context, blocks <-chan string) (<-chan *Node, <-chan error) {
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

func (rg *rootGeneratorPipeline) worker(ctx context.Context, wg *sync.WaitGroup, blocks <-chan string, rootc chan<- *Node, errc chan<- error) {
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
			select {
			case <-ctx.Done():
				return
			case rootc <- root:
			}
		}
	}
}
