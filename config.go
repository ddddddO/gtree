package gtree

import "context"

// Definition of optional functions for Node.
type NodeOption func(*Node)

// WithDuplicationAllowed is an optional function that ensures Add method call creates and returns a new node even if a node with the same text already exists at the same hierarchy. It can be specified in the NewRoot function.
func WithDuplicationAllowed() NodeOption {
	return func(n *Node) {
		n.allowDuplicates = true
	}
}

type config struct {
	lastBranch string
	midBranch  string
	hLine      string
	vLine      string

	// todo: refactor
	lastNodeFormat        *branchFormat
	intermedialNodeFormat *branchFormat

	massive        bool
	ctx            context.Context
	encode         encode
	dryrun         bool
	fileExtensions []string
	targetDir      string
	strictVerify   bool

	noUseIterOfSimpleOutput bool
}

func newConfig(options []Option) *config {
	c := &config{
		lastBranch:            "└",
		midBranch:             "├",
		hLine:                 "──",
		vLine:                 "│",
		lastNodeFormat:        &branchFormat{},
		intermedialNodeFormat: &branchFormat{},
		massive:               false,
		encode:                encodeDefault,
		targetDir:             ".",
	}
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt(c)
	}
	c.lastNodeFormat.directly = c.lastBranch + c.hLine
	c.lastNodeFormat.indirectly = "    "
	c.intermedialNodeFormat.directly = c.midBranch + c.hLine
	c.intermedialNodeFormat.indirectly = c.vLine + "   "
	return c
}

// Option is functional options pattern
type Option func(*config)

// WithLastBranch allows you to change the branch parts.
func WithLastBranch(s string) Option {
	return func(c *config) {
		c.lastBranch = s
	}
}

// WithMidBranch allows you to change the branch parts.
func WithMidBranch(s string) Option {
	return func(c *config) {
		c.midBranch = s
	}
}

// WithHLine allows you to change the branch parts.
func WithHLine(s string) Option {
	return func(c *config) {
		c.hLine = s
	}
}

// WithVLine allows you to change the branch parts.
func WithVLine(s string) Option {
	return func(c *config) {
		c.vLine = s
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

// Deprecated: This is for benchmark testing.
func WithNoUseIterOfSimpleOutput() Option {
	return func(c *config) {
		c.noUseIterOfSimpleOutput = true
	}
}
