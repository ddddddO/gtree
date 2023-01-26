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
		defaultSpreader: &defaultSpreader{},

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

		branches := ""
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				branches += ds.spreadBranch(root)
			}
		}
		if err := ds.write(w, branches); err != nil {
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

func (*defaultSpreader) write(w io.Writer, in string) error {
	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(in); err != nil {
		return err
	}
	return buf.Flush()
}

type colorizeSpreader struct {
	*defaultSpreader // NOTE: xxx

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

		ret := ""
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
				ret += fmt.Sprintf(
					"%s\n%s\n",
					cs.spreadBranch(root),
					cs.summary(),
				)
			}
		}
		if err := cs.write(w, ret); err != nil {
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
				jRoot := root.toJSONNode(nil)
				if err := enc.Encode(jRoot); err != nil {
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
