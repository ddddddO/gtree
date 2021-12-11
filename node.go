package gtree

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// Node is main struct for gtree.
type Node struct {
	Text      string `json:"value" yaml:"value" toml:"value"`
	hierarchy int
	index     int
	branch    string
	parent    *Node
	Children  []*Node `json:"children" yaml:"children" toml:"children"`
}

func newNode(text string, hierarchy, index int) *Node {
	return &Node{
		Text:      text,
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

func (n *Node) isDirectlyUnderParent(parent *Node) bool {
	return n.hierarchy == parent.hierarchy+1
}

func (n *Node) isLastOfHierarchy() bool {
	lastChildIndex := len(n.parent.Children) - 1
	return n.index == n.parent.Children[lastChildIndex].index
}

func (n *Node) isRoot() bool {
	return n.hierarchy == rootHierarchyNum
}

func (n *Node) getBranch() string {
	if n.isRoot() {
		return fmt.Sprintf("%s\n", n.Text)
	}
	return fmt.Sprintf("%s %s\n", n.branch, n.Text)
}

var (
	errEmptyText       = errors.New("empty text")
	errIncorrectFormat = errors.New("incorrect input format")
)

func (n *Node) validate() error {
	if len(n.Text) == 0 {
		return errEmptyText
	}
	if n.hierarchy == 0 {
		return errIncorrectFormat
	}
	return nil
}

type generateFunc func(row string) *Node

type spaceType string

const (
	twoSpaces  spaceType = "TWOSPACES"
	fourSpaces spaceType = "FOURSPACES"
	tabSpaces  spaceType = "TAB"
)

func decideGenerateFunc(space spaceType) generateFunc {
	switch space {
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
	rootHierarchyNum = 1
)

var (
	nodeIdx   int
	nodeIdxMu sync.Mutex
)

func generateFuncTab(row string) *Node {
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

	nodeIdxMu.Lock()
	defer nodeIdxMu.Unlock()

	nodeIdx++

	return newNode(text, hierarchy, nodeIdx)
}

func generateFuncTwoSpaces(row string) *Node {
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

	nodeIdxMu.Lock()
	defer nodeIdxMu.Unlock()

	nodeIdx++

	return newNode(text, hierarchy, nodeIdx)
}

func generateFuncFourSpaces(row string) *Node {
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

	nodeIdxMu.Lock()
	defer nodeIdxMu.Unlock()

	nodeIdx++

	return newNode(text, hierarchy, nodeIdx)
}
