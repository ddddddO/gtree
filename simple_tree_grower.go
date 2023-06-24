//go:build !wasm

package gtree

func newGrowerSimple(
	lastNodeFormat, intermedialNodeFormat branchFormat,
	enabledValidation bool,
) growerSimple {
	return &defaultGrowerSimple{
		lastNodeFormat:        lastNodeFormat,
		intermedialNodeFormat: intermedialNodeFormat,
		enabledValidation:     enabledValidation,
	}
}

func newNopGrowerSimple() growerSimple {
	return &nopGrowerSimple{}
}

type branchFormat struct {
	directly, indirectly string
}

type defaultGrowerSimple struct {
	lastNodeFormat        branchFormat
	intermedialNodeFormat branchFormat
	enabledValidation     bool
}

func (dg *defaultGrowerSimple) grow(roots []*Node) error {
	for _, root := range roots {
		if err := dg.assemble(root); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrowerSimple) assemble(current *Node) error {
	if err := dg.assembleBranch(current); err != nil {
		return err
	}

	for _, child := range current.children {
		if err := dg.assemble(child); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrowerSimple) assembleBranch(current *Node) error {
	current.clean() // 例えば、MkdirProgrammably funcでrootノードを使いまわすと、前回func実行時に形成されたノードの枝が残ったまま追記されてしまうため。

	dg.assembleBranchDirectly(current)

	// go back to the root to form a branch.
	tmpParent := current.parent
	if tmpParent != nil {
		for ; !tmpParent.isRoot(); tmpParent = tmpParent.parent {
			dg.assembleBranchIndirectly(current, tmpParent)
		}
	}

	dg.assembleBranchFinally(current, tmpParent)

	if dg.enabledValidation {
		return current.validatePath()
	}
	return nil
}

func (dg *defaultGrowerSimple) assembleBranchDirectly(current *Node) {
	if current == nil || current.isRoot() {
		return
	}

	current.setPath(current.name)

	if current.isLastOfHierarchy() {
		current.setBranch(current.branch(), dg.lastNodeFormat.directly)
	} else {
		current.setBranch(current.branch(), dg.intermedialNodeFormat.directly)
	}
}

func (dg *defaultGrowerSimple) assembleBranchIndirectly(current, parent *Node) {
	if current == nil || parent == nil || current.isRoot() {
		return
	}

	current.setPath(parent.name, current.path())

	if parent.isLastOfHierarchy() {
		current.setBranch(dg.lastNodeFormat.indirectly, current.branch())
	} else {
		current.setBranch(dg.intermedialNodeFormat.indirectly, current.branch())
	}
}

func (*defaultGrowerSimple) assembleBranchFinally(current, root *Node) {
	if current == nil {
		return
	}

	if root != nil {
		current.setPath(root.path(), current.path())
	}

	if current.isRoot() {
		current.setBranch(current.name, "\n")
	} else {
		current.setBranch(current.branch(), " ", current.name, "\n")
	}
}

func (dg *defaultGrowerSimple) enableValidation() {
	dg.enabledValidation = true
}

type nopGrowerSimple struct{}

func (*nopGrowerSimple) grow(_ []*Node) error { return nil }

func (*nopGrowerSimple) enableValidation() {}

var (
	_ growerSimple = (*defaultGrowerSimple)(nil)
	_ growerSimple = (*nopGrowerSimple)(nil)
)
