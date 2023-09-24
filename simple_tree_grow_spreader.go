//go:build !tinywasm

package gtree

import (
	"fmt"
	"io"
)

func newGrowSpreaderSimple(
	lastNodeFormat, intermedialNodeFormat branchFormat,
	enabledValidation bool,
) growSpreaderSimple {
	return &defaultGrowSpreaderSimple{
		defaultGrowerSimple: &defaultGrowerSimple{
			lastNodeFormat:        lastNodeFormat,
			intermedialNodeFormat: intermedialNodeFormat,
			enabledValidation:     enabledValidation,
		},
	}
}

type defaultGrowSpreaderSimple struct {
	*defaultGrowerSimple
	w io.Writer
}

func (dg *defaultGrowSpreaderSimple) growAndSpread(w io.Writer, roots []*Node) error {
	dg.w = w
	for _, root := range roots {
		if err := dg.assembleAndPrint(root); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrowSpreaderSimple) assembleAndPrint(current *Node) error {
	if err := dg.assembleBranch(current); err != nil {
		return err
	}

	ret := current.name + "\n"
	if !current.isRoot() {
		ret = current.branch() + " " + current.name + "\n"
	}
	fmt.Fprint(dg.w, ret)

	for _, child := range current.children {
		if err := dg.assembleAndPrint(child); err != nil {
			return err
		}
	}
	return nil
}

var (
	_ growSpreaderSimple = (*defaultGrowSpreaderSimple)(nil)
)
