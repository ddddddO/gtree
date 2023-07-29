//go:build !wasm

package gtree

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/fatih/color"
	toml "github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func newSpreaderSimple(encode encode) spreaderSimple {
	switch encode {
	case encodeJSON:
		return &jsonSpreaderSimple{}
	case encodeYAML:
		return &yamlSpreaderSimple{}
	case encodeTOML:
		return &tomlSpreaderSimple{}
	default:
		return &defaultSpreaderSimple{}
	}
}

type encode int

const (
	encodeDefault encode = iota
	encodeJSON
	encodeYAML
	encodeTOML
)

type defaultSpreaderSimple struct{}

func (ds *defaultSpreaderSimple) spread(w io.Writer, roots []*Node) error {
	branches := ""
	for _, root := range roots {
		branches += ds.spreadBranch(root)
	}
	return ds.write(w, branches)
}

func (*defaultSpreaderSimple) spreadBranch(current *Node) string {
	ret := current.name + "\n"
	if !current.isRoot() {
		ret = current.branch() + " " + current.name + "\n"
	}

	for _, child := range current.children {
		ret += (*defaultSpreaderSimple)(nil).spreadBranch(child)
	}
	return ret
}

func (*defaultSpreaderSimple) write(w io.Writer, in string) error {
	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(in); err != nil {
		return err
	}
	return buf.Flush()
}

type jsonSpreaderSimple struct{}

func (*jsonSpreaderSimple) spread(w io.Writer, roots []*Node) error {
	enc := json.NewEncoder(w)
	for _, root := range roots {
		jRoot := toFormattedNode(root, &jsonNode{Name: root.name})
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

func (jn *jsonNode) setChild(name string) {
	jn.Children = append(jn.Children, &jsonNode{Name: name})
}

func (jn *jsonNode) getChild(i int) sitter {
	return jn.Children[i]
}

type tomlSpreaderSimple struct{}

func (*tomlSpreaderSimple) spread(w io.Writer, roots []*Node) error {
	enc := toml.NewEncoder(w)
	for _, root := range roots {
		tRoot := toFormattedNode(root, &tomlNode{Name: root.name})
		if err := enc.Encode(tRoot); err != nil {
			return err
		}
	}
	return nil
}

type tomlNode struct {
	Name     string      `toml:"value"`
	Children []*tomlNode `toml:"children"`
}

func (tn *tomlNode) setChild(name string) {
	tn.Children = append(tn.Children, &tomlNode{Name: name})
}

func (tn *tomlNode) getChild(i int) sitter {
	return tn.Children[i]
}

type yamlSpreaderSimple struct{}

func (*yamlSpreaderSimple) spread(w io.Writer, roots []*Node) error {
	enc := yaml.NewEncoder(w)
	for _, root := range roots {
		yRoot := toFormattedNode(root, &yamlNode{Name: root.name})
		if err := enc.Encode(yRoot); err != nil {
			return err
		}
	}
	return nil
}

type yamlNode struct {
	Name     string      `yaml:"value"`
	Children []*yamlNode `yaml:"children"`
}

func (yn *yamlNode) setChild(name string) {
	yn.Children = append(yn.Children, &yamlNode{Name: name})
}

func (yn *yamlNode) getChild(i int) sitter {
	return yn.Children[i]
}

type sitter interface {
	setChild(string)
	getChild(int) sitter
}

func toFormattedNode[T sitter](parent *Node, fParent T) T {
	if !parent.hasChild() {
		return fParent
	}

	for i := range parent.children {
		fParent.setChild(parent.children[i].name)
		toFormattedNode(parent.children[i], fParent.getChild(i).(T))
	}

	return fParent
}

func newColorizeSpreaderSimple(fileExtensions []string) spreaderSimple {
	return &colorizeSpreaderSimple{
		defaultSpreaderSimple: &defaultSpreaderSimple{},

		fileConsiderer: newFileConsiderer(fileExtensions),
		fileColor:      color.New(color.Bold, color.FgHiCyan),
		fileCounter:    newCounter(),

		dirColor:   color.New(color.FgGreen),
		dirCounter: newCounter(),
	}
}

type colorizeSpreaderSimple struct {
	*defaultSpreaderSimple

	fileConsiderer *fileConsiderer
	fileColor      *color.Color
	fileCounter    *counter

	dirColor   *color.Color
	dirCounter *counter
}

func (cs *colorizeSpreaderSimple) spread(w io.Writer, roots []*Node) error {
	ret := ""
	for _, root := range roots {
		cs.fileCounter.reset()
		cs.dirCounter.reset()
		ret += fmt.Sprintf("%s\n%s\n", cs.spreadBranch(root), cs.summary())
	}
	return cs.write(w, ret)
}

func (cs *colorizeSpreaderSimple) spreadBranch(current *Node) string {
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

func (cs *colorizeSpreaderSimple) colorize(current *Node) string {
	if cs.fileConsiderer.isFile(current) {
		_ = cs.fileCounter.next()
		return cs.fileColor.Sprint(current.name)
	} else {
		_ = cs.dirCounter.next()
		return cs.dirColor.Sprint(current.name)
	}
}

func (cs *colorizeSpreaderSimple) summary() string {
	return fmt.Sprintf(
		"%d directories, %d files",
		cs.dirCounter.current(),
		cs.fileCounter.current(),
	)
}

var (
	_ spreaderSimple = (*defaultSpreaderSimple)(nil)
	_ spreaderSimple = (*jsonSpreaderSimple)(nil)
	_ spreaderSimple = (*yamlSpreaderSimple)(nil)
	_ spreaderSimple = (*tomlSpreaderSimple)(nil)
	_ spreaderSimple = (*colorizeSpreaderSimple)(nil)
)
