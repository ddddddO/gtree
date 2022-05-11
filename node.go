package gtree

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// Node is main struct for gtree.
type Node struct {
	name      string
	hierarchy uint
	index     uint
	brnch     branch
	parent    *Node
	children  []*Node
}

type branch struct {
	value string
	path  string
}

func newNode(name string, hierarchy, index uint) *Node {
	return &Node{
		name:      name,
		hierarchy: hierarchy,
		index:     index,
	}
}

func (n *Node) setParent(parent *Node) {
	n.parent = parent
}

func (n *Node) addChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) hasChild() bool {
	return len(n.children) > 0
}

func (n *Node) isDirectlyUnder(node *Node) bool {
	return n.hierarchy == node.hierarchy+1
}

func (n *Node) isLastOfHierarchy() bool {
	lastIdx := len(n.parent.children) - 1
	return n.index == n.parent.children[lastIdx].index
}

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
}

func (n *Node) cleanBranch() {
	n.brnch.value = ""
	n.brnch.path = ""
}

func (n *Node) setBranch(branchs ...string) {
	ret := ""
	for _, v := range branchs {
		ret += v
	}
	n.brnch.value = ret
}

func (n *Node) branch() string {
	return n.brnch.value
}

func (n *Node) prettyBranch() string {
	if n.isRoot() {
		return fmt.Sprintf("%s\n", n.name)
	}
	return fmt.Sprintf("%s %s\n", n.brnch.value, n.name)
}

func (n *Node) setPath(paths ...string) {
	n.brnch.path = filepath.Join(paths...)
}

func (n *Node) path() string {
	if n.isRoot() {
		return n.name
	}
	return n.brnch.path
}

func (n *Node) validatePath() error {
	invalidChars := "/" // TODO: ディレクトリ名に含めてはまずそうなものをここに追加する
	if strings.ContainsAny(n.name, invalidChars) {
		return fmt.Errorf("invalid node name: %s", n.name)
	}
	if !fs.ValidPath(n.path()) {
		return fmt.Errorf("invalid path: %s", n.path())
	}
	return nil
}
