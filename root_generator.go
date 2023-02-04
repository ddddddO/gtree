package gtree

import (
	"bufio"
	"context"
	"runtime/trace"
	"sync"
)

type rootGenerator struct {
	counter       *counter
	scanner       *bufio.Scanner
	nodeGenerator *nodeGenerator
}

func newRootGenerator(st spaceType) *rootGenerator {
	return &rootGenerator{
		nodeGenerator: newNodeGenerator(st),
	}
}

func (rg *rootGenerator) generate(ctx context.Context, lines <-chan *bufio.Scanner) (<-chan *Node, <-chan error) {
	rootsc := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			trace.StartRegion(ctx, "root generate").End()
			close(rootsc)
			close(errc)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func(wg *sync.WaitGroup, rootsc chan<- *Node, errc chan<- error) {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						return
					case sc, ok := <-lines:
						if !ok {
							return
						}

						// TODO: これ以降、不要な処理が結構あるかも
						var (
							nodes   *stack
							roots   = newStack()
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
								// TODO: このifブロック不要？
								if roots.size() > 0 {
									rootsc <- roots.pop()
								}
								roots.push(currentNode)

								counter.reset() // TODO: 不要？

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
						if err := sc.Err(); err != nil {
							errc <- err
							return
						}

						rootsc <- roots.pop() // rootを送出
					}
				}
			}(wg, rootsc, errc)
		}

		wg.Wait()
	}()

	return rootsc, errc
}
