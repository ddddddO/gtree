package main

import (
	"errors"

	"github.com/ddddddO/gtree"
	"github.com/urfave/cli/v2"
)

type spacesType uint

const (
	spacesTwo spacesType = 1 << iota
	spacesFour
)

type stateIndentation struct {
	spaces spacesType
}

func optionIndentation(c *cli.Context) (gtree.Option, error) {
	s := &stateIndentation{}
	if c.Bool("two-spaces") {
		s.spaces |= spacesTwo
	}
	if c.Bool("four-spaces") {
		s.spaces |= spacesFour
	}

	return s.decideOption()
}

func (s *stateIndentation) decideOption() (gtree.Option, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	switch s.spaces {
	case spacesTwo:
		return gtree.WithIndentTwoSpaces(), nil
	case spacesFour:
		return gtree.WithIndentFourSpaces(), nil
	}
	return nil, nil
}

const spacesTab = spacesType(0)

func (s *stateIndentation) validate() error {
	switch s.spaces {
	case spacesTab, spacesTwo, spacesFour:
		return nil
	}
	return errors.New(`choose either "two-spaces(ts)" or "four-spaces(fs)" or blank.`)
}
