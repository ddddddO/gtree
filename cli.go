package gtree

import (
	"io"
)

type CLI struct {
	stdout io.Writer
	stderr io.Writer // TODO: まだ使ってない.そもそもstdxxxという命名は止める。out, errでいいかな

	cfg *config
}

func NewCLI(stdout, stderr io.Writer) *CLI {
	return &CLI{
		stdout: stdout,
		stderr: stderr,
	}
}

// Output outputs a tree to w with r as Markdown format input.
func (c *CLI) Output(r io.Reader) error {
	rg := newRootGenerator(r, c.cfg.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	tree := newTree(c.cfg, roots)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(c.stdout)
}

// Mkdir makes directories.
func (c *CLI) Mkdir(r io.Reader) error {
	rg := newRootGenerator(r, c.cfg.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	tree := newTree(c.cfg, roots)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.mkdir()
}

// IndentTwoSpaces is for two spaces indent input.
func (c *CLI) IndentTwoSpaces() *CLI {
	c.cfg.space = spacesTwo
	return c
}

// IndentFourSpaces is for four spaces indent input.
func (c *CLI) IndentFourSpaces() *CLI {
	c.cfg.space = spacesFour
	return c
}

// FormatIntermedialNode is for branch format.
func (c *CLI) FormatIntermedialNode(directly, indirectly string) *CLI {
	c.cfg.intermedialNodeFormat.directly = directly
	c.cfg.intermedialNodeFormat.indirectly = indirectly
	return c
}

// FormatLastNode is for branch format.
func (c *CLI) FormatLastNode(directly, indirectly string) *CLI {
	c.cfg.lastNodeFormat.directly = directly
	c.cfg.lastNodeFormat.indirectly = indirectly
	return c
}

// EncodeJSON is for output json format.
func (c *CLI) EncodeJSON() *CLI {
	c.cfg.encode = encodeJSON
	return c
}

// EncodeYAML is for output yaml format.
func (c *CLI) EncodeYAML() *CLI {
	c.cfg.encode = encodeYAML
	return c
}

// EncodeTOML is for output toml format.
func (c *CLI) EncodeTOML() *CLI {
	c.cfg.encode = encodeTOML
	return c
}

// DryRun is for dry run. Detects node that is invalid for directory generation.
func (c *CLI) DryRun() *CLI {
	c.cfg.dryrun = true
	return c
}

// FileExtensions is for creating as a file instead of a directory.
func (c *CLI) FileExtensions(extensions []string) *CLI {
	c.cfg.fileExtensions = extensions
	return c
}
