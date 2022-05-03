package gtree

import (
	"strings"
)

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
	before, after, found := strings.Cut(row, "-")
	if !found {
		return newNode("", invalidHierarchyNum, idx)
	}

	tabCount := strings.Count(before, "	")
	if tabCount != len(before) {
		return newNode("", invalidHierarchyNum, idx)
	}

	hierarchy := uint(tabCount) + rootHierarchyNum
	text := strings.TrimPrefix(after, " ")
	return newNode(text, hierarchy, idx)
}

type twoSpacesStrategy struct{}

func (*twoSpacesStrategy) generate(row string, idx uint) *Node {
	before, after, found := strings.Cut(row, "-")
	if !found {
		return newNode("", invalidHierarchyNum, idx)
	}

	spaceCount := strings.Count(before, " ")
	if spaceCount != len(before) {
		return newNode("", invalidHierarchyNum, idx)
	}
	if spaceCount%2 != 0 {
		return newNode("", invalidHierarchyNum, idx)
	}

	hierarchy := uint(spaceCount/2) + rootHierarchyNum
	text := strings.TrimPrefix(after, " ")
	return newNode(text, hierarchy, idx)
}

type fourSpacesStrategy struct{}

func (*fourSpacesStrategy) generate(row string, idx uint) *Node {
	before, after, found := strings.Cut(row, "-")
	if !found {
		return newNode("", invalidHierarchyNum, idx)
	}

	spaceCount := strings.Count(before, " ")
	if spaceCount != len(before) {
		return newNode("", invalidHierarchyNum, idx)
	}
	if spaceCount%4 != 0 {
		return newNode("", invalidHierarchyNum, idx)
	}

	hierarchy := uint(spaceCount/4) + rootHierarchyNum
	text := strings.TrimPrefix(after, " ")
	return newNode(text, hierarchy, idx)
}
