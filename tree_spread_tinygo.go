//go:build wasm

package gtree

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/fatih/color"
)

// 関心事はtreeの出力
type spreader interface {
	spread(io.Writer, []*Node) error
}

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
		fileConsiderer:  newFileConsiderer(fileExtensions),
		colorFile:       color.New(color.Bold, color.FgHiCyan),
		colorDir:        color.New(color.FgGreen),
		counterFile:     newCounter(),
		counterDir:      newCounter(),
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

func (ds *defaultSpreader) spread(w io.Writer, roots []*Node) error {
	branches := ""
	for _, root := range roots {
		branches += ds.spreadBranch(root)
	}
	return ds.write(w, branches)
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
	fileConsiderer   *fileConsiderer
	colorFile        *color.Color
	colorDir         *color.Color
	counterFile      *counter
	counterDir       *counter
}

func (cs *colorizeSpreader) spread(w io.Writer, roots []*Node) error {
	ret := ""
	for _, root := range roots {
		cs.counterFile.reset()
		cs.counterDir.reset()

		ret += fmt.Sprintf("%s\n%s", cs.spreadBranch(root), cs.summary())
	}

	return cs.write(w, ret)
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
		_ = cs.counterFile.next()
		current.name = cs.colorFile.Sprint(current.name)
	} else {
		_ = cs.counterDir.next()
		current.name = cs.colorDir.Sprint(current.name)
	}
}

func (cs *colorizeSpreader) summary() string {
	return fmt.Sprintf("%d directories, %d files\n",
		cs.counterDir.current(),
		cs.counterFile.current(),
	)
}

type jsonSpreader struct{}

func (*jsonSpreader) spread(w io.Writer, roots []*Node) error {
	enc := json.NewEncoder(w)
	for _, root := range roots {
		jRoot := root.toJSONNode(nil)
		if err := enc.Encode(jRoot); err != nil {
			return err
		}
	}
	return nil
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
