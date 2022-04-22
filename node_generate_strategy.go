package gtree

type spaceType int

const (
	spacesTab spaceType = iota
	spacesTwo
	spacesFour
)

type nodeGenerateStrategy interface {
	generate(row string, idx uint) *Node
}

func newStrategy(st spaceType) nodeGenerateStrategy {
	switch st {
	case spacesTwo:
		return &twoSpacesStrategy{}
	case spacesFour:
		return &fourSpacesStrategy{}
	default:
		return &tabStrategy{}
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

type tabStrategy struct{}

func (*tabStrategy) generate(row string, idx uint) *Node {
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

type twoSpacesStrategy struct{}

func (*twoSpacesStrategy) generate(row string, idx uint) *Node {
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

type fourSpacesStrategy struct{}

func (*fourSpacesStrategy) generate(row string, idx uint) *Node {
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
