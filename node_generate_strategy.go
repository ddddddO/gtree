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

const (
	empty  = ""
	hyphen = "-"
	space  = " "
	tab    = "	"
)

const (
	invalidHierarchyNum uint = 0
	rootHierarchyNum    uint = 1
)

type tabStrategy struct{}

func (*tabStrategy) generate(row string, idx uint) *Node {
	before, after, found := strings.Cut(row, hyphen)
	if !found {
		return newNode(empty, invalidHierarchyNum, idx)
	}

	tabCount := strings.Count(before, tab)
	if tabCount != len(before) {
		return newNode(empty, invalidHierarchyNum, idx)
	}

	hierarchy := uint(tabCount) + rootHierarchyNum
	text := strings.TrimPrefix(after, space)
	return newNode(text, hierarchy, idx)
}

type twoSpacesStrategy struct{}

func (*twoSpacesStrategy) generate(row string, idx uint) *Node {
	before, after, found := strings.Cut(row, hyphen)
	if !found {
		return newNode(empty, invalidHierarchyNum, idx)
	}

	spaceCount := strings.Count(before, space)
	if spaceCount != len(before) {
		return newNode(empty, invalidHierarchyNum, idx)
	}
	if spaceCount%2 != 0 {
		return newNode(empty, invalidHierarchyNum, idx)
	}

	hierarchy := uint(spaceCount/2) + rootHierarchyNum
	text := strings.TrimPrefix(after, space)
	return newNode(text, hierarchy, idx)
}

type fourSpacesStrategy struct{}

func (*fourSpacesStrategy) generate(row string, idx uint) *Node {
	before, after, found := strings.Cut(row, hyphen)
	if !found {
		return newNode(empty, invalidHierarchyNum, idx)
	}

	spaceCount := strings.Count(before, space)
	if spaceCount != len(before) {
		return newNode(empty, invalidHierarchyNum, idx)
	}
	if spaceCount%4 != 0 {
		return newNode(empty, invalidHierarchyNum, idx)
	}

	hierarchy := uint(spaceCount/4) + rootHierarchyNum
	text := strings.TrimPrefix(after, space)
	return newNode(text, hierarchy, idx)
}
