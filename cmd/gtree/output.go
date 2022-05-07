package main

import (
	"errors"
	"io"

	"github.com/ddddddO/gtree"
	color "github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func outputWithValidation(in io.Reader, options []gtree.Option) error {
	options = append(options, gtree.WithDryRun())
	return output(color.Output, in, options)
}

func output(out io.Writer, in io.Reader, options []gtree.Option) error {
	return gtree.Output(out, in, options...)
}

type stateOutputFormat struct {
	encodeJSON bool
	encodeYAML bool
	encodeTOML bool
}

func newStateOutputFormat(c *cli.Context) *stateOutputFormat {
	return &stateOutputFormat{
		encodeJSON: c.Bool("json"),
		encodeYAML: c.Bool("yaml"),
		encodeTOML: c.Bool("toml"),
	}
}

func (s *stateOutputFormat) decideOption() (gtree.Option, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	switch {
	case s.encodeJSON:
		return gtree.WithEncodeJSON(), nil
	case s.encodeYAML:
		return gtree.WithEncodeYAML(), nil
	case s.encodeTOML:
		return gtree.WithEncodeTOML(), nil
	}
	return nil, nil
}

func (s *stateOutputFormat) validate() error {
	if s.encodeJSON && s.encodeYAML && s.encodeTOML {
		return errors.New(`choose either "json(j)" or "yaml(y)" or "toml(t)".`)
	}
	if s.encodeJSON && s.encodeYAML {
		return errors.New(`choose either "json(j)" or "yaml(y)".`)
	}
	if s.encodeJSON && s.encodeTOML {
		return errors.New(`choose either "json(j)" or "toml(t)".`)
	}
	if s.encodeTOML && s.encodeYAML {
		return errors.New(`choose either "toml(t)" or "yaml(y)".`)
	}
	return nil
}
