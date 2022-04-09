package gtree

import (
	"path/filepath"
)

type grower interface {
	grow([]*Node) error
	setDryRun(bool)
}

func newGrower(
	encode encode,
	lastNodeFormat, intermedialNodeFormat branchFormat,
	dryrunMode bool,
) grower {
	switch encode {
	case encodeDefault:
		return &defaultGrower{
			lastNodeFormat:        lastNodeFormat,
			intermedialNodeFormat: intermedialNodeFormat,
			dryrunMode:            dryrunMode,
		}
	default:
		return &noopGrower{}
	}
}

type defaultGrower struct {
	lastNodeFormat        branchFormat
	intermedialNodeFormat branchFormat
	dryrunMode            bool
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

	for _, child := range current.Children {
		if err := dg.assemble(child); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrower) assembleBranch(current *Node) error {
	if current.isRoot() {
		return nil
	}

	dg.assembleBranchDirectly(current)

	// go back to the root to form a branch
	tmpParent := current.parent
	for {
		if tmpParent.isRoot() {
			dg.assembleBranchFinally(current, tmpParent)

			if dg.dryrunMode {
				if err := current.validatePath(); err != nil {
					return err
				}
			}
			break
		}

		dg.assembleBranchIndirectly(current, tmpParent)

		tmpParent = tmpParent.parent
	}
	return nil
}

func (dg *defaultGrower) assembleBranchDirectly(current *Node) {
	current.branch.path = current.Name

	if current.isLastOfHierarchy() {
		current.branch.value += dg.lastNodeFormat.directly
	} else {
		current.branch.value += dg.intermedialNodeFormat.directly
	}
}

func (dg *defaultGrower) assembleBranchIndirectly(current, parent *Node) {
	current.branch.path = filepath.Join(parent.Name, current.branch.path)

	if parent.isLastOfHierarchy() {
		current.branch.value = dg.lastNodeFormat.indirectly + current.branch.value
	} else {
		current.branch.value = dg.intermedialNodeFormat.indirectly + current.branch.value
	}
}

func (*defaultGrower) assembleBranchFinally(current, root *Node) {
	current.branch.path = filepath.Join(root.Name, current.branch.path)
}

// TODO: どうにかしたい
func (dg *defaultGrower) setDryRun(dryrun bool) {
	dg.dryrunMode = dryrun
}

type noopGrower struct{}

func (ng *noopGrower) grow(_ []*Node) error {
	return nil
}

func (ng *noopGrower) setDryRun(_ bool) {
}
