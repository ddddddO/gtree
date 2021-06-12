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

func newNode(row string, isTwoSpaces, isFourSpaces bool) *node {
	myselfNode := &node{}
	name := ""
	hierarchy := 1
	nodeIdx++

	spaceCnt := 0
	isPrevChar := false
	for _, r := range row {
		// https://ja.wikipedia.org/wiki/ASCII
		switch r {
		case 45: // -
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
		case 32: // space
			if isPrevChar {
				name += string(r)
				continue
			}

			spaceCnt++
		case 9: // tab
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
	if n.hierarchy == 1 {
		return n.name + fmt.Sprintln()
	}
	return n.branch + " " + n.name + fmt.Sprintln()
}
