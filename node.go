package gtree

import (
	"fmt"
	"io/fs"
	"path"
	"strings"
)

// Node is main struct for gtree.
type Node struct {
	name            string
	hierarchy       uint
	index           uint
	brnch           branch
	parent          *Node
	children        []*Node
	allowDuplicates bool
}

type branch struct {
	value string
	path  string
}

func newNode(name string, hierarchy, index uint, options ...NodeOption) *Node {
	n := &Node{
		name:      name,
		hierarchy: hierarchy,
		index:     index,
	}
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt(n)
	}
	return n
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

func (n *Node) findChildByText(text string) *Node {
	for _, child := range n.children {
		if text == child.name {
			return child
		}
	}
	return nil
}

func (n *Node) isDirectlyUnder(node *Node) bool {
	if node == nil {
		return false
	}
	return n.hierarchy == node.hierarchy+1
}

func (n *Node) isLastOfHierarchy() bool {
	if n.parent == nil {
		return false
	}

	lastIdx := len(n.parent.children) - 1
	return n.index == n.parent.children[lastIdx].index
}

const (
	rootHierarchyNum uint = 1
)

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
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

func (n *Node) setPath(paths ...string) {
	n.brnch.path = path.Join(paths...)
}

func (n *Node) validatePath() error {
	invalidChars := "/" // NOTE: ディレクトリ名に含めてはまずそうなものをここに追加する
	if strings.ContainsAny(n.name, invalidChars) {
		return fmt.Errorf("invalid node name: %s", n.name)
	}
	if !fs.ValidPath(n.path()) {
		return fmt.Errorf("invalid path: %s", n.path())
	}
	return nil
}

func (n *Node) path() string {
	if n.isRoot() {
		return n.name
	}
	return n.brnch.path
}

func (n *Node) clean() {
	n.setBranch("")
	n.setPath("")
}
