package gtree

import (
	"bufio"
	"encoding/json"
	"io"

	toml "github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

// 関心事はtreeを出力すること
type spreader interface {
	spread(io.Writer, []*Node) error
}

func newSpreader(encode encode) spreader {
	if s, ok := spreaders[encode]; ok {
		return s
	}
	return &defaultSpreader{}
}

// NOTE: 微妙。分岐は無くせるし追加は楽そうだが、グローバルに持ってしまってる。
var spreaders = map[encode]spreader{
	encodeDefault: &defaultSpreader{},
	encodeJSON:    &jsonSpreader{},
	encodeTOML:    &tomlSpreader{},
	encodeYAML:    &yamlSpreader{},
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
		branches += ds.spreadBranch(root, "")
	}
	return ds.write(w, branches)
}

func (*defaultSpreader) spreadBranch(current *Node, out string) string {
	out += current.prettyBranch()
	for _, child := range current.children {
		out = (*defaultSpreader)(nil).spreadBranch(child, out)
	}
	return out
}

func (*defaultSpreader) write(w io.Writer, in string) error {
	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(in); err != nil {
		return err
	}
	return buf.Flush()
}

type jsonSpreader struct{}

func (*jsonSpreader) spread(w io.Writer, roots []*Node) error {
	enc := json.NewEncoder(w)
	for _, root := range roots {
		jRoot := (*jsonSpreader)(nil).toJsonNode(&jsonNode{Name: root.name}, root.children)
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

func (*jsonSpreader) toJsonNode(jParent *jsonNode, children []*Node) *jsonNode {
	if len(children) == 0 {
		return nil
	}

	jChildren := make([]*jsonNode, len(children))
	for i := range children {
		jChildren[i] = &jsonNode{Name: children[i].name}
		(*jsonSpreader)(nil).toJsonNode(jChildren[i], children[i].children)
	}
	jParent.Children = jChildren

	return jParent
}

type tomlSpreader struct{}

func (*tomlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := toml.NewEncoder(w)
	for _, root := range roots {
		tRoot := (*tomlSpreader)(nil).toTomlNode(&tomlNode{Name: root.name}, root.children)
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

func (*tomlSpreader) toTomlNode(tParent *tomlNode, children []*Node) *tomlNode {
	if len(children) == 0 {
		return nil
	}

	tChildren := make([]*tomlNode, len(children))
	for i := range children {
		tChildren[i] = &tomlNode{Name: children[i].name}
		(*tomlSpreader)(nil).toTomlNode(tChildren[i], children[i].children)
	}
	tParent.Children = tChildren

	return tParent
}

type yamlSpreader struct{}

func (*yamlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := yaml.NewEncoder(w)
	for _, root := range roots {
		yRoot := (*yamlSpreader)(nil).toYamlNode(&yamlNode{Name: root.name}, root.children)
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

func (*yamlSpreader) toYamlNode(yParent *yamlNode, children []*Node) *yamlNode {
	if len(children) == 0 {
		return nil
	}

	yChildren := make([]*yamlNode, len(children))
	for i := range children {
		yChildren[i] = &yamlNode{Name: children[i].name}
		(*yamlSpreader)(nil).toYamlNode(yChildren[i], children[i].children)
	}
	yParent.Children = yChildren

	return yParent
}
