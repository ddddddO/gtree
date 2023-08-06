package gtree

import (
	"errors"
	"fmt"

	md "github.com/ddddddO/gtree/markdown"
)

type nodeGenerator struct {
	parser *md.Parser
}

func newNodeGenerator() *nodeGenerator {
	return &nodeGenerator{
		parser: md.NewParser(),
	}
}

var (
	errEmptyText = errors.New("empty text")
)

type inputFormatError struct {
	// n uint64 // markdownの行数もあとで
	row string
}

func (ie *inputFormatError) Error() string {
	return fmt.Sprintf("incorrect input format: %s", ie.row)
}

func (ng *nodeGenerator) generate(row string, idx uint) (*Node, error) {
	markdown, err := ng.parser.Parse(row)
	if err != nil {
		return nil, ng.handleErr(err, row)
	}

	return newNode(
		markdown.Text(),
		markdown.Hierarchy(),
		idx,
	), nil
}

func (*nodeGenerator) handleErr(err error, row string) error {
	switch err {
	case md.ErrEmptyText:
		return errEmptyText
	case md.ErrIncorrectFormat:
		return &inputFormatError{
			row: row,
		}
	case md.ErrBlankLine:
		return nil
	}
	return err
}
