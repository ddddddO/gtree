//go:build !tinywasm

package gtree

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"iter"

	"github.com/fatih/color"
	"github.com/goccy/go-yaml"
	toml "github.com/pelletier/go-toml/v2"
)

func newSpreaderSimple(encode encode) spreaderSimple {
	switch encode {
	case encodeJSON:
		return newJSONSpreaderSimple()
	case encodeYAML:
		return newYAMLSpreaderSimple()
	case encodeTOML:
		return newTOMLSpreaderSimple()
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

type defaultSpreaderSimple struct {
	w io.Writer
}

func (ds *defaultSpreaderSimple) spread(w io.Writer, roots []*Node) error {
	ds.w = w
	for _, root := range roots {
		ds.spreadBranch(root)
	}
	return nil
}

func (ds *defaultSpreaderSimple) spreadIter(w io.Writer, rootIter iter.Seq2[*Node, error]) iter.Seq[error] {
	return func(yield func(error) bool) {
		next, stop := iter.Pull2(rootIter)
		defer stop()

		ds.w = w
		for {
			root, err, ok := next()
			if !ok {
				return
			}
			if err != nil {
				yield(err)
				return
			}

			ds.spreadBranch(root)
		}
	}
}

func (ds *defaultSpreaderSimple) spreadBranch(current *Node) {
	ret := current.name + "\n"
	if !current.isRoot() {
		ret = current.branch() + " " + current.name + "\n"
	}
	fmt.Fprint(ds.w, ret)

	for _, child := range current.children {
		ds.spreadBranch(child)
	}
}

type formattedSpreaderSimple[T sitter] struct {
	formattedRoot func(string) T
	encode        func(io.Writer) func(any) error
}

func newJSONSpreaderSimple() *formattedSpreaderSimple[*jsonNode] {
	return &formattedSpreaderSimple[*jsonNode]{
		formattedRoot: func(name string) *jsonNode {
			return &jsonNode{Name: name}
		},
		encode: func(w io.Writer) func(any) error {
			return json.NewEncoder(w).Encode
		},
	}
}

func newYAMLSpreaderSimple() *formattedSpreaderSimple[*yamlNode] {
	return &formattedSpreaderSimple[*yamlNode]{
		formattedRoot: func(name string) *yamlNode {
			return &yamlNode{Name: name}
		},
		encode: func(w io.Writer) func(any) error {
			return yaml.NewEncoder(w).Encode
		},
	}
}

func newTOMLSpreaderSimple() *formattedSpreaderSimple[*tomlNode] {
	return &formattedSpreaderSimple[*tomlNode]{
		formattedRoot: func(name string) *tomlNode {
			return &tomlNode{Name: name}
		},
		encode: func(w io.Writer) func(any) error {
			return toml.NewEncoder(w).Encode
		},
	}
}

func (f *formattedSpreaderSimple[T]) spread(w io.Writer, roots []*Node) error {
	encode := f.encode(w)
	for _, root := range roots {
		fRoot := toFormattedNode(root, f.formattedRoot(root.name))
		if err := encode(fRoot); err != nil {
			return err
		}
	}
	return nil
}

func (f *formattedSpreaderSimple[T]) spreadIter(w io.Writer, rootIter iter.Seq2[*Node, error]) iter.Seq[error] {
	return func(yield func(error) bool) {
		encode := f.encode(w)
		next, stop := iter.Pull2(rootIter)
		defer stop()

		for {
			root, err, ok := next()
			if !ok {
				return
			}
			if err != nil {
				yield(err)
				return
			}

			fRoot := toFormattedNode(root, f.formattedRoot(root.name))
			if err := encode(fRoot); err != nil {
				yield(err)
				return
			}
		}

	}
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

func (cs *colorizeSpreaderSimple) spreadIter(w io.Writer, rootIter iter.Seq2[*Node, error]) iter.Seq[error] {
	return func(yield func(error) bool) {
		next, stop := iter.Pull2(rootIter)
		defer stop()

		for {
			root, err, ok := next()
			if !ok {
				return
			}
			if err != nil {
				yield(err)
				return
			}

			cs.fileCounter.reset()
			cs.dirCounter.reset()
			cs.write(w, fmt.Sprintf("%s\n%s\n", cs.spreadBranch(root), cs.summary()))
		}
	}
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

func (*colorizeSpreaderSimple) write(w io.Writer, in string) error {
	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(in); err != nil {
		return err
	}
	return buf.Flush()
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
	_ spreaderSimple = (*formattedSpreaderSimple[sitter])(nil)
	_ spreaderSimple = (*colorizeSpreaderSimple)(nil)
)
