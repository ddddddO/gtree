package gtree

type encode int

const (
	encodeDefault encode = iota
	encodeJSON
	encodeYAML
	encodeTOML
	encodeMsgPack
)

type config struct {
	formatLastNode        branchFormat
	formatIntermedialNode branchFormat

	space  spaceType
	encode encode
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

// IndentTwoSpaces returns function for two spaces indent input.
func IndentTwoSpaces() OptFn {
	return func(c *config) error {
		c.space = twoSpaces
		return nil
	}
}

// IndentFourSpaces returns function for four spaces indent input.
func IndentFourSpaces() OptFn {
	return func(c *config) error {
		c.space = fourSpaces
		return nil
	}
}

// BranchFormatIntermedialNode returns function for branch format.
func BranchFormatIntermedialNode(directly, indirectly string) OptFn {
	return func(c *config) error {
		c.formatIntermedialNode.directly = directly
		c.formatIntermedialNode.indirectly = indirectly
		return nil
	}
}

// BranchFormatLastNode returns function for branch format.
func BranchFormatLastNode(directly, indirectly string) OptFn {
	return func(c *config) error {
		c.formatLastNode.directly = directly
		c.formatLastNode.indirectly = indirectly
		return nil
	}
}

// EncodeJSON returns function for output json format.
func EncodeJSON() OptFn {
	return func(c *config) error {
		c.encode = encodeJSON
		return nil
	}
}

// EncodeYAML returns function for output yaml format.
func EncodeYAML() OptFn {
	return func(c *config) error {
		c.encode = encodeYAML
		return nil
	}
}

// EncodeTOML returns function for output toml format.
func EncodeTOML() OptFn {
	return func(c *config) error {
		c.encode = encodeTOML
		return nil
	}
}

// EncodeMsgPack returns function for output message pack format.
func EncodeMsgPack() OptFn {
	return func(c *config) error {
		c.encode = encodeMsgPack
		return nil
	}
}
