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

func (rg *rootGenerator) generate(ctx context.Context, lines <-chan string) (<-chan *Node, <-chan error) {
	rootc := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(rootc)
			close(errc)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func(wg *sync.WaitGroup, rootc chan<- *Node, errc chan<- error) {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						return
					case ret, ok := <-lines:
						if !ok {
							return
						}

						var (
							sc      = bufio.NewScanner(strings.NewReader(ret))
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

						errc <- sc.Err() // handlePipelineErrでキャッチ後、ctx.Cancel()されるため、Doneする
						rootc <- root    // rootを送出
					}
				}
			}(wg, rootc, errc)
		}

		wg.Wait()
	}()

	return rootc, errc
}
