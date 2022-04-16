package gtree

import (
	"path/filepath"
)

// 関心事は各ノードの枝を組み立てること
type grower interface {
	grow([]*Node) error
	enableValidation()
}

func newGrower(
	encode encode,
	lastNodeFormat, intermedialNodeFormat branchFormat,
	enabledValidation bool,
) grower {
	if encode != encodeDefault {
		return &noopGrower{}
	}
	return &defaultGrower{
		lastNodeFormat:        lastNodeFormat,
		intermedialNodeFormat: intermedialNodeFormat,
		enabledValidation:     enabledValidation,
	}
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
	if err := dg.assemblebrnch(current); err != nil {
		return err
	}

	for _, child := range current.Children {
		if err := dg.assemble(child); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrower) assemblebrnch(current *Node) error {
	if current.isRoot() {
		return nil
	}

	dg.assemblebrnchDirectly(current)

	// go back to the root to form a brnch
	tmpParent := current.parent
	for {
		if tmpParent.isRoot() {
			dg.assemblebrnchFinally(current, tmpParent)

			if dg.enabledValidation {
				if err := current.validatePath(); err != nil {
					return err
				}
			}
			break
		}

		dg.assemblebrnchIndirectly(current, tmpParent)

		tmpParent = tmpParent.parent
	}
	return nil
}

func (dg *defaultGrower) assemblebrnchDirectly(current *Node) {
	current.brnch.path = current.Name

	if current.isLastOfHierarchy() {
		current.brnch.value += dg.lastNodeFormat.directly
	} else {
		current.brnch.value += dg.intermedialNodeFormat.directly
	}
}

func (dg *defaultGrower) assemblebrnchIndirectly(current, parent *Node) {
	current.brnch.path = filepath.Join(parent.Name, current.brnch.path)

	if parent.isLastOfHierarchy() {
		current.brnch.value = dg.lastNodeFormat.indirectly + current.brnch.value
	} else {
		current.brnch.value = dg.intermedialNodeFormat.indirectly + current.brnch.value
	}
}

func (*defaultGrower) assemblebrnchFinally(current, root *Node) {
	current.brnch.path = filepath.Join(root.Name, current.brnch.path)
}

func (dg *defaultGrower) enableValidation() {
	dg.enabledValidation = true
}

type noopGrower struct{}

func (*noopGrower) grow(_ []*Node) error { return nil }

func (*noopGrower) enableValidation() {}
