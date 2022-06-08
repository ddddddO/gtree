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

// 関心事はtreeを出力すること
type spreader interface {
	spread(io.Writer, []*Node) error
}

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
	ret := current.prettyBranch()
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
	branches := ""
	for _, root := range roots {
		branches += cs.spreadBranch(root)
	}

	ret := fmt.Sprintf("%s\n%s", branches, cs.summary())
	return cs.write(w, ret)
}

func (cs *colorizeSpreader) spreadBranch(current *Node) string {
	cs.colorize(current)
	ret := current.prettyBranch()
	for _, child := range current.children {
		ret += cs.spreadBranch(child)
	}
	return ret
}

func (cs *colorizeSpreader) colorize(current *Node) {
	if cs.fileConsiderer.nodeIsFile(current) {
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

type tomlSpreader struct{}

func (*tomlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := toml.NewEncoder(w)
	for _, root := range roots {
		tRoot := root.toTOMLNode(nil)
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

func (*yamlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := yaml.NewEncoder(w)
	for _, root := range roots {
		yRoot := root.toYAMLNode(nil)
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
