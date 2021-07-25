package gtree

import (
	"fmt"

	"github.com/pkg/errors"
)

type Node struct {
	text      string
	hierarchy int
	index     int // 上からscanしたときの順番
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

func (n *Node) isLastNodeOfHierarchy() bool {
	lastChildIndex := len(n.parent.children) - 1
	return n.index == n.parent.children[lastChildIndex].index
}

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
}

func (n *Node) buildBranch() string {
	if n.isRoot() {
		return fmt.Sprintf("%s\n", n.text)
	}
	return fmt.Sprintf("%s %s\n", n.branch, n.text)
}

var (
	ErrEmptyText       = errors.New("empty text")
	ErrIncorrectFormat = errors.New("incorrect input format")
)

func (n *Node) validate() error {
	if len(n.text) == 0 {
		return ErrEmptyText
	}
	if n.hierarchy == 0 {
		return ErrIncorrectFormat
	}
	return nil
}

type nodeGenerator interface {
	generate(row string) *Node
}

type nodeGeneratorTab struct{}
type nodeGeneratorTwoSpaces struct{}
type nodeGeneratorFourSpaces struct{}

func newNodeGenerator(conf Config) nodeGenerator {
	if conf.IsTwoSpaces {
		return &nodeGeneratorTwoSpaces{}
	}
	if conf.IsFourSpaces {
		return &nodeGeneratorFourSpaces{}
	}
	return &nodeGeneratorTab{}
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

func (*nodeGeneratorTab) generate(row string) *Node {
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

func (*nodeGeneratorTwoSpaces) generate(row string) *Node {
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

func (*nodeGeneratorFourSpaces) generate(row string) *Node {
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
