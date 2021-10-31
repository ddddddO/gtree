package gtree

import (
	"fmt"

	"github.com/pkg/errors"
)

// Node is main struct for gtree.
type Node struct {
	text      string
	hierarchy int
	index     int
	branch    string
	parent    *Node
	children  []*Node
}

func newNode(text string, hierarchy, index int) *Node {
	return &Node{
		text:      text,
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

func (n *Node) isDirectlyUnderParent(parent *Node) bool {
	return n.hierarchy == parent.hierarchy+1
}

func (n *Node) isLastOfHierarchy() bool {
	lastChildIndex := len(n.parent.children) - 1
	return n.index == n.parent.children[lastChildIndex].index
}

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
}

func (n *Node) getBranch() string {
	if n.isRoot() {
		return fmt.Sprintf("%s\n", n.text)
	}
	return fmt.Sprintf("%s %s\n", n.branch, n.text)
}

var (
	errEmptyText       = errors.New("empty text")
	errIncorrectFormat = errors.New("incorrect input format")
)

func (n *Node) validate() error {
	if len(n.text) == 0 {
		return errEmptyText
	}
	if n.hierarchy == 0 {
		return errIncorrectFormat
	}
	return nil
}

type generateFunc func(row string) *Node

func decideGenerateFunc(conf *config) generateFunc {
	if conf.isTwoSpaces {
		return generateFuncIfTwoSpaces
	}
	if conf.isFourSpaces {
		return generateFuncIfFourSpaces
	}
	return generateFuncIfTab
}

// https://ja.wikipedia.org/wiki/ASCII
const (
	hyphen = 45
	space  = 32
	tab    = 9
)

const (
	rootHierarchyNum = 1
)

var nodeIdx int

func generateFuncIfTab(row string) *Node {
	nodeIdx++

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

	return newNode(text, hierarchy, nodeIdx)
}

func generateFuncIfTwoSpaces(row string) *Node {
	nodeIdx++

	var (
		text      = ""
		hierarchy = rootHierarchyNum
	)
	var (
		spaceCnt   = 0
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

	return newNode(text, hierarchy, nodeIdx)
}

func generateFuncIfFourSpaces(row string) *Node {
	nodeIdx++

	var (
		text      = ""
		hierarchy = rootHierarchyNum
	)
	var (
		spaceCnt   = 0
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

	return newNode(text, hierarchy, nodeIdx)
}
