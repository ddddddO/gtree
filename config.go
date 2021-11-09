package gtree

import (
	"github.com/pkg/errors"
)

var errInvalidOption = errors.New("invalid option")

type config struct {
	isTwoSpaces  bool
	isFourSpaces bool

	formatLastNode        branchFormat
	formatIntermedialNode branchFormat

	selectedMode mode
}

type mode int

const (
	modeDefault mode = iota
	modeJSON
)

func newConfig(optFns ...optFn) (*config, error) {
	c := &config{
		formatLastNode: branchFormat{
			directly:   "└──",
			indirectly: "    ",
		},
		formatIntermedialNode: branchFormat{
			directly:   "├──",
			indirectly: "│   ",
		},
		selectedMode: modeDefault,
	}
	for _, opt := range optFns {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.isTwoSpaces && c.isFourSpaces {
		return nil, errInvalidOption
	}
	return c, nil
}

// optFn is functional options pattern
type optFn func(*config) error

// IndentTwoSpaces returns function for two spaces indent input.
func IndentTwoSpaces() optFn {
	return func(c *config) error {
		c.isTwoSpaces = true
		return nil
	}
}

// IndentFourSpaces returns function for four spaces indent input.
func IndentFourSpaces() optFn {
	return func(c *config) error {
		c.isFourSpaces = true
		return nil
	}
}

// BranchFormatIntermedialNode returns function for branch format.
func BranchFormatIntermedialNode(directly, indirectly string) optFn {
	return func(c *config) error {
		c.formatIntermedialNode.directly = directly
		c.formatIntermedialNode.indirectly = indirectly
		return nil
	}
}

// BranchFormatLastNode returns function for branch format.
func BranchFormatLastNode(directly, indirectly string) optFn {
	return func(c *config) error {
		c.formatLastNode.directly = directly
		c.formatLastNode.indirectly = indirectly
		return nil
	}
}

// ModeJSON returns function for json output.
func ModeJSON() optFn {
	return func(c *config) error {
		c.selectedMode = modeJSON
		return nil
	}
}
