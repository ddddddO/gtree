package gtree

import (
	"fmt"
	"io/fs"
	"path"
	"strings"
)

// Node is main struct for gtree.
type Node struct {
	value           string
	hierarchy       uint
	brnch           branch
	parent          *Node
	children        []*Node
	allowDuplicates bool
}

type branch struct {
	value strings.Builder
	path  string
}

func newNode(value string, hierarchy uint, options ...NodeOption) *Node {
	n := &Node{
		value:     value,
		hierarchy: hierarchy,
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
	tmp := make([]*Node, len(n.children)+1)
	copy(tmp, n.children)
	tmp[len(n.children)] = child
	n.children = tmp
}

func (n *Node) hasChild() bool {
	return len(n.children) > 0
}

func (n *Node) findChildByText(text string) *Node {
	for _, child := range n.children {
		if text == child.value {
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
	return n == n.parent.children[lastIdx]
}

const (
	rootHierarchyNum uint = 1
)

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
}

func (n *Node) setBranch(branchs ...string) {
	n.brnch.value.Reset()
	for _, v := range branchs {
		_, _ = n.brnch.value.WriteString(v)
	}
}

func (n *Node) appendBranch(branch string) {
	n.brnch.value.Grow(len(branch))
	n.brnch.value.WriteString(branch)
}

func (n *Node) prependBranch(branch string) {
	tmp := strings.Builder{}
	tmp.Grow(len(branch) + n.brnch.value.Len())
	tmp.WriteString(branch)
	tmp.WriteString(n.brnch.value.String())
	n.brnch.value = tmp
}

func (n *Node) branch() string {
	return n.brnch.value.String()
}

func (n *Node) setPath(paths ...string) {
	n.brnch.path = path.Join(paths...)
}

func (n *Node) validatePath() error {
	invalidChars := "/" // NOTE: ディレクトリ名に含めてはまずそうなものをここに追加する
	if strings.ContainsAny(n.value, invalidChars) {
		return fmt.Errorf("invalid node value: %s", n.value)
	}
	if !fs.ValidPath(n.path()) {
		return fmt.Errorf("invalid path: %s", n.path())
	}
	return nil
}

func (n *Node) path() string {
	if n.isRoot() {
		return n.value
	}
	return n.brnch.path
}

func (n *Node) clean() {
	n.setBranch("")
	n.setPath("")
}
