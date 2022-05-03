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
	invalidHierarchyNum uint = 0
	rootHierarchyNum    uint = 1
)

type tabStrategy struct{}

func (*tabStrategy) generate(row string, idx uint) *Node {
	var (
		hierarchy     = rootHierarchyNum
		startText     = 2
		containHyphen = false
	)

BREAK:
	for _, r := range row {
		switch r {
		case hyphen:
			containHyphen = true
			break BREAK
		case tab:
			hierarchy++
			startText++
		default: // directry or file text char
			hierarchy = invalidHierarchyNum
			break BREAK
		}
	}

	text := ""
	if !containHyphen {
		return newNode(text, invalidHierarchyNum, idx)
	}
	if hierarchy == invalidHierarchyNum {
		return newNode(text, hierarchy, idx)
	}
	if startText < len(row) {
		text = row[startText:len(row)]
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
		spaceCnt       = uint(0)
		existsPrevChar = false
		isRoot         = false
	)

	for i, r := range row {
		switch r {
		case hyphen:
			if i == 0 {
				isRoot = true
				continue
			}
			if existsPrevChar {
				text += string(r)
				existsPrevChar = true
				continue
			}
			if spaceCnt%2 == 0 {
				hierarchy += spaceCnt / 2
			}
			existsPrevChar = false
		case space:
			if existsPrevChar {
				text += string(r)
				existsPrevChar = true
				continue
			}
			spaceCnt++
		default: // directry or file text char
			text += string(r)
			existsPrevChar = true
		}
	}

	hierarchy = decideHierarchy(isRoot, hierarchy)

	return newNode(text, hierarchy, idx)
}

type fourSpacesStrategy struct{}

func (*fourSpacesStrategy) generate(row string, idx uint) *Node {
	var (
		text      = ""
		hierarchy = rootHierarchyNum
	)
	var (
		spaceCnt       = uint(0)
		existsPrevChar = false
		isRoot         = false
	)

	for i, r := range row {
		switch r {
		case hyphen:
			if i == 0 {
				isRoot = true
				continue
			}
			if existsPrevChar {
				text += string(r)
				existsPrevChar = true
				continue
			}
			if spaceCnt%4 == 0 {
				hierarchy += spaceCnt / 4
			}
			existsPrevChar = false
		case space:
			if existsPrevChar {
				text += string(r)
				existsPrevChar = true
				continue
			}
			spaceCnt++
		default: // directry or file text char
			text += string(r)
			existsPrevChar = true
		}
	}

	hierarchy = decideHierarchy(isRoot, hierarchy)

	return newNode(text, hierarchy, idx)
}

func decideHierarchy(isRoot bool, hierarchy uint) uint {
	if !isRoot && hierarchy == rootHierarchyNum {
		return invalidHierarchyNum
	}
	return hierarchy
}
