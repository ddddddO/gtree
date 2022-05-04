package gtree

import (
	"bufio"
	"errors"
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
		currentNode := rg.strategy.generate(rg.scanner.Text(), rg.counter.next())
		if err := currentNode.validate(); err != nil {
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

var (
	errEmptyText       = errors.New("empty text")
	errIncorrectFormat = errors.New("incorrect input format")
)

func (n *Node) validate() error {
	if n.hierarchy == 0 {
		return errIncorrectFormat
	}
	if len(n.name) == 0 {
		return errEmptyText
	}
	return nil
}
