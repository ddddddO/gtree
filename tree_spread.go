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
	for _, child := range current.Children {
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
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}

type tomlSpreader struct{}

func (*tomlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := toml.NewEncoder(w)
	for _, root := range roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}

type yamlSpreader struct{}

func (*yamlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := yaml.NewEncoder(w)
	for _, root := range roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}
