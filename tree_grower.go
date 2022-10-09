package gtree

func newGrower(
	lastNodeFormat, intermedialNodeFormat branchFormat,
	enabledValidation bool,
) grower {
	return &defaultGrower{
		lastNodeFormat:        lastNodeFormat,
		intermedialNodeFormat: intermedialNodeFormat,
		enabledValidation:     enabledValidation,
	}
}

func newNopGrower() grower {
	return &nopGrower{}
}

type branchFormat struct {
	directly, indirectly string
}

type defaultGrower struct {
	lastNodeFormat        branchFormat
	intermedialNodeFormat branchFormat
	enabledValidation     bool
}

func (dg *defaultGrower) grow(roots []*Node) error {
	for _, root := range roots {
		if err := dg.assemble(root); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrower) assemble(current *Node) error {
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

func (dg *defaultGrower) assembleBranch(current *Node) error {
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

func (dg *defaultGrower) assembleBranchDirectly(current *Node) {
	if current == nil {
		return
	}
	if current.isRoot() {
		return
	}

	current.setPath(current.name)

	if current.isLastOfHierarchy() {
		current.setBranch(current.branch(), dg.lastNodeFormat.directly)
	} else {
		current.setBranch(current.branch(), dg.intermedialNodeFormat.directly)
	}
}

func (dg *defaultGrower) assembleBranchIndirectly(current, parent *Node) {
	if current == nil || parent == nil {
		return
	}
	if current.isRoot() {
		return
	}

	current.setPath(parent.name, current.path())

	if parent.isLastOfHierarchy() {
		current.setBranch(dg.lastNodeFormat.indirectly, current.branch())
	} else {
		current.setBranch(dg.intermedialNodeFormat.indirectly, current.branch())
	}
}

func (*defaultGrower) assembleBranchFinally(current, root *Node) {
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

func (dg *defaultGrower) enableValidation() {
	dg.enabledValidation = true
}

type nopGrower struct{}

func (*nopGrower) grow(_ []*Node) error { return nil }

func (*nopGrower) enableValidation() {}

var (
	_ grower = (*defaultGrower)(nil)
	_ grower = (*nopGrower)(nil)
)
