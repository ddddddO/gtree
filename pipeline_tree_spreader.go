//go:build !tinywasm

package gtree

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/goccy/go-yaml"
	toml "github.com/pelletier/go-toml/v2"
)

func newSpreaderPipeline(encode encode) spreaderPipeline {
	switch encode {
	case encodeJSON:
		return newJSONSpreaderPipeline()
	case encodeYAML:
		return newYAMLSpreaderPipeline()
	case encodeTOML:
		return newTOMLSpreaderPipeline()
	default:
		return &defaultSpreaderPipeline{
			defaultSpreaderSimple: &defaultSpreaderSimple{},
		}
	}
}

type defaultSpreaderPipeline struct {
	*defaultSpreaderSimple
	sync.Mutex
}

const workerSpreadNum = 10

func (ds *defaultSpreaderPipeline) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(errc)
		}()

		ds.w = w
		wg := &sync.WaitGroup{}
		for range workerSpreadNum {
			wg.Add(1)
			go ds.worker(ctx, wg, roots, errc)
		}
		wg.Wait()
	}()

	return errc
}

func (ds *defaultSpreaderPipeline) worker(ctx context.Context, wg *sync.WaitGroup, roots <-chan *Node, _ chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}

			ds.Lock()
			ds.spreadBranch(root)
			ds.Unlock()
		}
	}
}

type formattedSpreaderPipeline[T sitter] struct {
	formattedRoot func(string) T
	encode        func(io.Writer) func(any) error
}

func newJSONSpreaderPipeline() *formattedSpreaderPipeline[*jsonNode] {
	return &formattedSpreaderPipeline[*jsonNode]{
		formattedRoot: func(name string) *jsonNode {
			return &jsonNode{Name: name}
		},
		encode: func(w io.Writer) func(any) error {
			return json.NewEncoder(w).Encode
		},
	}
}

func newYAMLSpreaderPipeline() *formattedSpreaderPipeline[*yamlNode] {
	return &formattedSpreaderPipeline[*yamlNode]{
		formattedRoot: func(name string) *yamlNode {
			return &yamlNode{Name: name}
		},
		encode: func(w io.Writer) func(any) error {
			return yaml.NewEncoder(w).Encode
		},
	}
}

func newTOMLSpreaderPipeline() *formattedSpreaderPipeline[*tomlNode] {
	return &formattedSpreaderPipeline[*tomlNode]{
		formattedRoot: func(name string) *tomlNode {
			return &tomlNode{Name: name}
		},
		encode: func(w io.Writer) func(any) error {
			return toml.NewEncoder(w).Encode
		},
	}
}

func (f *formattedSpreaderPipeline[T]) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		encode := f.encode(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				if err := encode(toFormattedNode(root, f.formattedRoot(root.name))); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}

func newColorizeSpreaderPipeline(fileExtensions []string) spreaderPipeline {
	return &colorizeSpreaderPipeline{
		colorizeSpreaderSimple: newColorizeSpreaderSimple(fileExtensions).(*colorizeSpreaderSimple),
	}
}

type colorizeSpreaderPipeline struct {
	*colorizeSpreaderSimple
}

func (cs *colorizeSpreaderPipeline) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		bw := bufio.NewWriter(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				cs.fileCounter.reset()
				cs.dirCounter.reset()

				if _, err := bw.WriteString(
					fmt.Sprintf(
						"%s\n%s\n",
						cs.spreadBranch(root),
						cs.summary()),
				); err != nil {
					errc <- err
					return
				}
			}
			if err := bw.Flush(); err != nil {
				errc <- err
				return
			}
		}
	}()

	return errc
}

var (
	_ spreaderPipeline = (*defaultSpreaderPipeline)(nil)
	_ spreaderPipeline = (*formattedSpreaderPipeline[sitter])(nil)
	_ spreaderPipeline = (*colorizeSpreaderPipeline)(nil)
)
