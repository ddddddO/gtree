package main

import (
	"io"

	"github.com/ddddddO/gtree"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

type outputFormat uint

const (
	outputFormatUndefined outputFormat = iota
	outputFormatStdout
	outputFormatJSON
	outputFormatYAML
	outputFormatTOML
)

func decideOutputFormat(c *cli.Context) (outputFormat, error) {
	if c.Bool("json") && c.Bool("yaml") && c.Bool("toml") {
		return outputFormatUndefined, errors.New(`choose either "json(j)" or "yaml(y)" or "toml(t)".`)
	}
	if c.Bool("json") && c.Bool("yaml") {
		return outputFormatUndefined, errors.New(`choose either "json(j)" or "yaml(y)".`)
	}
	if c.Bool("json") && c.Bool("toml") {
		return outputFormatUndefined, errors.New(`choose either "json(j)" or "toml(t)".`)
	}
	if c.Bool("toml") && c.Bool("yaml") {
		return outputFormatUndefined, errors.New(`choose either "toml(t)" or "yaml(y)".`)
	}

	outputFormat := outputFormatStdout
	switch {
	case c.Bool("json"):
		outputFormat = outputFormatJSON
	case c.Bool("yaml"):
		outputFormat = outputFormatYAML
	case c.Bool("toml"):
		outputFormat = outputFormatTOML
	}
	return outputFormat, nil
}

func output(out io.Writer, in io.Reader, indentation indentation, outputFormat outputFormat, dryrun bool, extensions []string) error {
	var options []gtree.OptFn

	switch indentation {
	case indentationTS:
		options = append(options, gtree.WithIndentTwoSpaces())
	case indentationFS:
		options = append(options, gtree.WithIndentFourSpaces())
	}

	switch outputFormat {
	case outputFormatJSON:
		options = append(options, gtree.WithEncodeJSON())
	case outputFormatYAML:
		options = append(options, gtree.WithEncodeYAML())
	case outputFormatTOML:
		options = append(options, gtree.WithEncodeTOML())
	}

	if dryrun {
		options = append(options, gtree.WithDryRun())
	}

	options = append(options, gtree.WithFileExtension(extensions))

	return gtree.Output(out, in, options...)
}
