package gtree

import (
	"errors"
	"strings"
)

type spaceType int

const (
	spacesTab spaceType = iota
	spacesTwo
	spacesFour
)

type nodeGenerateStrategy interface {
	generate(row string, idx uint) (*Node, error)
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
	rootHierarchyNum uint = 1
)

var (
	errEmptyText       = errors.New("empty text")
	errIncorrectFormat = errors.New("incorrect input format")
)

type tabStrategy struct{}

func (*tabStrategy) generate(row string, idx uint) (*Node, error) {
	before, after, found := strings.Cut(row, hyphen)
	if !found {
		return nil, errIncorrectFormat
	}

	tabCount := strings.Count(before, tab)
	if tabCount != len(before) {
		return nil, errIncorrectFormat
	}

	hierarchy := uint(tabCount) + rootHierarchyNum
	text := strings.TrimPrefix(after, space)
	if len(text) == 0 {
		return nil, errEmptyText
	}

	return newNode(text, hierarchy, idx), nil
}

type twoSpacesStrategy struct{}

func (*twoSpacesStrategy) generate(row string, idx uint) (*Node, error) {
	before, after, found := strings.Cut(row, hyphen)
	if !found {
		return nil, errIncorrectFormat
	}

	spaceCount := strings.Count(before, space)
	if spaceCount != len(before) {
		return nil, errIncorrectFormat
	}
	if spaceCount%2 != 0 {
		return nil, errIncorrectFormat
	}

	hierarchy := uint(spaceCount/2) + rootHierarchyNum
	text := strings.TrimPrefix(after, space)
	if len(text) == 0 {
		return nil, errEmptyText
	}

	return newNode(text, hierarchy, idx), nil
}

type fourSpacesStrategy struct{}

func (*fourSpacesStrategy) generate(row string, idx uint) (*Node, error) {
	before, after, found := strings.Cut(row, hyphen)
	if !found {
		return nil, errIncorrectFormat
	}

	spaceCount := strings.Count(before, space)
	if spaceCount != len(before) {
		return nil, errIncorrectFormat
	}
	if spaceCount%4 != 0 {
		return nil, errIncorrectFormat
	}

	hierarchy := uint(spaceCount/4) + rootHierarchyNum
	text := strings.TrimPrefix(after, space)
	if len(text) == 0 {
		return nil, errEmptyText
	}

	return newNode(text, hierarchy, idx), nil
}
