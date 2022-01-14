package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
			Name:        "file",
			Aliases:     []string{"f"},
			Usage:       "Markdown file path.",
			DefaultText: "stdin",
		},
		&cli.BoolFlag{
			Name:        "two-spaces",
			Aliases:     []string{"ts"},
			Usage:       "Markdown is Two Spaces indentation.",
			DefaultText: "tab spaces",
		},
		&cli.BoolFlag{
			Name:        "four-spaces",
			Aliases:     []string{"fs"},
			Usage:       "Markdown is Four Spaces indentation.",
			DefaultText: "tab spaces",
		},
	}

	outputFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:        "json",
			Aliases:     []string{"j"},
			Usage:       "Output JSON format.",
			DefaultText: "stdout",
		},
		&cli.BoolFlag{
			Name:        "yaml",
			Aliases:     []string{"y"},
			Usage:       "Output YAML format.",
			DefaultText: "stdout",
		},
		&cli.BoolFlag{
			Name:        "toml",
			Aliases:     []string{"t"},
			Usage:       "Output TOML format.",
			DefaultText: "stdout",
		},
		&cli.BoolFlag{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "Watching markdown file.",
		},
	}

	mkdirFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "dry-run",
			Aliases: []string{"d", "dr"},
			Usage:   "Dry run. Detects node that is invalid for directory generation. The order of the output and made directories does not always match.",
		},
		&cli.StringSliceFlag{
			Name:    "extension",
			Aliases: []string{"e", "ext"},
			Usage:   "Specified extension will be created as file.",
		},
	}

	app := &cli.App{
		Name:  "gtree",
		Usage: "This CLI outputs tree or makes directories from markdown.",
		Commands: []*cli.Command{
			{
				Name:    "output",
				Aliases: []string{"o", "out"},
				Usage:   "Output tree from markdown. Let's try 'gtree template | gtree output'. Output format is stdout or yaml or toml or json. Default stdout.",
				Flags:   append(commonFlags, outputFlags...),
				Action:  actionOutput,
			},
			{
				Name:    "mkdir",
				Aliases: []string{"m"},
				Usage:   "Make directories from markdown. It is possible to dry run. Let's try 'gtree template | gtree mkdir -e .go -e .md -e makefile'.",
				Flags:   append(commonFlags, mkdirFlags...),
				Action:  actionMkdir,
			},
			{
				Name:    "template",
				Aliases: []string{"t", "tmpl"},
				Usage:   "Output markdown template.",
				// Flags: NOTE: prepare various templates.
				Action: actionTemplate,
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
		if err := output(os.Stdout, os.Stdin, indentation, outputFormat, notDryrun, nil); err != nil {
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

		if err := output(os.Stdout, file, indentation, outputFormat, notDryrun, nil); err != nil {
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

				_ = output(os.Stdout, file, indentation, outputFormat, notDryrun, nil)
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

func actionMkdir(c *cli.Context) error {
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

	extensions := c.StringSlice("extension")
	if c.Bool("dry-run") {
		dryrun := true
		if err := output(os.Stdout, in, indentation, outputFormatStdout, dryrun, extensions); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	}

	if err := mkdir(in, indentation, extensions); err != nil {
		return cli.Exit(err, 1)
	}

	return nil
}

func mkdir(in io.Reader, indentation indentation, extensions []string) error {
	var options []gtree.OptFn

	switch indentation {
	case indentationTS:
		options = append(options, gtree.WithIndentTwoSpaces())
	case indentationFS:
		options = append(options, gtree.WithIndentFourSpaces())
	}

	options = append(options, gtree.WithFileExtension(extensions))

	return gtree.Mkdir(in, options...)
}

var template = strings.TrimLeft(`
- gtree
	- cmd
		- gtree
			- main.go
	- testdata
		- sample1.md
		- sample2.md
	- makefile
	- tree.go
`, "\n")

func actionTemplate(c *cli.Context) error {
	fmt.Print(template)
	return nil
}
