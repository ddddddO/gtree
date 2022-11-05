package gtree

import (
	"errors"

	md "github.com/ddddddO/gtree/markdown"
)

type spaceType int

const (
	spacesTab spaceType = iota
	spacesTwo
	spacesFour
)

type nodeGenerator struct {
	parser *md.Parser
}

func newNodeGenerator(st spaceType) *nodeGenerator {
	var p *md.Parser
	switch st {
	case spacesTwo:
		p = md.NewParser(2)
	case spacesFour:
		p = md.NewParser(4)
	default:
		p = md.NewParser(1)
	}
	return &nodeGenerator{
		parser: p,
	}
}

var (
	errEmptyText       = errors.New("empty text")
	errIncorrectFormat = errors.New("incorrect input format")
)

func (ng *nodeGenerator) generate(row string, idx uint) (*Node, error) {
	markdown, err := ng.parser.Parse(row)
	if err != nil {
		return nil, ng.handleErr(err)
	}

	return newNode(
		markdown.Text(),
		markdown.Hierarchy(),
		idx,
	), nil
}

func (*nodeGenerator) handleErr(err error) error {
	switch err {
	case md.ErrEmptyText:
		return errEmptyText
	case md.ErrIncorrectFormat, md.ErrBlankLine:
		return errIncorrectFormat
	}
	return err
}
