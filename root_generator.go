package gtree

import (
	"bufio"
	"context"
	"io"
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

func (rg *rootGenerator) generate(_ context.Context) (<-chan *Node, <-chan error) {
	rootsc := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer close(rootsc)
		defer close(errc)

		var (
			nodes          *stack
			roots          = newStack()
			isNotFirstRoot bool
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
				if isNotFirstRoot {
					rootsc <- roots.pop()
				}

				rg.counter.reset()
				roots.push(currentNode)
				nodes = newStack()
				nodes.push(currentNode)
				isNotFirstRoot = true
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
