package main

import (
	"github.com/ddddddO/gtree"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

type stateIndentation struct {
	spacesTwo  bool
	spacesFour bool
}

func newStateIndentation(c *cli.Context) *stateIndentation {
	return &stateIndentation{
		spacesTwo:  c.Bool("two-spaces"),
		spacesFour: c.Bool("four-spaces"),
	}
}

func (s *stateIndentation) decideOption() (gtree.Option, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	if s.spacesTwo {
		return gtree.WithIndentTwoSpaces(), nil
	}
	if s.spacesFour {
		return gtree.WithIndentFourSpaces(), nil
	}
	return nil, nil
}

func (s *stateIndentation) validate() error {
	if s.spacesTwo && s.spacesFour {
		return errors.New(`choose either "two-spaces(ts)" or "four-spaces(fs)".`)
	}
	return nil
}
