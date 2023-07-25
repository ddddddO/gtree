package gtree

import "context"

type config struct {
	lastNodeFormat        branchFormat
	intermedialNodeFormat branchFormat

	space          spaceType
	massive        bool
	ctx            context.Context
	encode         encode
	dryrun         bool
	fileExtensions []string
	targetDir      string
	strictVerify   bool
}

func newConfig(options []Option) *config {
	c := &config{
		lastNodeFormat: branchFormat{
			directly:   "└──",
			indirectly: "    ",
		},
		intermedialNodeFormat: branchFormat{
			directly:   "├──",
			indirectly: "│   ",
		},
		space:     spacesTab,
		massive:   false,
		encode:    encodeDefault,
		targetDir: ".",
	}
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt(c)
	}
	return c
}

// Option is functional options pattern
type Option func(*config)

// WithIndentTwoSpaces returns function for two spaces indent input.
func WithIndentTwoSpaces() Option {
	return func(c *config) {
		c.space = spacesTwo
	}
}

// WithIndentFourSpaces returns function for four spaces indent input.
func WithIndentFourSpaces() Option {
	return func(c *config) {
		c.space = spacesFour
	}
}

// WithBranchFormatIntermedialNode returns function for branch format.
func WithBranchFormatIntermedialNode(directly, indirectly string) Option {
	return func(c *config) {
		c.intermedialNodeFormat.directly = directly
		c.intermedialNodeFormat.indirectly = indirectly
	}
}

// WithBranchFormatLastNode returns function for branch format.
func WithBranchFormatLastNode(directly, indirectly string) Option {
	return func(c *config) {
		c.lastNodeFormat.directly = directly
		c.lastNodeFormat.indirectly = indirectly
	}
}

// WithMassive returns function for large amount roots.
func WithMassive(ctx context.Context) Option {
	return func(c *config) {
		c.massive = true

		if ctx == nil {
			ctx = context.Background()
		}
		c.ctx = ctx
	}
}

// WithEncodeJSON returns function for output json format.
func WithEncodeJSON() Option {
	return func(c *config) {
		c.encode = encodeJSON
	}
}

// WithEncodeYAML returns function for output yaml format.
func WithEncodeYAML() Option {
	return func(c *config) {
		c.encode = encodeYAML
	}
}

// WithEncodeTOML returns function for output toml format.
func WithEncodeTOML() Option {
	return func(c *config) {
		c.encode = encodeTOML
	}
}

// WithDryRun returns function for dry run. Detects node that is invalid for directory generation.
func WithDryRun() Option {
	return func(c *config) {
		c.dryrun = true
	}
}

// WithFileExtensions returns function for creating as a file instead of a directory.
func WithFileExtensions(extensions []string) Option {
	return func(c *config) {
		c.fileExtensions = extensions
	}
}

// WithTargetDir returns function for specifying directory. Default is current directory.
func WithTargetDir(dir string) Option {
	return func(c *config) {
		c.targetDir = dir
	}
}

// WithStrictVerify returns function for verifing directory strictly.
func WithStrictVerify() Option {
	return func(c *config) {
		c.strictVerify = true
	}
}
