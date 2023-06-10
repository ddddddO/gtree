package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
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
			DefaultText: "tree",
		},
		&cli.BoolFlag{
			Name:        "yaml",
			Aliases:     []string{"y"},
			Usage:       "Output YAML format.",
			DefaultText: "tree",
		},
		&cli.BoolFlag{
			Name:        "toml",
			Aliases:     []string{"t"},
			Usage:       "Output TOML format.",
			DefaultText: "tree",
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

	templateFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "description",
			Aliases: []string{"desc"},
			Usage:   "Show gtree description.",
		},
	}

	app := &cli.App{
		Name:  "gtree",
		Usage: "This CLI outputs tree or makes directories from markdown.",
		Commands: []*cli.Command{
			{
				Name:    "output",
				Aliases: []string{"o", "out"},
				Usage:   "Output tree from markdown. Let's try 'gtree template | gtree output'. Output format is tree or yaml or toml or json. Default tree.",
				Flags:   append(commonFlags, outputFlags...),
				Before:  notExistArgs,
				Action:  actionOutput,
			},
			{
				Name:    "mkdir",
				Aliases: []string{"m"},
				Usage:   "Make directories(and files) from markdown. It is possible to dry run. Let's try 'gtree template | gtree mkdir -e .go -e .md -e makefile'.",
				Flags:   append(commonFlags, mkdirFlags...),
				Before:  notExistArgs,
				Action:  actionMkdir,
			},
			{
				Name:    "template",
				Aliases: []string{"t", "tmpl"},
				Usage:   "Output markdown template.",
				Flags:   templateFlags,
				Before:  notExistArgs,
				Action:  actionTemplate,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Output gtree version.",
				Before:  notExistArgs,
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

func notExistArgs(c *cli.Context) error {
	if c.NArg() != 0 {
		return errors.New("command line contains unnecessary arguments")
	}
	return nil
}

func actionOutput(c *cli.Context) error {
	oi, err := optionIndentation(c)
	if err != nil {
		return exitErrOpts(err)
	}
	oo, err := optionOutput(c)
	if err != nil {
		return exitErrOpts(err)
	}
	options := []gtree.Option{oi, oo}

	markdownPath := c.Path("file")
	if isInputStdin(markdownPath) {
		if err := output(os.Stdin, options); err != nil {
			return exitErrOutput(err)
		}
		return nil
	}

	if c.Bool("watch") {
		if err := outputContinuously(markdownPath, options); err != nil {
			return exitErrOutput(err)
		}
	} else {
		f, err := os.Open(markdownPath)
		if err != nil {
			return exitErrOpen(err)
		}
		defer f.Close()

		if err := output(f, options); err != nil {
			return exitErrOutput(err)
		}
	}

	return nil
}

func actionMkdir(c *cli.Context) error {
	var (
		in  = os.Stdin
		err error
	)
	if !isInputStdin(c.Path("file")) {
		in, err = os.Open(c.Path("file"))
		if err != nil {
			return exitErrOpen(err)
		}
		defer in.Close()
	}

	oi, err := optionIndentation(c)
	if err != nil {
		return exitErrOpts(err)
	}
	oe := gtree.WithFileExtensions(c.StringSlice("extension"))
	options := []gtree.Option{oi, oe}

	if c.Bool("dry-run") {
		if err := outputWithValidation(in, options); err != nil {
			return exitErrOutput(err)
		}
		return nil
	}

	if err := mkdir(in, options); err != nil {
		return exitErrMkdir(err)
	}

	return nil
}

func isInputStdin(path string) bool {
	return path == "" || path == "-"
}

func actionTemplate(c *cli.Context) error {
	if c.Bool("description") {
		fmt.Print(description)
		return nil
	}

	fmt.Print(strings.TrimLeft(template, "\n"))
	return nil
}
