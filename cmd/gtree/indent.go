package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

type indentation uint

const (
	indentationUndefined indentation = iota
	indentationTab
	indentationTS
	indentationFS
)

func decideIndentation(c *cli.Context) (indentation, error) {
	if c.Bool("two-spaces") && c.Bool("four-spaces") {
		return indentationUndefined, errors.New(`choose either "two-spaces(ts)" or "four-spaces(fs)".`)
	}

	indentation := indentationTab
	if c.Bool("two-spaces") {
		indentation = indentationTS
	}
	if c.Bool("four-spaces") {
		indentation = indentationFS
	}
	return indentation, nil
}
