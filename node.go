package gentree

import "fmt"

type node struct {
	name      string
	hierarchy int
	index     int // 上からscanしたときの順番
	branch    string
	parent    *node
	children  []*node
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

func newNode(row string, isTwoSpaces, isFourSpaces bool) *node {
	nodeIdx++

	var (
		myselfNode = &node{}
		name       = ""
		hierarchy  = rootHierarchyNum
	)
	var (
		spaceCnt   = 0
		isPrevChar = false
	)

	for _, r := range row {
		switch r {
		case hyphen:
			if isPrevChar {
				name += string(r)
				continue
			}
			if isTwoSpaces && spaceCnt%2 == 0 {
				tmp := spaceCnt / 2
				hierarchy += tmp
			}
			if isFourSpaces && spaceCnt%4 == 0 {
				tmp := spaceCnt / 4
				hierarchy += tmp
			}
			isPrevChar = false
		case space:
			if isPrevChar {
				name += string(r)
				continue
			}
			spaceCnt++
		case tab:
			hierarchy++
		default: // directry or file name char
			name += string(r)
			isPrevChar = true
		}
	}

	myselfNode.name = name
	myselfNode.hierarchy = hierarchy
	myselfNode.index = nodeIdx
	return myselfNode
}

func (n *node) buildBranch() string {
	if n.hierarchy == rootHierarchyNum {
		return fmt.Sprintf("%s\n", n.name)
	}
	return fmt.Sprintf("%s %s\n", n.branch, n.name)
}
