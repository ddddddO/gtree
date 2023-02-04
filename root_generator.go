package gtree

import (
	"bufio"
	"context"
	"io"
	"runtime/trace"
)

type rootGenerator struct {
	counter       *counter
	scanner       *bufio.Scanner
	nodeGenerator *nodeGenerator
}

func newRootGenerator(r io.Reader, st spaceType) *rootGenerator {
	return &rootGenerator{
		counter:       newCounter(),
		scanner:       bufio.NewScanner(r),
		nodeGenerator: newNodeGenerator(st),
	}
}

func (rg *rootGenerator) generate(ctx context.Context) (<-chan *Node, <-chan error) {
	rootsc := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			trace.StartRegion(ctx, "root generate").End()
			close(rootsc)
			close(errc)
		}()

		var (
			nodes *stack
			roots = newStack()
		)
		for rg.scanner.Scan() {
			currentNode, err := rg.nodeGenerator.generate(rg.scanner.Text(), rg.counter.next())
			if err != nil {
				errc <- err
				return
			}
			if currentNode == nil {
				continue
			}

			if currentNode.isRoot() {
				if roots.size() > 0 {
					rootsc <- roots.pop()
				}
				roots.push(currentNode)

				rg.counter.reset()

				nodes = newStack()
				nodes.push(currentNode)
				continue
			}

			if nodes == nil {
				errc <- errNilStack
				return
			}

			nodes.dfs(currentNode)
		}
		if err := rg.scanner.Err(); err != nil {
			errc <- err
			return
		}

		rootsc <- roots.pop() // 最後のrootを送出
	}()

	return rootsc, errc
}
