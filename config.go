package gtree

type encode int

const (
	encodeDefault encode = iota
	encodeJSON
	encodeYAML
	encodeTOML
)

type config struct {
	formatLastNode        branchFormat
	formatIntermedialNode branchFormat

	space          spaceType
	encode         encode
	dryrun         bool
	fileExtensions []string
}

func newConfig(OptFns ...OptFn) (*config, error) {
	c := &config{
		formatLastNode: branchFormat{
			directly:   "└──",
			indirectly: "    ",
		},
		formatIntermedialNode: branchFormat{
			directly:   "├──",
			indirectly: "│   ",
		},
		space:  tabSpaces,
		encode: encodeDefault,
	}
	for _, opt := range OptFns {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// OptFn is functional options pattern
type OptFn func(*config) error

// WithIndentTwoSpaces returns function for two spaces indent input.
func WithIndentTwoSpaces() OptFn {
	return func(c *config) error {
		c.space = twoSpaces
		return nil
	}
}

// WithIndentFourSpaces returns function for four spaces indent input.
func WithIndentFourSpaces() OptFn {
	return func(c *config) error {
		c.space = fourSpaces
		return nil
	}
}

// WithBranchFormatIntermedialNode returns function for branch format.
func WithBranchFormatIntermedialNode(directly, indirectly string) OptFn {
	return func(c *config) error {
		c.formatIntermedialNode.directly = directly
		c.formatIntermedialNode.indirectly = indirectly
		return nil
	}
}

// WithBranchFormatLastNode returns function for branch format.
func WithBranchFormatLastNode(directly, indirectly string) OptFn {
	return func(c *config) error {
		c.formatLastNode.directly = directly
		c.formatLastNode.indirectly = indirectly
		return nil
	}
}

// WithEncodeJSON returns function for output json format.
func WithEncodeJSON() OptFn {
	return func(c *config) error {
		c.encode = encodeJSON
		return nil
	}
}

// WithEncodeYAML returns function for output yaml format.
func WithEncodeYAML() OptFn {
	return func(c *config) error {
		c.encode = encodeYAML
		return nil
	}
}

// WithEncodeTOML returns function for output toml format.
func WithEncodeTOML() OptFn {
	return func(c *config) error {
		c.encode = encodeTOML
		return nil
	}
}

// WithDryRun returns function for dry run. Detects node that is invalid for directory generation.
func WithDryRun() OptFn {
	return func(c *config) error {
		c.dryrun = true
		return nil
	}
}

// WithFileExtension returns function for creating as a file instead of a directory.
func WithFileExtension(extensions []string) OptFn {
	return func(c *config) error {
		c.fileExtensions = extensions
		return nil
	}
}
