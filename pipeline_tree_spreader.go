//go:build !wasm

package gtree

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/fatih/color"
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
		return &defaultSpreaderPipeline{}
	}
}

func newColorizeSpreaderPipeline(fileExtensions []string) spreaderPipeline {
	return &colorizeSpreaderPipeline{
		fileConsiderer: newFileConsiderer(fileExtensions),
		fileColor:      color.New(color.Bold, color.FgHiCyan),
		fileCounter:    newCounter(),

		dirColor:   color.New(color.FgGreen),
		dirCounter: newCounter(),
	}
}

type defaultSpreaderPipeline struct {
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

		errc <- bw.Flush()
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

			errc <- err
		}
	}
}

func (*defaultSpreaderPipeline) spreadBranch(current *Node) string {
	ret := current.name + "\n"
	if !current.isRoot() {
		ret = current.branch() + " " + current.name + "\n"
	}

	for _, child := range current.children {
		ret += (*defaultSpreaderPipeline)(nil).spreadBranch(child)
	}
	return ret
}

type colorizeSpreaderPipeline struct {
	fileConsiderer *fileConsiderer
	fileColor      *color.Color
	fileCounter    *counter

	dirColor   *color.Color
	dirCounter *counter
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

func (cs *colorizeSpreaderPipeline) spreadBranch(current *Node) string {
	ret := ""
	if current.isRoot() {
		ret = cs.colorize(current) + "\n"
	} else {
		ret = current.branch() + " " + cs.colorize(current) + "\n"
	}

	for _, child := range current.children {
		ret += cs.spreadBranch(child)
	}
	return ret
}

func (cs *colorizeSpreaderPipeline) colorize(current *Node) string {
	if cs.fileConsiderer.isFile(current) {
		_ = cs.fileCounter.next()
		return cs.fileColor.Sprint(current.name)
	} else {
		_ = cs.dirCounter.next()
		return cs.dirColor.Sprint(current.name)
	}
}

func (cs *colorizeSpreaderPipeline) summary() string {
	return fmt.Sprintf(
		"%d directories, %d files",
		cs.dirCounter.current(),
		cs.fileCounter.current(),
	)
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

var (
	_ spreaderPipeline = (*defaultSpreaderPipeline)(nil)
	_ spreaderPipeline = (*colorizeSpreaderPipeline)(nil)
	_ spreaderPipeline = (*jsonSpreaderPipeline)(nil)
	_ spreaderPipeline = (*yamlSpreaderPipeline)(nil)
	_ spreaderPipeline = (*tomlSpreaderPipeline)(nil)
)
