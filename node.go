package gtree

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Node is main struct for gtree.
type Node struct {
	Name      string `json:"value" yaml:"value" toml:"value"`
	hierarchy uint
	index     uint
	brnch     branch
	parent    *Node
	Children  []*Node `json:"children" yaml:"children" toml:"children"`
}

type branch struct {
	value string
	path  string
}

func newNode(name string, hierarchy, index uint) *Node {
	return &Node{
		Name:      name,
		hierarchy: hierarchy,
		index:     index,
	}
}

func (n *Node) setParent(parent *Node) {
	n.parent = parent
}

func (n *Node) addChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) isDirectlyUnder(node *Node) bool {
	return n.hierarchy == node.hierarchy+1
}

func (n *Node) isLastOfHierarchy() bool {
	lastIdx := len(n.parent.Children) - 1
	return n.index == n.parent.Children[lastIdx].index
}

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
}

func (n *Node) prettyBranch() string {
	if n.isRoot() {
		return fmt.Sprintf("%s\n", n.Name)
	}
	return fmt.Sprintf("%s %s\n", n.brnch.value, n.Name)
}

func (n *Node) branch() string {
	return n.brnch.value
}

func (n *Node) setBranch(branchs ...string) {
	ret := ""
	for _, v := range branchs {
		ret += v
	}
	n.brnch.value = ret
}

func (n *Node) path() string {
	if n.isRoot() {
		return n.Name
	}
	return n.brnch.path
}

func (n *Node) setPath(paths ...string) {
	n.brnch.path = filepath.Join(paths...)
}

func (n *Node) hasChild() bool {
	return len(n.Children) > 0
}

func (n *Node) validatePath() error {
	invalidChars := "/" // TODO: ディレクトリ名に含めてはまずそうなものをここに追加する
	if strings.ContainsAny(n.Name, invalidChars) {
		return errors.Errorf("invalid node name: %s", n.Name)
	}
	if !fs.ValidPath(n.path()) {
		return errors.Errorf("invalid path: %s", n.path())
	}
	return nil
}
