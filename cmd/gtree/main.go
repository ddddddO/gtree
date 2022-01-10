package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/ddddddO/gtree"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	commonFlags := []cli.Flag{
		&cli.PathFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "Input Markdown file path.",
		},
		&cli.BoolFlag{
			Name:    "two-spaces",
			Aliases: []string{"ts"},
			Usage:   "Markdown is Two Spaces indentation.",
		},
		&cli.BoolFlag{
			Name:    "four-spaces",
			Aliases: []string{"fs"},
			Usage:   "Markdown is Four Spaces indentation.",
		},
	}

	outputFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "json",
			Aliases: []string{"j"},
			Usage:   "Output JSON format.",
		},
		&cli.BoolFlag{
			Name:    "yaml",
			Aliases: []string{"y"},
			Usage:   "Output YAML format.",
		},
		&cli.BoolFlag{
			Name:    "toml",
			Aliases: []string{"t"},
			Usage:   "Output TOML format.",
		},
		&cli.BoolFlag{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "Watching Markdown file.",
		},
	}

	generateFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "dry-run",
			Aliases: []string{"d", "dr"},
			Usage:   "Dry run.",
		},
	}

	app := &cli.App{
		Name:  "gtree",
		Usage: "This CLI outputs tree or generates directories.",
		Commands: []*cli.Command{
			{
				Name:    "output",
				Aliases: []string{"o", "out"},
				Usage:   "Output tree, stdout or yaml or toml or json. Default stdout.",
				Flags:   append(commonFlags, outputFlags...),
				Action:  actionOutput,
			},
			{
				Name:    "generate",
				Aliases: []string{"g", "gen"},
				Usage:   "Generate directories. It is possible to dry run.",
				Flags:   append(commonFlags, generateFlags...),
				Action:  actionGenerate,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Output gtree version.",
				Action: func(c *cli.Context) error {
					fmt.Printf("gtree version %s / revision %s\n", Version, Revision)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func actionOutput(c *cli.Context) error {
	const notDryrun = false

	indentation, err := decideIndentation(c)
	if err != nil {
		return cli.Exit(err, 1)
	}

	outputFormat, err := decideOutputFormat(c)
	if err != nil {
		return cli.Exit(err, 1)
	}

	markdownPath := c.Path("file")
	if markdownPath == "" || markdownPath == "-" {
		if err := execute(os.Stdout, os.Stdin, indentation, outputFormat, notDryrun); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	}

	if !c.Bool("watch") {
		file, err := os.Open(markdownPath)
		if err != nil {
			return cli.Exit(err, 1)
		}
		defer file.Close()

		if err := execute(os.Stdout, file, indentation, outputFormat, notDryrun); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	}

	// watching markdown file
	ticker := time.NewTicker(1 * time.Second)
	var preFileModTime time.Time
	for range ticker.C {
		func() {
			file, err := os.Open(markdownPath)
			// FIXME: ?
			if err != nil {
				fmt.Errorf("%+v", err)
				os.Exit(1)
			}
			defer file.Close()

			fileInfo, err := file.Stat()
			// FIXME: ?
			if err != nil {
				fmt.Errorf("%+v", err)
				os.Exit(1)
			}

			if fileInfo.ModTime() != preFileModTime {
				preFileModTime = fileInfo.ModTime()

				_ = execute(os.Stdout, file, indentation, outputFormat, notDryrun)
			}
		}()
	}

	return nil
}

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

func execute(out io.Writer, in io.Reader, indentation indentation, outputFormat outputFormat, dryrun bool) error {
	var options []gtree.OptFn

	switch indentation {
	case indentationTS:
		options = append(options, gtree.IndentTwoSpaces())
	case indentationFS:
		options = append(options, gtree.IndentFourSpaces())
	}

	switch outputFormat {
	case outputFormatJSON:
		options = append(options, gtree.EncodeJSON())
	case outputFormatYAML:
		options = append(options, gtree.EncodeYAML())
	case outputFormatTOML:
		options = append(options, gtree.EncodeTOML())
	}

	if dryrun {
		options = append(options, gtree.WithDryRun())
	}

	return gtree.Execute(out, in, options...)
}

func actionGenerate(c *cli.Context) error {
	indentation, err := decideIndentation(c)
	if err != nil {
		return cli.Exit(err, 1)
	}

	markdownPath := c.Path("file")
	in := os.Stdin
	if markdownPath != "" && markdownPath != "-" {
		in, err = os.Open(markdownPath)
		if err != nil {
			return cli.Exit(err, 1)
		}
		defer in.Close()
	}

	// NOTE: この時に、無効なディレクトリ名がある判定もしたい
	if c.Bool("dry-run") {
		dryrun := true
		if err := execute(os.Stdout, in, indentation, outputFormatStdout, dryrun); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	}

	if err := generate(in, indentation); err != nil {
		return cli.Exit(err, 1)
	}

	return nil
}

func generate(in io.Reader, indentation indentation) error {
	var options []gtree.OptFn

	switch indentation {
	case indentationTS:
		options = append(options, gtree.IndentTwoSpaces())
	case indentationFS:
		options = append(options, gtree.IndentFourSpaces())
	}

	return gtree.Generate(in, options...)
}
