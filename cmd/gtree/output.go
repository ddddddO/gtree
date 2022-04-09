package main

import (
	"io"

	"github.com/ddddddO/gtree"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func outputDryrun(out io.Writer, in io.Reader, indentation gtree.Option, extensions []string) error {
	return output(out, in, indentation, nil, true, extensions)
}

func outputNotDryrun(out io.Writer, in io.Reader, indentation, outputFormat gtree.Option) error {
	return output(out, in, indentation, outputFormat, false, nil)
}

func output(out io.Writer, in io.Reader, indentation gtree.Option, outputFormat gtree.Option, dryrun bool, extensions []string) error {
	options := []gtree.Option{gtree.WithFileExtensions(extensions)}
	if indentation != nil {
		options = append(options, indentation)
	}
	if outputFormat != nil {
		options = append(options, outputFormat)
	}
	if dryrun {
		options = append(options, gtree.WithDryRun())
	}

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
