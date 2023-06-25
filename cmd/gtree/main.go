package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ddddddO/gtree"
	"github.com/fatih/color"
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
			Usage:       "specify the path to markdown file.",
			DefaultText: "stdin",
		},
		&cli.BoolFlag{
			Name:        "two-spaces",
			Aliases:     []string{"ts"},
			Usage:       "set this option when the markdown indent is 2 spaces.",
			DefaultText: "tab spaces",
		},
		&cli.BoolFlag{
			Name:        "four-spaces",
			Aliases:     []string{"fs"},
			Usage:       "set this option when the markdown indent is 4 spaces.",
			DefaultText: "tab spaces",
		},
		&cli.BoolFlag{
			Name:    "massive",
			Aliases: []string{"m"},
			Usage:   "set this option when there are very many blocks of markdown.",
		},
	}

	outputFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:        "json",
			Aliases:     []string{"j"},
			Usage:       "set this option when outputting JSON.",
			DefaultText: "tree",
		},
		&cli.BoolFlag{
			Name:        "yaml",
			Aliases:     []string{"y"},
			Usage:       "set this option when outputting YAML.",
			DefaultText: "tree",
		},
		&cli.BoolFlag{
			Name:        "toml",
			Aliases:     []string{"t"},
			Usage:       "set this option when outputting TOML.",
			DefaultText: "tree",
		},
		&cli.BoolFlag{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "follow changes in markdown file.",
		},
	}

	mkdirFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "dry-run",
			Aliases: []string{"d", "dr"},
			Usage: "dry run. detects node that is invalid for directory generation.\n" +
				"the order of the output and made directories does not always match.",
		},
		&cli.StringSliceFlag{
			Name:    "extension",
			Aliases: []string{"e", "ext"},
			Usage: "set this option if you want to create file instead of directory.\n" +
				"for example, if you want to generate files with \".go\" extension: \"-e .go\"",
		},
	}

	templateFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "description",
			Aliases: []string{"desc"},
			Usage:   "show gtree CLI description.",
		},
	}

	webFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "wsl",
			Usage: "specify this option if you are using WSL.",
		},
	}

	gocodeFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "godeps-to-tree",
			Usage: "output Go program to convert Go package dependency list to tree.",
		},
	}

	green := color.New(color.FgHiGreen).SprintFunc()

	app := &cli.App{
		Name: "gtree",
		Usage: "This CLI generates directory trees and the directories itself using Markdown.\n" +
			fmt.Sprintf("The symbols that can be used in Markdown are '%s', '%s', '%s', and '%s'.", green("-"), green("+"), green("*"), green("#")),
		Version: fmt.Sprintf("%s / revision %s", Version, Revision),
		Commands: []*cli.Command{
			{
				Name:    "output",
				Aliases: []string{"o", "out"},
				Usage: "Outputs tree from markdown.\n" +
					"Let's try 'gtree template | gtree output'.",
				Flags:  append(commonFlags, outputFlags...),
				Before: notExistArgs,
				Action: actionOutput,
			},
			{
				Name:    "mkdir",
				Aliases: []string{"m"},
				Usage: "Makes directories and files from markdown. It is possible to dry run.\n" +
					"Let's try 'gtree template | gtree mkdir -e .go -e .md -e Makefile'.",
				Flags:  append(commonFlags, mkdirFlags...),
				Before: notExistArgs,
				Action: actionMkdir,
			},
			{
				Name:    "template",
				Aliases: []string{"t", "tmpl"},
				Usage:   "Outputs markdown template. Use it to try out gtree CLI.",
				Flags:   templateFlags,
				Before:  notExistArgs,
				Action:  actionTemplate,
			},
			{
				Name:    "web",
				Aliases: []string{"w", "www"},
				Usage:   "Opens \"Tree Maker\" in your browser and shows the URL in terminal.",
				Flags:   webFlags,
				Before:  notExistArgs,
				Action:  actionWeb,
			},
			{
				Name:    "gocode",
				Aliases: []string{"gc", "code"},
				Usage:   "Outputs a sample Go program calling \"gtree\" package.",
				Flags:   gocodeFlags,
				Before:  notExistArgs,
				Action:  actionGoCode,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Prints the version.",
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
	if c.Bool("massive") {
		options = append(options, gtree.WithMassive())
	}

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
	if c.Bool("massive") {
		options = append(options, gtree.WithMassive())
	}

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
		return description.println()
	}

	return directory.println()
}

const treeMakerURL = "https://ddddddo.github.io/gtree/"

func actionWeb(c *cli.Context) error {
	_ = openWeb(treeMakerURL, c.Bool("wsl"))
	fmt.Printf("See: %s\n", treeMakerURL)
	return nil
}

func actionGoCode(c *cli.Context) error {
	if c.Bool("godeps-to-tree") {
		return goDependencesToTree.println()
	}

	return findToTree.println()
}
