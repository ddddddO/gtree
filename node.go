package gentree

import "fmt"

type node struct {
	name      string
	hierarchy int // 階層
	index     int // 上から何番目のノードか
	branch    string
	parent    *node
	children  []*node
}

var nodeIdx int

// https://ja.wikipedia.org/wiki/ASCII
const (
	hyphen = 45
	space  = 32
	tab    = 9
)

const (
	rootHierarchyNum = 1
)

func newNode(row string, isTwoSpaces, isFourSpaces bool) *node {
	myselfNode := &node{}
	name := ""
	hierarchy := rootHierarchyNum
	nodeIdx++

	spaceCnt := 0
	isPrevChar := false
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
