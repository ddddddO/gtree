package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ddddddO/gtree"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func output(in io.Reader, options []gtree.Option) error {
	return gtree.Output(os.Stdout, in, options...)
}

func outputWithValidation(in io.Reader, options []gtree.Option) error {
	options = append(options, gtree.WithDryRun())
	return gtree.Output(color.Output, in, options...)
}

const intervalms = 500 * time.Millisecond

func outputContinuously(markdownPath string, options []gtree.Option) error {
	ticker := time.NewTicker(intervalms)
	defer ticker.Stop()
	var preFileModTime time.Time
	for range ticker.C {
		if err := func() error {
			f, err := os.Open(markdownPath)
			if err != nil {
				return err
			}
			defer f.Close()

			fi, err := f.Stat()
			if err != nil {
				return err
			}

			if fi.ModTime() != preFileModTime {
				preFileModTime = fi.ModTime()
				_ = output(f, options)
				fmt.Println()
			}
			return nil
		}(); err != nil {
			return err
		}
	}
	return nil
}

type encodeType uint

const (
	encodeJSON encodeType = 1 << iota
	encodeYAML
	encodeTOML
)

type stateOutputFormat struct {
	encode encodeType
}

func optionOutput(c *cli.Context) (gtree.Option, error) {
	s := &stateOutputFormat{}

	if c.Bool("json") {
		s.encode |= encodeJSON
	}
	if c.Bool("yaml") {
		s.encode |= encodeYAML
	}
	if c.Bool("toml") {
		s.encode |= encodeTOML
	}

	return s.decideOption()
}

func (s *stateOutputFormat) decideOption() (gtree.Option, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	switch s.encode {
	case encodeJSON:
		return gtree.WithEncodeJSON(), nil
	case encodeYAML:
		return gtree.WithEncodeYAML(), nil
	case encodeTOML:
		return gtree.WithEncodeTOML(), nil
	}
	return nil, nil
}

const encodeDefault = encodeType(0)

func (s *stateOutputFormat) validate() error {
	switch s.encode {
	case encodeDefault, encodeJSON, encodeYAML, encodeTOML:
		return nil
	}
	return errors.New(`choose either "json(j)" or "yaml(y)" or "toml(t)" or blank.`)
}
