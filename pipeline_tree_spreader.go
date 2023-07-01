//go:build !wasm

package gtree

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	toml "github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func newSpreaderPipeline(encode encode) spreaderPipeline {
	switch encode {
	case encodeJSON:
		return &jsonSpreaderPipeline{}
	case encodeYAML:
		return &yamlSpreaderPipeline{}
	case encodeTOML:
		return &tomlSpreaderPipeline{}
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

		bw := bufio.NewWriter(w)
		wg := &sync.WaitGroup{}
		for i := 0; i < workerSpreadNum; i++ {
			wg.Add(1)
			go ds.worker(ctx, wg, bw, roots, errc)
		}
		wg.Wait()

		if err := bw.Flush(); err != nil {
			errc <- err
		}
	}()

	return errc
}

func (ds *defaultSpreaderPipeline) worker(ctx context.Context, wg *sync.WaitGroup, bw *bufio.Writer, roots <-chan *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			ret := ds.spreadBranch(root)

			ds.Lock()
			_, err := bw.WriteString(ret)
			ds.Unlock()

			if err != nil {
				errc <- err
			}
		}
	}
}

type jsonSpreaderPipeline struct{}

func (*jsonSpreaderPipeline) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := json.NewEncoder(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				if err := enc.Encode(root.toJSONNode(nil)); err != nil {
					errc <- err
					return
				}
			}
		}
	}()

	return errc
}

type tomlSpreaderPipeline struct{}

func (*tomlSpreaderPipeline) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := toml.NewEncoder(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				if err := enc.Encode(root.toTOMLNode(nil)); err != nil {
					errc <- err
					return
				}
			}
		}
	}()

	return errc
}

type yamlSpreaderPipeline struct{}

func (*yamlSpreaderPipeline) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := yaml.NewEncoder(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				if err := enc.Encode(root.toYAMLNode(nil)); err != nil {
					errc <- err
					return
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
	_ spreaderPipeline = (*jsonSpreaderPipeline)(nil)
	_ spreaderPipeline = (*yamlSpreaderPipeline)(nil)
	_ spreaderPipeline = (*tomlSpreaderPipeline)(nil)
	_ spreaderPipeline = (*colorizeSpreaderPipeline)(nil)
)
