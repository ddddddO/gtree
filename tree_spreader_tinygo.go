//go:build wasm

package gtree

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/fatih/color"
)

func newSpreader(encode encode) spreader {
	switch encode {
	case encodeJSON:
		return &jsonSpreader{}
	default:
		return &defaultSpreader{}
	}
}

func newColorizeSpreader(fileExtensions []string) spreader {
	return &colorizeSpreader{
		fileConsiderer: newFileConsiderer(fileExtensions),
		fileColor:      color.New(color.Bold, color.FgHiCyan),
		fileCounter:    newCounter(),

		dirColor:   color.New(color.FgGreen),
		dirCounter: newCounter(),
	}
}

type encode int

const (
	encodeDefault encode = iota
	encodeJSON
	encodeYAML
	encodeTOML
)

type defaultSpreader struct{}

func (ds *defaultSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
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
				if _, err := bw.WriteString(ds.spreadBranch(root)); err != nil {
					errc <- err
					return
				}
			}
		}
		if err := bw.Flush(); err != nil {
			errc <- err
			return
		}
	}()

	return errc
}

func (*defaultSpreader) spreadBranch(current *Node) string {
	ret := current.branch()
	for _, child := range current.children {
		ret += (*defaultSpreader)(nil).spreadBranch(child)
	}
	return ret
}

type colorizeSpreader struct {
	fileConsiderer *fileConsiderer
	fileColor      *color.Color
	fileCounter    *counter

	dirColor   *color.Color
	dirCounter *counter
}

func (cs *colorizeSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
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
		}
		if err := bw.Flush(); err != nil {
			errc <- err
			return
		}
	}()

	return errc
}

func (cs *colorizeSpreader) spreadBranch(current *Node) string {
	cs.colorize(current)
	ret := current.branch()
	for _, child := range current.children {
		ret += cs.spreadBranch(child)
	}
	return ret
}

func (cs *colorizeSpreader) colorize(current *Node) {
	if cs.fileConsiderer.isFile(current) {
		_ = cs.fileCounter.next()
		current.name = cs.fileColor.Sprint(current.name)
	} else {
		_ = cs.dirCounter.next()
		current.name = cs.dirColor.Sprint(current.name)
	}
}

func (cs *colorizeSpreader) summary() string {
	return fmt.Sprintf(
		"%d directories, %d files",
		cs.dirCounter.current(),
		cs.fileCounter.current(),
	)
}

type jsonSpreader struct{}

func (*jsonSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
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

type jsonNode struct {
	Name     string      `json:"value"`
	Children []*jsonNode `json:"children"`
}

func (parent *Node) toJSONNode(jParent *jsonNode) *jsonNode {
	if jParent == nil {
		jParent = &jsonNode{Name: parent.name}
	}
	if !parent.hasChild() {
		return jParent
	}

	jParent.Children = make([]*jsonNode, len(parent.children))
	for i := range parent.children {
		jParent.Children[i] = &jsonNode{Name: parent.children[i].name}
		_ = parent.children[i].toJSONNode(jParent.Children[i])
	}

	return jParent
}

var (
	_ spreader = (*defaultSpreader)(nil)
	_ spreader = (*colorizeSpreader)(nil)
	_ spreader = (*jsonSpreader)(nil)
)
