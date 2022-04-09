package gtree

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"strings"

	"github.com/pkg/errors"
)

// Node is main struct for gtree.
type Node struct {
	Name      string `json:"value" yaml:"value" toml:"value"`
	hierarchy uint
	index     uint
	branch    branch
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

func (n *Node) isDirectlyUnderNode(node *Node) bool {
	return n.hierarchy == node.hierarchy+1
}

func (n *Node) isLastOfHierarchy() bool {
	lastIdx := len(n.parent.Children) - 1
	return n.index == n.parent.Children[lastIdx].index
}

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
}

func (n *Node) getBranch() string {
	if n.isRoot() {
		return fmt.Sprintf("%s\n", n.Name)
	}
	return fmt.Sprintf("%s %s\n", n.branch.value, n.Name)
}

func (n *Node) getPath() string {
	if n.isRoot() {
		return n.Name
	}
	return n.branch.path
}

func (n *Node) hasChild() bool {
	return len(n.Children) > 0
}

func generateRoots(r io.Reader, space spaceType) ([]*Node, error) {
	var (
		scanner          = bufio.NewScanner(r)
		stack            *stack
		counter          = newCounter()
		generateNodeFunc = space.decideGenerateFunc()
		roots            []*Node
	)

	for scanner.Scan() {
		currentNode := generateNodeFunc(scanner.Text(), counter.next())
		if err := currentNode.validate(); err != nil {
			return nil, err
		}

		if currentNode.isRoot() {
			counter.reset()
			roots = append(roots, currentNode)
			stack = newStack()
			stack.push(currentNode)
			continue
		}

		if stack == nil {
			return nil, errNilStack
		}

		stack.dfs(currentNode)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return roots, nil
}

var (
	errEmptyText       = errors.New("empty text")
	errIncorrectFormat = errors.New("incorrect input format")
)

func (n *Node) validate() error {
	if len(n.Name) == 0 {
		return errEmptyText
	}
	if n.hierarchy == 0 {
		return errIncorrectFormat
	}
	return nil
}

func (n *Node) validatePath() error {
	invalidChars := "/" // TODO: ディレクトリ名に含めてはまずそうなものをここに追加する
	if strings.ContainsAny(n.Name, invalidChars) {
		return errors.Errorf("invalid node name: %s", n.Name)
	}
	if !fs.ValidPath(n.branch.path) {
		return errors.Errorf("invalid path: %s", n.branch.path)
	}
	return nil
}

type generateFunc func(row string, idx uint) *Node

type spaceType int

const (
	tabSpaces spaceType = iota
	twoSpaces
	fourSpaces
)

func (st spaceType) decideGenerateFunc() generateFunc {
	switch st {
	case twoSpaces:
		return generateFuncTwoSpaces
	case fourSpaces:
		return generateFuncFourSpaces
	default:
		return generateFuncTab
	}
}

// https://ja.wikipedia.org/wiki/ASCII
const (
	hyphen = 45
	space  = 32
	tab    = 9
)

const (
	rootHierarchyNum = uint(1)
)

func generateFuncTab(row string, idx uint) *Node {
	var (
		text      = ""
		hierarchy = rootHierarchyNum
	)
	var (
		isPrevChar = false
		isRoot     = false
	)

	for i, r := range row {
		switch r {
		case hyphen:
			if i == 0 {
				isRoot = true
			}
			if isPrevChar {
				text += string(r)
				continue
			}
			isPrevChar = false
		case space:
			if isPrevChar {
				text += string(r)
				continue
			}
		case tab:
			hierarchy++
		default: // directry or file text char
			text += string(r)
			isPrevChar = true
		}
	}
	if !isRoot && hierarchy == rootHierarchyNum {
		hierarchy = 0
	}

	return newNode(text, hierarchy, idx)
}

func generateFuncTwoSpaces(row string, idx uint) *Node {
	var (
		text      = ""
		hierarchy = rootHierarchyNum
	)
	var (
		spaceCnt   = uint(0)
		isPrevChar = false
	)

	for _, r := range row {
		switch r {
		case hyphen:
			if isPrevChar {
				text += string(r)
				continue
			}
			if spaceCnt%2 == 0 {
				hierarchy += spaceCnt / 2
			}
			isPrevChar = false
		case space:
			if isPrevChar {
				text += string(r)
				continue
			}
			spaceCnt++
		default: // directry or file text char
			text += string(r)
			isPrevChar = true
		}
	}

	return newNode(text, hierarchy, idx)
}

func generateFuncFourSpaces(row string, idx uint) *Node {
	var (
		text      = ""
		hierarchy = rootHierarchyNum
	)
	var (
		spaceCnt   = uint(0)
		isPrevChar = false
	)

	for _, r := range row {
		switch r {
		case hyphen:
			if isPrevChar {
				text += string(r)
				continue
			}
			if spaceCnt%4 == 0 {
				hierarchy += spaceCnt / 4
			}
			isPrevChar = false
		case space:
			if isPrevChar {
				text += string(r)
				continue
			}
			spaceCnt++
		default: // directry or file text char
			text += string(r)
			isPrevChar = true
		}
	}

	return newNode(text, hierarchy, idx)
}
