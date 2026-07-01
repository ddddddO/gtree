//go:build !tinywasm

package gtree

import (
	"io"
)

func newGrowSpreaderSimple(
	lastNodeFormat, intermedialNodeFormat *branchFormat,
) growSpreaderSimple {
	return &defaultGrowSpreaderSimple{
		defaultGrowerSimple: &defaultGrowerSimple{
			lastNodeFormat:        lastNodeFormat,
			intermedialNodeFormat: intermedialNodeFormat,
			enabledValidation:     false,
		},
	}
}

type defaultGrowSpreaderSimple struct {
	*defaultGrowerSimple
	w io.Writer
}

func (dgs *defaultGrowSpreaderSimple) growAndSpread(w io.Writer, roots []*Node) error {
	dgs.w = w
	for _, root := range roots {
		if err := dgs.assembleAndPrint(root); err != nil {
			return err
		}
	}
	return nil
}

func (dgs *defaultGrowSpreaderSimple) assembleAndPrint(current *Node) error {
	if err := dgs.assembleBranch(current); err != nil {
		return err
	}

	if current.isRoot() {
		_, _ = io.WriteString(dgs.w, current.value)
		_, _ = io.WriteString(dgs.w, "\n")
	} else {
		_, _ = io.WriteString(dgs.w, current.branch())
		_, _ = io.WriteString(dgs.w, " ")
		_, _ = io.WriteString(dgs.w, current.value)
		_, _ = io.WriteString(dgs.w, "\n")
	}

	for _, child := range current.children {
		if err := dgs.assembleAndPrint(child); err != nil {
			return err
		}
	}
	return nil
}

// 以降、simple_tree_grower.go の assembleBranch系メソッドとの違いは、setPathの有無
// こっちのメソッドはOutputFromRootからプログラム的に呼び出されるもので、パフォーマンスを気にする必要がある
// そして、Node.setPathの有無でベンチマークを取ると、ない方が圧倒的に改善がみられた
// setPathが必要なのはMkdir系関数なので、こちらのメソッドにはsetPath(と関連の処理)は不要なため落としている
func (dgs *defaultGrowSpreaderSimple) assembleBranch(current *Node) error {
	current.clean()

	dgs.assembleBranchDirectly(current)

	// go back to the root to form a branch.
	tmpParent := current.parent
	if tmpParent != nil {
		for ; !tmpParent.isRoot(); tmpParent = tmpParent.parent {
			dgs.assembleBranchIndirectly(current, tmpParent)
		}
	}

	return nil
}

func (dgs *defaultGrowSpreaderSimple) assembleBranchDirectly(current *Node) {
	if current == nil || current.isRoot() {
		return
	}

	if current.isLastOfHierarchy() {
		current.appendBranch(dgs.lastNodeFormat.directly)
	} else {
		current.appendBranch(dgs.intermedialNodeFormat.directly)
	}
}

func (dgs *defaultGrowSpreaderSimple) assembleBranchIndirectly(current, parent *Node) {
	if current == nil || parent == nil || current.isRoot() {
		return
	}

	if parent.isLastOfHierarchy() {
		current.prependBranch(dgs.lastNodeFormat.indirectly)
	} else {
		current.prependBranch(dgs.intermedialNodeFormat.indirectly)
	}
}

var (
	_ growSpreaderSimple = (*defaultGrowSpreaderSimple)(nil)
)
