package gtree

import (
	"bufio"
	"io"
)

type rootGenerator struct {
	counter  *counter
	scanner  *bufio.Scanner
	strategy nodeGenerateStrategy
}

func newRootGenerator(r io.Reader, st spaceType) *rootGenerator {
	return &rootGenerator{
		counter:  newCounter(),
		scanner:  bufio.NewScanner(r),
		strategy: newStrategy(st),
	}
}

func (rg *rootGenerator) generate() ([]*Node, error) {
	var (
		stack *stack
		roots []*Node
	)

	for rg.scanner.Scan() {
		currentNode, err := rg.strategy.generate(rg.scanner.Text(), rg.counter.next())
		if err != nil {
			return nil, err
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

	if err := rg.scanner.Err(); err != nil {
		return nil, err
	}
	return roots, nil
}
