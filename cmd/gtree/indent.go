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

func newStateIndentation(c *cli.Context) *stateIndentation {
	s := &stateIndentation{}

	if c.Bool("two-spaces") {
		s.spaces |= spacesTwo
	}
	if c.Bool("four-spaces") {
		s.spaces |= spacesFour
	}

	return s
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

func (s *stateIndentation) validate() error {
	switch s.spaces {
	case spacesType(0), spacesTwo, spacesFour:
		return nil
	}
	return errors.New(`choose either "two-spaces(ts)" or "four-spaces(fs)" or blank.`)
}
