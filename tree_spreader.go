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

func newSpreader(encode encode) spreader {
	switch encode {
	case encodeJSON:
		return &jsonSpreader{}
	case encodeYAML:
		return &yamlSpreader{}
	case encodeTOML:
		return &tomlSpreader{}
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

func (ds *defaultSpreader) spread(w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		branches := ""
		for root := range roots {
			branches += ds.spreadBranch(root)
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

func (cs *colorizeSpreader) spread(w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		ret := ""
		for root := range roots {
			cs.fileCounter.reset()
			cs.dirCounter.reset()
			ret += fmt.Sprintf("%s\n%s", cs.spreadBranch(root), cs.summary())
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
		"%d directories, %d files\n",
		cs.dirCounter.current(),
		cs.fileCounter.current(),
	)
}

type jsonSpreader struct{}

func (*jsonSpreader) spread(w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := json.NewEncoder(w)
		for root := range roots {
			jRoot := root.toJSONNode(nil)
			if err := enc.Encode(jRoot); err != nil {
				errc <- err
				return
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

type tomlSpreader struct{}

func (*tomlSpreader) spread(w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := toml.NewEncoder(w)
		for root := range roots {
			tRoot := root.toTOMLNode(nil)
			if err := enc.Encode(tRoot); err != nil {
				errc <- err
				return
			}
		}
	}()

	return errc
}

type tomlNode struct {
	Name     string      `toml:"value"`
	Children []*tomlNode `toml:"children"`
}

func (parent *Node) toTOMLNode(tParent *tomlNode) *tomlNode {
	if tParent == nil {
		tParent = &tomlNode{Name: parent.name}
	}
	if !parent.hasChild() {
		return tParent
	}

	tParent.Children = make([]*tomlNode, len(parent.children))
	for i := range parent.children {
		tParent.Children[i] = &tomlNode{Name: parent.children[i].name}
		_ = parent.children[i].toTOMLNode(tParent.Children[i])
	}

	return tParent
}

type yamlSpreader struct{}

func (*yamlSpreader) spread(w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := yaml.NewEncoder(w)
		for root := range roots {
			yRoot := root.toYAMLNode(nil)
			if err := enc.Encode(yRoot); err != nil {
				errc <- err
				return
			}
		}
	}()

	return errc
}

type yamlNode struct {
	Name     string      `yaml:"value"`
	Children []*yamlNode `yaml:"children"`
}

func (parent *Node) toYAMLNode(yParent *yamlNode) *yamlNode {
	if yParent == nil {
		yParent = &yamlNode{Name: parent.name}
	}
	if !parent.hasChild() {
		return yParent
	}

	yParent.Children = make([]*yamlNode, len(parent.children))
	for i := range parent.children {
		yParent.Children[i] = &yamlNode{Name: parent.children[i].name}
		_ = parent.children[i].toYAMLNode(yParent.Children[i])
	}

	return yParent
}

var (
	_ spreader = (*defaultSpreader)(nil)
	_ spreader = (*colorizeSpreader)(nil)
	_ spreader = (*jsonSpreader)(nil)
	_ spreader = (*yamlSpreader)(nil)
	_ spreader = (*tomlSpreader)(nil)
)
