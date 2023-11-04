//go:build !tinywasm

package gtree

// WalkerNode is used in user-defined function that can be executed with Walk/WalkProgrammably function.
type WalkerNode struct {
	name   string
	branch string
	row    string
	level  uint
	path   string
}

// Name returns name of node in completed tree structure.
func (wn *WalkerNode) Name() string {
	return wn.name
}

// Branch returns branch of node in completed tree structure.
func (wn *WalkerNode) Branch() string {
	return wn.branch
}

// Row returns row of node in completed tree structure.
func (wn *WalkerNode) Row() string {
	return wn.row
}

// Level returns level of node in completed tree structure.
func (wn *WalkerNode) Level() uint {
	return wn.level
}

// Path returns path of node in completed tree structure.
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
		row:    row,
		level:  current.hierarchy,
		path:   current.path(),
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
