package main

import (
	"io"

	"github.com/ddddddO/gtree"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func decideOutputFormat(c *cli.Context) (gtree.OptFn, error) {
	if c.Bool("json") && c.Bool("yaml") && c.Bool("toml") {
		return nil, errors.New(`choose either "json(j)" or "yaml(y)" or "toml(t)".`)
	}
	if c.Bool("json") && c.Bool("yaml") {
		return nil, errors.New(`choose either "json(j)" or "yaml(y)".`)
	}
	if c.Bool("json") && c.Bool("toml") {
		return nil, errors.New(`choose either "json(j)" or "toml(t)".`)
	}
	if c.Bool("toml") && c.Bool("yaml") {
		return nil, errors.New(`choose either "toml(t)" or "yaml(y)".`)
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

func output(out io.Writer, in io.Reader, indentation gtree.OptFn, outputFormat gtree.OptFn, dryrun bool, extensions []string) error {
	var options []gtree.OptFn
	if indentation != nil {
		options = append(options, indentation)
	}
	if outputFormat != nil {
		options = append(options, outputFormat)
	}
	if dryrun {
		options = append(options, gtree.WithDryRun())
	}
	options = append(options, gtree.WithFileExtension(extensions))

	return gtree.Output(out, in, options...)
}
