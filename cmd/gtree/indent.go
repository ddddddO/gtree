package main

import (
	"github.com/ddddddO/gtree"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func decideIndentation(c *cli.Context) (gtree.OptFn, error) {
	if c.Bool("two-spaces") && c.Bool("four-spaces") {
		return nil, errors.New(`choose either "two-spaces(ts)" or "four-spaces(fs)".`)
	}

	if c.Bool("two-spaces") {
		return gtree.WithIndentTwoSpaces(), nil
	}
	if c.Bool("four-spaces") {
		return gtree.WithIndentFourSpaces(), nil
	}
	return nil, nil
}
