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

func decideOutputFormat(c *cli.Context) (gtree.Option, error) {
	if err := validateOutputFormat(c); err != nil {
		return nil, err
	}

	switch {
	case c.Bool("json"):
		return gtree.WithEncodeJSON(), nil
	case c.Bool("yaml"):
		return gtree.WithEncodeYAML(), nil
	case c.Bool("toml"):
		return gtree.WithEncodeTOML(), nil
	}
	return nil, nil
}

func validateOutputFormat(c *cli.Context) error {
	if c.Bool("json") && c.Bool("yaml") && c.Bool("toml") {
		return errors.New(`choose either "json(j)" or "yaml(y)" or "toml(t)".`)
	}
	if c.Bool("json") && c.Bool("yaml") {
		return errors.New(`choose either "json(j)" or "yaml(y)".`)
	}
	if c.Bool("json") && c.Bool("toml") {
		return errors.New(`choose either "json(j)" or "toml(t)".`)
	}
	if c.Bool("toml") && c.Bool("yaml") {
		return errors.New(`choose either "toml(t)" or "yaml(y)".`)
	}
	return nil
}
