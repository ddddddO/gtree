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

func optionOutput(c *cli.Context) (gtree.Option, error) {
	switch c.String("format") {
	case "json":
		return gtree.WithEncodeJSON(), nil
	case "yaml":
		return gtree.WithEncodeYAML(), nil
	case "toml":
		return gtree.WithEncodeTOML(), nil
	case "":
		return nil, nil
	default:
		return nil, errors.New(`specify either "json" or "yaml" or "toml"`)
	}
}
