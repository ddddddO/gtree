package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/ddddddO/gtree"
	"github.com/goccy/go-yaml"
	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v3"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	app := &cli.Command{
		Name:    "xtree",
		Usage:   "This CLI uses {JSON|YAML|TOML} to generate directory tree",
		Version: fmt.Sprintf("%s / revision %s", Version, Revision),
		Commands: []*cli.Command{
			{
				Name:    "output",
				Aliases: []string{"o", "out"},
				Usage: "Outputs tree from {JSON|YAML|TOML}.\n" +
					"Let's try 'cat {.json|.yaml|.toml} | xtree output'.",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "omit-index",
						Aliases: []string{"omit", "o"},
						Usage:   "set this option when you do not want to display array indices.",
					},
					&cli.BoolFlag{
						Name:    "allow-duplicate",
						Aliases: []string{"a"},
						Usage:   "set this option when you want to allow duplicate node names at the same level.",
					},
				},
				Before: notExistArgs,
				Action: actionOutput,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Prints the version.",
				Before:  notExistArgs,
				Action: func(ctx context.Context, c *cli.Command) error {
					fmt.Printf("xtree version %s / revision %s\n", Version, Revision)
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func actionOutput(ctx context.Context, c *cli.Command) error {
	root := gtree.NewRoot(".")
	if c.Bool("allow-duplicate") {
		root = gtree.NewRoot(".", gtree.WithDuplicationAllowed())
	}
	omitIndex := c.Bool("omit-index")

	return output(os.Stdout, os.Stdin, root, omitIndex)
}

func output(w io.Writer, r io.Reader, root *gtree.Node, omitIndex bool) error {
	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	data, err := parseData(input)
	if err != nil {
		return err
	}
	walk(root, "", data, omitIndex)

	return gtree.OutputFromRoot(w, root)
}

func parseData(data []byte) (any, error) {
	var v any

	if err := json.Unmarshal(data, &v); err == nil {
		return v, nil
	}
	if err := toml.Unmarshal(data, &v); err == nil {
		return v, nil
	}
	if err := yaml.Unmarshal(data, &v); err == nil {
		return v, nil
	}

	return nil, errors.New("data is in an unsupported format or is invalid")
}

func walk(parent *gtree.Node, key string, value any, omitIndex bool) {
	switch v := value.(type) {
	case map[string]any:
		node := parent
		if key != "" {
			node = parent.Add(key)
		}

		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			walk(node, k, v[k], omitIndex)
		}

	case map[any]any:
		node := parent
		if key != "" {
			node = parent.Add(key)
		}
		for k, val := range v {
			walk(node, fmt.Sprintf("%v", k), val, omitIndex)
		}

	case []any:
		node := parent
		if key != "" {
			node = parent.Add(key)
		}

		for i, item := range v {
			if omitIndex {
				walk(node, "", item, omitIndex)
			} else {
				indexNode := node.Add(fmt.Sprintf("[%d]", i))
				walk(indexNode, "", item, omitIndex)
			}
		}

	default:
		val := strings.ReplaceAll(fmt.Sprintf("%v", v), "\n", "\\n")
		if key != "" {
			keyNode := parent.Add(key)
			keyNode.Add(val)
		} else {
			parent.Add(val)
		}
	}
}

func notExistArgs(ctx context.Context, c *cli.Command) (context.Context, error) {
	if c.NArg() != 0 {
		return nil, errors.New("command line contains unnecessary arguments")
	}
	return ctx, nil
}
