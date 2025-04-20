package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ddddddO/gtree"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	commonFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Aliases:     []string{"f"},
			Usage:       "specify the path to markdown file.",
			DefaultText: "stdin",
		},
	}

	outputFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "massive",
			Aliases: []string{"m"},
			Usage:   "set this option when there are very many blocks of markdown.",
		},
		&cli.DurationFlag{
			Name:    "massive-timeout",
			Aliases: []string{"mt"},
			Usage:   "set this option if you want to set a timeout.",
			Action: func(_ context.Context, _ *cli.Command, v time.Duration) error {
				if v <= 0 {
					return errors.New("the timeout value should be greater than 0")
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:  "format",
			Usage: `set this option when specifying output format. "json", "yaml", "toml"`,
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
			Aliases: []string{"d"},
			Usage:   "dry run. detects node that is invalid for directory generation. the order of the output and made directories does not always match.",
		},
		&cli.StringSliceFlag{
			Name:    "extension",
			Aliases: []string{"e"},
			Usage:   "set this option if you want to create file instead of directory. for example, if you want to generate files with \".go\" extension: \"-e .go\"",
		},
		&cli.StringFlag{
			Name:        "target-dir",
			Usage:       "set this option if you want to specify the directory you want to make directory.",
			DefaultText: "current directory",
		},
	}

	verifyFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "target-dir",
			Usage:       "set this option if you want to specify the directory you want to verify.",
			DefaultText: "current directory",
		},
		&cli.BoolFlag{
			Name:        "strict",
			Usage:       "set this option if you want strict directory match validation.",
			DefaultText: "non strict",
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

	green := color.New(color.FgHiGreen).SprintFunc()

	app := &cli.Command{
		Name: "gtree",
		Usage: "This CLI uses Markdown to generate directory trees and directories itself, and also verifies directories." + "\n" +
			fmt.Sprintf("The symbols that can be used in Markdown are '%s', '%s', '%s', and '%s'.", green("-"), green("+"), green("*"), green("#")) + "\n" +
			"Within Markdown, indentation represents hierarchy. The indentation can be whatever you specify, but use the same pattern.",
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
				Name:    "verify",
				Aliases: []string{"vf"},
				Usage: "Verifies tree structure represented in markdown by comparing it with existing directories.\n" +
					"Let's try 'gtree template | gtree verify'.",
				Flags:  append(commonFlags, verifyFlags...),
				Before: notExistArgs,
				Action: actionVerify,
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
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Prints the version.",
				Before:  notExistArgs,
				Action: func(ctx context.Context, c *cli.Command) error {
					fmt.Printf("gtree version %s / revision %s\n", Version, Revision)
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

func notExistArgs(ctx context.Context, c *cli.Command) (context.Context, error) {
	if c.NArg() != 0 {
		return nil, errors.New("command line contains unnecessary arguments")
	}
	return ctx, nil
}

func actionOutput(ctx context.Context, c *cli.Command) error {
	oo, err := optionOutput(c)
	if err != nil {
		return exitErrOpts(err)
	}
	var om gtree.Option
	if c.Bool("massive") {
		om = gtree.WithMassive(context.Background())
	}
	if c.Duration("massive-timeout") > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), c.Duration("massive-timeout"))
		defer cancel()
		om = gtree.WithMassive(ctx)
	}
	options := []gtree.Option{oo, om}

	markdownPath := c.String("file")
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

func actionMkdir(ctx context.Context, c *cli.Command) error {
	var (
		in  = os.Stdin
		err error
	)
	if !isInputStdin(c.String("file")) {
		in, err = os.Open(c.String("file"))
		if err != nil {
			return exitErrOpen(err)
		}
		defer in.Close()
	}

	options := []gtree.Option{gtree.WithTargetDir(c.String("target-dir")), gtree.WithFileExtensions(c.StringSlice("extension"))}
	if c.Bool("massive") {
		options = append(options, gtree.WithMassive(context.Background()))
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

func actionVerify(ctx context.Context, c *cli.Command) error {
	var (
		in  = os.Stdin
		err error
	)
	if !isInputStdin(c.String("file")) {
		in, err = os.Open(c.String("file"))
		if err != nil {
			return exitErrOpen(err)
		}
		defer in.Close()
	}

	options := []gtree.Option{gtree.WithTargetDir(c.String("target-dir"))}
	if c.Bool("strict") {
		options = append(options, gtree.WithStrictVerify())
	}

	if err := verify(in, options); err != nil {
		return exitErrVerify(err)
	}
	return nil
}

func actionTemplate(ctx context.Context, c *cli.Command) error {
	if c.Bool("description") {
		return description.println()
	}
	return directory.println()
}

const treeMakerURL = "https://ddddddo.github.io/gtree/"

func actionWeb(ctx context.Context, c *cli.Command) error {
	_ = openWeb(treeMakerURL, c.Bool("wsl"))
	fmt.Printf("See: %s\n", treeMakerURL)
	return nil
}
