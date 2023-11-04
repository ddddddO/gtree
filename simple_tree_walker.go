//go:build !tinywasm

package gtree

type WalkerNode struct {
	name   string
	branch string
	row    string
	path   string
}

func (wn *WalkerNode) Name() string {
	return wn.name
}

func (wn *WalkerNode) Branch() string {
	return wn.branch
}

func (wn *WalkerNode) Row() string {
	return wn.row
}

func (wn *WalkerNode) Path() string {
	return wn.path
}

func newWalkerSimple() walkerSimple {
	return &defaultWalkerSimple{}
}

type defaultWalkerSimple struct{}

func (dw *defaultWalkerSimple) walk(roots []*Node, cb func(*WalkerNode) error) error {
	for _, root := range roots {
		if err := dw.walkNode(root, cb); err != nil {
			return err
		}
	}
	return nil
}

func (dw *defaultWalkerSimple) walkNode(current *Node, cb func(*WalkerNode) error) error {
	row := current.name
	if !current.isRoot() {
		row = current.branch() + " " + current.name
	}

	wn := &WalkerNode{
		name:   current.name,
		branch: current.branch(),
		path:   current.path(),
		row:    row,
	}
	if err := cb(wn); err != nil {
		return err
	}

	for _, child := range current.children {
		if err := dw.walkNode(child, cb); err != nil {
			return err
		}
	}
	return nil
}
