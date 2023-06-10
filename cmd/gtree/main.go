package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ddddddO/gtree"
	"github.com/urfave/cli/v2"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

const (
	exitCodeErrOpts = iota + 1
	exitCodeErrOutput
	exitCodeErrOpen
	exitCodeErrMkdir
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
		return cli.Exit(err, exitCodeErrOpts)
	}
	oo, err := optionOutput(c)
	if err != nil {
		return cli.Exit(err, exitCodeErrOpts)
	}
	options := []gtree.Option{oi, oo}

	markdownPath := c.Path("file")
	if isInputStdin(markdownPath) {
		if err := output(os.Stdin, options); err != nil {
			return cli.Exit(err, exitCodeErrOutput)
		}
		return nil
	}

	if !c.Bool("watch") {
		f, err := os.Open(markdownPath)
		if err != nil {
			return cli.Exit(err, exitCodeErrOpen)
		}
		defer f.Close()

		if err := output(f, options); err != nil {
			return cli.Exit(err, exitCodeErrOutput)
		}
		return nil
	}

	if err := watchMarkdownAndOutput(markdownPath, options); err != nil {
		return cli.Exit(err, exitCodeErrOutput)
	}

	return nil
}

func isInputStdin(path string) bool {
	return path == "" || path == "-"
}

const intervalms = 500 * time.Millisecond

func watchMarkdownAndOutput(markdownPath string, options []gtree.Option) error {
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

func actionMkdir(c *cli.Context) error {
	markdownPath := c.Path("file")
	in := os.Stdin
	var err error
	if !isInputStdin(markdownPath) {
		in, err = os.Open(markdownPath)
		if err != nil {
			return cli.Exit(err, exitCodeErrOpen)
		}
		defer in.Close()
	}

	oi, err := optionIndentation(c)
	if err != nil {
		return cli.Exit(err, exitCodeErrOpts)
	}
	oe := gtree.WithFileExtensions(c.StringSlice("extension"))
	options := []gtree.Option{oi, oe}

	if c.Bool("dry-run") {
		if err := outputWithValidation(in, options); err != nil {
			return cli.Exit(err, exitCodeErrOutput)
		}
		return nil
	}

	if err := mkdir(in, options); err != nil {
		return cli.Exit(err, exitCodeErrMkdir)
	}

	return nil
}

const template = `
- gtree
	- cmd
		- gtree
			- main.go
	- testdata
		- sample1.md
		- sample2.md
	- makefile
	- tree.go
`

const description = "- # Description\n" +
	"	- Output tree from markdown or programmatically.\n" +
	"		- Output format is tree|yaml|toml|json.\n" +
	"		- Default tree.\n" +
	"	- Make directories from markdown or programmatically.\n" +
	"		- It is possible to dry run.\n" +
	"		- You can use `-e` flag to make specified extensions as file.\n" +
	"	- Output a markdown template that can be used with either `output` subcommand or `mkdir` subcommand.\n" +
	"	- Provide CLI, Go library and Web.\n"

func actionTemplate(c *cli.Context) error {
	if c.Bool("description") {
		fmt.Print(strings.TrimLeft(description, "\n"))
		return nil
	}

	fmt.Print(strings.TrimLeft(template, "\n"))
	return nil
}
