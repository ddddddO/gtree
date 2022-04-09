package gtree

type config struct {
	lastNodeFormat        branchFormat
	intermedialNodeFormat branchFormat

	space          spaceType
	encode         encode
	dryrun         bool
	fileExtensions []string
}

func newConfig(options ...Option) (*config, error) {
	c := &config{
		lastNodeFormat: branchFormat{
			directly:   "└──",
			indirectly: "    ",
		},
		intermedialNodeFormat: branchFormat{
			directly:   "├──",
			indirectly: "│   ",
		},
		space:  spacesTab,
		encode: encodeDefault,
	}
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Option is functional options pattern
type Option func(*config) error

// WithIndentTwoSpaces returns function for two spaces indent input.
func WithIndentTwoSpaces() Option {
	return func(c *config) error {
		c.space = spacesTwo
		return nil
	}
}

// WithIndentFourSpaces returns function for four spaces indent input.
func WithIndentFourSpaces() Option {
	return func(c *config) error {
		c.space = spacesFour
		return nil
	}
}

// WithBranchFormatIntermedialNode returns function for branch format.
func WithBranchFormatIntermedialNode(directly, indirectly string) Option {
	return func(c *config) error {
		c.intermedialNodeFormat.directly = directly
		c.intermedialNodeFormat.indirectly = indirectly
		return nil
	}
}

// WithBranchFormatLastNode returns function for branch format.
func WithBranchFormatLastNode(directly, indirectly string) Option {
	return func(c *config) error {
		c.lastNodeFormat.directly = directly
		c.lastNodeFormat.indirectly = indirectly
		return nil
	}
}

// WithEncodeJSON returns function for output json format.
func WithEncodeJSON() Option {
	return func(c *config) error {
		c.encode = encodeJSON
		return nil
	}
}

// WithEncodeYAML returns function for output yaml format.
func WithEncodeYAML() Option {
	return func(c *config) error {
		c.encode = encodeYAML
		return nil
	}
}

// WithEncodeTOML returns function for output toml format.
func WithEncodeTOML() Option {
	return func(c *config) error {
		c.encode = encodeTOML
		return nil
	}
}

// WithDryRun returns function for dry run. Detects node that is invalid for directory generation.
func WithDryRun() Option {
	return func(c *config) error {
		c.dryrun = true
		return nil
	}
}

// WithFileExtensions returns function for creating as a file instead of a directory.
func WithFileExtensions(extensions []string) Option {
	return func(c *config) error {
		c.fileExtensions = extensions
		return nil
	}
}
